package main

import(
	"fmt"
	"context"


)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}

	err = s.db.DeleteFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete feeds: %w", err)
	}

	err = s.db.DeleteFollowFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete follow feeds: %w", err)
	}
	
	fmt.Println("Database reset successfully!")
	return nil
}