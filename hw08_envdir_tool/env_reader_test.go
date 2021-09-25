package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testDir = "./testdata/env"

func TestReadDir(t *testing.T) {
	t.Run("чтение только файлов в реальной директории", func(t *testing.T) {
		checkKeysExist := []string{"BAR", "EMPTY", "FOO", "HELLO", "UNSET"}
		checkKeysNotExist := []string{"testfolder", "BAD=FILE"}

		envs, err := ReadDir(testDir)

		require.Nil(t, err)

		for _, keyExist := range checkKeysExist {
			_, ok := envs[keyExist]
			require.True(t, ok, "должен быть ключ "+keyExist)
		}

		for _, keyNotExist := range checkKeysNotExist {
			_, ok := envs[keyNotExist]
			require.False(t, ok, "не должно быть ключа "+keyNotExist)
		}

		unrealDir := "./unrealDir"
		_, err = ReadDir(unrealDir)
		require.Error(t, err, "должна быть ошибка на чтение не существующей директории "+unrealDir)
	})
}
