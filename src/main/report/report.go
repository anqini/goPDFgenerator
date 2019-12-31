package report

import (
	"github.com/jung-kurt/gofpdf"
	"log"
	"strconv"
)
// report data struct
type ReportData struct {
	Name       string
	ReportDate string
	Country    string
	Gpa        string
	Courses    []Course
}


// read the data as string to prevent float conversion
type Course struct {
	Name  string
	Credits     int
	Percentage       string
	LetterGrade string
	Gradepts    string
}


/*
// test json data
{
	"name": "Anqi Ni",
	"reportDate": "10/23/2019",
	"country":    "United States",
	"gpa":        "2.98",
	"courses": [{"name":"Hello World","credits":3,"letterGrade":"B-","percentage":"83.66","gradepts":"3"},{"name":"Linear Algebra","credits":3,"letterGrade":"B-","percentage":"83.66","gradepts":"3"},{"name":"Linear Algebra","credits":3,"letterGrade":"B-","percentage":"83.66","gradepts":"3"},{"name":"Linear Algebra","credits":3,"letterGrade":"B-","percentage":"83.66","gradepts":"3"},{"name":"Linear Algebra","credits":3,"letterGrade":"B-","percentage":"83.66","gradepts":"3"},{"name":"Linear Algebra","credits":3,"letterGrade":"B-","percentage":"83.66","gradepts":"3"}]
}
*/

// test go struct data
var reportd ReportData = ReportData{
	Name:       "Anqi Ni",
	ReportDate: "10/23/2019",
	Country:    "United States",
	Gpa:        "2.98",
	Courses: []Course{
		Course{
			Name:  "Social Anthropology",
			Credits:     3,
			Percentage:       "83.66",
			LetterGrade: "B",
			Gradepts:    "3.00"},
		Course{
			Name:  "Computer Science 177",
			Credits:     6,
			Percentage:       "93.66",
			LetterGrade: "A",
			Gradepts:    "$.00"},
		Course{
			Name:  "Computer Science 180",
			Credits:     3,
			Percentage:       "83.66",
			LetterGrade: "B",
			Gradepts:    "3.00"},
		Course{
			Name:  "Social Anthropology",
			Credits:     3,
			Percentage:       "83.66",
			LetterGrade: "B",
			Gradepts:    "3.00"}}}



func GenReport(reportd ReportData) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	reportInfo(pdf, reportd)
	printTable(pdf, reportd)
	lastPage(pdf)
	err := pdf.OutputFileAndClose("public/hello.pdf")
	if err != nil {
		log.Fatal("file is not created...")
	}
}
func reportInfo(pdf *gofpdf.Fpdf, reportd ReportData) {
	pdf.AddPage()
	pdf.Image("public/logo.png", 10, 20, 80, 0, false, "", 0, "")
	pdf.Ln(30)
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(15, 15, "GPA Report")
	pdf.Ln(10)
	// Name section
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(0, 15, "Name: ")
	pdf.SetX(50)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 15, reportd.Name)
	pdf.Ln(7)
	// report ID
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(0, 15, "Report ID")
	pdf.SetX(50)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 15, "107373")
	pdf.Ln(7)
	// Report Date
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(0, 15, "Report Date")
	pdf.SetX(50)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 15, reportd.ReportDate)
	pdf.Ln(7)
	// Country
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(0, 15, "Country")
	pdf.SetX(50)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 15, reportd.Country)
	pdf.Ln(7)
	// Grading Scale
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(0, 15, "Grading Scale")
	pdf.SetX(50)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 15, "Most Common")
	pdf.Ln(7)
	// GPA
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(0, 15, "GPA")
	pdf.SetX(50)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 15, reportd.Gpa)
}

func printTable(pdf *gofpdf.Fpdf, reportd ReportData) {
	// Table header
	pdf.SetFont("Arial", "B", 10)
	pdf.Ln(-1)
	pdf.CellFormat(15, 6, "", "1", 0, "", false, 0, "")
	pdf.CellFormat(50, 6, "Course Name", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(30, 6, "Credits/Hours", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(30, 6, "Grade", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(30, 6, "Letter Grade", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(30, 6, "Grade Points", "1", 0, "CM", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	for i, v := range reportd.Courses {
		pdf.Ln(-1)
		var width float64 = 8
		pdf.CellFormat(15, width, strconv.Itoa(i), "1", 0, "CM", false, 0, "")
		pdf.CellFormat(50, width, v.Name, "1", 0, "CM", false, 0, "")
		pdf.CellFormat(30, width, strconv.Itoa(v.Credits), "1", 0, "CM", false, 0, "")
		pdf.CellFormat(30, width, v.Percentage, "1", 0, "CM", false, 0, "")
		pdf.CellFormat(30, width, v.LetterGrade, "1", 0, "CM", false, 0, "")
		pdf.CellFormat(30, width, v.Gradepts, "1", 0, "CM", false, 0, "")
	}
}

func lastPage(pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 15)
	pdf.Cell(80, 15, "Letter grade and the numerical equivalents used for this calculator")
	pdf.Ln(20)
	// Times 12
	pdf.SetFont("Times", "", 12)
	// Output justified text
	txt := "Grade point average (GPA) is a commonly used indicator of an individual's academic achievement in school. It is the average of the grades attained in each course, taking course credit into consideration. Grading systems vary in different countries, or even schools. This calculator accepts letter grades as well as numerical inputs. These letter grades and percentile grades are translated into numerical values as shown below."
	pdf.MultiCell(0, 5, txt, "", "", false)
	// Line break
	pdf.Ln(10)

	txtArr := []string{"A+ = 4.3 grade points",
		"A = 4 grade points",
		"A- = 3.7 grade points",
		"B+ = 3.3 grade points",
		"B = 3 grade points",
		"B- = 2.7 grade points",
		"C+ = 2.3 grade points",
		"C = 2 grade points",
		"C- = 1.7 grade points",
		"D+ = 1.3 grade points",
		"D = 1 grade point",
		"D- = 0.7 grade points",
		"F = 0 grade points",
		"P (pass), NP (not pass), I (incomplete), W (withdrawal) will be ignored."}
	left := (210.0 - 4*40) / 2
	// pdf.SetX(left)
	// for _, str := range header {
	// 	pdf.CellFormat(40, 7, str, "1", 0, "", false, 0, "")
	// }
	for i, v := range txtArr {
		pdf.Ln(-1)
		pdf.SetX(left)
		if i == 0 {
			pdf.CellFormat(150, 6, v, "LTR", 0, "", false, 0, "")
		} else if i == len(txtArr)-1 {
			pdf.CellFormat(150, 6, v, "LBR", 0, "", false, 0, "")
		} else {
			pdf.CellFormat(150, 6, v, "LR", 0, "", false, 0, "")
		}
	}
	pdf.Ln(20)
	txt = "Most schools, colleges, and universities in the United States use a grading system based on the letters above, though E is sometimes used instead of F. Grading systems do differ however based on what constitutes an A or B, and some do not include grades such as an A+ or a B-. Others may attribute more weight to certain courses, and thus whatever grade is attained in the course will have a larger effect on overall GPA. The calculator can account for this based on the number of credits attributed to a course, where credit is the \"weighting\" of the course, as shown in the examples below."
	pdf.MultiCell(0, 5, txt, "", "", false)
	pdf.Ln(-1)
	pdf.Image("public/table1.png", 10, 210, 80, 0, false, "", 0, "")
	pdf.Image("public/table2.png", 100, 210, 90, 0, false, "", 0, "")

	pdf.Image("public/logo.png", 140, 280, 60, 0, false, "", 0, "")
}
