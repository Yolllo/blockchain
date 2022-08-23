package observer

import "errors"

// ErrEmptyObserversList signals that the list of observers is empty
var ErrEmptyObserversList = errors.New("empty observers list")

// ErrShardNotAvailable signals that the specified shard ID cannot be found in internal maps
var ErrShardNotAvailable = errors.New("the specified shard ID does not exist in proxy's configuration")

// ErrWrongObserversConfiguration signals an invalid observers configuration
var ErrWrongObserversConfiguration = errors.New("wrong observers configuration")
