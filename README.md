# Gator
CLI RSS feed aggregator 

## Dependencies 

- Go 
- Postgres

Setup: 
 - create a config file in your home directory, ~/.gatorconfig.json, with the following content:
 ```bash
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

# Commands 

register <name>
    registers new user and logs them in 

login <name>
    logs in user. no authentication

users
    prints list of users 

reset 
    deletes all users and feeds

addfeed
    adds rss feed
    `gator addfeed <"rss title">  <"RSS URL">'

follow 
    lets you follow a rss feed already exists in DB
    `gator follow <rss url>

unfollow <rss url>
    drops following feed

following 
    lists all feeds you are following

agg <time duration>
    fetchs feeds every interval you set: 1s 1m 1h 
    terminal open while running. `crtl c` to stop 

browse <limit>
    browse set of rss feeds that you follow 





