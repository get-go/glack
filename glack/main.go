package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/get-go/glack"
)

var token = flag.String("token", "", "Token for Slack API")
var saveToken = flag.Bool("save-token", false, "Save the slack token in ~/.glack")
var channel = flag.String("channel", "#general", "Slack channel to send message to")
var username = flag.String("username", "Glack", "Name of the bot user to send as")
var icon = flag.String("icon", ":shoe:", "Emoji icon for the message")

func main() {
	flag.Parse()

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if *token == "" {
		file, err := os.Open(usr.HomeDir + "/.glack")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Token was not set.\n%+v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if len(lines) <= 0 {
			fmt.Fprintln(os.Stderr, "Token was not set. (b)")
			os.Exit(1)
		}
		*token = string(lines[0])
	}

	if *saveToken {
		file, err := os.Create(usr.HomeDir + "/.glack")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Token was not saved.\n%+v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		w := bufio.NewWriter(file)
		fmt.Fprintln(w, *token)
		err = w.Flush()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Token was not saved.\n%+v\n", err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, "Token was saved.")
	}

	if flag.NArg() < 1 && *saveToken {
		return
	} else if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "No message text provided.")
		os.Exit(1)
	}

	c := glack.New(*token)
	m := glack.Message{Channel: *channel, Message: flag.Arg(0), Username: *username, Icon: *icon}

	id, err := c.Send(&m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Send Command Failed:\n%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Glack Message sent! Id: %v\n", id)
	os.Exit(0)
}
