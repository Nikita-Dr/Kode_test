package controller

import (
	"Kode_test/internal/domain/note/model"
	"Kode_test/pkg/api/response"
	"Kode_test/pkg/logger/sl"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
)

type NoteUsecase interface {
	CreateNote(noteDto model.NoteDTO) error
	GetNotes() ([]model.NoteDTO, error)
}

type NoteController struct {
	noteUsecase NoteUsecase
}

func New(log *slog.Logger, noteUsecase NoteUsecase) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "controller - NoteController - New"
		log = log.With(
			slog.String("cp", op),
		)

		var noteDto model.NoteDTO

		err := render.DecodeJSON(request.Body, &noteDto)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(writer, request, response.Error("failed to decode request body"))

			return
		}

		log.Info("request body decoded", slog.Any("requested", request.Body))

	}
}
