package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Unzip(filepath, dest string) {
	cmd := exec.Command("unzip", filepath, "-d", dest)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err, string(out))
		os.Exit(1)
	}
}

func Make(filepath string) {
	cmd := exec.Command("make", "-C", filepath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func MakeClean(filepath string) {
	cmd := exec.Command("make", "clean", "-C", filepath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Cp(from, to string) {
	cmd := exec.Command("cp", from, to)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Cd(filepath string) {
	err := os.Chdir(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Mv(from, to string) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("mv %s %s; 2>/dev/null", from, to))
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err, ":", string(out))
		os.Exit(1)
	}

	//fmt.Println(out)
}

func RunBashScript(filepath string, timeout int) string {
	cmd := exec.Command("bash", filepath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "Failure"
	}

	var timer *time.Timer
	timer = time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		timer.Stop()
		cmd.Process.Kill()
	})

	return string(out)
}

func FindFolders(filepath string) []string {
	cmd := exec.Command("find", filepath, "-maxdepth", "1", "!", "-path", filepath, "-type", "d")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ret []string
	if string(out) == "" {
		return ret
	}

	return strings.Split(string(out), "\n")
}

func FindFileType(filepath, ftype, to string) []string {
	cmd := exec.Command("/usr/bin/find", filepath, "-name", ftype, "-exec", "cat", "{}", "\\;")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err, " : ", string(out))
		os.Exit(1)
	}

	var ret []string
	if string(out) == "" {
		return ret
	}

	return strings.Split(string(out), "\n")
}

func Rm(filepath string) {
	cmd := exec.Command("rm", "-rf", filepath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Cat(filepath string) string {
	cmd := exec.Command("cat", filepath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(out)
}

func Head(filepath string, numlines int) string {
	cmd := exec.Command("head", fmt.Sprintf("-%d", numlines), filepath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(out)
}
