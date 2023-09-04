package validator

import (
	"errors"
	"testing"
)

type cases struct {
	name, field, value, param string
	expected                  error
}

func TestRule_Required(t *testing.T) {
	var testcases = []cases{
		{
			name:     "valid",
			field:    "name",
			value:    "John Doe",
			expected: nil,
		},
		{
			name:     "not valid",
			field:    "name",
			expected: errors.New("name is required"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := required(tc.field, tc.value, tc.param)
			if tc.expected != nil {
				if tc.expected.Error() != err.Error() {
					t.Errorf("error want %v got %v", tc.expected, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error %v", err)
				}
			}
		})
	}
}
