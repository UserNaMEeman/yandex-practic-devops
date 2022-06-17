package handler

// import (
// 	"net/http"
// 	"testing"
// )

// func TestHandleGuage(t *testing.T) {
// 	type want struct {
// 		response    string
// 		contentType string
// 	}
// 	tests := []struct {
// 		name string
// 		want want
// 	}{
// 		{
// 			name: "Normal request",
// 			want: want{
// 				response:    nil,
// 				contentType: "text/plain",
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			request := httptest.Newrequest(http.MethodPost, "/update/guage/Alloc/12334", nil)
// 			w := httptest.NewRecorder()
// 			h := http.HandlerFunc(handlers.StatusHandler)
// 			h.ServeHTTP(w, request)
// 			res := w.Result()
// 			// if res.StatusCode != tt.want.code{
// 			// 	t.Errorf("Expected status code %v, got %v", tt.want.code, w.Code)
// 			// }
// 			if res.Header.Get("Content-Type") != tt.want.contentType {
// 				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
// 			}
// 		})
// 	}
// }
