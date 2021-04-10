package util

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"time"
)

type SSHClient struct {
	IP       string
	Username string
	Password string
	Port     int
}

func (cli *SSHClient) RunCmd(cmd string, timeout time.Duration) (string, error) {
	config := &ssh.ClientConfig{
		User:            cli.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(cli.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
	}
	c, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", cli.IP, cli.Port), config)
	if err != nil {
		return "", err
	}
	defer func() { _ = c.Close() }()

	s, err := c.NewSession()
	if err != nil {
		return "", err
	}
	defer func() { _ = s.Close() }()

	res, err := s.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}

	return string(res), err
}
