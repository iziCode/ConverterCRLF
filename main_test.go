package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllFilesFromPath(t *testing.T) {
	t.Run("correct input", func(t *testing.T) {
		countFiles := len(GetAllFilesFromPath("test"))
		assert.Equal(t, 4, countFiles)
	})

	t.Run("wrong input", func(t *testing.T) {
		//...
	})

}
