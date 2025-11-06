package document

import (
	"github.com/luckymark84/gooxml/schema/soo/wml"
)

type StructuredDocument struct {
	d  *Document
	pr *wml.CT_SdtPr
	x  DropDownListInterface
}

func (s StructuredDocument) DropDownList() []*wml.CT_SdtListItem {
	return s.pr.Choice.DropDownList.ListItem
}

func (s StructuredDocument) SelectByValue(value string) {
	found := false
	for _, item := range s.DropDownList() {
		if *item.ValueAttr == value {
			found = true
			break
		}
	}
	if !found {
		return
	}
	s.x.SelectByValue(value)
}

type DropDownListInterface interface {
	ListItem() []*wml.CT_SdtListItem
	SelectByValue(value string)
}
