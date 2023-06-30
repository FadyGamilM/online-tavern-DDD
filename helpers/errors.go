package helpers

type ErrInvalidPersonName struct {
	ErrMsg string
}

func (e ErrInvalidPersonName) Error() string {
	return e.ErrMsg
}

type ErrInvalidPersonAge struct {
	ErrMsg string
}

func (e ErrInvalidPersonAge) Error() string {
	return e.ErrMsg
}

type ErrMissingEssentialData struct {
	ErrMsg string
}

func (e ErrMissingEssentialData) Error() string {
	return e.ErrMsg
}
