package generate

import (
	"fmt"

	"Athena/utils"
)

type encode struct{}

type generate struct {
	FileName string
	Data     *utils.Package

	encoder *encode
}

// Generate new generate
func Generate(data *utils.Package) *generate {
	return &generate{
		FileName: fmt.Sprintf("%s.md", data.PackageName),
		Data:     data,
	}
}
