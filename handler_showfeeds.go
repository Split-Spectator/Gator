package main

import (
	"fmt"
	"context"
)


 func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.ListFeeds(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("\nArticle: %s\n | Url: %s\n | User: %s\n", feed.Name, feed.Url, feed.UserName)
	}

	return nil
}