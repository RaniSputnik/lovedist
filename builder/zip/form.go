package zip

import (
	zip_impl "archive/zip"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// ArchiveMultipartFormFiles archives files uploaded from a multipart form
func ArchiveMultipartFormFiles(in []*multipart.FileHeader, writer io.Writer, progress ProgressFunc) error {
	zipWriter := zip_impl.NewWriter(writer)

	for _, f := range in {
		parts := strings.Split(f.Filename, string(os.PathSeparator))
		archivePath := filepath.Join(parts[1:]...)

		if progress != nil {
			progress(archivePath)
		}

		file, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			_ = file.Close()
		}()

		zipFileWriter, err := zipWriter.Create(archivePath)
		if err != nil {
			return err
		}

		if _, err = io.Copy(zipFileWriter, file); err != nil {
			return err
		}
	}

	return zipWriter.Close()
}
