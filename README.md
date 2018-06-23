# goC2
goC2 is a command and control implant with call back server written entirely in the 
Go programming language.

An 'implant' is a glorified name for a reverse shell which is established on a 
target endpoint which establishes a communication path to a pre-determined 'server' 
listening for inbound connections.  Once the connection is established, commands can 
be sent to the implant and executed on the target endpoint.

Instead of just calling this a reverse shell, the robustness of the wrapped reverse 
shell has features commonly found in 'agents' in the commercialized world of 
security tools, therefore 'implant' and 'agent' are interchangeable names. 

## Implant Features
* Auto host discovery, enumeration and data collection
* Unique implant UUIDv4 generated per install
* Cross-platform build (Windows, Mac OS, Linux)
* and more 

## Server Features
* TBD

## Building the Implant
To build an implant, a Makefile was created to make it easy to build and clean 
binaries. 

Environmental variables set within the Makefile are customized per the official 
instructions within the Go Documentation here: 
https://golang.org/doc/install/source#environment
