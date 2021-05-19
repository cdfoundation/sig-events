package cmd

import (
	"context"
	cde "github.com/cdfoundation/sig-events/cde/pkg/cdf/events"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(taskRunCmd)
	taskRunCmd.AddCommand(taskRunStartedCmd)
	taskRunCmd.AddCommand(taskRunFinishedCmd)

	taskRunCmd.PersistentFlags().StringVarP(&taskRunId, "id", "i", "", "Task Run Id")
	taskRunCmd.PersistentFlags().StringVarP(&taskRunName, "name", "n", "", "Task Run's Name")
	taskRunCmd.PersistentFlags().StringVarP(&taskRunPipelineId, "pipelineid", "p", "", "Task Run's Pipeline Id")

	pipelineRunCmd.PersistentFlags().StringToStringVarP(&taskRunData, "data", "d", map[string]string{}, "Task Run's Data")
}

var taskRunCmd = &cobra.Command{
	Use:   "taskrun",
	Short: "Emit TaskRun related Events",
	Long:  `Emit TaskRun related CloudEvent`,
}

var (
	taskRunId     string
	taskRunName   string
	taskRunPipelineId string
	taskRunData   map[string]string
)

var taskRunStartedCmd = &cobra.Command{
	Use:   "started",
	Short: "Emit TaskRun Started Event",
	Long:  `Emit TaskRun Started CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateTaskRunEvent(cde.TaskRunStartedEventV1, taskRunId,
			taskRunName, taskRunPipelineId, taskRunData)

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

var taskRunFinishedCmd = &cobra.Command{
	Use:   "finished",
	Short: "Emit TaskRun Finished Event",
	Long:  `Emit TaskRun Finished CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateTaskRunEvent(cde.TaskRunFinishedEventV1, taskRunId,
			taskRunName, taskRunPipelineId, taskRunData)

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

