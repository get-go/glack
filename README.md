# Glack

Send simple messages to slack from the command line.

```shell
# send a message to #general
$ glack -token "xxxx-xxxxxxxxx-xxxx" "Hello World"
Glack Message sent! Id: 1234567890.000000

# send a message to #random with a penguin icon and title
$ glack -token "xxxx-xxxxxxxxx-xxxx" \
    -channel "#random" \
    -icon ":penguin:" \
    -username "Penguin" \
    "Hello from a silly penguin"
Glack Message sent! Id: 1234567890.000000

# tired of typing out that token? Save to ~/.glack
$ glack -token "xxxx-xxxxxxxxx-xxxx" --save-token
Token was saved.

$ glack -channel "#random" "Token Free!"
Glack Message sent! Id: 1234567890.000000

# Read from stdin, and send a message per line
$ echo "Hello" > test.txt
$ echo "World" >> test.txt
$ cat test.txt | glack
Glack Message sent! Id: 1234567890.000000
Glack Message sent! Id: 1234567890.000000

# Parse message object as JSON
$ glack -json '{"channel":"#random","icon":":arrow:","username":"The Arrow","message":"Watch Out"}'
Glack Message sent! Id: 1234567890.000000

```

You can use several types of tokens, but the easiest is to get a personal token from the [Slack API](https://api.slack.com/web) documentation. If you want to send a message to yourself you can make a private channel or use `@{username}` as the channel name, replacing with your own username.

## Install

```shell
$ go get -u github.com/get-go/glack
$ go install github.com/get-go/glack/...
# And if your go bin isn't in your path
$ export PATH=$PATH:$GOPATH/bin
```
