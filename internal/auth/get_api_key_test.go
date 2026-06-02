package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type testCase struct {
		header  http.Header
		want    string
		wantErr bool
	}

	headerValidFormat := http.Header{}
	headerValidFormat.Add("Authorization", "ApiKey secret_token")
	headerInvalidFormat := http.Header{}
	headerInvalidFormat.Add("Authorization", "Invalid Format")
	headerNoAuth := http.Header{}

	tests := map[string]testCase{
		"valid API Key": {
			header:  headerValidFormat,
			want:    "secret_token",
			wantErr: false,
		},
		"invalid API Key": {
			header:  headerInvalidFormat,
			want:    "",
			wantErr: true,
		},
		"no Auth Header": {
			header:  headerNoAuth,
			want:    "",
			wantErr: true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.header)

			if tc.wantErr && err == nil {
				t.Error("expected error, got nil")
			}

			if got != tc.want {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}
