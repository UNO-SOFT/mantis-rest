// Code generated by go-swagger; DO NOT EDIT.

package issues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IssueFileDeleteReader is a Reader for the IssueFileDelete structure.
type IssueFileDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueFileDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIssueFileDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIssueFileDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIssueFileDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewIssueFileDeleteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /issues/:id/files] issueFileDelete", response, response.Code())
	}
}

// NewIssueFileDeleteNoContent creates a IssueFileDeleteNoContent with default headers values
func NewIssueFileDeleteNoContent() *IssueFileDeleteNoContent {
	return &IssueFileDeleteNoContent{}
}

/*
IssueFileDeleteNoContent describes a response with status code 204, with default header values.

File deleted successfully
*/
type IssueFileDeleteNoContent struct {
}

// IsSuccess returns true when this issue file delete no content response has a 2xx status code
func (o *IssueFileDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue file delete no content response has a 3xx status code
func (o *IssueFileDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue file delete no content response has a 4xx status code
func (o *IssueFileDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue file delete no content response has a 5xx status code
func (o *IssueFileDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this issue file delete no content response a status code equal to that given
func (o *IssueFileDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the issue file delete no content response
func (o *IssueFileDeleteNoContent) Code() int {
	return 204
}

func (o *IssueFileDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /issues/:id/files][%d] issueFileDeleteNoContent ", 204)
}

func (o *IssueFileDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /issues/:id/files][%d] issueFileDeleteNoContent ", 204)
}

func (o *IssueFileDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueFileDeleteForbidden creates a IssueFileDeleteForbidden with default headers values
func NewIssueFileDeleteForbidden() *IssueFileDeleteForbidden {
	return &IssueFileDeleteForbidden{}
}

/*
IssueFileDeleteForbidden describes a response with status code 403, with default header values.

Access denied
*/
type IssueFileDeleteForbidden struct {
}

// IsSuccess returns true when this issue file delete forbidden response has a 2xx status code
func (o *IssueFileDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue file delete forbidden response has a 3xx status code
func (o *IssueFileDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue file delete forbidden response has a 4xx status code
func (o *IssueFileDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue file delete forbidden response has a 5xx status code
func (o *IssueFileDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this issue file delete forbidden response a status code equal to that given
func (o *IssueFileDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the issue file delete forbidden response
func (o *IssueFileDeleteForbidden) Code() int {
	return 403
}

func (o *IssueFileDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /issues/:id/files][%d] issueFileDeleteForbidden ", 403)
}

func (o *IssueFileDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /issues/:id/files][%d] issueFileDeleteForbidden ", 403)
}

func (o *IssueFileDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueFileDeleteNotFound creates a IssueFileDeleteNotFound with default headers values
func NewIssueFileDeleteNotFound() *IssueFileDeleteNotFound {
	return &IssueFileDeleteNotFound{}
}

/*
IssueFileDeleteNotFound describes a response with status code 404, with default header values.

Issue or file doesn't exist
*/
type IssueFileDeleteNotFound struct {
}

// IsSuccess returns true when this issue file delete not found response has a 2xx status code
func (o *IssueFileDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue file delete not found response has a 3xx status code
func (o *IssueFileDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue file delete not found response has a 4xx status code
func (o *IssueFileDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue file delete not found response has a 5xx status code
func (o *IssueFileDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue file delete not found response a status code equal to that given
func (o *IssueFileDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue file delete not found response
func (o *IssueFileDeleteNotFound) Code() int {
	return 404
}

func (o *IssueFileDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /issues/:id/files][%d] issueFileDeleteNotFound ", 404)
}

func (o *IssueFileDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /issues/:id/files][%d] issueFileDeleteNotFound ", 404)
}

func (o *IssueFileDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueFileDeleteServiceUnavailable creates a IssueFileDeleteServiceUnavailable with default headers values
func NewIssueFileDeleteServiceUnavailable() *IssueFileDeleteServiceUnavailable {
	return &IssueFileDeleteServiceUnavailable{}
}

/*
IssueFileDeleteServiceUnavailable describes a response with status code 503, with default header values.

Mantis Offline
*/
type IssueFileDeleteServiceUnavailable struct {
}

// IsSuccess returns true when this issue file delete service unavailable response has a 2xx status code
func (o *IssueFileDeleteServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue file delete service unavailable response has a 3xx status code
func (o *IssueFileDeleteServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue file delete service unavailable response has a 4xx status code
func (o *IssueFileDeleteServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue file delete service unavailable response has a 5xx status code
func (o *IssueFileDeleteServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this issue file delete service unavailable response a status code equal to that given
func (o *IssueFileDeleteServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the issue file delete service unavailable response
func (o *IssueFileDeleteServiceUnavailable) Code() int {
	return 503
}

func (o *IssueFileDeleteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /issues/:id/files][%d] issueFileDeleteServiceUnavailable ", 503)
}

func (o *IssueFileDeleteServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /issues/:id/files][%d] issueFileDeleteServiceUnavailable ", 503)
}

func (o *IssueFileDeleteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
