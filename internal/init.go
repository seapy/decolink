package internal

import (
	"fmt"

	"github.com/spf13/viper"
)

// Init ...
func Init(args []string) {
	parser := viper.Get("parser").(string)
	if parser == "chrome" {
		parserChrome := ParserChrome{}
		fmt.Println(parserChrome.GetLink(args[0]))
	} else {
		parserGoquery := ParserGoquery{}
		fmt.Println(parserGoquery.GetLink(args[0]))
	}
}
