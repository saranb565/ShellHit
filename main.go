package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	genai "github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"golang.org/x/sys/windows"
	"google.golang.org/api/option"
)

// Load API key from .env
func loadAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("API_KEY")
}

var apiKey string

func init() {
	apiKey = loadAPIKey()
}

var (
	green   = color.New(color.FgGreen).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	red     = color.New(color.FgRed).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
)

func isAdmin() bool {
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid,
	)
	if err != nil {
		return false
	}
	token := windows.GetCurrentProcessToken()
	isAdmin, _ := token.IsMember(sid)
	return isAdmin
}

func runAsAdmin() {
	verb := "runas"
	exe, err := os.Executable()
	if err != nil {
		fmt.Println(red("Failed to get executable path:"), err)
		return
	}
	cmd := exec.Command("cmd", "/c", "start", verb, exe)
	err = cmd.Run()
	if err != nil {
		fmt.Println(red("Failed to restart with admin privileges:"), err)
	}
	os.Exit(0)
}

var dir string

func main() {
	if !isAdmin() {
		fmt.Println(yellow("🔹 Requesting admin privileges..."))
		runAsAdmin()
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(magenta("🎯 Welcome to Saran's AI ShellHit Terminal 🎯"))
	dir, _ = os.Getwd()

	for {
		fmt.Printf(blue("%s> "), dir)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, red("Error reading input:", err))
			continue
		}

		err = processInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, red("Error:"), err)
		}
	}
}

func processInput(input string) error {
	input = strings.TrimSpace(input)
	args := strings.Fields(input)

	if len(args) == 0 {
		return nil
	}

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("⚠️ " + yellow("Path required"))
		}
		path := strings.Join(args[1:], " ")
		err := os.Chdir(path)
		if err != nil {
			return fmt.Errorf(red("cd: %w"), err)
		}
		dir, _ = os.Getwd()
		return nil

	case "pwd":
		fmt.Println(green(dir))
		return nil

	case "ls":
		files, err := os.ReadDir(dir)
		if err != nil {
			return fmt.Errorf(red("ls: %w"), err)
		}
		for _, file := range files {
			fmt.Println(cyan(file.Name()))
		}
		return nil

	case "mkdir":
		if len(args) < 2 {
			return errors.New(yellow("⚠️ Path required"))
		}
		path := strings.Join(args[1:], " ")
		err := os.Mkdir(filepath.Join(dir, path), 0755)
		if err != nil {
			return fmt.Errorf(red("mkdir: %w"), err)
		}
		return nil

	case "rm":
		if len(args) < 2 {
			return errors.New(yellow("⚠️ Path required"))
		}
		path := strings.Join(args[1:], " ")
		err := os.Remove(filepath.Join(dir, path))
		if err != nil {
			return fmt.Errorf(red("rm: %w"), err)
		}
		return nil

	case "exit":
		fmt.Println(magenta("👋 Exiting AI ShellHit Terminal..."))
		os.Exit(0)
	}

	//AI
	command, err := getCommandFromAI(input)
	if err != nil {
		fmt.Println(red("AI: Not possible"))
		return nil
	}

	if command == "0" {
		fmt.Println(yellow("⚠️ AI: Sorry, I can't process that request."))
		return nil
	}

	if strings.HasPrefix(command, "cd ") {
		path := strings.TrimPrefix(command, "cd ")
		err := os.Chdir(path)
		if err != nil {
			fmt.Println(red("AI: Failed to change directory:"), err)
			return nil
		}
		dir, _ = os.Getwd()
		return nil
	}

	fmt.Println(green("🛠️ AI Executing:"), cyan(command))
	cmd := exec.Command("cmd", "/C", command)
	cmd.Dir = dir
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func getCommandFromAI(input string) (string, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Printf(red("Failed to create client: %v"), err)
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	prompt := fmt.Sprintf("Convert this user request into a Windows command. Give only the command and no other text. If such a Windows command doesn't exist just give 0: %s", input)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf(red("API request failed: %w"), err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New(red("No valid command found in AI response"))
	}

	part, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return "", errors.New(red("Unexpected response format from AI"))
	}

	command := strings.TrimSpace(string(part))
	command = strings.TrimPrefix(command, "```batch")
	command = strings.TrimPrefix(command, "```")
	command = strings.TrimSuffix(command, "```")
	command = strings.ReplaceAll(command, "`", "")

	return strings.TrimSpace(command), nil
}
