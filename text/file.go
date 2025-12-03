package text

import (
	"bufio"
	"io"
	"os"
)

func ReadFileByLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	raw := bufio.NewScanner(f)
	in := make([]string, 0, len(raw.Bytes()))
	for raw.Scan() {
		in = append(in, raw.Text())
	}

	return in, nil
}

func ReadFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	in, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(in), nil
}
