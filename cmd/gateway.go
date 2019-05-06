package cmd

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/owenliu1122/grpc-gateway-example/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// gatewayCmd represents the server command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Start up proxy server",
	Run:   gatewayProc,
}

func gatewayProc(cmd *cobra.Command, args []string) {

	// opts := loadApplicationOptions()

	logger := log.New()

	logger.Info("Start gatewayProc")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx, mux, "127.0.0.1:30001", opts)
	handleInitError("gateway proxy", err)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		handleInitError("ListenAndServe", err)
	}
}

func init() {
	rootCmd.AddCommand(gatewayCmd)
}
