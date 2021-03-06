package handler

import (
	"testing"
)

func TestHandleMetric(t *testing.T) {
	type want struct {
		statusCode  int
		response    string
		contentType string
	}
	// tests := []struct {
	// 	name    string
	// 	request string
	// 	want    want
	// }{
	// 	{
	// 		name:    "200",
	// 		request: "/update/gauge/testGauge/100",
	// 		want: want{
	// 			statusCode:  200,
	// 			response:    "OK",
	// 			contentType: "text/plain; charset=utf-8",
	// 		},
	// 	},
	// 	{
	// 		name:    "json",
	// 		request: "/value/",
	// 		want: want{
	// 			statusCode:  200,
	// 			response:    "OK",
	// 			contentType: "text/plain; charset=utf-8",
	// 		},
	// },
	// {
	// 	name:    "400",
	// 	request: "/update/test/testCounter/100",
	// 	want: want{
	// 		statusCode:  400,
	// 		response:    "OK",
	// 		contentType: "text/plain; charset=utf-8",
	// 	},
	// },
	// {
	// 	name:    "500",
	// 	request: "/update/counter/testCounter/none",
	// 	want: want{
	// 		statusCode:  500,
	// 		response:    "OK",
	// 		contentType: "text/plain; charset=utf-8",
	// 	},
	// },
	// {
	// 	name:    "501",
	// 	request: "/update/counter/testCounter/ 100",
	// 	want: want{
	// 		statusCode:  501,
	// 		response:    "OK",
	// 		contentType: "text/plain; charset=utf-8",
	// 	},
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		request := httptest.NewRequest(http.MethodPost, tt.request, nil)
	// 		request.Header.Set("Content-Type", "text/plain")
	// 		w := httptest.NewRecorder()
	// 		// h := http.HandlerFunc(HandleMetric)
	// 		// h.ServeHTTP(w, request)
	// 		result := w.Result()
	// 		assert.Equal(t, result.StatusCode, tt.want.statusCode)
	// 		assert.Equal(t, result.Header.Get("Content-Type"), tt.want.contentType)
	// 		result.Body.Close()
	// 	})
	// }
}
