package pdf_client

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/dslipak/pdf"
	"io"
)

type (
	PDF interface {
		PDFG() *wkhtmltopdf.PDFGenerator
		PDFR() *pdf.Reader
		ReadPDFContent(f io.ReaderAt, size int64) error
	}

	PdfClient struct {
		pdfg *wkhtmltopdf.PDFGenerator
		pdfr *pdf.Reader
	}

	PdfFileInfo struct {
		NumPage int
	}
)

func NewPdfClient() (*PdfClient, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	return &PdfClient{
		pdfg: pdfg,
	}, nil
}

func (c *PdfClient) PDFG() *wkhtmltopdf.PDFGenerator {
	return c.pdfg
}

func (c *PdfClient) PDFR() *pdf.Reader {
	return c.pdfr
}

func (c *PdfClient) ReadPDFContent(f io.ReaderAt, size int64) error {
	reader, err := pdf.NewReader(f, size)
	if err != nil {
		return err
	}
	reader.NumPage()
	return nil
}
