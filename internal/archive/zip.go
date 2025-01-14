package archive
import (
	"archive/zip"
	"bytes"
	"io"
)

func ZipFile(buffer *bytes.Buffer, w io.Writer, fileName string) error {
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	file, err := zipWriter.Create(fileName)
	if err != nil {
		return err
	}
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}