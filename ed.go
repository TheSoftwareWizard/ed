package main
import (
	"fmt"
	"flag"
	"os"
	"os/exec"
	"bufio"
	"strings"
)

var is_system, is_verbose bool

func main(){

	// Define usage preamble
	flag.Usage = func(){
		fmt.Fprintln(os.Stderr, "Ed: A meta editor compiler")
		fmt.Fprintln(os.Stderr, "Usage: specifiy a editing space (file location) to edit as the very last argument")
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}
	
	// Flags
	// // System
	flag.BoolVar(&is_system, "s", false, "edit system config")
	flag.BoolVar(&is_system, "system", false, "edit system config")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0{
		fmt.Printf("ed: error: no flags or editing space specified, see '--help'\n")
		os.Exit(1)
	}
	editing_space := args[0]

	switch editing_space{
	// check if flag is specified, but not editing space, then return more specific error, (c-1d)
	// only for flags that modify editing space (c+1u)
		case "-s", "--s", "-system", "--system":
			fmt.Printf("ed: error: -s --system: no specified system directory, see '--help system'\n")
			os.Exit(1)
	}

	if is_system {
		editing_space = find_system_file(editing_space)
		edit(true, editing_space)

	} else {
		edit(false, editing_space) 
	}
}


// open editor with or without sudo in an editing space
func edit(doas_root bool, editing_space string){
	// note: doas_root is not related to 'doas' but is just literally 'do as'
	var cmd_name string // define sub-scope changed variables on func scope (r-1d)
	var cmd *exec.Cmd // (r+1u)

	editor := strings.Fields(grab_env("EDITOR", ask_for{do_ask: true, prompt: "ed: $EDITOR is not set, please enter what editor to use (e.g. 'nano', 'vi', etc): "}))

	if doas_root {
		var cmd_args []string
		sudoer := strings.Fields(grab_env("SUDOER", ask_for{do_ask: true, prompt: "ed: $SUDOER is not set, please enter what sudoer command to use (e.g. 'sudo', 'doas'): "}))

		cmd_args = append(cmd_args, sudoer[1:]...)
		cmd_args = append(cmd_args, editor...)
		cmd_args = append(cmd_args, editing_space)

		cmd_name = strings.Join(sudoer, " ") + " " + strings.Join(editor, " ") + " " + editing_space
		cmd = exec.Command(sudoer[0], cmd_args...)

	} else {
		cmd_name = strings.Join(editor, " ") + " " + editing_space
		cmd_args := append(editor[1:], editing_space)

		cmd = exec.Command(editor[0], cmd_args...)
	}

	// set editor context to program context
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run() // run cmd
	if err != nil {
		fmt.Printf("ed: error running editor command '%v': %v\n", cmd_name, err)
	}
	os.Exit(0)
}

// find a file in system config
func find_system_file(keyword string) string {
	system_config_path := grab_env("SYSCONF", ask_for{do_ask: true, prompt: "ed: $SYSCONF is not set, please enter your system configuration path (e.g. '/etc/nixos/'): "})
	if !strings.HasSuffix(system_config_path, "/"){ // if no trailing slash, add it
		system_config_path += "/"
	}
	
	entries, err := os.ReadDir(system_config_path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ed: error reading system config directory: ", err)
		os.Exit(1)
	}
	for _, e := range entries {
		if e.IsDir(){
			continue // implement subdir searching and editing one day
		}
		entry_name := e.Name()
		if strings.HasPrefix(entry_name, keyword) {
			return system_config_path + entry_name
		}
	}

	fmt.Printf("ed: error: File starting with '%s' does not exist in '%s'.\n", keyword, system_config_path)
	os.Exit(1)
	return "f8f06a50-8abf-4e0f-ab5d-664befbfdb0f" // this is impossible
}

// a struct containing info about whether you should prompt and what the prompt should be if yes
type ask_for struct {
	do_ask bool
	prompt string
}

// grab enviroment variable contents util function
func grab_env(env_name string, predef_ask_input ...ask_for) (env string) {
	// establish default for ask_input, which controls whether or not to prompt if env is empty
	ask_input := ask_for{ do_ask: false, prompt: "", } // set the ask_input default to false, no prompt
	if len(predef_ask_input) > 0 {
		ask_input = predef_ask_input[0]
	}

	env = os.Getenv(env_name)
	if ask_input.do_ask && env == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(ask_input.prompt)
		input, _ := reader.ReadString('\n')
		env = strings.TrimSpace(input)
	}
	return env
}
