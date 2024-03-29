// Code generated by go-swagger; DO NOT EDIT.

package issues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IssueNoteDeleteReader is a Reader for the IssueNoteDelete structure.
type IssueNoteDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueNoteDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIssueNoteDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIssueNoteDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIssueNoteDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewIssueNoteDeleteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /issues/:id/notes] issueNoteDelete", response, response.Code())
	}
}

// NewIssueNoteDeleteNoContent creates a IssueNoteDeleteNoContent with default headers values
func NewIssueNoteDeleteNoContent() *IssueNoteDeleteNoContent {
	return &IssueNoteDeleteNoContent{}
}

/*
IssueNoteDeleteNoContent describes a response with status code 204, with default header values.

Note deleted successfully
*/
type IssueNoteDeleteNoContent struct {
}

// IsSuccess returns true when this issue note delete no content response has a 2xx status code
func (o *IssueNoteDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue note delete no content response has a 3xx status code
func (o *IssueNoteDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue note delete no content response has a 4xx status code
func (o *IssueNoteDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue note delete no content response has a 5xx status code
func (o *IssueNoteDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this issue note delete no content response a status code equal to that given
func (o *IssueNoteDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the issue note delete no content response
func (o *IssueNoteDeleteNoContent) Code() int {
	return 204
}

func (o *IssueNoteDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /issues/:id/notes][%d] issueNoteDeleteNoContent ", 204)
}

func (o *IssueNoteDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /issues/:id/notes][%d] issueNoteDeleteNoContent ", 204)
}

func (o *IssueNoteDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueNoteDeleteForbidden creates a IssueNoteDeleteForbidden with default headers values
func NewIssueNoteDeleteForbidden() *IssueNoteDeleteForbidden {
	return &IssueNoteDeleteForbidden{}
}

/*
IssueNoteDeleteForbidden describes a response with status code 403, with default header values.

Access denied
*/
type IssueNoteDeleteForbidden struct {
}

// IsSuccess returns true when this issue note delete forbidden response has a 2xx status code
func (o *IssueNoteDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue note delete forbidden response has a 3xx status code
func (o *IssueNoteDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue note delete forbidden response has a 4xx status code
func (o *IssueNoteDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue note delete forbidden response has a 5xx status code
func (o *IssueNoteDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this issue note delete forbidden response a status code equal to that given
func (o *IssueNoteDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the issue note delete forbidden response
func (o *IssueNoteDeleteForbidden) Code() int {
	return 403
}

func (o *IssueNoteDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /issues/:id/notes][%d] issueNoteDeleteForbidden ", 403)
}

func (o *IssueNoteDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /issues/:id/notes][%d] issueNoteDeleteForbidden ", 403)
}

func (o *IssueNoteDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueNoteDeleteNotFound creates a IssueNoteDeleteNotFound with default headers values
func NewIssueNoteDeleteNotFound() *IssueNoteDeleteNotFound {
	return &IssueNoteDeleteNotFound{}
}

/*
IssueNoteDeleteNotFound describes a response with status code 404, with default header values.

Issue or note doesn't exist
*/
type IssueNoteDeleteNotFound struct {
}

// IsSuccess returns true when this issue note delete not found response has a 2xx status code
func (o *IssueNoteDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue note delete not found response has a 3xx status code
func (o *IssueNoteDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue note delete not found response has a 4xx status code
func (o *IssueNoteDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue note delete not found response has a 5xx status code
func (o *IssueNoteDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue note delete not found response a status code equal to that given
func (o *IssueNoteDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue note delete not found response
func (o *IssueNoteDeleteNotFound) Code() int {
	return 404
}

func (o *IssueNoteDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /issues/:id/notes][%d] issueNoteDeleteNotFound ", 404)
}

func (o *IssueNoteDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /issues/:id/notes][%d] issueNoteDeleteNotFound ", 404)
}

func (o *IssueNoteDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueNoteDeleteServiceUnavailable creates a IssueNoteDeleteServiceUnavailable with default headers values
func NewIssueNoteDeleteServiceUnavailable() *IssueNoteDeleteServiceUnavailable {
	return &IssueNoteDeleteServiceUnavailable{}
}

/*
IssueNoteDeleteServiceUnavailable describes a response with status code 503, with default header values.

Mantis Offline
*/
type IssueNoteDeleteServiceUnavailable struct {
}

// IsSuccess returns true when this issue note delete service unavailable response has a 2xx status code
func (o *IssueNoteDeleteServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue note delete service unavailable response has a 3xx status code
func (o *IssueNoteDeleteServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue note delete service unavailable response has a 4xx status code
func (o *IssueNoteDeleteServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue note delete service unavailable response has a 5xx status code
func (o *IssueNoteDeleteServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this issue note delete service unavailable response a status code equal to that given
func (o *IssueNoteDeleteServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the issue note delete service unavailable response
func (o *IssueNoteDeleteServiceUnavailable) Code() int {
	return 503
}

func (o *IssueNoteDeleteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /issues/:id/notes][%d] issueNoteDeleteServiceUnavailable ", 503)
}

func (o *IssueNoteDeleteServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /issues/:id/notes][%d] issueNoteDeleteServiceUnavailable ", 503)
}

func (o *IssueNoteDeleteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
