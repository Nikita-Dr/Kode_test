package model

import "Kode_test/internal/domain/note/entity"

type NoteDTO struct {
	Id   int    `json:"id"`
	Note string `json:"note"`
}

func GetNoteDTO(note entity.Note) NoteDTO {
	return NoteDTO{
		Id:   note.Id,
		Note: note.Note,
	}
}

func GetNoteListDTO(noteList []entity.Note) []NoteDTO {
	noteListDTO := make([]NoteDTO, 0, len(noteList))
	for _, note := range noteList {
		noteListDTO = append(noteListDTO, GetNoteDTO(note))
	}
	return noteListDTO
}
