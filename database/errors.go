package database

import "errors"
// Errors
var (
	// ErrNoID is returned when no ID field or id tag is found in the struct.
	ErrNoID = errors.New("missing struct tag id or ID field")

	// ErrBadType is returned when a method receives an unexpected value type.
	ErrBadType = errors.New("provided data must be a struct or a pointer to struct")

	// ErrAlreadyExists is returned uses when trying to set an existing value on a field that has a unique index.
	ErrAlreadyExists = errors.New("already exists")

	// ErrUnknownTag is returned when an unexpected tag is specified.
	ErrUnknownTag = errors.New("unknown tag")

	// ErrIncompleteStructure is return when Some fields of an object are not assigned
	ErrIncompleteStructure = errors.New("Incomplete structure")

	// ErrSlicePtrNeeded is returned when an unexpected value is given, instead of a pointer to struct.
	ErrStructPtrNeeded = errors.New("provided target must be a pointer to struct")

	// ErrSlicePtrNeeded is returned when an unexpected value is given, instead of a pointer.
	ErrPtrNeeded = errors.New("provided target must be a pointer to a valid variable")

	// ErrNotFound is returned when the specified record is not saved in the bucket.
	ErrNotFound = errors.New("not found")
)
