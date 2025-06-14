package main

import (
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"Gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByURL(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	ffRow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(ffRow.UserName, ffRow.FeedName) //this line causing error
	return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	

	feedFollows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("* %s\n", ff.FeedName)
	}

	return nil
}


func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}


func handlerUnfollow(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	url := cmd.Args[0]

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	
	err := s.db.DeleteFeed(ctx, database.DeleteFeedParams{
		UserID: user.ID,
		Url:    url,
	})

	if err != nil {
		return err
	}

	fmt.Println("Feed unfollow successfully")

	return nil
}

 
