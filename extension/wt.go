package extension

import (
	"flag"
	"os"
	"os/exec"

	"github.com/Meonako/Aniko/config"

	"github.com/Meonako/go-logger/v2"
)

var (
	skipCheck = flag.Bool("skip-check", false, "Skip check & open Windows Terminal")
)

const (
	PREFIX string = "Error launching windows terminal: "

	RUN   int8 = 0
	BUILD int8 = 1
)

func init() {
	flag.Parse()
}

func init() {
	if !config.Config().OPEN_IN_WINDOWS_TERMINAL || *skipCheck {
		return
	}

	if _, err := os.Stat("Aniko.exe"); os.IsNotExist(err) {
		switch config.Config().IF_NOT_FOUND_BINARY {
		case RUN:
			justRun()
		case BUILD:
			buildAndRun()
		}
		return
	}

	run()
}

func buildAndRun() {
	err := exec.Command("cmd", "/C", "go", "build", ".").Run()
	if err != nil {
		logger.ToTerminalRed(PREFIX, "Can't build binary.")
		return
	}

	run()
}

func run() {
	dir, err := os.Getwd()
	if err != nil {
		logger.ToTerminalRed(PREFIX, err)
		return
	}

	err = exec.Command("cmd", "/C", "wt", "-w", "0", "new-tab", "-p", "Power Shell", "-d", dir, ".\\Aniko", "-skip-check").Run()
	if err != nil {
		logger.ToTerminalRed(PREFIX, err)
		return
	}
	os.Exit(0)
}

func justRun() {
	dir, err := os.Getwd()
	if err != nil {
		logger.ToTerminalRed(PREFIX, err)
		return
	}

	err = exec.Command("cmd", "/C", "wt", "-w", "0", "new-tab", "-p", "Power Shell", "-d", dir, "go", "run", ".", "-skip-check").Run()
	if err != nil {
		logger.ToTerminalRed(PREFIX, err)
		return
	}
	os.Exit(0)
}

// NOT USABLE
//
// func openWT(command ...string) {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		logger.ToTerminalRed(PREFIX, err)
// 		return
// 	}

// 	err = exec.Command("cmd", "/C", "wt", "-w", "0", "new-tab", "-p", "Power Shell", "-d", dir, command...).Run()
// 	if err != nil {
// 		logger.ToTerminalRed(PREFIX, err)
// 		return
// 	}
// }
