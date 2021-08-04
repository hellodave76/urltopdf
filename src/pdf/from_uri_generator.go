package pdf

import (
	"io"
)

// FromURIGenerator represents a PDF writer that takes a URI as its input
type FromURIGenerator struct {
	ChromeRunner ChromeRunnerInterface
	FileWriter   io.Writer
}

func (fromURIGenerator *FromURIGenerator) Generate() error {

	// Generate the PDF data
	buf, err := fromURIGenerator.ChromeRunner.GeneratePDFData()

	if err != nil {
		return err
	}

	// Write the resultant data to the output file
	_, err = fromURIGenerator.FileWriter.Write(buf)

	return err
}
