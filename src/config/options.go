package config

import (
	"errors"
	"flag"
	"net/url"
	"os"
	"path/filepath"
)

// Options represents the application IO and flags
type Options struct {
	IgnoreCertError       bool
	WithLandscape         bool
	WithPreferCSSPageSize bool
	WithPrintBackground   bool
	URITarget             string
	OutputTarget          string
}

// Init obtains and sets the options and flags
func (options *Options) Init() error {

	// Read optional flags
	flagset := flag.NewFlagSet("", flag.ExitOnError)
	flagset.BoolVar(&options.IgnoreCertError, "ignore-certificate-error", false, "Ignore certificate error")
	flagset.BoolVar(&options.IgnoreCertError, "ice", false, "Ignore certificate error (shorthand)")
	flagset.BoolVar(&options.WithPreferCSSPageSize, "prefer-css-page-size", false, "Prefer CSS page size")
	flagset.BoolVar(&options.WithPreferCSSPageSize, "pcps", false, "Prefer CSS page size (shorthand)")
	flagset.BoolVar(&options.WithPrintBackground, "include-print-background", false, "Include print background")
	flagset.BoolVar(&options.WithPrintBackground, "ipb", false, "Include print background (shorthand)")
	flagset.BoolVar(&options.WithLandscape, "with-landscape", false, "Landscape")
	flagset.BoolVar(&options.WithLandscape, "wls", false, "Landscape (shorthand)")
	err := flagset.Parse(os.Args[1:])

	if err != nil {
		return err
	}

	// Ensure an input URI and output path exist
	args := flagset.Args()
	const argCount = 2

	if len(args) != argCount {
		return errors.New("invalid number of arguments provided")
	}

	options.URITarget = args[0]
	options.OutputTarget = args[1]

	// Check that the first argument is a valid URI
	parsedURL, err := url.Parse(options.URITarget)

	if err != nil {
		return err
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return errors.New("URI provided is not valid")
	}

	// Make sure the second argument has the correct file extension
	const PDFext = ".pdf"
	if filepath.Ext(options.OutputTarget) != PDFext {
		options.OutputTarget += PDFext
	}

	// Check that the second argument is a valid local path
	_, err = os.Stat(filepath.Dir(options.OutputTarget))

	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("output path provided does not exist")
		}

		return err
	}

	return nil
}
