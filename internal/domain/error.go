package domain

import (
	"errors"
)

var ErrorDuplicatedPrimaryKey = errors.New("duplicated primary key")
var ErrorEntryNotFound = errors.New("entry not found")
var ErrorInternalServerError = errors.New("internal server error")

var Error404 = errors.New("not found")
var Error400 = errors.New("bad request")
var Error500 = errors.New("internal server error")
