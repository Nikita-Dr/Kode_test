package main

import (
	"Kode_test/config"
	"Kode_test/internal/domain/note/usecase"
	"Kode_test/internal/infrastucture/repository"
	"Kode_test/pkg/logger/sl"
	"Kode_test/pkg/storage/postgres"
	"github.com/go-chi/chi/v5"
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

	notesRepo := repository.NewNoteRepository(storage)
	_ = notesRepo

	//err = notesRepo.CreateNote(entity.Note{Id: 2, Note: "note 2"})
	//if err != nil {
	//	log.Error("failed to create note", sl.Err(err))
	//}
	//
	//notes := []entity.Note{}
	//notes, err = notesRepo.GetNotes()
	//fmt.Println(notes)

	noteUC := usecase.NewNoteUseCase(notesRepo)

	//noteListDTO := []model.NoteDTO{}
	//noteListDTO, err = noteUC.GetNotes()
	//fmt.Println(noteListDTO)

	//TODO init router

	router := chi.NewRouter()

	_ = noteUC
	_ = router

	//TODO run server
}
