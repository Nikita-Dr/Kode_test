package entity

type Note struct {
	Id     int
	Note   string
	UserId int
}

func NoteFromDTO(note string, userId int) Note {
	return Note{Note: note, UserId: userId}
}

func (n *Note) UpdateNote(note string) {
	n.Note = note
}
