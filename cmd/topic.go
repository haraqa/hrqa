package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func topicFlag() (string, string, string, string) {
	return "topic", "t", "", "topic to create (required)"
}

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Manage topics",
	Long:  `"hrqa topic" requires exactly 1 argument.`,
}

func init() {
	topicListCmd.Flags().String("prefix", "", "prefix to match topics to")
	topicListCmd.Flags().String("suffix", "", "suffix to match topics to")
	topicListCmd.Flags().StringP("regex", "r", "", "regexp to match topics to")

	topicDeleteCmd.Flags().StringP(topicFlag())
	must(topicDeleteCmd.MarkFlagRequired("topic"))

	topicCreateCmd.Flags().StringP(topicFlag())
	must(topicCreateCmd.MarkFlagRequired("topic"))

	topicOffsetsCmd.Flags().StringP(topicFlag())
	must(topicOffsetsCmd.MarkFlagRequired("topic"))

	topicCmd.AddCommand(topicListCmd, topicCreateCmd, topicDeleteCmd, topicOffsetsCmd)
	rootCmd.AddCommand(topicCmd)
}

// topicCreateCmd represents the create command
var topicCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new topic",
	Example: `  hrqa topic create -t hello`,
	Long:    `Create a new topic.`,
	Run: func(cmd *cobra.Command, args []string) {
		vfmt := newVerbose(cmd)

		// ge† flags
		topic, err := cmd.Flags().GetString("topic")
		must(err)

		// setup client connection
		client := newConnection(cmd, vfmt)
		defer client.Close()

		err = client.CreateTopic(context.Background(), []byte(topic))
		if err != nil {
			fmt.Printf("Unable to create topic %q: %q\n", topic, err.Error())
			os.Exit(1)
		}
	},
}

// topicDeleteCmd represents the delete command
var topicDeleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a topic",
	Example: `  hrqa topic delete -t hello`,
	Long:    `Delete a topic.`,
	Run: func(cmd *cobra.Command, args []string) {
		vfmt := newVerbose(cmd)

		// ge† flags
		topic, err := cmd.Flags().GetString("topic")
		must(err)

		// setup client connection
		client := newConnection(cmd, vfmt)
		defer client.Close()

		vfmt.Printf("Deleting topic %q\n", topic)
		err = client.DeleteTopic(context.Background(), []byte(topic))
		if err != nil {
			fmt.Printf("Unable to delete topic %q: %q\n", topic, err.Error())
			os.Exit(1)
		}
		vfmt.Printf("Deleted topic %q\n", topic)
	},
}

// topicOffsetsCmd represents the create command
var topicOffsetsCmd = &cobra.Command{
	Use:     "offsets",
	Short:   "Get the min and max offsets of a topic",
	Example: `  hrqa topic offsets -t hello`,
	Long:    `Get the min and max offsets of a topic.`,
	Run: func(cmd *cobra.Command, args []string) {
		vfmt := newVerbose(cmd)

		// ge† flags
		topic, err := cmd.Flags().GetString("topic")
		must(err)

		// setup client connection
		client := newConnection(cmd, vfmt)
		defer client.Close()

		min, max, err := client.Offsets(context.Background(), []byte(topic))
		if err != nil {
			fmt.Printf("Unable to get topic offsets for %q: %q\n", topic, err.Error())
			os.Exit(1)
		}
		fmt.Printf("min: %d, max: %d\n", min, max)
	},
}

// topicListCmd represents the list command
var topicListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all topics",
	Example: `  hrqa topic list`,
	Long:    `List all topics.`,
	Run: func(cmd *cobra.Command, args []string) {
		vfmt := newVerbose(cmd)

		// setup client connection
		client := newConnection(cmd, vfmt)
		defer client.Close()

		prefix, err := cmd.Flags().GetString("prefix")
		must(err)
		suffix, err := cmd.Flags().GetString("suffix")
		must(err)
		regex, err := cmd.Flags().GetString("regex")
		must(err)

		topics, err := client.ListTopics(context.Background(), prefix, suffix, regex)
		if err != nil {
			fmt.Printf("Unable to list topics: %q\n", err.Error())
			os.Exit(1)
		}
		if len(topics) == 0 {
			fmt.Println("No topics found.")
			os.Exit(0)
		}
		fmt.Println("Found topics:")
		for i := range topics {
			fmt.Printf("\t%s\n", topics[i])
		}
	},
}
