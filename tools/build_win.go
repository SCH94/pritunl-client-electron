package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	err := os.Chdir("tuntap_win")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("go", "build", "-v", "-a", "-o tuntap.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	err = os.Chdir(filepath.Join("..", "service"))
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("go", "get", "-u", "-f")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("go", "build", "-v", "-a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	err = os.Chdir(filepath.Join("..", "client"))
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("npm", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command(".\\node_modules\\.bin\\electron-rebuild")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command(".\\node_modules\\.bin\\electron-rebuild",
		".\\",
		"pritunl",
		"--platform=win32",
		"--arch=ia32",
		"--version=0.28.3",
		"--icon=www\\img\\logo.ico",
		"--out=..\\build\\win",
		"--prune",
		"--asar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
