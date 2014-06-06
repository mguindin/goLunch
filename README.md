goLunch
=======
![Travis build status](https://travis-ci.org/mguindin/goLunch.svg?branch=master)
Go version of nodeLunch
You need to have a yelp [API key](http://www.yelp.com/developers/getting_started) in a file called ```yelp_key```
in the ```bin``` directory with the ```lunch``` binary.
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
   
./lunch s -h
NAME:
   search - Search for lunch

USAGE:
   command search [command options] [arguments...]

DESCRIPTION:


OPTIONS:
   --debug              Output request URL
   --cuisine 'random'   cuisine to search for
   --radius '0.5'       Radius for search
   --location '10021'   Location to search from
   --choice '1'         Choice in selection
```
   
Example usage:
```bash
./lunch s --location 10021 --cuisine chinese --choice 2 --radius 0.1
You will be having Cafe Evergreen, which is located at 1367 1st Ave.
Cafe Evergreen has a rating of 3.5
People are saying: This review is only for the dine-in experience, not take out or delivery. Cafe Evergreen has easily the best dim sum of any place I've eaten at in...
```
