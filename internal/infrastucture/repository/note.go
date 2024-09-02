package repository

import (
	"Kode_test/internal/domain/note/entity"
	"Kode_test/pkg/storage/postgres"
	"fmt"
)

type NoteRepository struct {
	db *postgres.Storage
}

func NewNoteRepository(db *postgres.Storage) *NoteRepository {
	return &NoteRepository{db: db}
}

func (n *NoteRepository) CreateNote(note entity.Note) error {
	stmt, err := n.db.Prepare("INSERT INTO notes(note, user_id) VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("NoteRepository - CreateNote - Prepare:%v", err)
	}
	_, err = stmt.Exec(note.Note, note.UserId)
	if err != nil {
		return fmt.Errorf("NoteRepository - CreateNote - Exec:%v", err)
	}
	return nil
}

func (n *NoteRepository) GetNotes(userId int) ([]entity.Note, error) {
	notes := []entity.Note{}

	rows, err := n.db.Query("SELECT id, note FROM notes WHERE user_id = $1", userId)
	if err != nil {
		return nil, fmt.Errorf("NoteRepository - GetNotes - Query:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		note := entity.Note{}
		err = rows.Scan(&note.Id, &note.Note)
		if err != nil {
			return nil, fmt.Errorf("NoteRepository - GetNotes - Scan:%v", err)
		}
		notes = append(notes, note)
	}
	return notes, nil
}
