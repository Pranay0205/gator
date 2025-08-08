# Gator RSS Feed Aggregator 📰

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

### 3. Goose (Database Migration Tool)
Install goose for running database migrations:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Database Setup

### 1. Create Database and User

Connect to PostgreSQL as a superuser and run:

```sql
-- Create database
CREATE DATABASE gator_db;

-- Create user with password
CREATE USER gator_user WITH PASSWORD 'your_secure_password';

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE gator_db TO gator_user;

-- Connect to the gator_db database first
\c gator_db

-- Grant schema privileges
GRANT ALL ON SCHEMA public TO gator_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO gator_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO gator_user;
```

### 2. Run Database Migrations

After cloning/downloading the project source code, navigate to the project directory and run:

```bash
# Navigate to project directory
cd path/to/gator

# Run migrations
goose -dir sql/schema postgres "postgres://gator_user:your_secure_password@localhost:5432/gator_db?sslmode=disable" up
```

## Installation

### Install the Gator CLI

```bash
go install github.com/Pranay0205/gator@latest
```

This will install the `gator` binary to your `$GOPATH/bin` directory. Make sure this directory is in your `PATH`.

### Verify Installation

Check that gator is installed correctly:
```bash
gator users  # Should show "Usage: users" or similar if no config exists yet
```

## Configuration

### Create the configuration file

Gator looks for a configuration file named `.gatorconfig.json` in your home directory.

**On Linux/Mac:**
```bash
echo '{"db_url": "postgres://gator_user:your_secure_password@localhost:5432/gator_db?sslmode=disable", "current_user_name": ""}' > ~/.gatorconfig.json
```

**On Windows:**
Create the file `C:\Users\YourUsername\.gatorconfig.json` with this content:

```json
{
  "db_url": "postgres://gator_user:your_secure_password@localhost:5432/gator_db?sslmode=disable",
  "current_user_name": ""
}
```

Replace connection details with your actual values:
- `gator_user`: Your PostgreSQL username
- `your_secure_password`: Your PostgreSQL password  
- `localhost:5432`: Your PostgreSQL host and port
- `gator_db`: Your database name

## Quick Start 🚀

Get up and running in under 2 minutes:

```bash
# 1. Register yourself as a user
gator register john

# 2. Add your first feed
gator addfeed "Hacker News" "https://hnrss.org/frontpage"

# 3. Start fetching feeds (runs in background)
gator agg 1m &

# 4. Wait a minute, then browse posts
gator browse 3
```

Expected output after step 4:
```
=== Post 1 ===
Feed Name: Hacker News
Title: Some interesting tech article
Published: Jan 8, 2025
URL: https://news.ycombinator.com/item?id=123456
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
Output example:
```
* alice
* bob (current)
* charlie
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
Examples:
```bash
gator browse 10  # Shows 10 most recent posts
gator browse     # Shows 2 most recent posts (default)
```

### Feed Aggregation

**Start automatic feed fetching:**
```bash
gator agg <duration>
```

⚠️ **Important**: This command runs continuously and will keep fetching feeds until you stop it (Ctrl+C).

Examples:
```bash
gator agg 30s    # Fetch feeds every 30 seconds
gator agg 5m     # Fetch feeds every 5 minutes  
gator agg 1h     # Fetch feeds every hour
```

To run in background:
```bash
gator agg 10m &  # Runs in background
```

### Utility Commands

**Reset database (delete all users):**
```bash
gator reset
```

## Troubleshooting

### Common Issues

**1. "command not found: gator"**
- Ensure `$GOPATH/bin` is in your PATH
- Try `go env GOPATH` to see your GOPATH
- Add `export PATH=$PATH:$(go env GOPATH)/bin` to your shell profile

**2. "couldn't get the current user details"**
- Make sure you've registered a user: `gator register yourusername`
- Check your config file exists: `cat ~/.gatorconfig.json`

**3. "error connecting to the database"**
- Verify PostgreSQL is running: `pg_isready`
- Test your connection string manually
- Check your database credentials in `.gatorconfig.json`

**4. "failed to parse xml"**
- The RSS feed might be temporarily down
- Try a different feed URL
- Some feeds require specific User-Agent headers (gator uses "gator")

**5. Permission denied errors**
- Check that your PostgreSQL user has the right permissions
- Ensure the config file is readable: `ls -la ~/.gatorconfig.json`

### Database Connection Test

Test your database connection:
```bash
# Using psql
psql "postgres://gator_user:your_password@localhost:5432/gator_db"
```

### Config File Format Issues

Your `.gatorconfig.json` should look exactly like this:
```json
{
  "db_url": "postgres://username:password@host:port/database?sslmode=disable",
  "current_user_name": ""
}
```

Common mistakes:
- Missing quotes around values
- Trailing commas
- Wrong file location

## Example Workflow

Here's a complete example from setup to browsing feeds:

```bash
# 1. Set up database (already done in setup)

# 2. Create config file
echo '{"db_url": "postgres://gator_user:mypass@localhost/gator_db?sslmode=disable", "current_user_name": ""}' > ~/.gatorconfig.json

# 3. Register and start using
gator register alice
gator addfeed "Go Blog" "https://blog.golang.org/feed.atom"
gator addfeed "Hacker News" "https://hnrss.org/frontpage"

# 4. Follow some feeds
gator follow "https://blog.golang.org/feed.atom"

# 5. Start aggregation (in background)
gator agg 5m &

# 6. Wait a few minutes, then browse
sleep 60
gator browse 5

# 7. See what you're following
gator following
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