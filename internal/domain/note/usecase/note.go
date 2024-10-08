package usecase

import (
	"Kode_test/internal/domain/note/entity"
	"Kode_test/internal/domain/note/model"
	"fmt"
)

type NoteRepository interface {
	CreateNote(note entity.Note) error
	GetNotes(userId int) ([]entity.Note, error)
}

type TextValidator interface {
	ValidateText(inputText string) (string, error)
}

type NoteUseCase struct {
	noteRepo      NoteRepository
	textValidator TextValidator
}

func NewNoteUseCase(noteRepo NoteRepository, textValidator TextValidator) *NoteUseCase {
	return &NoteUseCase{
		noteRepo:      noteRepo,
		textValidator: textValidator,
	}
}

func (u *NoteUseCase) CreateNote(noteDTO model.CreateNoteDTO) error {
	note := entity.NoteFromDTO(noteDTO.Note, noteDTO.UserId)

	verifiedText, err := u.textValidator.ValidateText(note.Note)
	if err != nil {
		return fmt.Errorf("usecase - NoteUseCase - CheckText: %w", err)
	}
	note.UpdateNote(verifiedText)

	if err = u.noteRepo.CreateNote(note); err != nil {
		return fmt.Errorf("usecase - NoteUseCase - CreateNote: %w", err)
	}
	return nil
}

func (u *NoteUseCase) GetNotes(userId int) ([]model.ResponseNoteDTO, error) {
	noteEntityList, err := u.noteRepo.GetNotes(userId)
	if err != nil {
		return nil, fmt.Errorf("usecase - NoteUseCase - GetNotes: %w", err)
	}
	return model.GetNoteListDTO(noteEntityList), nil
}
