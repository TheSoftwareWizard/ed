# ed
a tool i build to edit my nixos config (there's probably a better way)

### requirements
- linux kernel
- any sudoer application, e.g. sudo, doas, etc
- any editor application, e.g. nano, micro, vi, emacs, vim, neovim, helix, etc

it is highly reccomended to use NixOS however it techincally functions without it, although it would function as no more than an alias. '-s' and '--system' will NOT work on any other distro

### install
1. download the latest release binary
2. (optional) add the binary to your PATH or a PATH location

### usage
ed is a very simple application

#### use case 1
to edit system config do:
```
ed -s <KEYWORD>
```
in this case ed will search through /etc/nixos/ for any files starting with your <KEYWORD> and open them.
it will prompt you for you password running through your sudoer application

example:
```
ed -s gra
```
This command will open my editor to any file starting with 'gra' in '/etc/nixos/', in this case 'graphics.nix'

#### use case 2
to edit a file do:
```
ed file_name.txt
```

### enviroment variables
using this application will be much easier if 2 specific enviroment variables are set:<br/>
EDITOR and SUDOER_CMD<br/>
set EDITOR to your editor command, e.g. nano, micro, vim, nvim, hx, etc<br/>
set SUDOER_CMD to your sudoer command, e.g. sudo, doas, etc

if these enviroment variables are not set it will prompt you for them every time you run ed

### flags
'-s' allows you to edit your system config, long-form alt is '--system'<br/>
'-v' enables verbose output, long-form alt is '--verbose'

### compiling from source
to compile from source clone this repo then build 'ed.go' with ```go build ed.go```<br/>
it should be quick because it's a kb sized file

### to do
- expand verbose output
- make a man entry | instructions reffering to the man entry exist in the code, should quickly resolve the lack of one




