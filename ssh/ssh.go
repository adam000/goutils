package ssh

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type Client struct {
	conn ssh.Client
}

func NewClient(sshKeyFile, username, hostname string) (Client, error) {
	key, err := ioutil.ReadFile(sshKeyFile)
	if err != nil {
		return Client{}, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return Client{}, err
	}
	authMethod := ssh.PublicKeys(signer)

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			authMethod,
		},
	}

	conn, err := ssh.Dial("tcp", hostname+":22", config)
	if err != nil {
		return Client{}, err
	}

	return Client{*conn}, nil
}

// RunOne runs a single command on a given client.
func (c *Client) RunOne(args ...string) error {
	session, err := c.conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	err = session.Run(strings.Join(args, " "))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RunOneInteractive(args ...string) error {
	session, err := c.conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return fmt.Errorf("request for pseudo terminal failed: %s", err)
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	terminal.MakeRaw(0)

	err = session.Run(strings.Join(args, " "))
	if err != nil {
		return err
	}

	return nil
}
