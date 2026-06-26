package main
import (
	"fmt"
	"flag"
	"os"
	"os/exec"
	"bufio"
	"strings"
	"time"
)

var isSystem, isVerbose bool

func main(){

	// Define usage preamble
	flag.Usage = func(){
		fmt.Fprintln(os.Stderr, "Ed: A meta editor compiler")
		fmt.Fprintln(os.Stderr, "Usage: specifiy a location to edit as last argument")
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}
	
	// System flags
	flag.BoolVar(&isSystem, "s", false, "edit system config")
	flag.BoolVar(&isSystem, "system", false, "edit system config")

	flag.BoolVar(&isVerbose, "v", false, "verbose output")
	flag.BoolVar(&isVerbose, "verbose", false, "verbose output")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0{
		fmt.Printf("ed: error: no flags or location specified, see '--help'\n")
		os.Exit(1)
	}
	location := args[0]

	switch location{
		case "-s", "--s", "-system", "--system":
			fmt.Printf("ed: error: -s --system: no specified system directory, see '--help system'\n")
			os.Exit(1)
	}

	// If isSystem 
	if isSystem {
		location = "/etc/nixos/" + getSysFile(location)
		edit(true, location)
	} else {
		edit(false, location) // if location is not a flag, perform edit() on it		
	}

	// Check location
}

func edit(sudoer bool, location string){
	// If we made it here, everything is fine
	var editor, sudoer_cmd string
	var cmd *exec.Cmd

	editor = os.Getenv("EDITOR")
	if editor == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("No editor enviroment variable set (see 'man ed env_editor') please enter one (ex: 'nano'): ")
		input, _ := reader.ReadString('\n')
		editor = strings.TrimSpace(input)
	}

	if sudoer {
		sudoer_cmd = os.Getenv("SUDOER_CMD")
		if sudoer_cmd == "" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("No sudoer enviroment variable set (see 'man ed env_sudoer-cmd') please enter one (ex: 'sudo'): ")
			input, _ := reader.ReadString('\n')
			sudoer_cmd = strings.TrimSpace(input)
		
		}
		if isVerbose { fmt.Printf("Command: %s %s %s\n", sudoer_cmd, editor, location) }
		cmd = exec.Command(sudoer_cmd, editor, location)

	} else {
		if isVerbose { fmt.Printf("Command: %s %s\n", editor, location) }
		cmd = exec.Command(editor, location)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if isVerbose { time.Sleep(1500 * time.Millisecond) }
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}
	os.Exit(0)
}

func getSysFile(keyword string) string {
	sysPath := "/etc/nixos"
	entries, err := os.ReadDir(sysPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ed: error: ", err)
		os.Exit(1)
	}
	for _, e := range entries {
		if e.IsDir(){
			continue
		}
		entryName := e.Name()
		if strings.HasPrefix(entryName, keyword) {
			return entryName
		}
	}

	fmt.Printf("ed: error: File starting with '%s' does not exist in '/etc/nixos/'.\n", keyword)
	os.Exit(1)
	return "ed: Fatal unknown error at getSysFile() stage, this shouldn't happen...?"
}
