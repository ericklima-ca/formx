package main

import (
	"log"
	"os"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)
	err := m.OutputFileAndClose("temp/div_rhino_fruit.pdf")
	if err != nil {
		log.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}
	log.Println("PDF saved successfully")
}
