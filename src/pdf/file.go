package pdf

import (
	"os"
)

// File handles writing the PDF to disk, and implements the io.Writer interface
type File struct {
	Filename string
}

// Write the File to disk
func (f *File) Write(p []byte) (n int, err error) {

	osFile, err := os.Create(f.Filename)

	if err != nil {
		return 0, err
	}

	return osFile.Write(p)
}
