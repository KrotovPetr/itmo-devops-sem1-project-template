package utils

import (
	"archive/tar"
	"bytes"
	"errors"
	"io"
	"path/filepath"
)

func UntarFile(r io.Reader) (io.ReadCloser, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	tarReader := tar.NewReader(bytes.NewReader(data))

	for header, err := tarReader.Next(); err != io.EOF; {
		if err != nil {
			return nil, err
		}

		if isValidCSVFile(header.Name) {
			return io.NopCloser(tarReader), nil
		}
	}

	return nil, errors.New("CSV file not found")
}

func isValidCSVFile(filename string) bool {
	return filepath.Ext(filename) == ".csv" && filepath.Base(filename)[0] != '.'
}
