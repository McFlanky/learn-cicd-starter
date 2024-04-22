package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		// Valid auth header with api key included
		{
			name:    "Valid Authorization header",
			headers: http.Header{"Authorization": []string{"ApiKey my_api_key"}},
			want:    "my_api_key",
			wantErr: nil,
		},
		// Empty auth header
		{
			name:    "Empty Authorization header",
			headers: http.Header{"Authorization": []string{""}},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		// Malformed auth header
		{
			name:    "Malformed Authorization header",
			headers: http.Header{"Authorization": []string{"Bearer Token"}},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
