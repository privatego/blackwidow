package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Get Project Root Directory
func GetBaseDirectory() string {
	file, _ := exec.LookPath(os.Args[0])
	fpath, _ := filepath.Abs(file)
	name := filepath.Base(file)
	appDir := strings.TrimRight(fpath, name) + "bin/"
	fmt.Println("[BaseDir]", appDir)
	return appDir
}

func InitializeWidow() {
	now := time.Now()
	logfileName := fmt.Sprintf("f%d%d%d_%d%d%d.log",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	setupLogFile(logfileName)
}

// Create log file
func setupLogFile(fileName string) {
	appDir := GetBaseDirectory()
	logFile, err := os.OpenFile(appDir+fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend|0666)
	if err == nil {
		log.SetOutput(logFile)
		log.Println("\n\n\n")
		log.SetFlags(log.Flags() | log.Lshortfile)
		log.Println("[Info] Server Started ! ")
	}
}
