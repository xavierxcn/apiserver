package token

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSign(t *testing.T) {
	type args struct {
		ctx    *gin.Context
		c      Context
		secret string
	}
	tests := []struct {
		name            string
		args            args
		wantTokenString string
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ctx:    nil,
				c:      Context{"root"},
				secret: "123",
			},
			wantTokenString: "",
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTokenString, err := Sign(tt.args.ctx, tt.args.c, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTokenString != tt.wantTokenString {
				t.Errorf("Sign() gotTokenString = %v, want %v", gotTokenString, tt.wantTokenString)
			}
		})
	}
}
