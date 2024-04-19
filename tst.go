package andres_ruiz_golang

import "testing"

// True calls t.Error() If expr is false
func True(t *testing.T, expr bool, fmtMsg string, vals ...interface{}) {
	t.Helper()
	if !expr {
		t.Errorf(fmtMsg, vals...)
	}
}
