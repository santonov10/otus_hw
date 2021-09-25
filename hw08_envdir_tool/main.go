package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatal(`Должно быть минимум 2 аргумента при вызове: 
					1 - путь к папке с файлами с переменными окружения
					2 - вызываемая программа`)
	}
	envDir := args[1]
	execute := args[2:]
	env, err := ReadDir(envDir)
	if err != nil || len(env) == 0 {
		log.Fatal(`Ошибка чтения данных из папки`, err)
	}

	os.Exit(RunCmd(execute, env))
}
