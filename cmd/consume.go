package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/haraqa/haraqa"
	"github.com/spf13/cobra"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		vfmt := newVerbose(cmd)

		// ge† flags
		topic, err := cmd.Flags().GetString("topic")
		must(err)
		offset, err := cmd.Flags().GetInt64("offset")
		must(err)
		maxBatchSize, err := cmd.Flags().GetInt64("max")
		must(err)

		// setup client connection
		client := newConnection(cmd, vfmt)
		defer client.Close()

		// send consume message
		ctx := context.Background()
		buf := haraqa.NewConsumeBuffer()

		vfmt.Printf("Consuming from the topic %q\n", topic)
		msgs, err := client.Consume(ctx, []byte(topic), offset, maxBatchSize, buf)
		if err != nil {
			fmt.Printf("Unable to consume message(s) from %q: %q\n", topic, err.Error())
			client.Close()
			os.Exit(1)
		}

		// print messages to stdout
		for _, msg := range msgs {
			fmt.Println(string(msg))
		}
	},
}

func init() {
	consumeCmd.Flags().StringP(topicFlag())
	must(consumeCmd.MarkFlagRequired("topic"))
	consumeCmd.Flags().Int64P("offset", "o", -1, "offset to consume from, -1 for message from the last available offset")
	consumeCmd.Flags().Int64P("max", "m", 100, "maximum number of messages to consume")
	rootCmd.AddCommand(consumeCmd)
}
