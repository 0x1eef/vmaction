package ssh

import (
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
		conn, err := _ssh.Dial("tcp", ip, &config)
		if err != nil {
			attempts++
			if attempts >= max {
				return nil, err
			}
			time.Sleep(1 * time.Second)
		} else {
			return conn.NewSession()
		}
	}
}
