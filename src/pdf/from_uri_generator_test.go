package pdf

import (
	"errors"
	"testing"
)

type MockChromeRunner struct{}

func (mockChromeRunner *MockChromeRunner) GeneratePDFData() (buf []byte, err error) {
	return buf, nil
}

type MockChromeRunnerFail struct{}

func (mockChromeRunnerFail *MockChromeRunnerFail) GeneratePDFData() (buf []byte, err error) {
	return buf, errors.New("chromeRunner failed")
}

type MockFile struct{}

func (mockFile *MockFile) Write(p []byte) (n int, err error) {
	return 1, nil
}

type MockFileFail struct{}

func (mockFileFail *MockFileFail) Write(p []byte) (n int, err error) {
	return 0, errors.New("file writer failed")
}

func TestGenerate(t *testing.T) {

	generator := FromURIGenerator{
		ChromeRunner: &MockChromeRunner{},
		FileWriter:   &MockFile{},
	}

	err := generator.Generate()

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGenerateFileFail(t *testing.T) {

	generator := FromURIGenerator{
		ChromeRunner: &MockChromeRunner{},
		FileWriter:   &MockFileFail{},
	}

	err := generator.Generate()

	if err.Error() != "file writer failed" {
		t.Errorf("expected 'file writer failed error', got %s", err.Error())
	}
}

func TestGenerateChromeRunnerFail(t *testing.T) {

	generator := FromURIGenerator{
		ChromeRunner: &MockChromeRunnerFail{},
		FileWriter:   &MockFile{},
	}

	err := generator.Generate()

	if err.Error() != "chromeRunner failed" {
		t.Errorf("expected 'chromeRunner failed error', got %s", err.Error())
	}
}
