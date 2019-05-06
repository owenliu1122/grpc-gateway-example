package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/owenliu1122/grpc-gateway-example/pb"

	"github.com/owenliu1122/grpc-gateway-example/client"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var clientCmdHost string

// serverCmd represents the server command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start up message notification client",
	Run:   clientProc,
}

func clientProc(cmd *cobra.Command, args []string) {

	log.Debug("Start clientProc")
	opts := loadApplicationOptions()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cli, err := client.NewClient(opts.CLI)
	handleInitError("client", err)

	// go func(ctx context.Context) {
	for {
		// select {
		// case <-ctx.Done():
		// 	return
		// default:
		req := &pb.HelloWorldReq{}
		_, _ = fmt.Scanln(&req.Name)

		resp, err := cli.HelloWorld(ctx, req)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("resp: %v\n", resp.Echo)
		time.Sleep(1 * time.Second)
		// }

	}
	// }(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Debug("Exit clientProc")
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
