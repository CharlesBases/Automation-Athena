package main

import (
	"fmt"
	"os"

	"Athena/generate"
	"Athena/utils"
)

func main() {
	utils.Plugin()
	run()
	utils.WriteResponse(os.Stdout)
}

// run generate doc
func run() {
	for _, x := range utils.Information.Data {
		utils.Push(fmt.Sprintf("%s.md", x.PackageName), generate.Generate(x).Markdown())
	}
}
