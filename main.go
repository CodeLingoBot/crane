package main

import (
	"fmt"
	"os"
	"os/exec"
)

var verbose bool

func main() {
	if len(os.Args) == 1 {
		displayHelp()
	} else {
		// On panic, recover the error and display it
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("ERROR: ", err)
			}
		}()
		// Set verbosity
		setVerbosity()
		// Read Cranefile
		container := readCranefile("Cranefile")
		// Run subcommand
		switch os.Args[1] {
		case "provision":
			fmt.Println("Provisioning...")
			container.provision()
			break

		case "run":
			fmt.Println("Running...")
			// "Entry" container is attachable
			container.run(true)
			break

		case "rm":
			fmt.Println("Removing...")
			container.rm()
			break

		case "kill":
			fmt.Println("Killing...")
			container.kill()
			break

		case "start":
			fmt.Println("Starting...")
			container.start()
			break

		case "stop":
			fmt.Println("Stopping...")
			container.stop()
			break

		case "help":
			displayHelp()
			break

		default:
			fmt.Println("Command not found. See available commands with `crane help`.")
			break
		}
	}
}

func executeCommand(name string, args []string) {
	if verbose {
		fmt.Printf("%v\n", args)
	}
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
	if !cmd.ProcessState.Success() {
		panic(cmd.ProcessState.String()) // pass the error?
	}
}

func setVerbosity() {
	if len(os.Args) == 3 {
		if os.Args[2] == "--verbose" {
			verbose = true
		}
	}
}

func displayHelp() {
	fmt.Println("crane - Lift containers with ease")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tcrane command")
	fmt.Println("")
	fmt.Println("The commands are:")
	fmt.Println("")
	fmt.Println("\tprovision\tBuild or pull containers")
	fmt.Println("\trun\t\tRun containers (linking them)")
	fmt.Println("\tkill\t\tKill containers")
	fmt.Println("\trm\t\tRemove containers")
	fmt.Println("\tstart\t\tStart containers")
	fmt.Println("\tstop\t\tStop containers")
	fmt.Println("")
	fmt.Println("See the docker commands for more information.")
	fmt.Println("")
}
