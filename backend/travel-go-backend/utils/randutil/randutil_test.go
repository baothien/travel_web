package randutil

import (
	"testing"
)

func TesGenReferralCode(t *testing.T) {
	str := RandomString(6)
	if len(str) != 6 {
		t.Errorf("fail %s", str)
	}
}
