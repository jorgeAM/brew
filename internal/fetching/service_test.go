package fetching

import (
	"errors"
	"testing"

	"github.com/jorgeAM/brew/internal/storage/inmem"
	"github.com/stretchr/testify/assert"
)

func TestFetchByID(t *testing.T) {
	repo := inmem.NewRepository()
	service := NewService(repo)

	tests := map[string]struct {
		input    int
		expected int
		err      error
	}{
		"valid beer": {
			input:    127,
			expected: 127,
			err:      nil,
		},
		"not found beer": {
			input: 128,
			err:   errors.New("beer not found"),
		},
		"no id passed": {
			err: errors.New("arguments missing"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			beer, err := service.FetchByID(test.input)

			if test.err != nil {
				assert.Error(t, err)
			}

			if test.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, test.expected, beer.ID)
		})
	}
}
