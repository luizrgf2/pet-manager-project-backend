package core_errors

import "fmt"

type ErroBase struct {
	Message string
	Code    uint
}

func (e *ErroBase) Error() string {
	return fmt.Sprintf("Error code %d : %s", e.Code, e.Message)
}
