package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// ProcessKiller 인터페이스는 프로세스를 종료하는 기능을 정의합니다.
type ProcessKiller interface {
	KillProcess(port string) error
}

// UnixProcessKiller UnixProcessKiller Unix 계열 운영체제에서 프로세스를 종료합니다.
type UnixProcessKiller struct{}

func (killer UnixProcessKiller) KillProcess(port string) error {
	command := fmt.Sprintf("lsof -t -i:%s", port)
	fmt.Println(command)
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`\d+`)
	pidStr := re.FindString(string(output))
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return err
	}

	cmd = exec.Command("kill", "-9", fmt.Sprintf("%d", pid))
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("Killed process with PID %d\n", pid)
	return nil
}

// WindowsProcessKiller WindowsProcessKiller Windows 운영체제에서 프로세스를 종료합니다.
type WindowsProcessKiller struct{}

func (killer WindowsProcessKiller) KillProcess(port string) error {
	// Windows에서 프로세스 종료 로직 구현
	return nil
}

func main() {
	port := input()
	var killer ProcessKiller

	// 운영체제에 따라 적절한 ProcessKiller 구현체 선택
	if isUnix() {
		killer = UnixProcessKiller{}
	} else {
		killer = WindowsProcessKiller{}
	}

	err := killer.KillProcess(port)
	if err != nil {
		fmt.Printf("Failed to kill process: %v\n", err)
	}
}

// isUnix 함수는 현재 운영체제가 Unix 계열인지 확인합니다.
func isUnix() bool {
	// 운영체제 판별 로직 구현
	osName := os.Getenv("OS")
	// os name contains "Windows" if the OS is Windows
	if strings.Contains(osName, "Windows") {
		return false
	}
	return true
}

func input() string {
	// 입력 받기
	fmt.Print("Enter port for kill process: ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return ""
	}
	return input
}
