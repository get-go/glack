# Glack

Send simple messages to slack from the command line.

```shell
# send a message to #general
$ glack -token "xxxx-xxxxxxxxx-xxxx" "Hello World"
# send a message to #random with a penguin icon and title
$ glack -token "xxxx-xxxxxxxxx-xxxx" \
    -channel "#random" \
    -icon ":penguin:" \
    -username "Penguin" \
    "Hello from a silly penguin"
# tired of typing out that token?
$ glack -token "xxxx-xxxxxxxxx-xxxx" --save-token
# Token was saved to "~/.glack"
$ glack -channel "#random" "Token Free!"
```

You can use several types of tokens, but the easiest is to get a personal token from the [Slack API](https://api.slack.com/web) documentation. If you want to send a message to yourself you can make a private channel or use `@{username}` as the channel name, replacing with your own username.

## Install

```shell
$ go get -u github.com/get-go/glack
$ go install github.com/get-go/glack
# And if your go bin isn't in your path
$ export PATH=$PATH:$GOPATH/bin
