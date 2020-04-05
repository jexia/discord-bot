# Jexia Discord Bot
A Discord bot covering all of the functions needed on Jexia's Discord server.

| Variables  | Description |
| ------------- | ------------- |
| `token` | The token provided by Discord to authenticate the gateway and API requests  |
| `prefix`  | The value prefixing the commands that are called though a Discord message  |

| Endpoints  | Description | Events Supported |
| ------------- | ------------- | ------------- |
| `/github` | This is the endpoint for receiving GitHub's webhook payload events. Should be added as if it was a webhook. | `release` |

| Commands  | Description | Permissions Required |
| ------------- | ------------- | ------------- |
| `ping` | This command simply returns the time taken to respond to an event. | *none* |
