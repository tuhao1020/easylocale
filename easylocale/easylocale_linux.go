package easylocale

import (
	"os"
	"strings"
)

var langEnv = []string{"LC_ALL", "LANG"}

// CurrentLocale get locale from environment
func CurrentLocale() (string, error) {
	var locale string

	for _, env := range langEnv {
		locale = os.Getenv(env)
		if locale == "" {
			continue
		}
	}

	if locale == "" {
		return "", errors.New("can't detect locale info from environment")
	}

	return strings.ReplaceAll(strings.Split(locale, ".")[0], "_", "-"), nil
}

// MustCurrentLocale get locale through system call, panic if failed
func MustCurrentLocale() string {
	locale, err := CurrentLocale()
	if err != nil {
		panic(err)
	}
	return locale
}
