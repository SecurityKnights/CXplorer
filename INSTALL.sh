# Installing Go
sudo apt-get install golang

# Copying Custom Modules
sudo mkdir $HOME/go/src/CXplorer
sudo cp -r modules/ $HOME/go/src/CXplorer

# Installing Required C Libraries
sudo apt-get install libpcap-dev

# Installing Required Go Packages
go get github.com/google/gopacket
go get golang.org/x/net/bpf
go get golang.org/x/sys/unix

# Copying Wordlist
sudo mkdir /usr/local/CXplorer
sudo cp wordlist.txt /usr/local/CXplorer/

# Building Tool
go build -o CXplorer

# Moving Executable to Bin
sudo cp CXplorer /usr/local/bin/
