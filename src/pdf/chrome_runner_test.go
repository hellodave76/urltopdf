package pdf

import (
	"bytes"
	"context"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/hellodave76/urltopdf/src/config"
)

type MockChromedp struct{}

func (mockChromedp MockChromedp) Run(ctx context.Context, actions ...chromedp.Action) error {
	var err error

	return err
}

func TestGeneratePDFData(t *testing.T) {

	options := &config.Options{
		IgnoreCertError: true,
		URITarget:       "https://www.google.co.uk",
	}

	chromeRunner := ChromeRunner{
		Options:              options,
		ChromeDevToolsRunner: MockChromedp{}.Run,
	}
	buf, err := chromeRunner.GeneratePDFData()

	if err != nil {
		t.Error(err.Error())
	}

	if !bytes.Equal(buf, nil) {
		t.Error("expected response with more than 0 bytes")
	}
}
