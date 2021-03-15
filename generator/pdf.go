// package generator

// import (
// 	"bytes"
// 	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
// 	"html/template"
// 	"log"
// )

// type RequestPDF struct {
// 	body string
// }

// func NewRequestPdf(body string) *RequestPDF {
// 	return &RequestPDF{body: body}
// }

// func (r *RequestPDF) ParseTemplate(filename string, data interface{}) error {
// 	t, err := template.ParseFiles(filename)
// 	if err != nil {
// 		return err
// 	}

// 	buf := new(bytes.Buffer)
// 	if err = t.Execute(buf, data); err != nil {
// 		return err
// 	}
// 	r.body = buf.String()
// 	return nil
// }

// func (r, *RequestPDF) GeneratePDF(pdfPath string) (bool, error) {
// 	pdfg, err := wkhtmltopdf.NewPDFGenerator()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return true, nil
// }
