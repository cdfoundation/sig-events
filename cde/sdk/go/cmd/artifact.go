package cmd

import (
	"context"
	cde "github.com/cdfoundation/sig-events/cde/pkg/cdf/events"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(artifactCmd)
	artifactCmd.AddCommand(artifactPackagedCmd)
	artifactCmd.AddCommand(artifactPublishedCmd)


	artifactCmd.PersistentFlags().StringVarP(&artifactId, "id", "i", "", "Artifact Id")
	artifactCmd.PersistentFlags().StringVarP(&artifactName, "name", "n", "", "Artifact Name")
	artifactCmd.PersistentFlags().StringVarP(&artifactVersion, "version", "v", "", "Artifact Version")
	artifactCmd.PersistentFlags().StringToStringVarP(&artifactData, "data", "d", map[string]string{}, "Artifact Data")
}

var artifactCmd = &cobra.Command{
	Use:   "artifact",
	Short: "Emit Artifact related Events",
	Long:  `Emit Artifact related CloudEvent`,
}

var (
	artifactId     string
	artifactName   string
	artifactVersion string
	artifactData   map[string]string
)

var artifactPackagedCmd = &cobra.Command{
	Use:   "packaged",
	Short: "Emit Artifact Packaged Event",
	Long:  `Emit Artifact Packaged CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateArtifactEvent(cde.ArtifactPackagedEventV1, artifactId,
			artifactName, artifactVersion, artifactData)

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

var artifactPublishedCmd = &cobra.Command{
	Use:   "finished",
	Short: "Emit Artifact Published Event",
	Long:  `Emit Artifact Published  CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateArtifactEvent(cde.ArtifactPublishedEventV1, artifactId,
			artifactName, artifactVersion, artifactData)

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

