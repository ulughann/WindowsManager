package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func (a *App) CheckChoco() string {
	return checkChoco()
}

func (a *App) InstallApp(app string) string {
	cmd := exec.Command("powershell", "-Command", "choco install -y "+app)

	// Hide the PowerShell window:
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing Chocolatey installation script:", err)
		fmt.Println("Output:", string(output))
		return "false"
	}

	fmt.Println("Chocolatey installation completed successfully.")
	return checkChoco()
}

func checkChoco() string {
	cmd := exec.Command("powershell", "-Command", "choco --version")

	// Hide the PowerShell window:
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing Chocolatey version check:", err)
		fmt.Println("Output:", string(output))
		return "Error"
	}
	return string(output)
}

func InstallChoco() string {
	if !isAdmin() {
		fmt.Println("Administrative privileges are required to install Chocolatey.")
		fmt.Println("Please run this program as an administrator.")
		return "false"
	}

	cmd := exec.Command("powershell", "-windowstyle hidden", "-Command", `
        Set-ExecutionPolicy Bypass -Scope Process -Force;
        [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072;
        iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))`)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing Chocolatey installation script:", err)
		fmt.Println("Output:", string(output))
		return "false"
	}

	fmt.Println("Chocolatey installation completed successfully.")
	return checkChoco()
}

func isAdmin() bool {
	cmd := exec.Command("powershell", "-Command", "$currentPrincipal = New-Object Security.Principal.WindowsPrincipal([Security.Principal.WindowsIdentity]::GetCurrent()); $currentPrincipal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error checking administrative privileges:", err)
		return false
	}
	return output[0] == "True"[0]
}
