package main 

import (
	"fmt"
	"os"
	"context"
	"Gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
		return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
			if err != nil {
				fmt.Print("Failed to get current user from DB")
				os.Exit(1)
			}	
		return handler(s, cmd, user)
		}
	}