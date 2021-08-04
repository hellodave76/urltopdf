package pdf

import (
	"context"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/hellodave76/urltopdf/src/config"
)

type ChromeRunnerInterface interface {
	GeneratePDFData() (buf []byte, err error)
}

// ChromeRunner is a wrapper around chromedp.run
type ChromeRunner struct {
	Options              *config.Options
	ChromeDevToolsRunner func(ctx context.Context, actions ...chromedp.Action) error
}

// Generates the PDF data
func (chromeRunner *ChromeRunner) GeneratePDFData() (buf []byte, err error) {

	// Set up Chrome DevTools with options
	devToolsOptions := chromedp.DefaultExecAllocatorOptions[:]

	if chromeRunner.Options.IgnoreCertError {
		devToolsOptions = append(devToolsOptions, chromedp.Flag("ignore-certificate-errors", true))
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), devToolsOptions...)

	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	defer cancel()

	// Compile a list of instructions for DevTools to execute
	navigateToURL := chromedp.Navigate(chromeRunner.Options.URITarget)

	printToPDF := chromedp.ActionFunc(func(ctx context.Context) error {
		buf, _, err = page.PrintToPDF().
			WithPreferCSSPageSize(chromeRunner.Options.WithPreferCSSPageSize).
			WithPrintBackground(chromeRunner.Options.WithPrintBackground).
			WithLandscape(chromeRunner.Options.WithLandscape).
			Do(ctx)
		return err
	})

	devToolsTasks := chromedp.Tasks{navigateToURL, printToPDF}

	// Execute the tasks
	err = chromeRunner.ChromeDevToolsRunner(ctx, devToolsTasks)

	return buf, err
}
