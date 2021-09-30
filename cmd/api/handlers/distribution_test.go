package handlers

import (
	"encoding/json"
	"github/ibanezv/minesweeper-API/internal/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_PatchDistribution(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		distribution models.Distribution
		codeExpected int
	}{
		{
			name:         "patch distribution success",
			url:          "/distribution/1",
			distribution: models.Distribution{GameID: 1, RowNumber: 2, ColNumber: 2, Value: "select"},
			codeExpected: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, _ := json.Marshal(tt.distribution)
			req := httptest.NewRequest(http.MethodPatch, tt.url, strings.NewReader(string(e)))
			vars := map[string]string{
				"id": "1",
			}
			req = mux.SetURLVars(req, vars)

			w := httptest.NewRecorder()
			h := NewHandlerDistributions(distributionServiceMock{})
			h.PatchDistribution(w, req)
			resp := w.Result()
			defer resp.Body.Close()
			assert.Equal(t, tt.codeExpected, resp.StatusCode)
		})
	}
}
