package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	//"net"
	"net/http"
	//"net/url"
	"os"
	"os/signal"
	//"strings"
	"syscall"
	//"time"
)

func clearScreen() {
    fmt.Print("\033[2J\033[;H")
}

func printBanner() {
	ascii := `
	██▓███   ██▀███   ▒█████  ▒██   ██▒  ▄████  ▒█████  
	▓██░  ██▒▓██ ▒ ██▒▒██▒  ██▒▒▒ █ █ ▒░ ██▒ ▀█▒▒██▒  ██▒
	▓██░ ██▓▒▓██ ░▄█ ▒▒██░  ██▒░░  █   ░▒██░▄▄▄░▒██░  ██▒
	▒██▄█▓▒ ▒▒██▀▀█▄  ▒██   ██░ ░ █ █ ▒ ░▓█  ██▓▒██   ██░
	▒██▒ ░  ░░██▓ ▒██▒░ ████▓▒░▒██▒ ▒██▒░▒▓███▀▒░ ████▓▒░
	▒▓▒░ ░  ░░ ▒▓ ░▒▓░░ ▒░▒░▒░ ▒▒ ░ ░▓ ░ ░▒   ▒ ░ ▒░▒░▒░ 
	░▒ ░       ░▒ ░ ▒░  ░ ▒ ▒░ ░░   ░▒ ░  ░   ░   ░ ▒ ▒░ 
	░░         ░░   ░ ░ ░ ░ ▒   ░    ░  ░ ░   ░ ░ ░ ░ ▒  
	░         ░ ░   ░    ░        ░     ░ ░
	`
	fmt.Printf("\033[1;38;5;135m%s\033[1;38;5;251m\n",ascii)
}

func printMenu() {
	options := []string{"socks4", "socks5", "http", "all"}
	for i, option := range options {
		fmt.Printf("\n\t\t\033[1;38;5;135m %d ~> \033[1;38;5;251m%s",i+1,option)
	}
}

func userCMD() {
	var cmd int
	for {
		fmt.Printf("\n\n\t\t\033[1;38;5;135m $ \033[1;38;5;251m%s\033[1;38;5;135m > \033[1;38;5;251m","Enter Number")
		if _,err := fmt.Scanf("%d", &cmd); err != nil || cmd <1 || cmd >4 {
			printAll()
			continue
		}
		break
	}
	handleRequest(cmd)
}

func printAll() {
	clearScreen()
	printBanner()
	printMenu()
}

func handleRequest(a int) {
	options := []string{"socks4", "socks5", "http", "all"}
	option := options[a-1]
	filename := option + ".txt"
	url := fmt.Sprintf("https://api.openproxylist.xyz/%s.txt", option)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("\033[1;38;5;9mError getting %s proxies:\033[0m \033[1;38;5;231m%s\033[0m\n", option, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("\033[1;38;5;9mError reading %s proxy response:\033[0m \033[1;38;5;231m%v\033[0m	",option,err)
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("\n\033[1;38;5;9mError writing to file:\033[0m \033[1;38;5;231munable to save proxies to \033[1;38;5;135m%s\033[0m",filename)
		return
	}
	defer file.Close()

	if _, err := file.Write(body); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("\n\t\033[1;38;5;135m # \033[1;38;5;251mSaved proxies to\033[1;38;5;135m %s",filename)
}

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		fmt.Printf("\033[1;38;5;220mLeaving so soon? Come back soon, friend!\033[0m")
		os.Exit(0)
	}()
	printAll()
	userCMD()
}
