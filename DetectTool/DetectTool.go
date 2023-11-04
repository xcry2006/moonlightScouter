package DetectTool

import (
	"fmt"
	"os/exec"
)

func DetectTool() {
	DetectCurl()
	DetectFTP()
	DetectLrzsz()
	DetectRsync()
	DetectWget()
}

func DetectWget() {
	cmd := exec.Command("wget", "-h")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found wget")
	} else {
		fmt.Println("Found wget")
	}
}

func DetectCurl() {
	cmd := exec.Command("curl", "-h")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found curl")
	} else {
		fmt.Println("Found curl")
	}
}

func DetectFTP() {
	cmd := exec.Command("ftp")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found ftp")
	} else {
		fmt.Println("Found ftp")
	}
}

/*func DetectSCP() {
	cmd := exec.Command("scp")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found scp")
	} else {
		fmt.Println("Found scp")
	}
}

/*func DetectSFTP() {
	cmd := exec.Command("sftp", "-h")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found sftp")
	} else {
		fmt.Println("Found sftp")
	}
}*/

func DetectRsync() {
	cmd := exec.Command("rsync", "--help")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Not found rsync")
	} else {
		fmt.Println("Found rsync")
	}
}

func DetectLrzsz() {
	cmd := exec.Command("rz", "-h")
	err := cmd.Run()
	if err != nil {
		cmd := exec.Command("sz", "-h")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Not found lrzsz")
		} else {
			fmt.Println("Foundn lrzsz")
		}
	} else {
		fmt.Println("Found lrzsz")
	}
}
