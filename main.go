package main

import (
	 _ "github.com/lib/pq"
	"log"
	"database/sql"
	"os"
	"Gator/internal/config"
	"Gator/internal/database"
)


type state struct {
	db  *database.Queries    
	cfg *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	dbQueries := database.New(db)

	programState := &state{
		db:   dbQueries,
		cfg: &cfg,
	}


	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", registerHandler)
	cmds.register("reset", deleteUsers)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("feeds", handlerFeeds)
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("addfeed", middlewareLoggedIn(handlerAddfeed))
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerListFeedFollows))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

   
}
