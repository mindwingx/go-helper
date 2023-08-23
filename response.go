package helper

import (
	"github.com/mindwingx/abstraction"
	"net/http"
	"time"
)

type (
	Valid string
	Err   string

	response struct {
		Success       bool                   `json:"success"`
		StatusCode    int                    `json:"status_code"`
		Status        string                 `json:"status"`
		Message       string                 `json:"message"`
		Data          map[string]interface{} `json:"data"`
		RepresentedAt string                 `json:"represented_at"`
	}
)

const (
	OK                    Valid = "Ok"
	Created               Valid = "Created"
	NoContent             Valid = "No Content"
	BadRequest            Err   = "Bad Request"
	UnAuthorized          Err   = "UnAuthorized"
	StatusForbidden       Err   = "status forbidden"
	NotFound              Err   = "Not Found"
	UnpronounceableEntity Err   = "Unpronounceable Entity"
	InternalServerError   Err   = "Internal Server Error"
)

var (
	Validates = map[Valid]int{
		OK:        http.StatusOK,
		Created:   http.StatusCreated,
		NoContent: http.StatusNoContent,
	}

	Errors = map[Err]int{
		BadRequest:            http.StatusBadRequest,
		UnAuthorized:          http.StatusUnauthorized,
		StatusForbidden:       http.StatusForbidden,
		NotFound:              http.StatusNotFound,
		UnpronounceableEntity: http.StatusUnprocessableEntity,
		InternalServerError:   http.StatusInternalServerError,
	}
)

// validate http status
func (e Valid) String() string {
	return string(e)
}

// error http status
func (e Err) String() string {
	return string(e)
}

func newResponse() *response {
	return &response{}
}

func SuccessResponse(
	c abstraction.AbstractCtx,
	status Valid,
	message string,
	data map[string]interface{},
) {
	res := newResponse()
	res.Success = true
	res.StatusCode = Validates[status]
	res.Status = status.String()
	res.Message = message
	res.Data = data
	res.RepresentedAt = time.Now().Format(time.RFC3339)

	if len(message) > 0 {
		res.Message = "Process done successfully"
	}

	c.JSON(res.StatusCode, res)
	return
}

func ErrorResponse(
	c abstraction.AbstractCtx,
	status Err,
	errMessage error,
	data map[string]interface{},
) {
	res := newResponse()
	res.Success = false
	res.StatusCode = Errors[status]
	res.Status = status.String()
	res.Message = errMessage.Error()
	res.Data = data
	res.RepresentedAt = time.Now().Format(time.RFC3339)

	if errMessage == nil {
		res.Message = "Process failed"
	}

	c.JSON(res.StatusCode, res)
	c.Abort()
	return
}
