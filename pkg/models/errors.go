package models

type UserExistsErr struct{}

func (e UserExistsErr) Error() string {
	return "user already exists"
}
