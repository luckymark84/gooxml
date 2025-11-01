package main

import "github.com/luckymark84/gooxml/document"

func main() {
	testdocx := "testdoc.docx"

	doc, _ := document.Open(testdocx)

	doc.SaveToFile("result.docx")
}
