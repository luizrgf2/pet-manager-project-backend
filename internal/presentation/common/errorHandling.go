package common

import (
	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	"github.com/luizrgf2/pet-manager-project-backend/internal/presentation/contracts"
)

type ErrorHandling[T any] struct {
}

func (e *ErrorHandling[T]) HandlingGenericError(err error) contracts.HTTPResponse[T] {
	if errBase, ok := err.(*core_errors.ErroBase); ok {
		return contracts.HTTPResponse[T]{
			Response:      nil,
			ErrorsMessage: []string{errBase.Message},
			Code:          errBase.Code,
		}
	}

	return contracts.HTTPResponse[T]{
		Response:      nil,
		ErrorsMessage: []string{err.Error()},
		Code:          500,
	}
}

func (e *ErrorHandling[T]) HandlingFieldicError(errMessages []string) contracts.HTTPResponse[T] {
	return contracts.HTTPResponse[T]{
		Response:      nil,
		ErrorsMessage: errMessages,
		Code:          400,
	}
}
