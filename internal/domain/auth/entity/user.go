package entity

type User struct {
	Id     int
	Email  string
	Passwd string
}

func UserFromDTO(email string, passwd string) User {
	return User{Email: email, Passwd: passwd}
}
