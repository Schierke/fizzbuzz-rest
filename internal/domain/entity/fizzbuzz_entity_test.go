package entity_test

import (
	"fizzbuzz/internal/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewEntity(t *testing.T) {

	type inputs struct {
		int1  int
		int2  int
		limit int
		str1  string
		str2  string
	}

	tests := []struct {
		name      string
		inputs    inputs
		wantError error
	}{
		{
			name: "OK",
			inputs: inputs{
				3,
				5,
				15,
				"fizz",
				"buzz",
			},
			wantError: nil,
		},
		{
			name: "Invalid string input",
			inputs: inputs{
				3,
				5,
				15,
				"",
				"buzz",
			},
			wantError: entity.ErrInvalidStringInput,
		},
		{
			name: "Invalid interger input",
			inputs: inputs{
				-1,
				-1,
				24,
				"fizz",
				"buzz",
			},
			wantError: entity.ErrInvalidIntegerInput,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := entity.NewFizzBuzzEntity(test.inputs.int1, test.inputs.int2, test.inputs.limit, test.inputs.str1, test.inputs.str2)

			if test.wantError != nil {
				assert.EqualError(t, err, test.wantError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_CalculateFizzBuzz(t *testing.T) {
	tests := []struct {
		name   string
		entity entity.FizzBuzz
		want   []string
	}{
		{
			name: "OK",
			entity: entity.FizzBuzz{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			want: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := test.entity.CalculateFizzBuzz()

			assert.Equal(t, test.want, got)
		})
	}

}
