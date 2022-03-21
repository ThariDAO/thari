package main

import (
	"github.com/ThariCommunity/thari/cmd/tharid/cmd"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	thari "github.com/ThariCommunity/thari/app"
	thariParams "github.com/ThariCommunity/thari/app/params"
)

func main() {
	thariParams.SetAddressPrefixes()

	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, thari.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
