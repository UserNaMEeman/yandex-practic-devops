package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestHandleMetric(t *testing.T) {
	type want struct {
		statusCode  int
		response    string
		contentType string
	}
	tests := []struct {
		name    string
		request string
		want    want
	}{
		{
			name: "Normal request",
			// request: "/update/guage/Alloc/3443",
			request: "/update/counter/testCounter/100",
			want: want{
				statusCode:  200,
				response:    "OK",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.request, nil)
			request.Header.Set("Content-Type", "text/plain")
			w := httptest.NewRecorder()
			h := http.HandlerFunc(HandleMetric)
			h.ServeHTTP(w, request)
			result := w.Result()
			assert.Equal(t, result.StatusCode, tt.want.statusCode)
			assert.Equal(t, result.Header.Get("Content-Type"), tt.want.contentType)
		})
	}
}
