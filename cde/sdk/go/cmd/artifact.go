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
	rootCmd.AddCommand(artifactCmd)
	artifactCmd.AddCommand(artifactPackagedCmd)
	artifactCmd.AddCommand(artifactPublishedCmd)

	artifactCmd.PersistentFlags().StringVarP(&artifactParams.ArtifactId, "id", "i", "", "Artifact Id")
	artifactCmd.PersistentFlags().StringVarP(&artifactParams.ArtifactName, "name", "n", "", "Artifact Name")
	artifactCmd.PersistentFlags().StringVarP(&artifactParams.ArtifactVersion, "version", "v", "", "Artifact Version")
	artifactCmd.PersistentFlags().StringToStringVarP(&artifactParams.ArtifactData, "data", "d", map[string]string{}, "Artifact Data")
}

var artifactCmd = &cobra.Command{
	Use:   "artifact",
	Short: "Emit Artifact related Events",
	Long:  `Emit Artifact related CloudEvent`,
}

var artifactParams = cde.ArtifactEventParams{}

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
		event, _ := cde.CreateArtifactEvent(cde.ArtifactPackagedEventV1, artifactParams)

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
		event, _ := cde.CreateArtifactEvent(cde.ArtifactPublishedEventV1, artifactParams)

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
