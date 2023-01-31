package runnablereadme

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func ExecuteMarkdown(filePath string) (result bool) {
	var commands []string
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error")
	}
	readme := string(content)
	formattedString := strings.Replace(readme, "```console", "CONSOLE_CODE_HERE", -1)
	newFormattedString := strings.Replace(formattedString, "\n", " ", -1)
	finalString := strings.Replace(newFormattedString, "```", "CONSOLE_CODE_HERE", -1)
	var re = regexp.MustCompile(`CONSOLE_CODE_HERE\s*(.*?)\s*CONSOLE_CODE_HERE`)
	for _, match := range re.FindAllString(finalString, -1) {
		finalCmd := strings.Replace(match, "CONSOLE_CODE_HERE", "", -1)
		cmd := strings.TrimSpace(finalCmd)
		commands = append(commands, cmd)
	}
	for i, cmd := range commands {
		fmt.Println("\nExecuting command number ", i+1)
		//s := strings.Split("cmd", " ")
		command := strings.Replace(cmd, "\\", " ", -1)
		out, err := exec.Command("bash", "-c", command).Output()
		if err != nil {
			fmt.Println("Error:", err)
			return false
		}
		fmt.Printf("Output: %s", out)
	}
	return true
}
