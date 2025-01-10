package conf

import (
	"github.com/ekhidirov/go-parser-php/pkg/errors"
	"github.com/ekhidirov/go-parser-php/pkg/version"
)

type Config struct {
	Version          *version.Version
	ErrorHandlerFunc func(e *errors.Error)
}
