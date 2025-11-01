package common

import "archive/zip"

type Package struct {
	Path string
}

type PackageRef struct {
	d             *DocBase
	relationships Relationships
	pkg           Package
	relID         string
}

func PackageFromFile(file *zip.File) Package {
	return Package{Path: file.Name}
}

func NewPackageRef(d *DocBase, relationships Relationships, pkg Package) PackageRef {
	return PackageRef{
		d:             d,
		relationships: relationships,
		pkg:           pkg,
	}
}
