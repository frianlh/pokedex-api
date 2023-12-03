package encrypt

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCompareHashAndPassword(t *testing.T) {
	// argument
	type args struct {
		hashedPassword *string
		password       *string
	}
	passwordMatch := "Unit Test"
	passwordNotMatch := "Integration Test"
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordMatch), bcrypt.MinCost)
	hashedPassword := string(encryptedPassword)

	// test case
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// success scenario: test password is match
		{
			name: "Success_Password_Is_Match",
			args: args{
				hashedPassword: &hashedPassword,
				password:       &passwordMatch,
			},
			wantErr: false,
		},
		// failed scenario: test password is not match
		{
			name: "Failed_Password_Is_Not_Match",
			args: args{
				hashedPassword: &hashedPassword,
				password:       &passwordNotMatch,
			},
			wantErr: true,
		},
	}

	// test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CompareHashAndPassword(tt.args.hashedPassword, tt.args.password)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
