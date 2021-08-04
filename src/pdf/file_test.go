package pdf

import (
	"crypto/rand"
	"fmt"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {

	filename := "test-write.pdf"

	file := File{
		Filename: filename,
	}

	buf := make([]byte, 4)
	_, err := rand.Read(buf)

	if err != nil {
		t.Error(err.Error())
	}

	n, err := file.Write(buf)
	defer os.Remove(filename)

	if err != nil {
		t.Error(err.Error())
	}

	if n != 4 {
		t.Errorf("expected 4 bytes written, got %d", n)
	}
}

func TestWriteFail(t *testing.T) {

	filename := "/non-existent-directory/test-write.pdf"

	file := File{
		Filename: filename,
	}

	buf := make([]byte, 4)
	_, err := rand.Read(buf)

	if err != nil {
		t.Error(err.Error())
	}

	n, err := file.Write(buf)

	if err.Error() != fmt.Sprintf("open %s: no such file or directory", filename) {
		t.Errorf("expected error 'no such file or directory' got %s", err.Error())
	}

	if n != 0 {
		t.Errorf("expected 0 bytes written, got %d", n)
	}
}
