package error

func NewErrMissingInput(pArgument, pInfo string) *ErrMissingInput {
	err := new(ErrMissingInput)
	err.Argument = pArgument
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}
