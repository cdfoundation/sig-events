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
	repositoryId   string
	repositoryName string
	repositoryURL  string
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
