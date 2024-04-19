package andres_ruiz_golang

import "testing"

// Ok calls t.Fatal(err) If err is not nil
func Ok(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)

	}
}

// True calls t.Error() If expr is false
func True(t *testing.T, expr bool, fmtMsg string, vals ...interface{}) {
	t.Helper()
	if !expr {
		t.Errorf(fmtMsg, vals...)
	}
}
