package easylocale

import (
	"testing"
)

func TestCurrentLocale(t *testing.T) {
	t.Parallel()
	_, err := CurrentLocale()
	if err != nil {
		t.Fail()
	}
}
