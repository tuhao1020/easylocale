package easylocale

import (
	"errors"
	"os"
)

var langEnv = []string{"LC_ALL", "LANG"}

func CurrentLocale() (string, error) {
	var locale string

	for _, env := range langEnv {
		locale = os.Getenv(env)
		if locale != "" {
			return localle, nil
		}
	}

	var err error
	if locale == "" {
		err = errors.New("can't detect locale info from environment")
	}

	return locale, err
}