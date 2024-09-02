package model

import "Kode_test/internal/domain/note/entity"

type CreateNoteDTO struct {
	Id     int
	Note   string `json:"note"`
	UserId int
}

type ResponseNoteDTO struct {
	Id   int
	Note string `json:"note"`
}

func GetNoteDTO(note entity.Note) ResponseNoteDTO {
	return ResponseNoteDTO{
		Id:   note.Id,
		Note: note.Note,
	}
}

func GetNoteListDTO(noteList []entity.Note) []ResponseNoteDTO {
	noteListDTO := make([]ResponseNoteDTO, 0, len(noteList))
	for _, note := range noteList {
		noteListDTO = append(noteListDTO, GetNoteDTO(note))
	}
	return noteListDTO
}
