package transport

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoFoundHandler(t *testing.T) {
	type requestOptions struct {
		method string
		path   string
	}

	testCases := []struct {
		name     string
		opt      requestOptions
		wantCode int
		wantBody string
	}{
		{
			name: "CASE-1",
			opt: requestOptions{
				method: http.MethodGet,
				path:   "/qwerty",
			},
			wantCode: http.StatusNotFound,
			wantBody: ``,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, _ := http.NewRequest(tc.opt.method, tc.opt.path, nil)
			w := httptest.NewRecorder()

			r.Header.Add("Content-Type", "application/json")
			testDeps.srvWeb.router.ServeHTTP(w, r)

			assert.Equal(t, tc.wantCode, w.Code)
			t.Log(w.Body.String())
		})
	}
}
