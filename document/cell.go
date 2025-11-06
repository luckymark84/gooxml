// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"fmt"
	"github.com/luckymark84/gooxml/schema/soo/wml"
)

// Cell is a table cell within a document (not a spreadsheet)
type Cell struct {
	d   *Document
	x   *wml.CT_Tc
	sdt *wml.CT_SdtCell
}

// X returns the inner wrapped XML type.
func (c Cell) X() *wml.CT_Tc {
	return c.x
}

// AddParagraph adds a paragraph to the table cell.
func (c Cell) AddParagraph() Paragraph {
	ble := wml.NewEG_BlockLevelElts()
	c.x.EG_BlockLevelElts = append(c.x.EG_BlockLevelElts, ble)
	bc := wml.NewEG_ContentBlockContent()
	ble.EG_ContentBlockContent = append(ble.EG_ContentBlockContent, bc)
	p := wml.NewCT_P()
	bc.P = append(bc.P, p)

	return Paragraph{c.d, p}
}

// AddTable adds a table to the table cell.
func (c Cell) AddTable() Table {
	ble := wml.NewEG_BlockLevelElts()
	c.x.EG_BlockLevelElts = append(c.x.EG_BlockLevelElts, ble)
	bc := wml.NewEG_ContentBlockContent()
	ble.EG_ContentBlockContent = append(ble.EG_ContentBlockContent, bc)
	tbl := wml.NewCT_Tbl()
	bc.Tbl = append(bc.Tbl, tbl)

	return Table{c.d, tbl}
}

// Properties returns the cell properties.
func (c Cell) Properties() CellProperties {
	if c.x.TcPr == nil {
		c.x.TcPr = wml.NewCT_TcPr()
	}
	return CellProperties{c.x.TcPr}
}

// Paragraphs returns the paragraphs defined in the cell.
func (c Cell) Paragraphs() []Paragraph {
	ret := []Paragraph{}
	for _, ble := range c.x.EG_BlockLevelElts {
		for _, cbc := range ble.EG_ContentBlockContent {
			for _, p := range cbc.P {
				ret = append(ret, Paragraph{c.d, p})
			}
		}
	}
	return ret
}

type SdtCell struct {
	s *wml.CT_SdtCell
}

func (s SdtCell) ListItem() []*wml.CT_SdtListItem {
	return s.s.SdtPr.Choice.DropDownList.ListItem
}

func (s SdtCell) SelectByValue(value string) {
	fmt.Printf("%+v\n", s.s.SdtContent.Tc[0].EG_BlockLevelElts[0].
		EG_ContentBlockContent[0].P[0].EG_PContent[1].EG_ContentRunContent[0].R.EG_RunInnerContent[0].T)
	crc := wml.NewEG_ContentRunContent()

	ric := wml.NewEG_RunInnerContent()
	ric.T = &wml.CT_Text{Content: value}

	crc.R = wml.NewCT_R()
	crc.R.EG_RunInnerContent = append(crc.R.EG_RunInnerContent, ric)

	p := &wml.CT_P{
		EG_PContent: []*wml.EG_PContent{
			{
				EG_ContentRunContent: []*wml.EG_ContentRunContent{crc},
			},
		},
	}

	cbc := wml.NewEG_ContentBlockContent()
	cbc.P = append(cbc.P, p)

	ble := wml.NewEG_BlockLevelElts()
	ble.EG_ContentBlockContent = append(ble.EG_ContentBlockContent, cbc)

	tc := &wml.CT_Tc{
		EG_BlockLevelElts: []*wml.EG_BlockLevelElts{ble},
	}

	if s.s.SdtContent.Tc == nil {
		s.s.SdtContent.Tc = append(s.s.SdtContent.Tc, tc)
		return
	}

	for _, stc := range s.s.SdtContent.Tc {
		for _, sble := range stc.EG_BlockLevelElts {
			for _, scbc := range sble.EG_ContentBlockContent {
				for _, pp := range scbc.P {
					for _, pc := range pp.EG_PContent {
						for _, scrc := range pc.EG_ContentRunContent {
							if scrc.R == nil {
								continue
							}
							for _, sric := range scrc.R.EG_RunInnerContent {
								if sric.T == nil {
									continue
								}
								sric.T.Content = value
							}
						}
					}
				}
			}
		}
	}

}

func (c Cell) DropDownListSdt() *StructuredDocument {
	if c.sdt != nil {
		sdtCell := SdtCell{c.sdt}
		return &StructuredDocument{d: c.d, x: sdtCell, pr: c.sdt.SdtPr}
	}
	return nil
}

func (c Cell) AddDropDownSdt(sdt *wml.CT_SdtCell) *StructuredDocument {
	c.sdt = sdt
	sdtCell := SdtCell{c.sdt}
	return &StructuredDocument{d: c.d, x: sdtCell}
}
