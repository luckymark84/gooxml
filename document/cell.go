// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"github.com/luckymark84/gooxml/schema/soo/wml"
	"sync"
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

type DropDownControl struct {
	mu      sync.Mutex
	pr      *wml.CT_SdtPr
	Tag     string
	content *wml.CT_Text
}

type Option struct {
	Label string
	Value string
}

func (d *DropDownControl) GetOptions() []*Option {
	var options []*Option
	for _, item := range d.pr.Choice.DropDownList.ListItem {
		option := Option{
			*item.DisplayTextAttr,
			*item.ValueAttr,
		}
		options = append(options, &option)
	}
	return options
}

func (d *DropDownControl) SelectByValue(value string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, option := range d.GetOptions() {
		if option.Value == value {
			d.content.Content = value
		}
	}
}

func (c Cell) DropDownListSdt() *StructuredDocument {
	if c.sdt != nil {
		return &StructuredDocument{d: c.d, x: c.sdt}
	}
	return nil
}

func (c Cell) AddDropDownSdt(sdt *wml.CT_SdtCell) *StructuredDocument {
	c.sdt = sdt
	return &StructuredDocument{d: c.d, x: c.sdt}
}

func (c Cell) ContentControls() *DropDownControl {
	if c.sdt != nil {
		var content *wml.CT_Text
		for _, ctCell := range c.sdt.SdtContent.Tc {
			for _, ble := range ctCell.EG_BlockLevelElts {
				for _, cbc := range ble.EG_ContentBlockContent {
					for _, p := range cbc.P {
						for _, pc := range p.EG_PContent {
							for _, crc := range pc.EG_ContentRunContent {
								if crc.R == nil {
									continue
								}
								for _, ric := range crc.R.EG_RunInnerContent {
									content = ric.T
									goto Label
								}
							}
						}
					}
				}
			}
		}
	Label:
		if content == nil {
			return nil
		}

		return &DropDownControl{pr: c.sdt.SdtPr, Tag: c.sdt.SdtPr.Tag.ValAttr, content: content}
	}
	return nil
}

func (c Cell) SetContentText(text string) error {
	for _, ble := range c.x.EG_BlockLevelElts {
		for _, cbc := range ble.EG_ContentBlockContent {
			for _, p := range cbc.P {
				for _, pc := range p.EG_PContent {
					for _, crc := range pc.EG_ContentRunContent {
						if crc.R == nil {
							continue
						}
						for _, ric := range crc.R.EG_RunInnerContent {
							ric.T.Content = text
							break
						}
					}
				}
			}
		}
	}
	return nil
}
