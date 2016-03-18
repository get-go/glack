package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/get-go/glack"
)

var showVersion = flag.Bool("version", false, "Show the version ("+glack.Version+")")
var token = flag.String("token", "", "Token for Slack API")
var saveToken = flag.Bool("save-token", false, "Save the slack token in ~/.glack")
var channel = flag.String("channel", "#general", "Slack channel to send message to")
var username = flag.String("username", "Glack", "Name of the bot user to send as")
var icon = flag.String("icon", ":shoe:", "Emoji icon for the message")
var filename = flag.String("upload-file", "", "Upload a file if specified.")
var quiet = flag.Bool("quiet", false, "Quiet the output to a minimum, just return message ID and errors")
var silent = flag.Bool("silent", false, "Silence all output, including message ID's and errors")

func getHomeDir() (dir string, err error) {
	usr, err := user.Current()
	return usr.HomeDir, err
}

func printVersion() {
	if *showVersion {
		fmt.Fprintln(os.Stdout, glack.Version)
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
			if !*silent {
				fmt.Fprintf(os.Stderr, "Token was not set.\n%+v\n", err)
			}
			os.Exit(1)
		}
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if len(lines) <= 0 {
			if !*silent {
				fmt.Fprintln(os.Stderr, "Token was not set. (b)")
			}
			os.Exit(1)
		}
		*token = string(lines[0])
	}

	if *saveToken {
		home, err := getHomeDir()
		file, err := os.Create(home + "/.glack")
		if err != nil {
			if !*silent {
				fmt.Fprintf(os.Stderr, "Token was not saved.\n%+v\n", err)
			}
			os.Exit(1)
		}
		defer file.Close()

		w := bufio.NewWriter(file)
		fmt.Fprintln(w, *token)
		err = w.Flush()
		if err != nil {
			if !*silent {
				fmt.Fprintf(os.Stderr, "Token was not saved.\n%+v\n", err)
			}
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, "Token was saved.")
	}

	c := glack.New(*token)

	if *filename != "" {
		c.UploadFile(*channel, *filename)
	} else if flag.NArg() >= 1 {
		sendMessage(&c, *channel, flag.Arg(0), *username, *icon)
	} else {
		//read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			// Account for empty lines, for now just ignore them
			text := scanner.Text()
			if len(text) > 0 {
				sendMessage(&c, *channel, scanner.Text(), *username, *icon)
			}
		}
	}

	os.Exit(0)
}

func sendMessage(client *glack.Client, channel, message, username, icon string) {
	m := glack.Message{Channel: channel, Message: message, Username: username, Icon: icon}
	_, msgID, err := client.Send(&m)
	if err != nil {
		if !*silent {
			fmt.Fprintf(os.Stderr, "Send Command Failed:\n%v\n", err)
		}
		os.Exit(1)
	}

	if !*silent && !*quiet {
		fmt.Fprintf(os.Stdout, "Glack Message sent! Id: %v\n", msgID)
	} else if !*silent {
		fmt.Fprintln(os.Stdout, msgID)
	}
}
