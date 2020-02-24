package main

import (
	"io/ioutil"
	"log"

	generator "github.com/Shasta/go-invoice-generator"
)

func main() {
	err := GenerateInvoicePDF()
	if err != nil {
		log.Println("Error generating the PDF")
	}
}

func GenerateInvoicePDF() error {
	doc, err := generator.New(generator.INVOICE, &generator.Options{
		TextTypeInvoice: "Invoice",
		AutoPrint:       true,
	})
	if err != nil {
		return err
	}

	doc.SetHeader(&generator.HeaderFooter{
		Text: "Some header text",
	})

	doc.SetFooter(&generator.HeaderFooter{
		Text:       "<center>Invoice to: Terms & Conditions:All payments must be made in the name of: SHASTA TECHNOLOGIES SL</center>",
		Pagination: true,
	})

	// doc.SetNumber("testnumber")
	// doc.SetVersion("test version")

	doc.SetDescription("Some text describing document")

	logoBytes, _ := ioutil.ReadFile("/home/osboxes/Pictures/shasta/logoShasta2.png")
	// cif := "B67394254"
	// email := "alex@shasta.me"
	doc.SetCompany(&generator.Contact{
		Name: "SHASTA TECHNOLOGIES SL",
		// Cif:   &cif,
		// Email: &email,
		Logo: &logoBytes, // Image as byte array, supported format: png, jpeg, gif
		Address: &generator.Address{
			Address:    "Rambla Catalunya 124",
			PostalCode: "08008",
			City:       "Barcelona",
			Country:    "Espana",
		},
	})

	// doc.SetDate("07-01-2020") If not specified will be filled automaticly
	doc.SetCustomer(&generator.Contact{
		Name: "Specter Labs",
		Address: &generator.Address{
			Address:    "89 Rue de Paris",
			PostalCode: "29200",
			City:       "Brest",
			Country:    "France",
		},
	})

	doc.AppendItem(&generator.Item{
		Name:     "Card Payin",
		UnitCost: "550",
		Quantity: "1",
		Tax: &generator.Tax{
			Percent: "21",
		},
	})

	// doc.AppendItem(&generator.Item{
	// 	Name:     "Item two",
	// 	UnitCost: "5.89",
	// 	Quantity: "11",
	// 	Tax: &generator.Tax{
	// 		Amount: "10",
	// 	},
	// })
	doc.SetRef("img_reference")

	pdf, err := doc.Build()

	if err != nil {
		return err
	}

	err = pdf.OutputFileAndClose("/home/osboxes/Documents/Invoices/invoice_test2.pdf")

	if err != nil {
		return err
	}
	return nil
}
