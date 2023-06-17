// Package stack implements a generic set.

// T must be comparable.

package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TEST",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet[string]()
			s.Insert("test")
			assert.Equal(t, true, s.Exist("test"))
			assert.Equal(t, false, s.Exist("test no"))
			s.Delete("test")
			assert.Equal(t, false, s.Exist("test"))
		})
	}
}
