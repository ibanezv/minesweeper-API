package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_GetUser(t *testing.T) {
	//given
	tests := []struct {
		name         string
		url          string
		idExpected   int
		codeExpected int
	}{
		{
			name:         "get user success",
			url:          "/users/1",
			idExpected:   1,
			codeExpected: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()
			userHandler := NewHandlerUsers(userServiceMock{})
			vars := map[string]string{
				"id": "1",
			}
			req = mux.SetURLVars(req, vars)

			//when
			userHandler.GetUser(w, req)
			resp := w.Result()
			defer resp.Body.Close()

			//then
			assert.Equal(t, tt.codeExpected, resp.StatusCode)
		})
	}
}
