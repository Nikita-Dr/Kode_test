package model

import "Kode_test/internal/domain/note/entity"

type NoteDTO struct {
	ID   int    `json:"id"`
	Note string `json:"note"`
}

func GetNoteDTO(note entity.Note) NoteDTO {
	return NoteDTO{
		ID:   note.Id,
		Note: note.Note,
	}
}

func GetNoteListDTO(noteList []entity.Note) []NoteDTO {
	noteListDTO := []NoteDTO{}
	for _, note := range noteList {
		noteListDTO = append(noteListDTO, GetNoteDTO(note))
	}
	return noteListDTO
}
