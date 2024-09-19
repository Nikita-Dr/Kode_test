package controller

import (
	"Kode_test/internal/domain/note/model"
	"Kode_test/pkg/api/response"
	"Kode_test/pkg/logger/sl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
)

type NoteUsecase interface {
	CreateNote(noteDto model.CreateNoteDTO) error
	GetNotes(userId int) ([]model.ResponseNoteDTO, error)
}

type NoteController struct {
	noteUsecase NoteUsecase
	log         *slog.Logger
}

func NewNoteController(handler *chi.Mux, noteUsecase NoteUsecase, log *slog.Logger) {
	controller := &NoteController{noteUsecase: noteUsecase, log: log}

	handler.Get("/notes", controller.getNotes)
	handler.Post("/notes", controller.createNote)
}

func (c *NoteController) getNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := c.noteUsecase.GetNotes(r.Context().Value("userID").(int))
	if err != nil {
		c.log.Error("failed to get notes", sl.Err(err))
		render.JSON(w, r, response.Error("failed to get notes"))
		return
	}
	render.JSON(w, r, notes)
}

func (c *NoteController) createNote(w http.ResponseWriter, r *http.Request) {
	var noteDto model.CreateNoteDTO

	err := render.DecodeJSON(r.Body, &noteDto)
	noteDto.UserId = r.Context().Value("userID").(int)

	if err != nil {
		c.log.Error("failed to decode request body", sl.Err(err))
		render.JSON(w, r, response.Error("failed to decode request body"))
		return
	}

	//c.log.Info("request body decoded", slog.Any("requested", r.Body))

	if err = c.noteUsecase.CreateNote(noteDto); err != nil {
		c.log.Error("failed to create note", sl.Err(err))
		render.JSON(w, r, response.Error("failed to create note"))
		return
	}

	render.JSON(w, r, response.OK())
}
