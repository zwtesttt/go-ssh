package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

type SshInfo struct {
	address  string
	port     int
	password string
}

func ConSsh(s *SshInfo) *ssh.Session {
	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.address, s.port), &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password(s.password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	// 建立新会话
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}

	return session
}
