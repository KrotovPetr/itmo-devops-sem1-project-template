package archive

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"path/filepath"
)

func UnzipFile(r io.Reader) (io.ReadCloser, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}

	return findCSVFile(zipReader)
}

func findCSVFile(zipReader *zip.Reader) (io.ReadCloser, error) {
	for _, file := range zipReader.File {
		if !(filepath.Ext(file.Name) == ".csv") {
			continue
		}

		rc, err := file.Open()
		if err != nil {
			return nil, err
		}

		return rc, nil
	}

	return nil, errors.New("CSV file not found in the archive")
}
