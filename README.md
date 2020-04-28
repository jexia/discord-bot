[![Go Report Card](https://goreportcard.com/badge/github.com/jexia/discord-bot?style=flat-square)](https://goreportcard.com/report/github.com/jexia/discord-bot)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/jexia/discord-bot)
[![Release](https://img.shields.io/github/release/jexia/discord-bot.svg?style=flat-square)](https://github.com/jexia/discord-bot/releases/latest)
[![Chat on Discord](https://img.shields.io/badge/chat-on%20discord-7289da.svg?sanitize=true)](https://chat.jexia.com)

# Jexia Discord Bot

A Discord bot covering all of the functions needed on Jexia's Discord server.

| Variables | Description                                                                | Example Value        |
| --------- | -------------------------------------------------------------------------- | -------------------- |
| `token`   | The token provided by Discord to authenticate the gateway and API requests | _Discord Token_      |
| `prefix`  | The value prefixing the commands that are called though a Discord message  | `!`                  |
| `channel` | **(Depreciated)** The channel ID where the event message will be sent      | _Discord Channel ID_ |
| `address` | The value of the API server                                                | `0.0.0.0:80`         |

| Endpoints | Description                                                                                                 | Events Supported |
| --------- | ----------------------------------------------------------------------------------------------------------- | ---------------- |
| `/github` | This is the endpoint for receiving GitHub's webhook payload events. Should be added as if it was a webhook. | `release`        |

| Commands | Description                                                        | Permissions Required |
| -------- | ------------------------------------------------------------------ | -------------------- |
| `ping`   | This command simply returns the time taken to respond to an event. | _none_               |
