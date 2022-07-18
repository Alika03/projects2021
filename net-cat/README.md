# net-cat

## Authors: Alika03

### Objectives

### How to run:

1. Run the **go build TCPChat.go** command in the root directory of the project
2. Run the **./TCPChat $port** or **./TCPChat** default port:8989

This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

   * NetCat, nc system command, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open TCP connections, send UDP packages, listen on arbitrary TCP and UDP ports and many more.

   * To see more information about NetCat inspect the manual man nc.

Project works in a similar way that the original NetCat works, it has the following features :

   * TCP connection between server and multiple clients (relation of 1 to many).
    * A name requirement to the client.
    * Control connections quantity.
    * Clients able to send messages to the chat.
    * It does not broadcast EMPTY messages from a client.
    * Messages sent and identified by the time that was sent and the user name of who sent the message, example : [2020-01-20 15:48:41][client.name]:[client.message]
    * If a Client joins the chat, all the previous messages sent to the chat is uploaded to the new Client.
    * If a Client connects to the server, the rest of the * Clients are informed by the server that the Client joined the group.
    * If a Client exits the chat, the rest of the Clients are informed by the server that the Client left.
    * All Clients receive the messages sent by other Clients.
    * If a Client leaves the chat, the rest of the Clients do not disconnect.
    * If there is no port specified, then sets as default the port 8989. Otherwise, program responds with usage message: [USAGE]: ./TCPChat $port

### Used Packages
* io
* log
* os
* fmt
* net
* sync
* time
* bufio
* errors
* strings
* reflect

* Manipulation of structures.
* Net-Cat
* TCP/UDP
   * TCP/UDP connection
   * TCP/UDP socket
* Go concurrency
    * Channels
    * Goroutines
* Mutexes
* IP and ports

### Usage

Server: 
\$ go run .
Listening on the port :8989
\$ go run . 2525
Listening on the port :2525
\$ go run . 2525 localhost
[USAGE]: ./TCPChat \$port
\$

Client:
\$ nc localhost *port

* Server answers the client with a linux logo and asks for their name, when connection is received
* It accepts connection with non-empty name

# Image
* SERVER
1. ![alt text](https://github.com/Alika03/net-cat/blob/b257e0b58c1fb02989957a5c6e658b13764399f9/image/Screenshot%20from%202021-11-15%2018-26-20.png?raw=true)
* CLIENT
2. ![alt text](https://github.com/Alika03/net-cat/blob/b257e0b58c1fb02989957a5c6e658b13764399f9/image/Screenshot%20from%202021-11-15%2018-27-14.png?raw=true)