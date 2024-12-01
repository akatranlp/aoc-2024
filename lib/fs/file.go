package fs

import (
	"bufio"
	"io"
	"os"
)

func ReadEntireFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GetScannerForLines(path string) (*bufio.Scanner, io.Closer, error) {
	file, error := os.Open(path)
	if error != nil {
		return nil, nil, error
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner, file, nil
}

func ApplyToLines(path string, f func(string)) error {
	scanner, closer, err := GetScannerForLines(path)
	if err != nil {
		return err
	}
	defer closer.Close()

	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "" {
			break
		}
		f(msg)
	}
	return nil
}
