package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/google/uuid"
	//"Gator/internal/config"
	"Gator/internal/database"
)

func handlerAddfeed(s *state, cmd command ) error {
	if len(cmd.Args) < 2 {
		log.Fatal("usage: Gator addfeed NAME URL")
	}
	ctx := context.Background()

	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		fmt.Print("Failed to get current user from DB")
		os.Exit(1)
	}

	_, err = fetchFeed(ctx, url)
	if err != nil {
		fmt.Printf("Failed to fetch feed: %v\n", err)
		os.Exit(1)
	}

	f, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		UserID:    user.ID,
		Url:       url,
		Name:      name,
	})
	if err != nil {
		return err
	}
	fmt.Println("Feed created successfully:")
	printFeed(f)
	fmt.Println()
	fmt.Println("=====================================")

	

	return nil
}
func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
