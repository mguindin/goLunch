goLunch
=======

Go version of nodeLunch

Made using codegangsta's [cli library](https://github.com/codegangsta/cli)
```bash
USAGE:
   lunch [global options] command [command options] [arguments...]

VERSION:
   0.1

COMMANDS:
   search, s    Search for lunch
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --generate-bash-completion
   --version, -v                print the version
   --help, -h                   show help
```
   
Example usage:
```bash
./lunch s --location 10021 --cuisine chinese --choice 2 --radius 0.1
You will be having Cafe Evergreen, which is located at 1367 1st Ave.
Cafe Evergreen has a rating of 3.5
People are saying: This review is only for the dine-in experience, not take out or delivery. Cafe Evergreen has easily the best dim sum of any place I've eaten at in...
```
