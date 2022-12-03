# Capstone Project
## Summary
For my capstone project, I have chosen to create a discord bot that lets you roll dice. 
You do this by using a custom slash command that I also created. The requests can be sent using
HTTP GET requests or directly through discord.

In this project, I will demonstrate the following:
- [Building a REST API in Go]
- Using Go's [sql package](https://pkg.go.dev/database/sql) and a [sqlite driver](https://github.com/mattn/go-sqlite3) to cache and fetch data. 
- Deploying my REST API to a [third party PaaS ] (https://railway.app/project/c9c1436a-bfc1-44f1-b75f-05f1bca24731)
- [Building a user friendly CLI client application]
- Error handling
- Testing in Go
- Working with Discord to create the bot.

##User Stories
### Number One
### As a user I would like to be able to select the type of dice, number of sides, and whether or not the roll is at advantage, disadvantage or neither.
**Acceptance Criteria**
The bot will have preselected options to choose from that will allow users to do all of the above.
### Number Two
### As a user I would like to build my own bot and use the rest API to communicate with the API.