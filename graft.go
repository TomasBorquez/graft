package main

import (
	"fmt"
	"github.com/TomasBorquez/graft/pkg"
)

func main() {
	graft.Config("build", func (p *graft.Project) {
		fmt.Println("Executed build")
	})
	
	graft.Config("start", func (p *graft.Project) {
		fmt.Println("Executed start")
	})
}
