package main

import (
	"fmt"
	"github.com/sosedoff/gitkit"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	fmt.Printf("Default repo dir is: %s\n", viper.GetString("repoDir"))

	// Configure git hooks
	hooks := &gitkit.HookScripts{
		PreReceive: `echo "Hello World!"`,
	}

	// Configure git service
	service := gitkit.New(gitkit.Config{
		Dir:        viper.GetString("repoDir"),
		AutoCreate: false,
		AutoHooks:  true,
		Hooks:      hooks,
	})

	// Configure git server. Will create git repos path if it does not exist.
	// If hooks are set, it will also update all repos with new version of hook scripts.
	if err := service.Setup(); err != nil {
		log.Fatal(err)
	}

	http.Handle("/", service)

	// Start HTTP server
	if err := http.ListenAndServe(":"+viper.GetString("port"), nil); err != nil {
		log.Fatal(err)
	}
}
