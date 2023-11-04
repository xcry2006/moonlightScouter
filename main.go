package main

import (
	"fmt"
	"moonlightScouter/DetectArch"
	"moonlightScouter/DetectOS"
	"moonlightScouter/DetectPL"
	"moonlightScouter/DetectTool"
)

func main() {
	fmt.Println("The operating system is", DetectOS.SysOS)
	fmt.Println("The system arch is", DetectArch.SysArch)
	fmt.Println("Found tools")
	DetectTool.DetectTool()
	fmt.Println("Found programming language")
	DetectPL.DetectPL()
}
