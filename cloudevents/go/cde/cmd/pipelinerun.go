package cmd

import (
	"context"
	cde "github.com/cdfoundation/sig-events/cde/pkg/cdf/events"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(pipelineRunCmd)
	pipelineRunCmd.AddCommand(pipelineRunStartedCmd)
	pipelineRunCmd.AddCommand(pipelineRunFinishedCmd)
	pipelineRunCmd.AddCommand(pipelineRunQueuedCmd)

	pipelineRunCmd.PersistentFlags().StringVarP(&pipelineRunId, "id", "i", "", "Pipeline Run Id")
	pipelineRunCmd.PersistentFlags().StringVarP(&pipelineRunName, "name", "n", "", "Pipeline Run's Name")
	pipelineRunCmd.PersistentFlags().StringVarP(&pipelineRunStatus, "status", "s", "", "Pipeline Run's Status")
	pipelineRunCmd.PersistentFlags().StringVarP(&pipelineRunURL, "url", "u", "", "Pipeline Run's URL")
	pipelineRunCmd.PersistentFlags().StringVarP(&pipelineRunErrors, "errors", "e", "", "Pipeline Run's Errors")
	pipelineRunCmd.PersistentFlags().StringToStringVarP(&pipelineRunData, "data", "d", map[string]string{}, "Pipeline Run's Data")
}

var pipelineRunCmd = &cobra.Command{
	Use:   "pipelinerun",
	Short: "Emit PipelineRun related Events",
	Long:  `Emit PipelineRun related CloudEvent`,
}

var (
	pipelineRunId     string
	pipelineRunName   string
	pipelineRunStatus string
	pipelineRunURL    string
	pipelineRunErrors string
	pipelineRunData   map[string]string
)

var pipelineRunStartedCmd = &cobra.Command{
	Use:   "started",
	Short: "Emit PipelineRun Started Event",
	Long:  `Emit PipelineRun Started CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreatePipelineRunEvent(cde.PipelineRunStartedEventV1, pipelineRunId,
			pipelineRunName, pipelineRunStatus, pipelineRunURL, pipelineRunErrors, pipelineRunData)

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

var pipelineRunFinishedCmd = &cobra.Command{
	Use:   "finished",
	Short: "Emit Pipeline Run Finished Event",
	Long:  `Emit Pipeline Run Finished CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreatePipelineRunEvent(cde.PipelineRunFinishedEventV1, pipelineRunId,
			pipelineRunName, pipelineRunStatus, pipelineRunURL, pipelineRunErrors, pipelineRunData)

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

var pipelineRunQueuedCmd = &cobra.Command{
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
		event, _ := cde.CreatePipelineRunEvent(cde.PipelineRunQueuedEventV1, pipelineRunId,
			pipelineRunName, pipelineRunStatus, pipelineRunURL, pipelineRunErrors, pipelineRunData)

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
