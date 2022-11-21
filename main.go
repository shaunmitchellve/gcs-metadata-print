package main

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// listFiles lists objects within specified bucket.
func listFiles(bucket string) error {
	// bucket := "bucket-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
					return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	it := client.Bucket(bucket).Objects(ctx, nil)
	for {
					attrs, err := it.Next()

					if err == iterator.Done {
									break
					}
					if err != nil {
									return fmt.Errorf("Bucket(%q).Objects: %v", bucket, err)
					}
					
					metadata := attrs.Metadata
					fmt.Println(attrs.Name)
					for key, value := range metadata {
						fmt.Printf("\t%v = %v\n", key, value)
					}
	}
	return nil
}

func main() {
	listFiles("sftp-files")
}