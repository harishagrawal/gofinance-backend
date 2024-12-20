package util

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	type args struct {
		headerValue string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Scenario 1: Valid Authorization Header",
			args:    args{headerValue: "Bearer testtoken"},
			wantErr: false,
		},
		{
			name:    "Scenario 2: Invalid Authorization Header Key",
			args:    args{headerValue: ""},
			wantErr: true,
		},
		{
			name:    "Scenario 3: Invalid JWT Token",
			args:    args{headerValue: "Bearer invalidtoken"},
			wantErr: true,
		},
		{
			name:    "Scenario 4: Null Context Passed",
			args:    args{headerValue: nil},
			wantErr: true,
		},
		{
			name:    "Scenario 5: Multiple Fields in Authorization Header",
			args:    args{headerValue: "Bearer testtoken testfield"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &http.Request{}
			request.Header = make(http.Header)
			request.Header.Set("Authorization", tt.args.headerValue)

			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request = request

			if err := GetTokenInHeaderAndVerify(c); (err != nil) != tt.wantErr {
				t.Errorf("GetTokenInHeaderAndVerify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

