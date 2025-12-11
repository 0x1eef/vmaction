package ssh

import (
	"fmt"
	"time"

	_ssh "golang.org/x/crypto/ssh"
)

type Session = _ssh.Session

var (
	max    = 100
	config = _ssh.ClientConfig{
		User: "root",
		Auth: []_ssh.AuthMethod{_ssh.Password("")},
	}
)

func Run(ip string) (*_ssh.Session, error) {
	attempts := 0
	for {
		conn, err := _ssh.Dial("tcp", fmt.Sprintf("%s:%s", ip, "22"), &config)
		if err != nil {
			attempts++
			if attempts >= max {
				return nil, err
			}
			time.Sleep(1 * time.Second)
			fmt.Printf("%v (%d/%d)\n", err, attempts, max)
		} else {
			return conn.NewSession()
		}
	}
}
