package groups

import "errors"

// ErrNilGinHandler signals that a nil gin handler has been provided
var ErrNilGinHandler = errors.New("nil gin handler")

// ErrEndpointAlreadyRegistered signals that the the provided endpoint path already exists
var ErrEndpointAlreadyRegistered = errors.New("endpoint already registered")

// ErrHandlerDoesNotExist signals that the requested handler does not exist
var ErrHandlerDoesNotExist = errors.New("handler does not exist")

// ErrWrongTypeAssertion signals that a wrong type assertion issue was found during the execution
var ErrWrongTypeAssertion = errors.New("wrong type assertion")
