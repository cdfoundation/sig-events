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

	cde "github.com/cdfoundation/sig-events/cde/pkg/cdf/events"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(taskRunCmd)
	taskRunCmd.AddCommand(taskRunStartedCmd)
	taskRunCmd.AddCommand(taskRunFinishedCmd)

	taskRunCmd.PersistentFlags().StringVarP(&taskRunId, "id", "i", "", "Task Run Id")
	taskRunCmd.PersistentFlags().StringVarP(&taskRunName, "name", "n", "", "Task Run's Name")
	taskRunCmd.PersistentFlags().StringVarP(&taskRunPipelineId, "pipelineid", "p", "", "Task Run's Pipeline Id")
	taskRunCmd.PersistentFlags().StringToStringVarP(&taskRunData, "data", "d", map[string]string{}, "Task Run's Data")
}

var taskRunCmd = &cobra.Command{
	Use:   "taskrun",
	Short: "Emit TaskRun related Events",
	Long:  `Emit TaskRun related CloudEvent`,
}

var (
	taskRunId         string
	taskRunName       string
	taskRunPipelineId string
	taskRunData       map[string]string
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
