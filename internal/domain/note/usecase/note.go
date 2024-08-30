package usecase

import (
	"Kode_test/internal/domain/note/entity"
	"Kode_test/internal/domain/note/model"
	"fmt"
)

type NoteRepository interface {
	CreateNote(note entity.Note) error
	GetNotes() ([]entity.Note, error)
}

type NoteUseCase struct {
	noteRepo NoteRepository
}

func NewNoteUseCase(noteRepo NoteRepository) *NoteUseCase {
	return &NoteUseCase{noteRepo: noteRepo}
}

func (u *NoteUseCase) CreateNote(noteDTO model.NoteDTO) error {
	note := entity.NoteFromDTO(noteDTO.Note)
	if err := u.noteRepo.CreateNote(note); err != nil {
		return fmt.Errorf("usecase - NoteUseCase - CreateNote: %w", err)
	}
	return nil
}

func (u *NoteUseCase) GetNotes() ([]model.NoteDTO, error) {
	noteEntityList, err := u.noteRepo.GetNotes()
	if err != nil {
		return nil, fmt.Errorf("usecase - NoteUseCase - GetNotes: %w", err)
	}
	return model.GetNoteListDTO(noteEntityList), nil
}
