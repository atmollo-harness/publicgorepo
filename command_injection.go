// for testing go vulnerabilities
// https://semgrep.dev/docs/cheat-sheets/go-command-injection#1-running-an-os-command
package main

import (
	"context"
	"os"
	"os/exec"
)

func osCommand1A() {
	ctx := context.Background()
	// Command example
	exec.Command("echo", "hello")

	// CommandContext example
	exec.CommandContext(ctx, "sleep", "5").Run()

	// Not vulnerable
	exec.Command("echo", "1; cat /etc/passwd")

	// Vunerable
	userInput := "echo 1 | cat /etc/passwd" // value supplied by user input
	_, _ = exec.Command("sh", "-c", userInput).Output()

	// Vulnerable
	userInput1 := "cat"         // value supplied by user input
	userInput2 := "/etc/passwd" // value supplied by user input
	_, _ = exec.Command(userInput1, userInput2).Output()
}

func osCommand1B() {
	cmd := &exec.Cmd{
		// Path is the path of the command to run.
		Path: "/bin/echo",
		// Args holds command line arguments, including the command itself as Args[0].
		Args:   []string{"echo", "hello world"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	cmd.Start()
	cmd.Wait()

	// Args can be also ommited and {Path} will be used by default
	cmd = &exec.Cmd{
		Path: "/bin/echo",
	}

	// Vulnerable
	userInput := "/pwn/exploit" // value supplied by user input
	cmd = &exec.Cmd{
		Path: userInput,
	}

	// Vulnerable
	userInput1 := "/bin/bash"                    // value supplied by user input
	userInput2 := []string{"bash", "exploit.sh"} // value supplied by user input
	cmd = &exec.Cmd{
		Path: userInput1,
		Args: userInput2,
	}
}

// func osCommand1C() {
// 	cmd := exec.Command("bash")
// 	// StdinPipe initialization
// 	cmdWriter, _ := cmd.StdinPipe()
// 	cmd.Start()
// 	// Vulnerability when `password` controlled by user input
// 	cmdInput := fmt.Sprintf("sshpass -p %s", password)
// 	// Writing to StdinPipe
// 	cmdWriter.Write([]byte(cmdInput + "\n"))
// 	cmd.Wait()
// }

// func osCommand1D() {
// 	// Exec invokes the execve(2) system call.
// 	syscall.Exec(binary, args, env)
// 	// ForkExec - combination of fork and exec, careful to be thread safe.
// 	syscall.ForkExec(binary, args, env)
// 	vulnerableCode()
// }

// func vulnerableCode(userInput string) {
// 	//  Vulnerable: Do not let `path` be defined by user input
// 	path, _ := exec.LookPath(userInput)
// 	args := []string{"ls", "-a", "-l", "-h"}
// 	env := os.Environ()
// 	execErr := syscall.Exec(path, args, env)
// }
