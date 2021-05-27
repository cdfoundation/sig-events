package cmd

import (
	"context"
	cde "github.com/cdfoundation/sig-events/cde/pkg/cdf/events"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.AddCommand(buildStartedCmd)
	buildCmd.AddCommand(buildFinishedCmd)
	buildCmd.AddCommand(buildQueuedCmd)

	buildCmd.PersistentFlags().StringVarP(&buildId, "id", "i", "", "Build Id")
	buildCmd.PersistentFlags().StringVarP(&buildName, "name", "n", "", "Build Name")
	buildCmd.PersistentFlags().StringVarP(&buildArtifactId, "artifact", "a", "", "Build's Artifact Id")
	buildCmd.PersistentFlags().StringToStringVarP(&buildData, "data", "d", map[string]string{}, "Build Data")
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Emit Build related Events",
	Long:  `Emit Build related CloudEvent`,
}

var (
	buildId     string
	buildName   string
	buildArtifactId string
	buildData   map[string]string
)

var buildStartedCmd = &cobra.Command{
	Use:   "started",
	Short: "Emit Build Started Event",
	Long:  `Emit Build Started CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateBuildEvent(cde.BuildStartedEventV1, buildId,
			buildName, buildArtifactId, buildData)

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

var buildFinishedCmd = &cobra.Command{
	Use:   "finished",
	Short: "Emit Build Finished Event",
	Long:  `Emit Build Finished CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateBuildEvent(cde.BuildFinishedEventV1, buildId,
			buildName, buildArtifactId, buildData)

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

var buildQueuedCmd = &cobra.Command{
	Use:   "queued",
	Short: "Emit PipelineRun Queued Event",
	Long:  `Emit PipelineRun Queued CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateBuildEvent(cde.BuildQueuedEventV1, buildId,
			buildName, buildArtifactId, buildData)

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
