package cmd

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/owenliu1122/grpc-gateway-example/pb"

	"github.com/owenliu1122/grpc-gateway-example/internal/controllers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start up grpc entrypoint server",
	Run:   serverProc,
}

func serverProc(cmd *cobra.Command, args []string) {

	opts := loadApplicationOptions()

	logger := log.New()

	logger.Info("Start serverProc")

	// grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.API.Port))
	handleInitError("net", err)

	gs := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time: 10 * time.Minute,
		}),
	)

	defer gs.GracefulStop()

	ctl := controllers.NewAPIServer(opts.API)
	pb.RegisterHelloWorldServer(gs, ctl)
	go gs.Serve(lis)

	ip, _ := getIntranetIP()
	logger.Infof("server started: %s:%d", ip, opts.API.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("Exit serverProc")
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
