package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	fromFilePath = "./testdata/input.txt"
	toFilePath   = "./testdata/tmpTest.txt"
)

func TestCopy(t *testing.T) {
	t.Run("полное и частичное копирование", func(t *testing.T) {
		defer os.Remove(toFilePath)
		simpleTests := []struct {
			equalFilePath string
			offset, limit int64
		}{
			{equalFilePath: "./testdata/out_offset0_limit0.txt", offset: 0, limit: 0},
			{equalFilePath: "./testdata/out_offset0_limit10.txt", offset: 0, limit: 10},
			{equalFilePath: "./testdata/out_offset0_limit1000.txt", offset: 0, limit: 1000},
			{equalFilePath: "./testdata/out_offset0_limit10000.txt", offset: 0, limit: 10000},
			{equalFilePath: "./testdata/out_offset100_limit1000.txt", offset: 100, limit: 1000},
			{equalFilePath: "./testdata/out_offset6000_limit1000.txt", offset: 6000, limit: 1000},
		}

		for _, tc := range simpleTests {
			tc := tc
			t.Run(tc.equalFilePath, func(t *testing.T) {
				Copy(fromFilePath, toFilePath, tc.offset, tc.limit)
				require.FileExists(t, toFilePath)
				require.FileExists(t, tc.equalFilePath)
				contentResult, _ := ioutil.ReadFile(toFilePath)
				contentTestFile, _ := ioutil.ReadFile(tc.equalFilePath)
				require.Equal(t, contentResult, contentTestFile)
			})
		}
	})

	t.Run("проверка на ошибки", func(t *testing.T) {
		defer os.Remove(toFilePath)

		err := Copy(fromFilePath, toFilePath, 99999999999999999, 0)
		require.Equal(t, err, ErrOffsetExceedsFileSize)

		err = Copy("./", toFilePath, 0, 0)
		require.NotNil(t, err)

		err = Copy(fromFilePath, "./", 0, 0)
		require.NotNil(t, err)

		err = Copy("", toFilePath, 0, 0)
		require.Equal(t, err, ErrNoFrom)

		err = Copy(fromFilePath, "", 0, 0)
		require.Equal(t, err, ErrNoToPath)
	})
}
