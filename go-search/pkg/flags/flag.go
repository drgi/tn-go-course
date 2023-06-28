package flags

import (
	"errors"
	"flag"
)

var (
	errParamRequired = errors.New("param for search required")
)

// Получение параметра флага
func RequireFlag(flagKey string) (string, error) {
	var param string
	flag.StringVar(&param, flagKey, "", "input word for search")
	flag.Parse()
	if param == "" {
		return param, errParamRequired
	}
	return param, nil
}
