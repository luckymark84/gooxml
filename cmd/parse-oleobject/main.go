package main

import (
	"flag"
	"github.com/luckymark84/gooxml"
	"github.com/luckymark84/gooxml/document"
)

func main() {
	var input string
	var output string
	flag.StringVar(&input, "input", "", "input file")
	flag.StringVar(&output, "output", "", "output file")
	flag.Parse()

	gooxml.DisableLogging()
	doc, _ := document.Open(input)

	tables := doc.Tables()

	tbl := tables[6]
	control := tbl.Rows()[1].Cells()[0].ContentControls()
	if control != nil {
		//options := control.GetOptions()
		//fmt.Println(control.Tag)
		control.SelectByValue("test")
	}
	doc.SaveToFile(output)

}
