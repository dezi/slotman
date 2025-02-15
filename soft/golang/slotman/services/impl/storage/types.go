package storage

import (
	"errors"
)

type Mode string

const (
	plain Mode = "plain"
)

const (
	datePartFormat = "20060102"
	timePartFormat = "150405-0700"
	dateTimeFormat = "20060102-150405-0700"
)

var errNoFilesFound = errors.New("no files found")
