package main

import (
	"log"

	"github.com/chromedp/chromedp"
	"github.com/hellodave76/urltopdf/src/config"
	"github.com/hellodave76/urltopdf/src/pdf"
)

func main() {

	// Obtain input URI, output path and options set with flags
	options := &config.Options{}
	err := options.Init()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Input URI: ", options.URITarget)
	log.Println("Output Path: ", options.OutputTarget)

	if options.IgnoreCertError {
		log.Println("Ignoring certificate errors...")
	}

	if options.WithPreferCSSPageSize {
		log.Println("Preferring CSS page size...")
	}

	if options.WithPrintBackground {
		log.Println("Including print background...")
	}

	if options.WithLandscape {
		log.Println("Using landscape orientation...")
	}

	// Generate the PDF
	file := &pdf.File{
		Filename: options.OutputTarget,
	}

	chromeRunner := &pdf.ChromeRunner{
		Options:              options,
		ChromeDevToolsRunner: chromedp.Run,
	}

	generator := &pdf.FromURIGenerator{
		ChromeRunner: chromeRunner,
		FileWriter:   file,
	}

	err = generator.Generate()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Complete")
}
