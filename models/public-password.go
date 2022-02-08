package models

type PublicPassword struct {
	ID          uint
	Password    string
	ClientName  string
	ContactName string
}

type PublicPasswords []*PublicPassword
