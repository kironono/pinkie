package handler

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealth_Show(t *testing.T) {
	type want struct {
		status int
		body   []byte
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Status OK",
			want: want{
				status: http.StatusOK,
				body:   []byte(`{"status":"ok"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodGet,
				"/health",
				bytes.NewBuffer([]byte("")),
			)

			h := &health{}
			h.Show(w, r)
			actual := w.Result()

			assert.Equal(t, tt.want.status, actual.StatusCode)
			body, err := io.ReadAll(actual.Body)
			require.NoError(t, err)
			assert.Equal(t, tt.want.body, body)
		})
	}
}
