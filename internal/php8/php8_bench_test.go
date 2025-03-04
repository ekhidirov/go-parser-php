package php8_test

import (
	"io/ioutil"
	"testing"

	"github.com/ekhidirov/go-parser-php/internal/php8"
	"github.com/ekhidirov/go-parser-php/pkg/conf"
	"github.com/ekhidirov/go-parser-php/pkg/version"
)

func BenchmarkPhp8(b *testing.B) {
	src, err := ioutil.ReadFile("test.php")

	if err != nil {
		b.Fatal("can not read test.php: " + err.Error())
	}

	for n := 0; n < b.N; n++ {
		config := conf.Config{
			Version: &version.Version{
				Major: 8,
				Minor: 8,
			},
		}
		lexer := php8.NewLexer(src, config)
		php8parser := php8.NewParser(lexer, config)
		php8parser.Parse()
	}
}
