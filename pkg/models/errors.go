package models

type UserExistsErr struct{}

func (e UserExistsErr) Error() string {
	return "user already exists"
}

type UserNotExistsErr struct{}

func (e UserNotExistsErr) Error() string {
	return "user not found"
}

type InvalidPasswordErr struct{}

func (e InvalidPasswordErr) Error() string {
	return "password is invalid"
}
