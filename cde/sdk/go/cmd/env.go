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
	envId      string
	envName    string
	envRepoUrl string
	envData    map[string]string
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
