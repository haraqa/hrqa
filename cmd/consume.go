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
	Short: "Consume messages from a topic",
	Long:  `Consume messages from a topic`,
	Run: func(cmd *cobra.Command, args []string) {
		vfmt := newVerbose(cmd)

		// ge† flags
		topic, err := cmd.Flags().GetString("topic")
		must(err)
		offset, err := cmd.Flags().GetInt64("offset")
		must(err)
		limit, err := cmd.Flags().GetInt64("limit")
		must(err)
		follow, err := cmd.Flags().GetBool("follow")
		must(err)

		// setup client connection
		client := newConnection(cmd, vfmt)
		defer client.Close()

		// start watching the topic
		ctx := context.Background()
		var watcher *haraqa.Watcher
		if follow {
			if offset < 0 {
				offset = 0
			}

			watcher, err = client.NewWatcher(ctx, []byte(topic))
			if err != nil {
				fmt.Printf("Unable to watch topic %q: %q", topic, err.Error())
				os.Exit(1)
			}
			defer watcher.Close()
		}

		buf := haraqa.NewConsumeBuffer()
		for {
			// send consume message
			vfmt.Printf("Consuming from the topic %q\n", topic)
			msgs, err := client.Consume(ctx, []byte(topic), offset, limit, buf)
			if err != nil {
				fmt.Printf("Unable to consume message(s) from %q: %q\n", topic, err.Error())
				os.Exit(1)
			}

			// print messages to stdout
			for _, msg := range msgs {
				fmt.Println(string(msg))
			}
			if !follow {
				break
			}
			offset += int64(len(msgs))
			if len(msgs) == 0 {
				// wait for a watch event
				<-watcher.Events()
			}
		}
	},
}

func init() {
	consumeCmd.Flags().StringP(topicFlag())
	must(consumeCmd.MarkFlagRequired("topic"))
	consumeCmd.Flags().Int64P("offset", "o", -1, "offset to consume from, -1 for message from the last available offset")
	consumeCmd.Flags().Int64P("limit", "l", 100, "maximum number of messages to consume per consume call")
	consumeCmd.Flags().BoolP("follow", "f", false, "follow the topic, continuously consume until ctrl+c is called")
	rootCmd.AddCommand(consumeCmd)
}
