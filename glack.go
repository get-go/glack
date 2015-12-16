package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/user"
)

var version = "1.0.0"
var showVersion = flag.Bool("version", false, "Show the version ("+version+")")
var token = flag.String("token", "", "Token for Slack API")
var saveToken = flag.Bool("save-token", false, "Save the slack token in ~/.glack")
var channel = flag.String("channel", "#general", "Slack channel to send message to")
var username = flag.String("username", "Glack", "Name of the bot user to send as")
var icon = flag.String("icon", ":shoe:", "Emoji icon for the message")

func getHomeDir() (dir string, err error) {
	usr, err := user.Current()
	return usr.HomeDir, err
}

func printVersion() {
	if *showVersion {
		fmt.Fprintln(os.Stdout, version)
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	printVersion()

	if *token == "" {
		home, err := getHomeDir()
		file, err := os.Open(home + "/.glack")
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
		home, err := getHomeDir()
		file, err := os.Create(home + "/.glack")
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

	c := New(*token)
	s := func(channel, message, username, icon string) {
		m := Message{Channel: channel, Message: message, Username: username, Icon: icon}
		id, err := c.Send(&m)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Send Command Failed:\n%v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "Glack Message sent! Id: %v\n", id)
	}

	if flag.NArg() >= 1 {
		s(*channel, flag.Arg(0), *username, *icon)
	} else {
		//read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			s(*channel, scanner.Text(), *username, *icon)
		}
	}

	os.Exit(0)
}
