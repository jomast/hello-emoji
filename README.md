# hello-emoji

A simple web server that returns a random emoji.

## Why?

Because returning "hello world" is boring and doesn't give
any indication if the content is being cached at some level.

## Variables

You can set the following environment variables

| Var     | Description                                  | Default |
| ------- | -------------------------------------------- | ------- |
| `COUNT` | The number of emojis to return               | `1`     |
| `PORTS` | A comma separated list of ports to listen on | `80`    |
