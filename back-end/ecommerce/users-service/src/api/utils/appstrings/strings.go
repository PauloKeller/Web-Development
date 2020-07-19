package appstrings

import "errors"

var ErrorNotFoundOnCache = errors.New("Element not found on Cache.")
var ErrorNotFoundOnDB = errors.New("Element not found on DB.")

var ErrorNotFound = errors.New("Element not found.")
var ErrorInsertUser = errors.New("Error creating user.")
