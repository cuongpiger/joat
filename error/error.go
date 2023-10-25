package error

import "fmt"

// ***************************************************** BaseError *****************************************************

type baseError struct {
	DefaultError string
	Info         string
}

func (s *baseError) Error() string {
	s.DefaultError = defaultErrorString
	return s.choseErrString()
}

func (s *baseError) SetInfo(pInfo string) IErrorBuilder {
	s.Info = pInfo
	return s
}

func (s *baseError) choseErrString() string {
	if s.Info != "" {
		return s.Info
	}
	return s.DefaultError
}

// ************************************************** ErrMissingInput **************************************************

type ErrMissingInput struct {
	baseError
	Argument string
}

func (s *ErrMissingInput) Error() string {
	s.DefaultError = fmt.Sprintf("missing input for argument [%s]", s.Argument)
	return s.choseErrString()
}
