package domain

type UserInsert struct {
	UUID     string
	Username string
	Name     string
	Password string
}

type User struct {
	UUID     string
	Username string
	Name     string
	Password string
}
