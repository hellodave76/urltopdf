package config

import (
	"os"
	"testing"
)

const outputTarget = "test-output.pdf"
const uri = "https://www.google.co.uk"

func TestInit(t *testing.T) {

	cases := []struct {
		Args []string
	}{
		{[]string{uri, outputTarget}},
		{[]string{"-ice", "-pcps", "-ipb", "-wls", uri, outputTarget}},
		{[]string{"-ice", "-pcps", "-ipb", uri, outputTarget}},
		{[]string{"-ice", "-pcps", "-wls", uri, outputTarget}},
		{[]string{"-ice", "-ipb", "-wls", uri, outputTarget}},
		{[]string{"-pcps", "-ipb", "-wls", uri, outputTarget}},
		{[]string{"-ice", "-pcps", uri, outputTarget}},
		{[]string{"-ice", "-wls", uri, outputTarget}},
		{[]string{"-ice", "-ipb", uri, outputTarget}},
		{[]string{"-pcps", "-ipb", uri, outputTarget}},
		{[]string{"-pcps", "-wls", uri, outputTarget}},
		{[]string{"-ipb", "-wls", uri, outputTarget}},
		{[]string{"-ice", uri, outputTarget}},
		{[]string{"-pcps", uri, outputTarget}},
		{[]string{"-ipb", uri, outputTarget}},
		{[]string{"-wls", uri, outputTarget}},
		{[]string{"--ignore-certificate-error", "--prefer-css-page-size", "--include-print-background", "--with-landscape", uri, outputTarget}},
		{[]string{uri, "test-output"}},
	}

	for _, c := range cases {
		args := append([]string{"urltopdf"}, c.Args...)

		iceSet, pcpsSet, ipbSet, wlsSet := false, false, false, false
		for _, arg := range args {
			if arg == "-ice" || arg == "--ignore-certificate-error" {
				iceSet = true
			}
			if arg == "-pcps" || arg == "--prefer-css-page-size" {
				pcpsSet = true
			}
			if arg == "-ipb" || arg == "--include-print-background" {
				ipbSet = true
			}
			if arg == "-wls" || arg == "--with-landscape" {
				wlsSet = true
			}
		}

		os.Args = args
		options := Options{}
		err := options.Init()

		if err != nil {
			t.Error(err.Error())
		}
		if iceSet != options.IgnoreCertError {
			t.Errorf("ignore Certificate Error flag set incorrectly, expected '%t', got '%t'", iceSet, options.IgnoreCertError)
		}
		if wlsSet != options.WithLandscape {
			t.Errorf("with Landscape flag set incorrectly, expected '%t', got '%t'", wlsSet, options.WithLandscape)
		}
		if ipbSet != options.WithPrintBackground {
			t.Errorf("include Print Background flag set incorrectly, expected '%t', got '%t'", ipbSet, options.WithPrintBackground)
		}
		if pcpsSet != options.WithPreferCSSPageSize {
			t.Errorf("with Prefer CSS Page Size flag set incorrectly, expected '%t', got '%t'", pcpsSet, options.WithPreferCSSPageSize)
		}
		if options.OutputTarget != outputTarget {
			t.Errorf("expected output path to be '%s', got '%s'", outputTarget, options.OutputTarget)
		}
	}
}

func TestInitFail(t *testing.T) {

	cases := []struct {
		Args  []string
		Error string
	}{
		{
			[]string{"urltopdf", uri},
			"invalid number of arguments provided",
		},
		{
			[]string{"urltopdf", "not-a-uri", outputTarget},
			"URI provided is not valid",
		},
		{
			[]string{"urltopdf", "s11*hp://www.google.co.uk", outputTarget},
			"parse \"s11*hp://www.google.co.uk\": first path segment in URL cannot contain colon",
		},
		{
			[]string{"urltopdf", uri, "not-a-real-directory/test-output.pdf"},
			"output path provided does not exist",
		},
	}

	for _, c := range cases {
		os.Args = c.Args
		options := Options{}
		err := options.Init()

		if err != nil {
			if err.Error() != c.Error {
				t.Errorf("expected '%s' error, got '%s'", c.Error, err.Error())
			}
		}
	}
}
