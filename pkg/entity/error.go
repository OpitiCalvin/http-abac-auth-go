package entity

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("not found")

// ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

// ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot be deleted")

// ErrClientAlreadySubscribedToProduct client is already subscribed to the product
var ErrClientAlreadySubscribedToProduct = errors.New("client already subscribed to the product")

// ErrClientAlreadyLinkedToPartner client already linked to the partner record
var ErrClientAlreadyLinkedToPartner = errors.New("client already linked to partner record")

// ErrClientNotLinkedToPartner client record is not linked to a/the partner record
var ErrClientNotLinkedToPartner = errors.New("client is not linked to a/the partner")
