package pdf_client

import "github.com/SebastiaanKlippert/go-wkhtmltopdf"

type PdfClient struct {
	pdfg *wkhtmltopdf.PDFGenerator
}

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
