/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"github.com/mitchellh/go-homedir"
	"github.com/ocoderr/bsc-tool/config"
	"github.com/ocoderr/bsc-tool/data"
	"github.com/ocoderr/bsc-tool/screens"
	"github.com/ocoderr/bsc-tool/utils"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func main() {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configFile := filepath.Join(home, "pancake-conf.yaml")
	if !Exist(configFile) {
		config.NewConfig(configFile)
	}
	viper.AddConfigPath(home)
	viper.SetConfigName("pancake-conf")

	viper.SetDefault("language", "EN")
	//viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
		if err := viper.Unmarshal(&config.CF); err != nil {
			panic(err)
		}
	}
	data.Client = utils.EstimateClient("bsc")
	go data.StartRefreshPrices()

	a := app.NewWithID("com.openwallet.pancakeswap")
	a.SetIcon(resourceLogoPng)
	//a.Settings().SetTheme(theme.LightTheme())
	a.Settings().SetTheme(&MyTheme{})
	screens.Master(a)

}
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
