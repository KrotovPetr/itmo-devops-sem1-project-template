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
	return findCSVInTar(tarReader)
}

func findCSVInTar(tarReader *tar.Reader) (io.ReadCloser, error) {
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if isCSVFile(header.Name) && !isHiddenFile(header.Name) {
			return io.NopCloser(tarReader), nil
		}
	}
	return nil, errors.New("CSV file not found in the archive")
}

func isHiddenFile(fileName string) bool {
	base := filepath.Base(fileName)
	return len(base) > 0 && base[0] == '.'
}
