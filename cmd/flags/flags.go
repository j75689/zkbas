package flags

import (
	"github.com/urfave/cli/v2"
)

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Aliases: []string{"f"},
		Usage:   "the config file",
	}
	ContractAddrFlag = &cli.StringFlag{
		Name:  "contractAddr",
		Usage: "the contract addresses file",
	}
	DSNFlag = &cli.StringFlag{
		Name:  "dsn",
		Usage: "data source name",
	}
	BSCTestNetworkRPCFlag = &cli.StringFlag{
		Name:  "testnet",
		Value: "https://data-seed-prebsc-1-s1.binance.org:8545/",
		Usage: "the rpc endpoint of bsc testnet",
	}
	LocalTestNetworkRPCFlag = &cli.StringFlag{
		Name:  "local",
		Value: "http://127.0.0.1:8545/",
		Usage: "the rpc endpoint of local net",
	}
	BlockHeightFlag = &cli.Int64Flag{
		Name:  "height",
		Usage: "block height",
	}
	ServiceNameFlag = &cli.StringFlag{
		Name:  "service",
		Usage: "service name(committer, witness)",
	}
	BatchSizeFlag = &cli.IntFlag{
		Name:  "batch",
		Value: 1000,
		Usage: "batch size for reading history record from the database",
	}
	PProfFlag = &cli.BoolFlag{
		Name:  "pprof",
		Value: false,
		Usage: "Enable the pprof HTTP server",
	}
	PProfPortFlag = &cli.IntFlag{
		Name:  "pprof.port",
		Usage: "pprof HTTP server listening port",
		Value: 6060,
	}
	PProfAddrFlag = &cli.StringFlag{
		Name:  "pprof.addr",
		Usage: "pprof HTTP server listening interface",
		Value: "127.0.0.1",
	}
)
