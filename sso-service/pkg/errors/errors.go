package errors

import "fmt"

// Processing unknown email error
type EmailNotFoundError struct{
	Email string
}

func (e *EmailNotFoundError) Error() string {
	return fmt.Sprintf("Email '%s' not found", e.Email)
}

// Processing wrong password error
type WrongPasswordError struct{
	Email string
}

func (e *WrongPasswordError) Error() string {
	return fmt.Sprintf("Wrong password for user '%s'", e.Email)
}