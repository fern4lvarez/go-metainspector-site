package metainspector

import (
	"testing"
)

func TestVersion(t *testing.T) {
	v := Version()
	if v != "0.2.4" {
		t.Errorf(msgFail, "Version", "0.2.4", v)
	}
}
