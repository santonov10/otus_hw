package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrNoFrom                = errors.New("отсутствует параметр fromPath")
	ErrNoToPath              = errors.New("отсутствует параметр toPath")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if fromPath == "" {
		return ErrNoFrom
	}
	if toPath == "" {
		return ErrNoToPath
	}

	fileFrom, err := os.OpenFile(fromPath, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}
	defer fileFrom.Close()

	fileFromInfo, err := fileFrom.Stat()
	if err != nil {
		return err
	}
	if offset > fileFromInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	fileWrite, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer fileWrite.Close()

	return copyProcess(fileFrom, fileWrite, offset, limit)
}

func copyProcess(fileFrom, toFile *os.File, offset, limit int64) error {
	readerFF := bufio.NewReader(fileFrom)
	fileFromInfo, err := fileFrom.Stat()
	if err != nil {
		return err
	}

	_, err = readerFF.Discard(int(offset))
	if err != nil {
		return err
	}

	ttlFileSize := fileFromInfo.Size()

	lastOffsetPos := offset + limit
	if lastOffsetPos > ttlFileSize || limit <= 0 {
		lastOffsetPos = ttlFileSize
	}

	buf := make([]byte, 1024)
	currentPos := offset
	var curProgress, ttlProgress int64
	ttlProgress = lastOffsetPos - offset
	for {
		readBytes, errRead := readerFF.Read(buf)

		if errors.Is(errRead, io.EOF) {
			break
		} else if errRead != nil {
			return errRead
		}

		currentPos += int64(readBytes)
		writeBytes := int64(readBytes)

		curProgress += int64(readBytes)
		writeProgressToConsole(curProgress, ttlProgress)

		if currentPos > lastOffsetPos {
			writeBytes = lastOffsetPos - (currentPos - int64(readBytes))
			_, errWrite := toFile.Write(buf[:writeBytes])
			if errWrite != nil {
				return errWrite
			}
			break
		}

		_, errWrite := toFile.Write(buf[:writeBytes])
		if errWrite != nil {
			return errWrite
		}
	}
	return nil
}

func writeProgressToConsole(cur, total int64) {
	if total > 0 {
		if cur > total {
			cur = total
		}

		progress := float64(cur) / float64(total) * 100
		fmt.Printf("\r скопировано на %.2f%%", progress)
		time.Sleep(50 * time.Millisecond) // задержка, чтобы увидеть работу прогресс бара
	}
}
