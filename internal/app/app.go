package app

import (
	"Kode_test/config"
	usecase2 "Kode_test/internal/domain/auth/usecase"
	"Kode_test/internal/domain/note/usecase"
	"Kode_test/internal/infrastucture/controller"
	middleware2 "Kode_test/internal/infrastucture/controller/middleware"
	"Kode_test/internal/infrastucture/repository"
	"Kode_test/pkg/jwt_auth"
	"Kode_test/pkg/logger/sl"
	"Kode_test/pkg/server"
	"Kode_test/pkg/storage/postgres"
	"Kode_test/pkg/yadex"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
	"os"
)

type App struct {
	db          *postgres.Storage
	http_server *server.Server
	handler     *chi.Mux
	jwtKey      string
	log         *slog.Logger
}

func NewApp(cfg *config.Config) App {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// pkg/logger/sl для вывода ошибок
	storage, err := postgres.New(cfg.DB)
	if err != nil {
		log.Error("failed to connect to database", sl.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	srv := server.NewHttpServer(cfg.Http, router)

	return App{
		db:          storage,
		http_server: srv,
		handler:     router,
		jwtKey:      cfg.Jwt,
		log:         log,
	}
}

func (app *App) Run() error {
	// Инициализация слоя репозитория
	notesRepo := repository.NewNoteRepository(app.db)
	userRepo := repository.NewUserRepository(app.db)

	//Инициализация пакетов
	textChecker := yadex.NewTextChecker()
	jwtAuth := jwt_auth.NewJwtAuth(app.jwtKey)

	// Инициализация слоя usecase
	textValidator := usecase.NewTextValidator(textChecker)
	userUC := usecase2.NewAuthUseCase(userRepo, jwtAuth)
	noteUC := usecase.NewNoteUseCase(notesRepo, textValidator)

	app.handler.Use(middleware.RequestID)
	app.handler.Use(middleware.Logger)
	app.handler.Use(middleware.Recoverer)

	app.handler.Use(middleware2.ParseToken(jwtAuth))

	// Инициализация слоя Controller
	controller.NewNoteController(app.handler, noteUC, app.log)
	controller.NewAuthController(app.handler, userUC, app.log)

	app.log.Info("starting server", slog.String("adress ", app.http_server.GetAddr()))
	if err := app.http_server.Start(); err != nil {
		app.log.Error("failed to start server", sl.Err(err))
	}

	return nil
}

func (app *App) Shutdown(ctx context.Context) error {
	if err := app.http_server.Stop(ctx); err != nil {
		return err
	}
	return nil
}
