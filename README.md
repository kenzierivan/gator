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

## Database setup (postgres)

1. Make sure Postgres is installed and running.
2. Create a config file in your home directory, the JSON file should have this structure:
    ```json
    {
        "db_url": "connection_string_goes_here",
        "current_user_name": "username_goes_here"
    }
    ```
3. Create a database for gator:
   ```bash
   create database <name of database>
   ```
4. Setup your Postgres connection URL, for example:
    ```bash
   postgres://postgres:yourpassword@localhost:5432/gator?sslmode=disable
   ```
5. Add connection URL to the config file.

## Commands

`register`

**Description:**  
Creates a new user in the database.  
**Usage:**  
```bash
gator register <name>
```

`login`

**Description:**  
Sets the current user in the config file.  
**Usage:**  
```bash
gator login <name>
```

`addfeed`

**Description:**  
Adds a new RSS feed to the database and associates it with the current user.  
**Usage:**  
```bash
gator addfeed <name> <url>
```

`agg`

**Description:**  
Fetches and stores new posts from all subscribed RSS feeds for the current user.  
**Usage:**  
```bash
gator agg <time_between_requests>
```

`browse`

**Description:**  
Lists out all the posts that the current user follows.  
**Usage:**  
```bash
gator browse
```
or
```bash
gator browse <limit>
```




