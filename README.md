# Cha - An unofficial CLI for the Change Platform

API Reference: https://docs.getchange.io

## Usage
```
Usage:
  cha [command]

Available Commands:
  climate     Draft or create Climate offsets
  donations   Find or create Donations
  help        Help about any command
  nonprofits  List or search Nonprofits

Flags:
  -h, --help   help for cha

Use "cha [command] --help" for more information about a command.
```

Note: You can set your `GOBIN` path and `go install` to make `cha` a global executable on your system.

## Example
Be sure to export `CHA_PUBLIC_KEY` and `CHA_SECRET_KEY` in order to authenticate requests. You can find these in your Change dashboard (https://api.getchange.io/).
```
export CHA_PUBLIC_KEY=YOUR_KEY_HERE
export CHA_SECRET_KEY=YOUR_KEY_HERE
```

Build the binary
```
go build
```

Search for `ALASKA` non-profit on page 2. Use ` ./cha nonprofits search --help` to see all available flags.
```
./cha nonprofits search -n ALASKA -p 2
```

## TODO
- Move flag vars to struct/map
- Error handling
- Interactive CLI (auto-populate ID's) shell using https://github.com/manifoldco/promptui
