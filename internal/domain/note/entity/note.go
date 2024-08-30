package entity

type Note struct {
	Id   int
	Note string
}

func NoteFromDTO(note string) Note {
	return Note{Note: note}
}
