package document

import "github.com/luckymark84/gooxml/schema/soo/wml"

type StructuredDocument struct {
	d  *Document
	pr *wml.CT_SdtPr
	c  *wml.CT_SdtContentRun
	x  *wml.CT_SdtRun
}

func (s StructuredDocument) X() *wml.CT_SdtRun {
	return s.x
}

func (s StructuredDocument) DropDownListItem() []*wml.CT_SdtListItem {
	return s.pr.Choice.DropDownList.ListItem
}

func (s StructuredDocument) SelectByValue(value string) {
	for _, v := range s.DropDownListItem() {
		if *v.ValueAttr != value {
			return
		}
	}
	for _, crc := range s.c.EG_ContentRunContent {
		for _, ric := range crc.R.EG_RunInnerContent {
			ric.T.Content = value
		}
	}
}
