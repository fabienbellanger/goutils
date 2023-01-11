package errors

// https://github.com/g8rswimmer/error-chain

import (
	"errors"
	"strings"
)

// Link represents an error
type Link struct {
	// Code uint   `json:"code"`  // Error code
	// Msg  string `json:"msg"`   // Error message
	// File string `json:"file"`  // Context file name
	// Line uint   `json:"line"`  // Context line number
	Err  error `json:"error"` // Error
	Next *Link `json:"next"`  // Next errors
}

// ErrorChain chains multiple errors
type ErrorChain struct {
	List *Link `json:"list"`
	Last *Link `json:"last"`
}

// New create a new error chain
func New() *ErrorChain {
	return &ErrorChain{}
}

// TODO: Improve output
func (e *ErrorChain) Error() string {
	errs := []string{}
	h := e.List
	for h != nil {
		errs = append(errs, h.Err.Error())
		h = h.Next
	}
	return strings.Join(errs, ": ")
}

// Unwrap gives the next error
func (e *ErrorChain) Unwrap() error {
	if e.List.Next == nil {
		return nil
	}
	ec := &ErrorChain{
		List: e.List.Next,
		Last: e.Last,
	}
	return ec
}

// Is compares to the target
func (e *ErrorChain) Is(target error) bool {
	return errors.Is(e.List.Err, target)
}

// Add will place another error in the chain
func (e *ErrorChain) Add(err error, code uint, msg string) {
	l := &Link{
		// Code: code,
		// Msg:  msg,
		Err: err,
	}
	if e.List == nil {
		e.List = l
		e.Last = l
		return
	}

	e.Last.Next = l
	e.Last = l
}

// Errors returns the errors in the chain
func (e *ErrorChain) Errors() []error {
	errs := []error{}
	l := e.List
	for l != nil {
		errs = append(errs, l.Err)
		l = l.Next
	}
	return errs
}
