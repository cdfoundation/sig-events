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
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.AddCommand(serviceDeployedCmd)
	serviceCmd.AddCommand(serviceUpgradedCmd)
	serviceCmd.AddCommand(serviceRolledbackCmd)
	serviceCmd.AddCommand(serviceRemovedCmd)
	serviceCmd.AddCommand(serviceCreatedCmd)
	serviceCmd.AddCommand(servicePublishedCmd)

	serviceCmd.PersistentFlags().StringVarP(&serviceEnvID, "envId", "e", "", "Environment Id where the Service is running")
	serviceCmd.PersistentFlags().StringVarP(&serviceName, "name", "n", "", "Service's Name")
	serviceCmd.PersistentFlags().StringVarP(&serviceVersion, "version", "v", "", "Service's Version")
	serviceCmd.PersistentFlags().StringVar(&serviceNamespace, "namespace", "", "Service's Namespace")
	serviceCmd.PersistentFlags().StringVarP(&serviceActiveRevisionName, "revision", "r", "", "Service's active revision")
	serviceCmd.PersistentFlags().StringToStringVarP(&serviceData, "data", "d", map[string]string{}, "Service's Data")
}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Emit Environment related Events",
	Long:  `Emit Environment related CloudEvent`,
}

var (
	serviceEnvID              string
	serviceName               string
	serviceVersion            string
	serviceNamespace          string
	serviceActiveRevisionName string
	serviceData               map[string]string
)

var serviceDeployedCmd = &cobra.Command{
	Use:   "deployed",
	Short: "Emit Service Deployed Event",
	Long:  `Emit Service Deployed CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateServiceEvent(cde.ServiceDeployedEventV1, cde.ServiceEventParams{
			ServiceEnvID:              serviceEnvID,
			ServiceVersion:            serviceVersion,
			ServiceName:               serviceName,
			ServiceNamespace:          serviceNamespace,
			ServiceActiveRevisionName: serviceActiveRevisionName,
			ServiceData:               serviceData,
		})

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

var serviceUpgradedCmd = &cobra.Command{
	Use:   "upgraded",
	Short: "Emit Service Upgraded Event",
	Long:  `Emit Service Upgraded CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateServiceEvent(cde.ServiceUpgradedEventV1, cde.ServiceEventParams{
			ServiceEnvID:              serviceEnvID,
			ServiceVersion:            serviceVersion,
			ServiceName:               serviceName,
			ServiceNamespace:          serviceNamespace,
			ServiceActiveRevisionName: serviceActiveRevisionName,
			ServiceData:               serviceData,
		})

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

var serviceRemovedCmd = &cobra.Command{
	Use:   "removed",
	Short: "Emit Service Removed Event",
	Long:  `Emit Service Removed CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateServiceEvent(cde.ServiceRemovedEventV1, cde.ServiceEventParams{
			ServiceEnvID:              serviceEnvID,
			ServiceVersion:            serviceVersion,
			ServiceName:               serviceName,
			ServiceNamespace:          serviceNamespace,
			ServiceActiveRevisionName: serviceActiveRevisionName,
			ServiceData:               serviceData,
		})

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

var serviceRolledbackCmd = &cobra.Command{
	Use:   "rolledback",
	Short: "Emit Service Rolledback Event",
	Long:  `Emit Service Rolledback CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateServiceEvent(cde.ServiceRolledbackEventV1, cde.ServiceEventParams{
			ServiceEnvID:              serviceEnvID,
			ServiceVersion:            serviceVersion,
			ServiceName:               serviceName,
			ServiceNamespace:          serviceNamespace,
			ServiceActiveRevisionName: serviceActiveRevisionName,
			ServiceData:               serviceData,
		})

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

var serviceCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Service created Event",
	Long:  `Emit Service created CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateServiceEvent(cde.ServiceCreatedV1, cde.ServiceEventParams{
			ServiceEnvID:              serviceEnvID,
			ServiceVersion:            serviceVersion,
			ServiceName:               serviceName,
			ServiceNamespace:          serviceNamespace,
			ServiceActiveRevisionName: serviceActiveRevisionName,
			ServiceData:               serviceData,
		})

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

var servicePublishedCmd = &cobra.Command{
	Use:   "published",
	Short: "Emit Service published Event",
	Long:  `Emit Service published CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateServiceEvent(cde.ServicePublishedV1, cde.ServiceEventParams{
			ServiceEnvID:              serviceEnvID,
			ServiceVersion:            serviceVersion,
			ServiceName:               serviceName,
			ServiceNamespace:          serviceNamespace,
			ServiceActiveRevisionName: serviceActiveRevisionName,
			ServiceData:               serviceData,
		})

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
