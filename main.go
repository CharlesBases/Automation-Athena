package main

import (
	"fmt"
	"os"

	"github.com/CharlesBases/protoc-gen-doc/generate"
	"github.com/CharlesBases/protoc-gen-doc/utils"
)

func main() {
	utils.Plugin()
	run()
	utils.WriteResponse(os.Stdout)
}

// run generate doc
func run() {
	for _, x := range utils.Information.Data {
		markdown := fmt.Sprintf("%s.md", x.PackageName)
		utils.Push(markdown, generate.New(markdown, x).Generate(generate.Markdown))

		html := fmt.Sprintf("%s.html", x.PackageName)
		utils.Push(html, generate.New(html, x).Generate(generate.HTML))
	}
}
