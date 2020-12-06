package errors

import "fmt"

const (
	// InvalidArgumentError is should be raised when a passed
	// argument is invalid for the intended use.
	InvalidArgumentError string = "Invalid argument %s"
	// InvalidOperationError is should be raised when a passed
	// when the runtime is in an invalid state to perform the operation..
	InvalidOperationError string = "Invalid operation %s"
	// MissingArgumentError should be raised when a required
	// flag or method parameter is missing. It includes an optional
	// format string so that users can specify the argument that is missing.
	MissingArgumentError string = "Required argument not found %s"
	// DirectoryNotFoundError should be raised when a directory cannot be found.
	// It includes an optional format string so users can specify the missing
	// directory path.
	DirectoryNotFoundError string = "Directory not found %s"
	// FileNotFoundError should be raised when a file cannot be found.
	// It includes an optional format string so users can specify the missing
	// file path.
	FileNotFoundError string = "File not found %s"
	// MemberNotFoundError should be raised when an attempt to access a
	// struct or hcl config member evaluates to nil. It includes an optional
	// format string so that users can specify the missing member.
	MemberNotFoundError string = "Member not found %s"
	// NotImplementedError should be the initial value of all new methods
	// when using a TDD approach to interface design. It includes an optional
	// format string so that users can specify the method name.
	NotImplementedError string = "Method not implemented %s"
	// NonFlagArgumentsError should be raised when non-flags arguments have
	// been passed but are not allowed.
	NonFlagArgumentsError string = "Should have no non-flag arguments."
	// Error is the catch all Error for unanticipated Errors
	// or Errors from external codes. It assumes you will pass the method name
	// that raised the Error as well as the error instance.
	Error string = "%s: %v"
)

// NewWithMessage creates an error from the method that threw the error, and a user
// defined error message. The probable use case is a validation failure or anticipated
// error.
func NewWithMessage(methodName, message string) error {
	return fmt.Errorf("%s: %v", methodName, message)
}

// NewWithError creates an error from the method that threw the error, and the error
// instance that was thrown. The probable use case is an unanticipated error.
func NewWithError(methodName string, err error) error {
	return fmt.Errorf("%s: %v", methodName, err)
}
