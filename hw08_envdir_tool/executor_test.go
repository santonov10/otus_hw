package main

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	// не знаю как протестировать в windows установку переменных окружения
	t.Run("проверка выполнения команд на windows", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			testNewFilePath := "./testdata/testfile.txt"
			defer os.Remove(testNewFilePath)

			cmds := []string{"cmd", "/C", "echo", "testdata", ">", testNewFilePath}
			RunCmd(cmds, nil)
			require.FileExists(t, testNewFilePath)
		}
	})
}
