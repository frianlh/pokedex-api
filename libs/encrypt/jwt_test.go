package encrypt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewWithClaims(t *testing.T) {
	// argument
	type args struct {
		claims jwt.MapClaims
		jwtKey string
	}

	// test case
	tests := []struct {
		name             string
		args             args
		wantSignedString string
		wantErr          bool
	}{
		// success scenario: test with nil claims
		{
			name: "Success_With_Nil_Claims",
			args: args{
				claims: nil,
				jwtKey: "Unit Testing",
			},
			wantSignedString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.bnVsbA.LV6FIaZ2TBJ3UXo3n446GkN2C4EGut0hz_w_HJ9c8Hg",
			wantErr:          false,
		},
		// success scenario: test with claims
		{
			name: "Success_With_Claims",
			args: args{
				claims: jwt.MapClaims{
					"id":      "79274b58-b7b9-4fac-9f8c-1b7b6b8ff01e",
					"name":    "Unit Test",
					"email":   "unit.test@gmail.com",
					"role_id": "6e3acdce-9b17-498e-aae4-ca8c92cd5b34",
					"exp":     1701434210,
				},
				jwtKey: "Unit Testing",
			},
			wantSignedString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVuaXQudGVzdEBnbWFpbC5jb20iLCJleHAiOjE3MDE0MzQyMTAsImlkIjoiNzkyNzRiNTgtYjdiOS00ZmFjLTlmOGMtMWI3YjZiOGZmMDFlIiwibmFtZSI6IlVuaXQgVGVzdCIsInJvbGVfaWQiOiI2ZTNhY2RjZS05YjE3LTQ5OGUtYWFlNC1jYThjOTJjZDViMzQifQ.Ggwq8uijBFxdtNR3QxtIYrHgnivNDGCeKEuojtkDsk8",
			wantErr:          false,
		},
	}

	// test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSignedString, err := NewWithClaims(tt.args.claims, tt.args.jwtKey)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, gotSignedString)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, gotSignedString)
				assert.Equal(t, gotSignedString, tt.wantSignedString)
			}
		})
	}
}

func TestParse(t *testing.T) {
	// argument
	type args struct {
		auth   string
		jwtKey string
	}
	mapClaims := jwt.MapClaims{
		"id":      "79274b58-b7b9-4fac-9f8c-1b7b6b8ff01e",
		"name":    "Unit Test",
		"email":   "unit.test@gmail.com",
		"role_id": "6e3acdce-9b17-498e-aae4-ca8c92cd5b34",
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	validToken, _ := NewWithClaims(mapClaims, "UnitTesting")
	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVuaXQudGVzdEBnbWFpbC5jb20iLCJleHAiOjE3MDE0MzQyMTAsImlkIjoiNzkyNzRiNTgtYjdiOS00ZmFjLTlmOGMtMWI3YjZiOGZmMDFlIiwibmFtZSI6IlVuaXQgVGVzdCIsInJvbGVfaWQiOiI2ZTNhY2RjZS05YjE3LTQ5OGUtYWFlNC1jYThjOTJjZDViMzQifQ.Ggwq8uijBFxdtNR3QxtIYrHgnivNDGCeKEuojtkDsk8"

	// test case
	tests := []struct {
		name       string
		args       args
		wantClaims jwt.MapClaims
		wantErr    bool
	}{
		// success scenario: test with valid token
		{
			name: "Success_With_Valid_Token",
			args: args{
				auth:   "Bearer " + validToken,
				jwtKey: "UnitTesting",
			},
			wantClaims: mapClaims,
			wantErr:    false,
		},
		// failed scenario: test with expired token
		{
			name: "Failed_With_Expired_Token",
			args: args{
				auth:   "Bearer " + expiredToken,
				jwtKey: "UnitTesting",
			},
			wantClaims: nil,
			wantErr:    true,
		},
		// failed scenario: test with invalid token format
		{
			name: "Failed_With_Invalid_Token",
			args: args{
				auth:   expiredToken,
				jwtKey: "UnitTesting",
			},
			wantClaims: nil,
			wantErr:    true,
		},
	}

	// test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			claims, err := Parse(tt.args.auth, tt.args.jwtKey)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, claims)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, claims["id"], tt.wantClaims["id"])
				assert.Equal(t, claims["name"], tt.wantClaims["name"])
				assert.Equal(t, claims["email"], tt.wantClaims["email"])
				assert.Equal(t, claims["role_id"], tt.wantClaims["role_id"])
			}
		})
	}
}
