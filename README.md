# Slack Bot with Age Calculator

This is a simple Slack bot written in Go using the `slacker` package. The bot allows users to calculate their seniority based on their birth year.


## Getting Started

To use this bot, you'll need to create a Slack app and obtain bot and app tokens. Follow the steps below to set up the bot:

1. Clone the repository:
   ```shell
   git clone https://github.com/jepon26/Slack-bot/
   
   
   
# Set the Slack bot token and app token as environment variables:
```shell
export SLACK_BOT_TOKEN="your-bot-token"
export SLACK_APP_TOKEN="your-app-token"
```


# Build and run the bot:

```shell
go build ./Slack-bot
```



# Usage:

### The bot supports a single command:

```shell
My old is <year>
```
  
  
### Replace <year> with your birth year. For example, to calculate your seniority if you were born in 2000, you would use the following command:
  
```shell
My age is 2000
```
  
  
  
  
