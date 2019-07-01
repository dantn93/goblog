package main

import (
	"flag"
	"fmt"

	"github.com/goblog/accountservice/config"
	"github.com/goblog/accountservice/dbclient"
	"github.com/goblog/accountservice/service"
	"github.com/spf13/viper"
)

var appName = "accountservice"

// Init function, runs before main()
func init() {
	// Read command line flags
	profile := flag.String("profile", "test", "Environment profile, something similar to spring profiles")
	configServerUrl := flag.String("configServerUrl", "http://configserver:8888", "Address to config server")
	configBranch := flag.String("configBranch", "master", "git branch to fetch configuration from")
	flag.Parse()
	// Pass the flag values into viper.
	viper.Set("profile", *profile)
	viper.Set("configServerUrl", *configServerUrl)
	viper.Set("configBranch", *configBranch)
}
func main() {
	fmt.Printf("Starting %v\n", appName)
	// NEW - load the config
	config.LoadConfigurationFromBranch(
		viper.GetString("configServerUrl"),
		appName,
		viper.GetString("profile"),
		viper.GetString("configBranch"))
	initializeBoltClient()
	service.StartWebServer(viper.GetString("server_port")) // NEW, use port from loaded config
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
