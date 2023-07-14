# Telegram Bot for Profanity Filter (Golang)

This project is a Telegram bot written in Golang using the telegram-bot-api. The main feature of this bot is its profanity filter functionality - when added to a group, it checks new messages for any profanity or banned words and deletes any message containing such language.

## Usage

Before running the bot, make sure to replace the "TELEGRAM_BOT_TOKEN" placeholder in the .env file with your actual bot token from the BotFather on Telegram.

You can supplement the forbidden words by writing them to the `./words.json` file

Example:

```json
{
  "forbidden_words": [
    "your", 
    "forbidden",
    "words"
  ]
}
```

To start the bot, run:

```sh
make start
```

To start in dev mode:

```sh
make dev
```
