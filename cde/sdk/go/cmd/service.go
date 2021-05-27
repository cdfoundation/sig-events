package cmd

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cde "github.com/cdfoundation/sig-events/cde/pkg/cdf/events"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.AddCommand(serviceDeployedCmd)
	serviceCmd.AddCommand(serviceUpgradedCmd)
	serviceCmd.AddCommand(serviceRolledbackCmd)
	serviceCmd.AddCommand(serviceRemovedCmd)

	serviceCmd.PersistentFlags().StringVarP(&serviceEnvId, "envId", "e", "", "Environment Id where the Service is running")
	serviceCmd.PersistentFlags().StringVarP(&serviceName, "name", "n", "", "Service's Name")
	serviceCmd.PersistentFlags().StringVarP(&serviceVersion, "version", "v", "", "Service's Version")
	serviceCmd.PersistentFlags().StringToStringVarP(&serviceData, "data", "d", map[string]string{}, "Service's Data")
}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Emit Environment related Events",
	Long:  `Emit Environment related CloudEvent`,
}

var (
	serviceEnvId string
	serviceName string
	serviceVersion string
	serviceData map[string]string
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
		event, _ := cde.CreateServiceEvent(cde.ServiceDeployedEventV1, serviceEnvId,
			serviceName, serviceVersion, serviceData)

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
		event, _ := cde.CreateServiceEvent(cde.ServiceUpgradedEventV1, serviceEnvId,
			serviceName, serviceVersion, serviceData)

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
		event, _ := cde.CreateServiceEvent(cde.ServiceRemovedEventV1, serviceEnvId,
			serviceName, serviceVersion, serviceData)

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
		event, _ := cde.CreateServiceEvent(cde.ServiceRolledbackEventV1, serviceEnvId,
			serviceName, serviceVersion, serviceData)

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
