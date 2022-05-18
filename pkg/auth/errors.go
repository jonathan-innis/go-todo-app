package auth

type InvalidTokenErr struct{}

func (e InvalidTokenErr) Error() string {
	return "token is invalid"
}

type UnauthorizedTokenErr struct{}

func (e UnauthorizedTokenErr) Error() string {
	return "token is anauthorized"
}
