// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/UNO-SOFT/mantis-rest/models"
)

// UserGetMeReader is a Reader for the UserGetMe structure.
type UserGetMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGetMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUserGetMeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUserGetMeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/me] userGetMe", response, response.Code())
	}
}

// NewUserGetMeOK creates a UserGetMeOK with default headers values
func NewUserGetMeOK() *UserGetMeOK {
	return &UserGetMeOK{}
}

/*
UserGetMeOK describes a response with status code 200, with default header values.

Success
*/
type UserGetMeOK struct {
	Payload *models.UserMeResponse
}

// IsSuccess returns true when this user get me o k response has a 2xx status code
func (o *UserGetMeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user get me o k response has a 3xx status code
func (o *UserGetMeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user get me o k response has a 4xx status code
func (o *UserGetMeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user get me o k response has a 5xx status code
func (o *UserGetMeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user get me o k response a status code equal to that given
func (o *UserGetMeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user get me o k response
func (o *UserGetMeOK) Code() int {
	return 200
}

func (o *UserGetMeOK) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] userGetMeOK  %+v", 200, o.Payload)
}

func (o *UserGetMeOK) String() string {
	return fmt.Sprintf("[GET /users/me][%d] userGetMeOK  %+v", 200, o.Payload)
}

func (o *UserGetMeOK) GetPayload() *models.UserMeResponse {
	return o.Payload
}

func (o *UserGetMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserMeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGetMeForbidden creates a UserGetMeForbidden with default headers values
func NewUserGetMeForbidden() *UserGetMeForbidden {
	return &UserGetMeForbidden{}
}

/*
UserGetMeForbidden describes a response with status code 403, with default header values.

Access denied
*/
type UserGetMeForbidden struct {
}

// IsSuccess returns true when this user get me forbidden response has a 2xx status code
func (o *UserGetMeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user get me forbidden response has a 3xx status code
func (o *UserGetMeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user get me forbidden response has a 4xx status code
func (o *UserGetMeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user get me forbidden response has a 5xx status code
func (o *UserGetMeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user get me forbidden response a status code equal to that given
func (o *UserGetMeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user get me forbidden response
func (o *UserGetMeForbidden) Code() int {
	return 403
}

func (o *UserGetMeForbidden) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] userGetMeForbidden ", 403)
}

func (o *UserGetMeForbidden) String() string {
	return fmt.Sprintf("[GET /users/me][%d] userGetMeForbidden ", 403)
}

func (o *UserGetMeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGetMeServiceUnavailable creates a UserGetMeServiceUnavailable with default headers values
func NewUserGetMeServiceUnavailable() *UserGetMeServiceUnavailable {
	return &UserGetMeServiceUnavailable{}
}

/*
UserGetMeServiceUnavailable describes a response with status code 503, with default header values.

Mantis Offline
*/
type UserGetMeServiceUnavailable struct {
}

// IsSuccess returns true when this user get me service unavailable response has a 2xx status code
func (o *UserGetMeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user get me service unavailable response has a 3xx status code
func (o *UserGetMeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user get me service unavailable response has a 4xx status code
func (o *UserGetMeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this user get me service unavailable response has a 5xx status code
func (o *UserGetMeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this user get me service unavailable response a status code equal to that given
func (o *UserGetMeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the user get me service unavailable response
func (o *UserGetMeServiceUnavailable) Code() int {
	return 503
}

func (o *UserGetMeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] userGetMeServiceUnavailable ", 503)
}

func (o *UserGetMeServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /users/me][%d] userGetMeServiceUnavailable ", 503)
}

func (o *UserGetMeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
