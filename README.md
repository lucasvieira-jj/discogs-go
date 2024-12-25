# Discogs Go

This test was developed using the principal API from Discogs, that's because the discogs not permitted to scrapping in the web site and make hard to develop the code for this, So I use the api for this test.

# First thing first
```bash
# Clone this repository
git clone https://github.com/lucasvieira-jj/discogs-go.git
```

## Prerequisites

Verify if you have golang installed correctly
- **Golang 1.23.3**

-- Go to the project folder

-- And then use this command to retrieve the project libs
go mod tidy

## Config
We need a initial config to run the code, because we are using the API from discogs to get data

1. We have a file in root that called '.env-example' remove the -example or create one file called .env with this config
   TOKEN=USE_A_TOKEN_FROM_DISCOGS_API
2. Go to the discogs api website and get your token
   https://www.discogs.com/settings/developers
3. Change the TOKEN variable into .env file

## How to execute
Inside the root folder and after run the command to get all go mod, exectute this command into your terminal **go run**