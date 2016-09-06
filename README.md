This Packer builder import a ovf into virtual box but does not delete and export the vm.
The goal is to speed up the process because we I did not need a intermediate image.

How to compile:

- godep restore 
- remove vendor directory from packer dependency (../packer)
- godep restore inside packer directory
- go build -o /path/to/file main.go (the path is usually ~/.packer.d/plugins as it is the packer plugins directory path)
