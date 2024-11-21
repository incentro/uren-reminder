# Slack Reminder Bot

This project is a Slack bot that sends reminder messages to a specified Slack channel. The bot reads messages from a JSON file, selects a random message, and posts it to the channel at a specified time.

## Features

- Reads messages from a JSON file
- Selects a random message to send
- Posts messages to a specified Slack channel
- Configurable via environment variables

## Prerequisites

- Go 1.18 or later
- A Slack workspace and a bot token
- A JSON file containing the messages

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/slack-reminder-bot.git
   cd slack-reminder-bot
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

## Configuration

Set the following environment variables:

- `SLACK_TOKEN`: Your Slack bot token
- `SLACK_CHANNEL_ID`: The ID of the Slack channel where the bot will post messages
- `AFAS_URL`: The URL to be included in the reminder message

## Usage

1. Create a JSON file named `messages.json` in the `input` directory with the following structure:

   ```json
   {
     "messages": [
       "Message 1",
       "Message 2",
       "Message 3"
     ]
   }
   ```

2. Run the bot:

   ```sh
   go run main.go
   ```

## Testing

To run the unit tests, use the following command:

```sh
go test ./...
```

## Project Structure

- `main.go`: The main entry point of the application
- `blocks.go`: Contains the function to create Slack message blocks
- `blocks_test.go`: Unit tests for `blocks.go`
- `day-utils_test.go`: Additional utility tests
- `input/messages.json`: JSON file containing the messages

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.