package easylocale

import "strings"

var langEnv = []string{"LC_ALL", "LANG"}

// CurrentLocale get locale from environment
func CurrentLocale() (string, error) {
	var locale string

	for _, env := range langEnv {
		locale = os.Getenv(env)
		if locale != "" {
			return strings.Split(locale, ".")[0], nil
		}
	}

	var err error
	if locale == "" {
		err = errors.New("can't detect locale info from environment")
	}

	return locale, err
}

// MustCurrentLocale get locale through system call, panic if failed
func MustCurrentLocale() string {
	locale, err := CurrentLocale()
	if err != nil {
		panic(err)
	}
	return locale
}
