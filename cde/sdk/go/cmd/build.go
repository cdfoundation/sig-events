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
	buildId         string
	buildName       string
	buildArtifactId string
	buildData       map[string]string
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
