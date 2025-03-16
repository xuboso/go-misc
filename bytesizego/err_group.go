package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func fetchURL(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, errors.New(resp.Status)
	}

	return resp, nil
}

func main() {
	urls := []string{
		"http://example1.com",
		"http://example2.com",
		"http://example3.com",
	}

	g, ctx := errgroup.WithContext(context.Background())
	for _, url := range urls {
		currentURL := url
		g.Go(func() error {
			resp, err := fetchURL(ctx, currentURL)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			fmt.Println(currentURL, "was fetched successfully")
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Encountered error: %v\n", err)
		return
	}

	fmt.Println("All URLs were fetched successfully")
}
