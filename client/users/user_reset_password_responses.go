// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UserResetPasswordReader is a Reader for the UserResetPassword structure.
type UserResetPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserResetPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserResetPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUserResetPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUserResetPasswordServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserResetPasswordOK creates a UserResetPasswordOK with default headers values
func NewUserResetPasswordOK() *UserResetPasswordOK {
	return &UserResetPasswordOK{}
}

/*UserResetPasswordOK handles this case with default header values.

Success
*/
type UserResetPasswordOK struct {
}

func (o *UserResetPasswordOK) Error() string {
	return fmt.Sprintf("[PUT /users/reset][%d] userResetPasswordOK ", 200)
}

func (o *UserResetPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserResetPasswordForbidden creates a UserResetPasswordForbidden with default headers values
func NewUserResetPasswordForbidden() *UserResetPasswordForbidden {
	return &UserResetPasswordForbidden{}
}

/*UserResetPasswordForbidden handles this case with default header values.

Access denied
*/
type UserResetPasswordForbidden struct {
}

func (o *UserResetPasswordForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/reset][%d] userResetPasswordForbidden ", 403)
}

func (o *UserResetPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserResetPasswordServiceUnavailable creates a UserResetPasswordServiceUnavailable with default headers values
func NewUserResetPasswordServiceUnavailable() *UserResetPasswordServiceUnavailable {
	return &UserResetPasswordServiceUnavailable{}
}

/*UserResetPasswordServiceUnavailable handles this case with default header values.

Mantis Offline
*/
type UserResetPasswordServiceUnavailable struct {
}

func (o *UserResetPasswordServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /users/reset][%d] userResetPasswordServiceUnavailable ", 503)
}

func (o *UserResetPasswordServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}