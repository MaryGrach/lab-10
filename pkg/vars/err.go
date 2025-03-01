package vars

import "errors"

var (
	ErrorAlreadyExists    = errors.New("Already exist")
	ErrorDBNotInitialized = errors.New("DB is not initialized")
)
