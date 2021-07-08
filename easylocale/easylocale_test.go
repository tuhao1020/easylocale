package easylocale

import (
	"fmt"
	"testing"
)

func TestCurrentLocale(t *testing.T) {
	t.Parallel()
	locale, err := CurrentLocale()
	if err != nil {
		t.Fail()
	}
	fmt.Println(locale)
}
