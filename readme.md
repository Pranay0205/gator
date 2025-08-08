# Gator RSS Feed Aggregator

A command-line RSS feed aggregator built in Go that allows you to follow, fetch, and browse RSS feeds from your terminal.

## Features

- User management (register, login, list users)
- RSS feed management (add feeds, list feeds, follow/unfollow)
- Automatic feed aggregation with configurable intervals
- Browse recent posts from followed feeds
- PostgreSQL database backend with UUID-based primary keys
- Concurrent-safe feed fetching

## Prerequisites

Before you can run Gator, you'll need to have the following installed on your system:

### 1. Go (version 1.24.4 or later)
Download and install Go from [https://golang.org/dl/](https://golang.org/dl/)

### 2. PostgreSQL
Install PostgreSQL from [https://www.postgresql.org/download/](https://www.postgresql.org/download/)

Make sure PostgreSQL is running and you have:
- A database created for Gator
- Connection details (host, port, username, password, database name)

## Installation

### Install the Gator CLI

```bash
go install github.com/Pranay0205/gator@latest
```

This will install the `gator` binary to your `$GOPATH/bin` directory. Make sure this directory is in your `PATH`.

## Configuration

### 1. Create the configuration file

Gator looks for a configuration file named `.gatorconfig.json` in your home directory.

Create this file with the following structure:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator_db?sslmode=disable",
  "current_user_name": ""
}
```

Replace the database URL with your actual PostgreSQL connection string:
- `username`: Your PostgreSQL username
- `password`: Your PostgreSQL password  
- `localhost:5432`: Your PostgreSQL host and port
- `gator_db`: Your database name

### 2. Set up the database schema

You'll need to run the database migrations to create the required tables. If you're using a migration tool like `goose`, run:

```bash
goose -dir sql/schema postgres "your-connection-string" up
```

## Usage

### User Management

**Register a new user:**
```bash
gator register <username>
```

**Login as an existing user:**
```bash
gator login <username>
```

**List all users:**
```bash
gator users
```

**View user details:**
```bash
gator user <username>
```

### Feed Management

**Add a new RSS feed:**
```bash
gator addfeed <feed-name> <feed-url>
```
Example:
```bash
gator addfeed "TechCrunch" "https://techcrunch.com/feed/"
```

**List all feeds:**
```bash
gator feeds
```

**Follow an existing feed:**
```bash
gator follow <feed-url>
```

**Unfollow a feed:**
```bash
gator unfollow <feed-url>
```

**View feeds you're following:**
```bash
gator following
```

### Content Browsing

**Browse recent posts:**
```bash
gator browse [limit]
```
Example:
```bash
gator browse 10  # Shows 10 most recent posts
gator browse     # Shows 2 most recent posts (default)
```

### Feed Aggregation

**Start automatic feed fetching:**
```bash
gator agg <duration>
```
Example:
```bash
gator agg 30s    # Fetch feeds every 30 seconds
gator agg 5m     # Fetch feeds every 5 minutes
gator agg 1h     # Fetch feeds every hour
```

This command will run continuously, fetching and updating feeds at the specified interval.

### Utility Commands

**Reset database (delete all users):**
```bash
gator reset
```

## Example Workflow

1. **Set up your environment:**
   ```bash
   # Create config file in your home directory
   echo '{"db_url": "postgres://user:pass@localhost/gator_db?sslmode=disable", "current_user_name": ""}' > ~/.gatorconfig.json
   ```

2. **Register and login:**
   ```bash
   gator register john
   gator login john
   ```

3. **Add some feeds:**
   ```bash
   gator addfeed "Hacker News" "https://hnrss.org/frontpage"
   gator addfeed "Go Blog" "https://blog.golang.org/feed.atom"
   ```

4. **Start aggregation in the background:**
   ```bash
   gator agg 10m &
   ```

5. **Browse your feeds:**
   ```bash
   gator browse 5
   ```

## Project Structure

- `main.go` - Entry point and command registration
- `commands.go` - Command handling framework
- `handler*.go` - Individual command handlers
- `internal/config/` - Configuration management
- `internal/database/` - Generated database queries (sqlc)
- `internal/rss/` - RSS feed fetching and parsing
- `sql/schema/` - Database migration files
- `sql/queries/` - SQL query definitions

## Dependencies

- [github.com/google/uuid](https://github.com/google/uuid) - UUID generation
- [github.com/lib/pq](https://github.com/lib/pq) - PostgreSQL driver
- [sqlc](https://sqlc.dev/) - Type-safe SQL code generation

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request
