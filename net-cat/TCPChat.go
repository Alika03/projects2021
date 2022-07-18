package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net-cat/function"
	"os"
	"time"

	"github.com/jroimartin/gocui"
)

var (
	clients = make(map[net.Conn]string)

	msg             = make(chan string)
	conns           = make(chan net.Conn)
	name            = make(chan string)
	words           = make([]string, 0)
	history         = make([]string, 0)
	dconns          = make(chan net.Conn)
	port            string
	amountofClients int = 0
)

func main() {
	// Checking arguments...
	if len(os.Args[1:]) > 1 {
		log.Println("[USAGE]: ./TCPChat $port")
		return
	} else if len(os.Args[1:]) == 0 {
		port = ":8989"
	} else {
		port = ":" + os.Args[1]
	}

	// Listeing port...
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}

	// Reading data from file...
	function.InitDatafromFile(history)

	// Save all data in file "chatdata" if not exists then create
	defer func() {
		f, err := os.OpenFile("data/chatdata.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
			return
		}
		for _, text := range words {
			f.Write([]byte(text))
		}
		f.Close()
		ln.Close()
	}()

	fmt.Println("Server is ...")

	// Reading Logo from file...
	logg, err := function.ReadLogoFromFile()
	if err != nil {
		log.Println(err)
		return
	}

	// Accepting client's connection
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			conns <- conn
			conn.Write(logg)
			go handleUserConnection(dconns, conn, name, msg, clients)
		}
	}()

	//Drawing server-interface...
	InitGraph(port)
}

func handleUserConnection(dconn chan net.Conn, conn net.Conn, name, msg chan string, clients map[net.Conn]string) {
	rd := bufio.NewReader(conn)
	conn.Write([]byte("[ENTER YOUR NAME]: "))
	nam, err := rd.ReadString('\n')
	if err != nil {
		return
	}
	for nam == "\n" || !function.CheckEmpty(nam) {
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		nam, err = rd.ReadString('\n')
		if err != nil {
			return
		}
	}
	nam = nam[:len(nam)-1]
	clients[conn] = nam
	function.LoadingHistory(conn, words, history)
	function.WriteClient(conn, nam, true, clients)
	name <- nam
	for {
		userInput, err := rd.ReadString('\n')
		if err != nil {
			dconn <- conn
			conn.Close()
			return
		}
		t := time.Now()
		if (len(userInput) == 1 && userInput == "\n") || !function.CheckEmpty(userInput) {
			id := fmt.Sprintf("[%v][%v]: ", t.Format("2006-01-02 15:04:05"), clients[conn])
			conn.Write([]byte("Empty line is not allowed!\n" + id))
			continue
		}
		id := fmt.Sprintf("[%v][%v]: %v", t.Format("2006-01-02 15:04:05"), clients[conn], userInput)
		for client := range clients {
			if client != conn {
				client.Write([]byte("\n" + id))
			}
		}
		msg <- id
	}

}

func server(g *gocui.Gui) {
	for {
		select {
		case con := <-conns:
			if amountofClients == 10 {
				con.Write([]byte("Sorry! the room is full\n"))
				con.Close()
				continue
			}
			amountofClients++
			g.Update(func(g *gocui.Gui) error {
				temp := con.RemoteAddr().String()
				v, err := g.View("v1")
				if err != nil {
					return err
				}
				v.Write([]byte("\033[36;1m" + temp + "\033[0m\n"))
				return nil
			})
		case nam := <-name:
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("v2")
				if err != nil {
					return err
				}
				v.Write([]byte("\033[32;1m" + nam + " has joined our chat...\033[0m\n"))
				v1, err := g.View("v3")
				if err != nil {
					return err
				}
				v1.Write([]byte("\033[32;1m" + nam + "\033[0m\n"))
				return nil
			})

		case msgs := <-msg:
			t := time.Now()
			if len(msgs) > 1 {
				words = append(words, msgs[1:])
			}
			for client, name := range clients {
				client.Write([]byte("[" + t.Format("2006-01-02 15:04:05") + "][" + name + "]: "))
			}
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("v2")
				if err != nil {
					return err
				}
				v.Write([]byte("\033[33;1m" + msgs + "\033[0m"))
				return nil
			})
		case dcon := <-dconns:
			nam := function.FindName(dcon, clients)
			amountofClients--
			function.WriteClient(dcon, nam, false, clients)
			delete(clients, dcon)
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("v2")
				if err != nil {
					return err
				}
				v.Write([]byte("\033[31;1m" + nam + " has left our chat...\033[0m\n"))
				v1, err := g.View("v3")
				if err != nil {
					return err
				}
				v1.Clear()
				for _, name := range clients {
					v1.Write([]byte("\033[32;1m" + name + "\033[0m\n"))
				}
				return nil
			})
		}
	}
}

// Drawing server-interface...
func InitGraph(port string) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	go server(g)

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Println(err)
		return
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Println("Please set the server in new window terminal")
		return
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if lcolumn1, err := g.SetView("v0", 1, 2, maxX/4, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		lcolumn1.Title = "SERVER"
		lcolumn1.Wrap = true
		lcolumn1.Write([]byte("\033[31;1mPort: " + port + "\033[0m\n"))
		lcolumn1.Write([]byte("\033[34;1mAuthor:\033[0m\n\033[35;1mAlika\033[0m\n"))
	}
	if lcolumn2, err := g.SetView("v1", 1, 11, maxX/4, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		lcolumn2.Wrap = true
		lcolumn2.Title = "IP-Address"
	}
	if midColumn, err := g.SetView("v2", maxX/4+1, 2, (maxX+maxX+maxX)/4, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		midColumn.Autoscroll = true
		midColumn.Title = "Chat"
		midColumn.Wrap = true
	}
	if rColumn, err := g.SetView("v3", (maxX+maxX+maxX)/4+1, 2, maxX-2, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		rColumn.Title = "Online"
		rColumn.Wrap = true
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
