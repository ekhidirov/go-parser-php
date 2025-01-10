package php5_test

import (
	"io/ioutil"
	"testing"

	"github.com/ekhidirov/go-parser-php/internal/php5"
	"github.com/ekhidirov/go-parser-php/internal/scanner"
	"github.com/ekhidirov/go-parser-php/pkg/conf"
	"github.com/ekhidirov/go-parser-php/pkg/version"
)

func BenchmarkPhp5(b *testing.B) {
	src, err := ioutil.ReadFile("test.php")
	if err != nil {
		b.Fatal("can not read test.php: " + err.Error())
	}

	for n := 0; n < b.N; n++ {
		config := conf.Config{
			Version: &version.Version{
				Major: 5,
				Minor: 6,
			},
		}
		lexer := scanner.NewLexer(src, config)
		php5parser := php5.NewParser(lexer, config)
		php5parser.Parse()
	}
}
