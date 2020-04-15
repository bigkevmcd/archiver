package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func archiveToS3Cmd() *cobra.Command {
	var bucketURL string
	cmd := &cobra.Command{
		Use:   "archive --bucket-url s3://bucket/prefix <file>",
		Short: "upload the provided file or files to S3.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, n := range args {
				log.Printf("uploading %s to %s\n", n, bucketURL)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&bucketURL, "bucket-url", "", "s3 bucket to upload to")
	cmd.MarkFlagRequired("bucket-url")
	return cmd
}
