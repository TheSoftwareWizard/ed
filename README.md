# ed
a tool i build to edit my nixos config (there's probably a better way)

#### Readme version 2.0 | latest is 2.0

### install
1. download the latest release binary for your architecture
2. add the binary to your PATH or a PATH location

### usage
ed is a very simple application

#### use case 1
to edit system config do:
```
ed -s <KEYWORD>
```
in this case ed will search through /etc/nixos/ for any files starting with your ```<KEYWORD>``` and open them. <br/>
your sudoer application will prompt you for your password to edit the file it found.

example:
```
ed -s gra
```
This command will open my editor to any file starting with 'gra' in my SYSCONF path or '/etc/nixos/', finding my 'graphics.nix' file

#### use case 2
to edit a file do:
```
ed file_name.txt
```
this will open your editor to that file relative to the location where the command was run

### enviroment variables
using this application will be much easier if 3 specific enviroment variables are set:<br/>
'EDITOR', 'SUDOER', and 'SYSCONF'<br/>
set EDITOR to your editor command, e.g. nano, vi, etc<br/>
set SUDOER to your sudoer command, e.g. sudo, doas, etc <br/>
set SYSCONF to your system configuration directory, e.g. '/etc/nixos/, etc

if these enviroment variables are not set it will prompt you for them every time you run ed

### flags
'-s' allows you to edit your system config, long-form alt is '--system'<br/>

### compiling from source
to compile from source clone this repo then build 'ed.go' with ```go build ed.go```<br/>

### to do
- [ ] Add multi system configuration path support
- [ ] Add searching of system configuration path subdirectories for system configuration files

## License

[![License: GPL v3](https://shields.io)](https://gnu.org)

This project is licensed under the GPLv3 License. See the [LICENSE](LICENSE) file for details.

Copyright (C) 2026 @TheSoftwareWizard
