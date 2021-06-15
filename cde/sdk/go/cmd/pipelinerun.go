/*
Copyright 2021 The CD Events SDK Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"
	"log"

	cde "github.com/cdfoundation/sig-events/cde/sdk/go/pkg/cdf/events"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
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
