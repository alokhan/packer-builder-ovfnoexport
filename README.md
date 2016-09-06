How to compile:

- godep restore 
- remove vendor directory from packer dependency (../packer)
- godep restore inside packer directory
- go build -o /path/to/file main.go (the path is usually ~/.packer.d/plugins as it is the packer plugins directory path)
