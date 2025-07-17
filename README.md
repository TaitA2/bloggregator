# Bloggregator - An RSS Feed Blog Aggregator

## Requirements
* Postgres
* Go

## Installation
### Install with Go
* Run the following command in your terminal:
`go install github.com/TaitA2/bloggregator@latest`

## Setup
* Manually create a config file in your home directory, ~/.gatorconfig.json, with the following content:
`{
  "db_url": "postgres://example"
}`

## Usage
* Run the program by running `bloggregator [command]` 

### Commands
* help                 - prints help message
* register [username]  - register a new user with the given username
* login [username]     - login as the given user
* users                - lists all registered users
* reset                - removes all registered users
* feeds                - lists all available feeds
* addfeed [name] [url] - saves the RSS feed of the given URL as the given name to the available feeds
* follow [feed name]   - adds the named feed to the current user's followed feeds
* unfollow [feed name] - removes the named feed from the current user's followed feeds
* following            - lists all of the current user's followed feeds
* agg [interval]       - aggregates all available feeds once per given interval (1s, 1m, 1h, etc.). Intended to be run in the background.
* browse [limit]       - prints [limit] aggregated posts from followed feeds
