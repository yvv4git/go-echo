package transport

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddressHandlerV1(t *testing.T) {
	testCases := []struct {
		name     string
		wantCode int
		wantBody string
	}{
		{
			name:     "CASE-1",
			wantCode: 200,
			wantBody: `{"host_address":"192.168.1.2"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, _ := http.NewRequest(http.MethodGet, "/v1/address", nil)
			w := httptest.NewRecorder()

			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			testDeps.srvWeb.router.ServeHTTP(w, r)

			assert.Equal(t, tc.wantCode, w.Code)
			assert.Equal(t, tc.wantBody, w.Body.String())
		})
	}
}
