package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) < 1 {
		return -1
	}
	command := cmd[0]
	args := cmd[1:]
	c := exec.Command(command, args...)
	c.Env = append(os.Environ(), preparedEnvSlice(env)...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		// linter error "type assertion on error will fail on wrapped errors. Use errors.As to check for specific errors (errorlint)"
		// не нашел как заменить выражение "err.(*exec.ExitError)" чтобы линтер не ругался
		if exitError, ok := err.(*exec.ExitError); ok { //nolint
			return exitError.ExitCode()
		}
		log.Fatal(err)
	}

	return 0
}

func preparedEnvSlice(env Environment) []string {
	res := make([]string, 0, len(env))
	for key, e := range env {
		envStr := key + "=" + e.Value
		res = append(res, envStr)
	}
	return res
}
