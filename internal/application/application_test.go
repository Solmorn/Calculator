package application_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Solmorn/Calculator/internal/application"
)

func TestEvaluateExpression(t *testing.T) {
	tests := []struct {
		expr      string
		expectErr bool
		expected  string
	}{
		{"5 / 0", true, ""},
		{"2 + 2", false, "4.000000"},
		{"(3 + 2) * (4 - 1", true, ""},
		{"10 / (2 - 2)", true, ""},
		{"(5 + 3))", true, ""},
	}

	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			body, _ := json.Marshal(application.Request{Expression: test.expr})
			req := httptest.NewRequest(http.MethodPost, "/evaluate", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			application.CalcHandler(w, req)

			res := w.Result()
			if test.expectErr {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
				}
			} else {
				if res.StatusCode != http.StatusOK {
					t.Errorf("expected status %d, got %d", http.StatusOK, res.StatusCode)
				}

				var result Result
				if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if result.Result != test.expected {
					t.Errorf("expected result %s, got %s", test.expected, result.Result)
				}
			}
		})
	}
}
