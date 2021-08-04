# urltopdf: a Go PDF generator

urltopdf is a command line tool to generate a PDF from a URL.
The heavy lifting is done by [Chrome DevTools Protocol clients and tools for Go](github.com/chromedp).

## Usage

    urltopdf [options] url filename

| Options | Shorthand | Description |
| --- | --- | --- |
| `-ignore-certificate-error` | `-ice` | Ignore certificate error |
| `-include-print-background` | `-ipb` | Include print background |
| `-prefer-css-page-size` | `-pcps` | Prefer CSS page size |
| `-with-landscape` | `-wls` | Landscape |

##### Example -

```
urltopdf -ice -ipb https://bbc.co.uk/sport sport.pdf
```
The `.pdf` extension will be added if not provided.