# Golang ChitChat

Run this to alter GOPATH:

`export GOPATH=$HOME/Desktop/MProjects/GoWebProgramming`

To build the chitchat app:

`go build -i -v -o myApp`


To Run after building:

`./myApp`

To Run without building:

`go run .`

To create a docker image:

`docker build user/chitchat .`

To Run from docker:

`docker run --rm --name chitchat -p 8080:8080 user/chtichat`
