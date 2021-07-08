package easylocale

import (
	"os/exec"
	"strings"
)

// CurrentLocale get locale through shell script "osascript -e 'user locale of (get system info)'"
func CurrentLocale() (string, error) {
	cmd := exec.Command("sh", "-c", "osascript -e 'user locale of (get system info)'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(strings.TrimSpace(string(output)), "_", "-"), nil
}

// MustCurrentLocale get locale through system call, panic if failed
func MustCurrentLocale() string {
	locale, err := CurrentLocale()
	if err != nil {
		panic(err)
	}
	return locale
}
