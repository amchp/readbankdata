package bankreaders

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    data, err := io.ReadAll(file)
    if err!= nil {
        return nil, err
    }
    return data, err
}

func parseCSV(data []byte, separator rune) (*csv.Reader, error) {
    reader := csv.NewReader(bytes.NewReader(data))
    reader.Comma = separator
    if _, err := reader.Read(); err != nil {
        return nil, err
    }
    return reader, nil
}
