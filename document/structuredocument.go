package document

import "github.com/luckymark84/gooxml/schema/soo/wml"

type StructuredDocument struct {
	d  *Document
	pr *wml.CT_SdtPr
	c  *wml.CT_SdtContentRun
}

//func (p Paragraph) DropDownList()  {
//
//}
