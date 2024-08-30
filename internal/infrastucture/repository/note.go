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
	stmt, err := n.db.Prepare("INSERT INTO notes(id, note) VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("repository - NoteRepository - CreateNote - Prepare:%v", err)
	}
	_, err = stmt.Exec(note.Id, note.Note)
	if err != nil {
		return fmt.Errorf("repository - NoteRepository - CreateNote - Exec:%v", err)
	}
	return nil
}

func (n *NoteRepository) GetNotes() ([]entity.Note, error) {
	notes := []entity.Note{}

	rows, err := n.db.Query("SELECT * FROM notes")
	if err != nil {
		return nil, fmt.Errorf("repository - NoteRepository - GetNotes - Query:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		note := entity.Note{}
		err := rows.Scan(&note.Id, &note.Note)
		if err != nil {
			return nil, fmt.Errorf("repository - NoteRepository - GetNotes - Scan:%v", err)
		}
		notes = append(notes, note)
	}
	return notes, nil
}
