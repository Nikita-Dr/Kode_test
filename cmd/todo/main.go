package main

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
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
	"os"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	//TODO узнать нахера я подключал pkg/logger/sl https://www.youtube.com/watch?v=rCJvW2xgnk0&t=3082s
	storage, err := postgres.New(cfg.DB)
	if err != nil {
		log.Error("failed to connect to database", sl.Err(err))
		os.Exit(1)
	}

	// Инициализация слоя репозитория
	notesRepo := repository.NewNoteRepository(storage)
	userRepo := repository.NewUserRepository(storage)

	//err = notesRepo.CreateNote(entity.Note{Id: 2, Note: "note 2"})
	//if err != nil {
	//	log.Error("failed to create note", sl.Err(err))
	//}
	//
	//notes := []entity.Note{}
	//notes, err = notesRepo.GetNotes()
	//fmt.Println(notes)

	//Инициализация пакетов
	textChecker := yadex.NewTextChecker()
	jwtAuth := jwt_auth.NewJwtAuth(cfg.Jwt)

	// Инициализация слоя usecase
	textValidator := usecase.NewTextValidator(textChecker)
	userUC := usecase2.NewAuthUseCase(userRepo, jwtAuth)
	noteUC := usecase.NewNoteUseCase(notesRepo, textValidator)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(middleware2.ParseToken(jwtAuth))

	// Инициализация слоя Controller
	controller.NewNoteController(router, noteUC, log)
	controller.NewAuthController(router, userUC, log)

	//TODO run server

	log.Info("starting server", slog.String("adress", cfg.Http.Host+":"+cfg.Http.Port))
	if err = server.NewHttpServer(cfg.Http, router).Start(); err != nil {
		log.Error("failed to start server", sl.Err(err))
	}

	log.Error("server stopped")
}
