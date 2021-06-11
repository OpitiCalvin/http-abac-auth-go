package entity

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("not found")

// ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

// ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot be deleted")

// ErrClientAlreadySubscribedToProduct
var ErrClientAlreadySubscribedToProduct = errors.New("client already subscribed to the product")
