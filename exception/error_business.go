package exception

type ErrorBusiness struct {
	Code         int
	ErrorMessage string
}

func (e ErrorBusiness) Error() string {
	return e.ErrorMessage
}

func NewErrorBusiness(code int, err error) ErrorBusiness {
	return ErrorBusiness{
		Code:         code,
		ErrorMessage: err.Error(),
	}
}

func PanicErrorBusiness(code int, err error) {
	if err != nil {
		panic(NewErrorBusiness(code, err))
	}
}
