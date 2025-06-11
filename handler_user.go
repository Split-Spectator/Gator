package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	//"Gator/internal/config"
	"Gator/internal/database"
	"os"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	ctx := context.Background()

	_, err := s.db.GetUser(ctx, cmd.Args[0])
	if err != nil {
		fmt.Println("User does not exist")
		os.Exit(1)
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func registerHandler(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("missing required argument USERNAME")
	}

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, cmd.Args[0])
	if err == nil {
		fmt.Println("User already exists")
		os.Exit(1)
	}

	u, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		fmt.Printf("Failed to create user %v\n", err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(u.Name)
	if err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}

	fmt.Printf("User Created %v\n", u)

	return nil
}