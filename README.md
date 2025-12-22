# Gator

Gator is an RSS feed aggregator in go.

## Installation

1. Make sure you have Go and Postgres installed.
2. Clone this repository and build the program:
   ```bash
   git clone https://github.com/kenzierivan/gator
   cd gator
   go build
   ./gator
   ```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
    {
        "db_url": "postgres://username:@localhost:5432/database?sslmode=disable",
        "current_user_name": "username_goes_here"
    }
```

Replace the values with your database connection string.


## Usage

`register`

Creates a new user:

```bash
gator register <name>
```

`login`

Sets the current user in the config file:

```bash
gator login <name>
```

`addfeed`

Adds a new RSS feed to the database and associates it with the current user: 

```bash
gator addfeed <name> <url>
```

`agg`

Fetches and stores new posts from all subscribed RSS feeds for the current user:

```bash
gator agg <time_between_requests>
```

`browse`

Lists out all the posts that the current user follows:

```bash
gator browse
```
or
```bash
gator browse <limit>
```

There are a few other commands you'll need as well:

- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database



