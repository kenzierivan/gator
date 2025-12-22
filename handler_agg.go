package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kenzierivan/gator/internal/database"
	"github.com/lib/pq"
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

	for _, item := range rssFeed.Channel.Item {
		publishedDate, err := time.Parse(
			"Mon, 02 Jan 2006 15:04:05 -0700",
    		item.PubDate,
		)
		if err != nil {
			return fmt.Errorf("couldn't parse published date: %w", err)
		}

		err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: item.Description,
			PublishedAt: publishedDate,
			FeedID: nextFeed.ID,
		})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				continue
			}
			return fmt.Errorf("couldn't create post: %w", err)
		}
		fmt.Printf("Found post: %s\n", item.Title)
	}

	return nil
}