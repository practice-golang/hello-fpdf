package main // import "hello-fpdf"

import (
	"fmt"
	"math"
	"strconv"

	"github.com/go-pdf/fpdf"
)

type PageDefinition struct {
	PageWidth      float64
	PageHeight     float64
	MarginHorizon  float64 // single side Horizontal margin. should use (MargineX x 2) value when actual area get.
	MarginVertical float64 // single side Vertical margin. should use (MargineX x 2) value when actual area get.
	ActualWidth    float64 // PageWidth - (MarginHorizon x 2)
	ActualHeight   float64 // PageHeight - (MarginVertical x 2)
}

var data [][]string = [][]string{
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
	{"2", "aa", "2.30"},
	{"5", "bb", "5.63"},
	{"10", "cc", "3.63"},
}

var pdf *fpdf.Fpdf

var pageDefs PageDefinition

var (
	gapY   float64 = 2.0
	lineHt float64 = 10.0

	// tableColumnsWidthBase []float64 = []float64{10.0, 75.0, 25.0, 40.0, 40.0} // 190
	tableColumnsWidthBase []float64 = []float64{7.26, 39.47, 11.16, 21.05, 21.05} // 100
	colWidth              []float64
)

func recalcTableColumnsWidth(tableColumnWidths []float64, pageDef PageDefinition) []float64 {
	result := make([]float64, 0)

	for _, w := range tableColumnWidths {
		result = append(result, w/100*pageDef.ActualWidth)
	}

	return result
}

func printTableHeader() {
	header := []string{"No", "설명", "수량", "단가 ($)", "총금액 ($)"}
	pdf.SetFontStyle("B")
	pdf.SetFillColor(200, 200, 200)

	colNumber := len(colWidth)
	for colJ := 0; colJ < colNumber; colJ++ {
		pdf.CellFormat(colWidth[colJ], lineHt, header[colJ], "1", 0, "CM", true, 0, "")
	}
	pdf.Ln(-1)
}

func printTableData() {
	var subtotal float64 = 0.0

	pdf.SetFontStyle("")
	_, pageSizeH := pdf.GetPageSize()
	lastLineY := pageSizeH - pageDefs.MarginVertical - lineHt

	for i := 0; i < len(data); i++ {
		val := data[i]
		if len(val) == 0 {
			continue
		}

		// Column 1 / 2 / 3: Unit / Description / Price per unit
		unit, _ := strconv.Atoi(val[0])
		desc := val[1]
		pricePerUnit, _ := strconv.ParseFloat(val[2], 64)
		pricePerUnit = math.Round(pricePerUnit*100) / 100
		totalPrice := float64(unit) * pricePerUnit

		subtotal += totalPrice

		pdf.CellFormat(colWidth[0], lineHt, fmt.Sprintf("%d", i+1), "1", 0, "CM", false, 0, "")
		pdf.CellFormat(colWidth[1], lineHt, desc, "1", 0, "LM", false, 0, "")
		pdf.CellFormat(colWidth[2], lineHt, fmt.Sprintf("%d", unit), "1", 0, "CM", false, 0, "")
		pdf.CellFormat(colWidth[3], lineHt, fmt.Sprintf("%.2f", pricePerUnit), "1", 0, "CM", false, 0, "")
		pdf.CellFormat(colWidth[4], lineHt, fmt.Sprintf("%.2f", totalPrice), "1", 0, "CM", false, 0, "")
		pdf.Ln(-1)

		if pdf.GetY() >= math.Floor(lastLineY*100)/100 {
			printTableHeader()
			pdf.SetFontStyle("")
		}
	}

	// Print subtotal
	pdf.SetFontStyle("B")
	leftIndent := 0.0
	for i := 0; i < 3; i++ {
		leftIndent += colWidth[i]
	}
	pdf.SetX(pageDefs.MarginHorizon + leftIndent)
	pdf.CellFormat(colWidth[3], lineHt, "금액 합계", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWidth[4], lineHt, fmt.Sprintf("%.2f", subtotal), "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)

	taxAmount := math.Round(subtotal*float64(20)) / 100
	pdf.SetX(pageDefs.MarginHorizon + leftIndent)
	pdf.CellFormat(colWidth[3], lineHt, "부가세", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWidth[4], lineHt, fmt.Sprintf("%.2f", taxAmount), "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)

	grandTotal := subtotal + taxAmount
	pdf.SetX(pageDefs.MarginHorizon + leftIndent)
	pdf.CellFormat(colWidth[3], lineHt, "지불 합계", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWidth[4], lineHt, fmt.Sprintf("%.2f", grandTotal), "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)
}

func main() {
	// pdf = fpdf.New("L", "mm", "A4", "")
	pdf = fpdf.New("P", "mm", "A4", "")

	pageSizeW, pageSizeH := pdf.GetPageSize()
	// fmt.Printf("Page size: %2f , %2f\n", pageSizeW, pageSizeH)

	pageDefs.PageWidth = pageSizeW
	pageDefs.PageHeight = pageSizeH
	pageDefs.MarginHorizon = 10
	pageDefs.MarginVertical = 20
	pageDefs.ActualWidth = pageDefs.PageWidth - pageDefs.MarginHorizon*2
	pageDefs.ActualHeight = pageDefs.PageHeight - pageDefs.MarginVertical*2

	colWidth = recalcTableColumnsWidth(tableColumnsWidthBase, pageDefs)

	// https://pkg.go.dev/github.com/jung-kurt/gofpdf#Fpdf.AddUTF8Font
	pdf.AddUTF8Font("gamtan", "", "fonts/gamtan_road_dotum-regular.ttf")
	pdf.AddUTF8Font("gamtan", "B", "fonts/gamtan_road_dotum-bold.ttf")
	pdf.AddUTF8Font("gamtan", "I", "fonts/gamtan_road_dotum-regular.ttf")
	pdf.AddUTF8Font("gamtan", "BI", "fonts/gamtan_road_dotum-bold.ttf")

	pdf.SetMargins(pageDefs.MarginHorizon, pageDefs.MarginVertical, pageDefs.MarginHorizon)
	pdf.AddPage()

	// Set footer for page
	pdf.SetFooterFunc(func() {
		pdf.SetY(-10)
		pdf.SetFont("Arial", "I", 14)
		pdf.CellFormat(0, 10, "- "+fmt.Sprint(pdf.PageNo())+" -", "", 0, "C", false, 0, "")
	})

	pageW, _ := pdf.GetPageSize()
	safeAreaW := pageW - 2*pageDefs.MarginHorizon

	pdf.ImageOptions("assets/gopher.png", 10, 10, 30, 40, false, fpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")

	pdf.SetFont("gamtan", "", 16)
	_, lineHeight := pdf.GetFontSize()
	lineBreak := lineHeight + float64(1)

	leftY := pdf.GetY() + lineHeight + gapY
	newY := leftY

	invoiceDetailW := float64(40)
	pdf.SetXY(safeAreaW/2+20, newY)
	pdf.Cell(invoiceDetailW, lineHeight, "Invoice No.:")
	pdf.Cell(invoiceDetailW, lineHeight, "a12345")
	pdf.Ln(lineBreak)
	pdf.SetX(safeAreaW/2 + 30)

	currentY := pdf.GetY() + 20
	pdf.SetY(currentY)
	pdf.Cell(40, 10, "n월 정산")
	pdf.Ln(-1)

	printTableHeader() // Table headers

	// Table data
	pdf.SetFontStyle("")
	printTableData()

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err)
	}
}
