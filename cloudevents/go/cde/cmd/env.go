package cmd

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cde "github.com/cdfoundation/sig-events/cde/pkg/cdf/events"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.AddCommand(envCreatedCmd)
	envCmd.AddCommand(envDeletedCmd)
	envCmd.AddCommand(envModifiedCmd)

	envCmd.PersistentFlags().StringVarP(&envId, "id", "i", "", "Environment's Id")
	envCmd.PersistentFlags().StringVarP(&envName, "name", "n", "", "Environment's Name")
	envCmd.PersistentFlags().StringVarP(&envRepoUrl, "repo", "r", "", "Environment's RepoUrl")
	envCmd.PersistentFlags().StringToStringVarP(&envData, "data", "d", map[string]string{}, "Environment's Data")
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Emit Environment related Events",
	Long:  `Emit Environment related CloudEvent`,
}

var (
	envId string
	envName string
	envRepoUrl string
	envData map[string]string
)

var envCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Env Created Event",
	Long:  `Emit Environment Created CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateEnvironmentEvent(cde.EnvironmentCreatedEventV1, envId, envName, envRepoUrl, envData)

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

var envDeletedCmd = &cobra.Command{
	Use:   "deleted",
	Short: "Emit Environment Deleted Event",
	Long:  `Emit Environment Deleted CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateEnvironmentEvent(cde.EnvironmentDeletedEventV1, envId, envName, envRepoUrl, envData)

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

var envModifiedCmd = &cobra.Command{
	Use:   "modified",
	Short: "Emit Environment Modified Event",
	Long:  `Emit Environment Modified CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateEnvironmentEvent(cde.EnvironmentModifiedEventV1, envId, envName, envRepoUrl, envData)

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


