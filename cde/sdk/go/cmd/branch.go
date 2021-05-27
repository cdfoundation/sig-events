package cmd

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cde "github.com/cdfoundation/sig-events/cde/pkg/cdf/events"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.AddCommand(branchCreatedCmd)
	branchCmd.AddCommand(branchDeletedCmd)

	branchCmd.PersistentFlags().StringVarP(&branchId, "id", "i", "", "Branch Id")
	branchCmd.PersistentFlags().StringVarP(&branchName, "name", "n", "", "Branch Name")
	branchCmd.PersistentFlags().StringVarP(&branchRepositoryId, "repoid", "r", "", "Branch Repository Id")
	branchCmd.PersistentFlags().StringToStringVarP(&branchData, "data", "d", map[string]string{}, "Branch Data")
}

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Emit Branch related Events",
	Long:  `Emit Branch related CloudEvent`,
}

var (
	branchId string
	branchName string
	branchRepositoryId string
	branchData map[string]string
)

var branchCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Branch Created Event",
	Long:  `Emit Branch Created CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateBranchEvent(cde.BranchCreatedEventV1, branchId,
			branchName, branchRepositoryId, branchData)

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


var branchDeletedCmd = &cobra.Command{
	Use:   "deleted",
	Short: "Emit Branch Deleted Event",
	Long:  `Emit Branch Deleted CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event, _ := cde.CreateBranchEvent(cde.BranchDeletedEventV1, branchId,
			branchName, branchRepositoryId, branchData)

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
