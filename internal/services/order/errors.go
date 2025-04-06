package order

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

const (
	ErrNotEnoughMoneyCode = "NotEnoughMoney"
)

var (
	ErrNotEnoughMoney = &Error{
		Code:    ErrNotEnoughMoneyCode,
		Message: "not enough money",
	}
)
