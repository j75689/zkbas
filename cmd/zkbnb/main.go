package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/bnb-chain/zkbnb/cmd/flags"
	"github.com/bnb-chain/zkbnb/service/apiserver"
	"github.com/bnb-chain/zkbnb/service/committer"
	"github.com/bnb-chain/zkbnb/service/monitor"
	"github.com/bnb-chain/zkbnb/service/prover"
	"github.com/bnb-chain/zkbnb/service/sender"
	"github.com/bnb-chain/zkbnb/service/witness"
	"github.com/bnb-chain/zkbnb/tools/dbinitializer"
	"github.com/bnb-chain/zkbnb/tools/recovery"

	"net/http"
	pprof "net/http/pprof"
)

// Build Info (set via linker flags)
var (
	gitCommit = "unknown"
	gitDate   = "unknown"
	version   = "unknown"
)

func main() {
	cli.VersionPrinter = func(ctx *cli.Context) {
		fmt.Println("Version:", ctx.App.Version)
		fmt.Println("Git Commit:", gitCommit)
		fmt.Println("Git Commit Date:", gitDate)
		fmt.Println("Architecture:", runtime.GOARCH)
		fmt.Println("Go Version:", runtime.Version())
		fmt.Println("Operating System:", runtime.GOOS)
	}

	app := &cli.App{
		Name:        "ZkBNB",
		HelpName:    "zkbnb",
		Version:     version,
		Description: "ZkRollup BNB Application Side Chain",
		Commands: []*cli.Command{
			// services
			{
				Name:  "prover",
				Usage: "Run prover service",
				Flags: []cli.Flag{
					flags.ConfigFlag,
					flags.MetricsEnabledFlag,
					flags.MetricsHTTPFlag,
					flags.MetricsPortFlag,
					flags.PProfEnabledFlag,
					flags.PProfAddrFlag,
					flags.PProfPortFlag,
				},
				Action: func(cCtx *cli.Context) error {
					if !cCtx.IsSet(flags.ConfigFlag.Name) {
						return cli.ShowSubcommandHelp(cCtx)
					}
					startMetricsServer(cCtx)
					return prover.Run(cCtx.String(flags.ConfigFlag.Name))
				},
			},
			{
				Name:  "witness",
				Usage: "Run witness service",
				Flags: []cli.Flag{
					flags.ConfigFlag,
					flags.MetricsEnabledFlag,
					flags.MetricsHTTPFlag,
					flags.MetricsPortFlag,
					flags.PProfEnabledFlag,
					flags.PProfAddrFlag,
					flags.PProfPortFlag,
				},
				Action: func(cCtx *cli.Context) error {
					if !cCtx.IsSet(flags.ConfigFlag.Name) {
						return cli.ShowSubcommandHelp(cCtx)
					}
					startMetricsServer(cCtx)
					return witness.Run(cCtx.String(flags.ConfigFlag.Name))
				},
			},
			{
				Name:  "monitor",
				Usage: "Run monitor service",
				Flags: []cli.Flag{
					flags.ConfigFlag,
					flags.MetricsEnabledFlag,
					flags.MetricsHTTPFlag,
					flags.MetricsPortFlag,
					flags.PProfEnabledFlag,
					flags.PProfAddrFlag,
					flags.PProfPortFlag,
				},
				Action: func(cCtx *cli.Context) error {
					if !cCtx.IsSet(flags.ConfigFlag.Name) {
						return cli.ShowSubcommandHelp(cCtx)
					}
					startMetricsServer(cCtx)
					return monitor.Run(cCtx.String(flags.ConfigFlag.Name))
				},
			},
			{
				Name: "committer",
				Flags: []cli.Flag{
					flags.ConfigFlag,
					flags.MetricsEnabledFlag,
					flags.MetricsHTTPFlag,
					flags.MetricsPortFlag,
					flags.PProfEnabledFlag,
					flags.PProfAddrFlag,
					flags.PProfPortFlag,
				},
				Usage: "Run committer service",
				Action: func(cCtx *cli.Context) error {
					if !cCtx.IsSet(flags.ConfigFlag.Name) {
						return cli.ShowSubcommandHelp(cCtx)
					}
					startMetricsServer(cCtx)
					return committer.Run(cCtx.String(flags.ConfigFlag.Name))
				},
			},
			{
				Name:  "sender",
				Usage: "Run sender service",
				Flags: []cli.Flag{
					flags.ConfigFlag,
					flags.MetricsEnabledFlag,
					flags.MetricsHTTPFlag,
					flags.MetricsPortFlag,
					flags.PProfEnabledFlag,
					flags.PProfAddrFlag,
					flags.PProfPortFlag,
				},
				Action: func(cCtx *cli.Context) error {
					if !cCtx.IsSet(flags.ConfigFlag.Name) {
						return cli.ShowSubcommandHelp(cCtx)
					}
					startMetricsServer(cCtx)
					return sender.Run(cCtx.String(flags.ConfigFlag.Name))
				},
			},
			{
				Name:  "apiserver",
				Usage: "Run apiserver service",
				Flags: []cli.Flag{
					flags.ConfigFlag,
					flags.MetricsEnabledFlag,
					flags.MetricsHTTPFlag,
					flags.MetricsPortFlag,
					flags.PProfEnabledFlag,
					flags.PProfAddrFlag,
					flags.PProfPortFlag,
				},
				Action: func(cCtx *cli.Context) error {
					if !cCtx.IsSet(flags.ConfigFlag.Name) {
						return cli.ShowSubcommandHelp(cCtx)
					}
					startMetricsServer(cCtx)
					return apiserver.Run(cCtx.String(flags.ConfigFlag.Name))
				},
			},
			// tools
			{
				Name:  "db",
				Usage: "Database tools",
				Subcommands: []*cli.Command{
					{
						Name:  "initialize",
						Usage: "Initialize DB tables",
						Flags: []cli.Flag{
							flags.ContractAddrFlag,
							flags.DSNFlag,
							flags.BSCTestNetworkRPCFlag,
							flags.LocalTestNetworkRPCFlag,
						},
						Action: func(cCtx *cli.Context) error {
							if !cCtx.IsSet(flags.ContractAddrFlag.Name) ||
								!cCtx.IsSet(flags.DSNFlag.Name) {
								return cli.ShowSubcommandHelp(cCtx)
							}

							return dbinitializer.Initialize(
								cCtx.String(flags.DSNFlag.Name),
								cCtx.String(flags.ContractAddrFlag.Name),
								cCtx.String(flags.BSCTestNetworkRPCFlag.Name),
								cCtx.String(flags.LocalTestNetworkRPCFlag.Name),
							)
						},
					},
				},
			},
			{
				Name:  "tree",
				Usage: "TreeDB tools",
				Subcommands: []*cli.Command{
					{
						Name:  "recovery",
						Usage: "Recovery treedb from the database",
						Flags: []cli.Flag{
							flags.ConfigFlag,
							flags.BlockHeightFlag,
							flags.ServiceNameFlag,
							flags.BatchSizeFlag,
						},
						Action: func(cCtx *cli.Context) error {
							if !cCtx.IsSet(flags.ServiceNameFlag.Name) ||
								!cCtx.IsSet(flags.BlockHeightFlag.Name) ||
								!cCtx.IsSet(flags.ConfigFlag.Name) {
								return cli.ShowSubcommandHelp(cCtx)
							}
							recovery.RecoveryTreeDB(
								cCtx.String(flags.ConfigFlag.Name),
								cCtx.Int64(flags.BlockHeightFlag.Name),
								cCtx.String(flags.ServiceNameFlag.Name),
								cCtx.Int(flags.BatchSizeFlag.Name),
							)
							return nil
						},
					},
				},
			},
		},
	}
	go func() {
		fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func startMetricsServer(ctx *cli.Context) {
	if !ctx.Bool(flags.PProfEnabledFlag.Name) && !ctx.Bool(flags.MetricsEnabledFlag.Name) {
		return
	}

	httpServer := http.NewServeMux()

	if ctx.Bool(flags.PProfEnabledFlag.Name) {
		httpServer.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
		httpServer.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
		httpServer.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		httpServer.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
		httpServer.Handle("/debug/pprof/block", pprof.Handler("block"))
		httpServer.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
	}
	if ctx.Bool(flags.MetricsEnabledFlag.Name) {
		httpServer.Handle("/debug/metrics", promhttp.Handler())
	}

	pprofAddress := fmt.Sprintf("%s:%d",
		ctx.String(flags.PProfAddrFlag.Name),
		ctx.Int(flags.PProfPortFlag.Name))
	metricsAddress := fmt.Sprintf("%s:%d",
		ctx.String(flags.MetricsHTTPFlag.Name),
		ctx.Int(flags.MetricsPortFlag.Name))

	if ctx.Bool(flags.PProfEnabledFlag.Name) && ctx.Bool(flags.MetricsEnabledFlag.Name) && metricsAddress == pprofAddress {
		go func() {
			logx.Info("Starting pprof server", "addr", fmt.Sprintf("http://%s/debug/pprof", pprofAddress))
			logx.Info("Starting metrics server", "addr", fmt.Sprintf("http://%s/debug/metrics", metricsAddress))
			if err := http.ListenAndServe(pprofAddress, httpServer); err != nil {
				logx.Error("Failure in running pprof and metrics server", "err", err)
			}
		}()
		return
	}

	if ctx.Bool(flags.MetricsEnabledFlag.Name) {
		go func() {
			logx.Info("Starting pprof server", "addr", fmt.Sprintf("http://%s/debug/pprof", pprofAddress))
			if err := http.ListenAndServe(pprofAddress, httpServer); err != nil {
				logx.Error("Failure in running pprof server", "err", err)
			}
		}()
	}

	if ctx.Bool(flags.MetricsEnabledFlag.Name) {
		go func() {
			logx.Info("Starting metrics server", "addr", fmt.Sprintf("http://%s/debug/metrics", metricsAddress))
			if err := http.ListenAndServe(metricsAddress, httpServer); err != nil {
				logx.Error("Failure in running metrics server", "err", err)
			}
		}()
	}

}
