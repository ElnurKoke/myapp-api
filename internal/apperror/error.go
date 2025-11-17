package apperror

import "errors"

var (
	ErrWrongPassword     = errors.New(" Wrong password ")
	ErrUserExists        = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user does not exist or password incorrect")
	ErrInvalidUserName   = errors.New("invalid username - your username should consist at least 6 characters")
	ErrInvalidEmail      = errors.New("invalid email")
	ErrPasswordDontMatch = errors.New("password didn't match")
	ErrShortPassword     = errors.New("incorrect password - your password should be a minimum of 8 characters and consist of at least:1 lower case letter, 1 upper case letter, 1 number, 1 special symbol")
)
