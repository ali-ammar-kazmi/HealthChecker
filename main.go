package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \nError: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable, \nFrom: %v\nTo: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}

func main() {
	app := &cli.App{
		Name:  "HealthChecker",
		Usage: "Tool to check website is UP/DOWN",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "port number to check",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if port == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
