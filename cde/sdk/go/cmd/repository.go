package cmd

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cde "github.com/cdfoundation/sig-events/cde/pkg/cdf/events"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(repositoryCmd)
	repositoryCmd.AddCommand(repositoryCreatedCmd)
	repositoryCmd.AddCommand(repositoryModifiedCmd)
	repositoryCmd.AddCommand(repositoryDeletedCmd)


	repositoryCmd.PersistentFlags().StringVarP(&repositoryId, "id", "i", "", "Repository Id")
	repositoryCmd.PersistentFlags().StringVarP(&repositoryName, "name", "n", "", "Repository's Name")
	repositoryCmd.PersistentFlags().StringVarP(&repositoryURL, "url", "u", "", "Repository's URL")
	repositoryCmd.PersistentFlags().StringToStringVarP(&repositoryData, "data", "d", map[string]string{}, "Repository's Data")
}

var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Emit Repository related Events",
	Long:  `Emit Repository related CloudEvent`,
}

var (
	repositoryId string
	repositoryName string
	repositoryURL string
	repositoryData map[string]string
)

var repositoryCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Repository Created Event",
	Long:  `Emit Repository Created CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateRepositoryEvent(cde.RepositoryCreatedEventV1, repositoryId,
			repositoryName, repositoryURL, repositoryData)

		event.SetSource(source)

		// Set a target.
		ctx := cloudevents.ContextWithTarget(context.Background(), CDE_SINK)

		// Send that Event.
		log.Printf("sending event %s\n", event)

		if result := c.Send(ctx, event); !cloudevents.IsACK(result) {
			log.Fatalf("failed to send, %v", result)
			return result
		}

		return nil
	},
}

var repositoryModifiedCmd = &cobra.Command{
	Use:   "modified",
	Short: "Emit Repository Modified Event",
	Long:  `Emit Repository Modified CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateRepositoryEvent(cde.RepositoryModifiedEventV1, repositoryId,
			repositoryName, repositoryURL, repositoryData)

		event.SetSource(source)

		// Set a target.
		ctx := cloudevents.ContextWithTarget(context.Background(), CDE_SINK)

		// Send that Event.
		log.Printf("sending event %s\n", event)

		if result := c.Send(ctx, event); !cloudevents.IsACK(result) {
			log.Fatalf("failed to send, %v", result)
			return result
		}

		return nil
	},
}

var repositoryDeletedCmd = &cobra.Command{
	Use:   "deleted",
	Short: "Emit Repository Deleted Event",
	Long:  `Emit Repository Deleted CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateRepositoryEvent(cde.RepositoryDeletedEventV1, repositoryId,
			repositoryName, repositoryURL, repositoryData)

		event.SetSource(source)

		// Set a target.
		ctx := cloudevents.ContextWithTarget(context.Background(), CDE_SINK)

		// Send that Event.
		log.Printf("sending event %s\n", event)

		if result := c.Send(ctx, event); !cloudevents.IsACK(result) {
			log.Fatalf("failed to send, %v", result)
			return result
		}

		return nil
	},
}
