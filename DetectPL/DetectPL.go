package DetectPL

import (
	"fmt"
	"os/exec"
)

func DetectPL() {
	DetectGo()
	DetectJava()
	DetectPerl()
	DetectPHP()
	DetectPython()
	DetectRuby()
	DetectRust()
}

func DetectPython() {
	cmd := exec.Command("python3")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found python3")
	} else {
		fmt.Println("Found python3")
		cmd := exec.Command("python2")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Not found python2")
		} else {
			fmt.Println("Found python2")
			cmd := exec.Command("python")
			err := cmd.Run()
			if err != nil {
				fmt.Println("Not found python")
			} else {
				fmt.Println("Found python")
			}
		}
	}
}

func DetectPerl() {
	cmd := exec.Command("perl", "-h")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found perl")
	} else {
		fmt.Println(("Found perl"))
	}
}

func DetectPHP() {
	cmd := exec.Command("php", "-h")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found php")
	} else {
		fmt.Println("Found php")
	}
}

func DetectJava() {
	cmd := exec.Command("java")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found java")
	} else {
		fmt.Println("Found java")
	}
}

func DetectRuby() {
	cmd := exec.Command("ruby", "-h")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found ruby")
	} else {
		fmt.Println("Found ruby")
	}
}

func DetectGo() {
	cmd := exec.Command("go", "help")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found go")
	} else {
		fmt.Println("Found go")
	}
}

func DetectRust() {
	cmd := exec.Command("rustc")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found rust")
	} else {
		fmt.Println("Found rust")
	}
}
