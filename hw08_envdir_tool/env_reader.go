package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value string
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	resEnv := Environment{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.Mode().IsDir() {
			continue
		}
		//- имя `S` не должно содержать `=`;
		if strings.Contains(f.Name(), "=") {
			continue
		}
		filePath := dir + "/" + f.Name()
		fileValue, err := getValueFromFile(filePath)
		if err != nil {
			continue
		}
		resEnv[f.Name()] = EnvValue{Value: fileValue}
	}

	return resEnv, nil
}

func getValueFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	reader := bufio.NewReader(file)
	line, err := reader.ReadString(byte('\n'))
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}
	line = strings.TrimRight(line, " \t\n\r")
	line = strings.ReplaceAll(line, "\x00", "\n")

	return line, nil
}
