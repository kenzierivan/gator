package main

import (
	"context"
	"fmt"
	"time"
)
func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenReqs, _ := time.ParseDuration(cmd.Args[0])
	
	ticker := time.NewTicker(timeBetweenReqs)
	fmt.Println("Collecting feeds every", timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch next feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for _, feed := range rssFeed.Channel.Item {
		fmt.Printf("* %s\n", feed.Title)
	}
	return nil
}