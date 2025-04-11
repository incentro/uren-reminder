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
- Install Docker for building and pushing container images 
- Install Docker Compose version 2.32.1 or later for spinning the application up as a container locally 

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
These configurations are also required when running the application as a docker container 

Create a `.env` file and copy the environment variables from `.env.dist` into it:
```shell
cp .env.dist .env
```

Configure the following environment variables with the appropriate values:

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

2. Run the bot directly:

   ```sh
   go run main.go
   ```

3. Build and run the bot as a container locally: 
   ```sh
   docker compose up --build
   ```

> [!NOTE]
> To remove the container and its attached volumes:
> ```sh
> docker compose down -v
> ```

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
- `docker-compose.yml`: the docker compose file used for spinning up the container

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.