package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kenzierivan/gator/internal/database"
)

func handlerbrowse(s *state, cmd command) error {
	if len(cmd.Args) != 0 && len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s || usage: %s <limit>", cmd.Name, cmd.Name)
	}

	limit := 2
	if len(cmd.Args) == 1 {
		selectedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("enter an integer: %w", err)
		}
		limit = selectedLimit
	}

	posts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}
	for _, post := range posts {
		printPost(post)
		fmt.Println("=====================================")
	}
	return nil
}

func printPost(post database.Post) {
	fmt.Printf("* Title:        %s\n", post.Title)
	fmt.Printf("* Description:  %s\n", post.Description)
	fmt.Printf("* Published at: %v\n", post.PublishedAt)
}