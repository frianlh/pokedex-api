package form

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQLInjector(t *testing.T) {
	// argument
	type args struct {
		input string
	}

	// test case
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		// success scenario: test with input no injection
		{
			name: "Success_Without_Injection",
			args: args{
				input: "Unit Testing",
			},
			wantOutput: "Unit Testing",
		},
		// success scenario: test with input injection
		{
			name: "Success_With_Injection",
			args: args{
				input: "$Unit Testing\n",
			},
			wantOutput: "Unit Testing",
		},
	}

	// test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput := SQLInjector(tt.args.input)
			assert.NotNil(t, gotOutput)
			assert.Equal(t, gotOutput, tt.wantOutput)
		})
	}
}
