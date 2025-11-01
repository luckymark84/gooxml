package common

type OLEObject struct {
	Path string
}

type OLEObjectRef struct {
	d         *DocBase
	rels      Relationships
	oleobject OLEObject
	relID     string
}

func NewOLEObjectRef(d *DocBase, r Relationships, ole OLEObject) OLEObjectRef {
	return OLEObjectRef{
		d:         d,
		rels:      r,
		oleobject: ole,
	}
}
