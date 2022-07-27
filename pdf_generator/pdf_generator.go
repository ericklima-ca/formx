package pdf_generator

import (
	"log"
	"os"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type data interface {
	GetData() [][]string
}

func BuildPDF(seed data) {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildDataList(m, seed)

	err := m.OutputFileAndClose("./temp/_tmp.pdf")
	if err != nil {
		log.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}
	log.Println("PDF saved successfully")
}

func buildHeading(m pdf.Maroto) {
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("General Data", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
		})
	})
}

func buildDataList(m pdf.Maroto, seed data) {
	tableHeadings := []string{"Field", "Value"}
	// contents := [][]string{
	// {"Name", "Amorim S/A"},
	// {"Email", "email@email.com"},
	// {"Phone", "+5599991234567"},
	// }
	contents := seed.GetData()

	lightPurpleColor := getLightPurpleColor()
	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Infos", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{5, 7},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{5, 7},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}
