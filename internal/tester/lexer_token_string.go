package tester

import (
	"testing"

	"github.com/ekhidirov/go-parser-php/internal/php8"
	"github.com/ekhidirov/go-parser-php/internal/scanner"
	"github.com/ekhidirov/go-parser-php/pkg/conf"
	"github.com/ekhidirov/go-parser-php/pkg/version"
	"gotest.tools/assert"
)

type LexerTokenStringTestSuite struct {
	t *testing.T

	Code     string
	Expected []string

	Version version.Version
}

func NewLexerTokenStringTestSuite(t *testing.T) *LexerTokenStringTestSuite {
	return &LexerTokenStringTestSuite{
		t: t,
		Version: version.Version{
			Major: 7,
			Minor: 4,
		},
	}
}

func (l *LexerTokenStringTestSuite) UsePHP8() {
	l.Version = version.Version{Major: 8, Minor: 0}
}

func (l *LexerTokenStringTestSuite) Run() {
	config := conf.Config{
		Version: &l.Version,
	}

	var lexer Lexer

	if l.Version.Less(&version.Version{Major: 8, Minor: 0}) {
		lexer = scanner.NewLexer([]byte(l.Code), config)
	} else {
		lexer = php8.NewLexer([]byte(l.Code), config)
	}

	for _, expected := range l.Expected {
		tkn := lexer.Lex()
		actual := string(tkn.Value)
		assert.DeepEqual(l.t, expected, actual)
	}
}
