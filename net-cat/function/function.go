package function

import (
	"bufio"
	"net"
	"os"
	"time"
)

func WriteClient(conn net.Conn, nam string, isJoinorLeft bool, clients map[net.Conn]string) {
	t := time.Now()
	if isJoinorLeft {
		for client, name := range clients {
			if client != conn {
				client.Write([]byte("\n\033[;32m" + nam + " has joined our chat...\033[0m\n[" + t.Format("2006-01-02 15:04:05") + "][" + name + "]: "))
			} else {
				client.Write([]byte("[" + t.Format("2006-01-02 15:04:05") + "][" + name + "]: "))
			}
		}
	} else {
		for client, name := range clients {
			if client != conn {
				client.Write([]byte("\n\033[;31m" + name + " has left our chat...\033[0m\n[" + t.Format("2006-01-02 15:04:05") + "][" + name + "]: "))
			}
		}
	}
}

func FindName(conn net.Conn, clients map[net.Conn]string) string {
	for client, name := range clients {
		if client == conn {
			return name
		}
	}
	return ""
}

func CheckEmpty(s string) bool {
	for _, char := range s {
		if char != ' ' && char != '\n' {
			return true
		}
	}
	return false
}

func InitDatafromFile(history []string) error {
	f, err := os.Open("data/chatdata.txt")
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewScanner(f)
	for buf.Scan() {
		history = append(history, buf.Text()+"\n")
	}
	return nil
}

func LoadingHistory(con net.Conn, words, history []string) {
	for _, hisMessenger := range words {
		history = append(history, hisMessenger)
	}
	if len(history) != 0 {
		for _, hisMessenger := range history {
			con.Write([]byte(hisMessenger))
		}
	}
}

func IsCheckLetters(word string) bool {
	if word == "" {
		return false
	}
	word = word[:len(word)-1]
	for _, char := range word {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}

func ReadLogoFromFile() ([]byte, error) {
	logo, err := os.Open("logo")
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(logo)
	size := buf.Size()
	logg := make([]byte, size)
	buf.Read(logg)
	return logg, nil
}
