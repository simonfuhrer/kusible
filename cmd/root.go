/*
Copyright © 2019 Michael Gruener

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

package cmd

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "kusible",
	Short: "Render and deploy kubernetes resources",
	Long:  `This is a CLI tool to render and deploy kubernetes resources to multiple clusters`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("json-log") {
			log.SetFormatter(&log.JSONFormatter{})
		}

		logLevel, err := log.ParseLevel(viper.GetString("log-level"))

		if err != nil {
			log.Fatal(err.Error())
		}

		log.SetLevel(logLevel)
	},
}

func init() {
	viper.SetEnvPrefix(appName)
	dashReplacer := strings.NewReplacer("-", "_", ".", "_")
	viper.SetEnvKeyReplacer(dashReplacer)
	viper.AutomaticEnv()

	rootCmd.PersistentFlags().StringP("log-level", "", log.WarnLevel.String(), "log level (trace,debug,info,warn/warning,error,fatal,panic)")
	rootCmd.PersistentFlags().BoolP("json-log", "", false, "log as json")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Suppress all normal output")
	rootCmd.PersistentFlags().StringP("inventory", "i", "inventory.yml", "Path to the inventory")
	rootCmd.PersistentFlags().StringP("group-vars-dir", "d", "group_vars", "Source directory to read from")
	rootCmd.PersistentFlags().StringP("ejson-privkey", "k", "", "EJSON private key")
	rootCmd.PersistentFlags().StringP("ejson-key-dir", "", "/opt/ejson/keys", "Directory containing EJSON keys")
	rootCmd.PersistentFlags().BoolP("skip-eval", "", false, "Skip spruce operator evaluation")
	rootCmd.PersistentFlags().StringSliceP("limit", "l", []string{}, "Limit selected groups")

	viper.BindPFlag("group-vars-dir", rootCmd.PersistentFlags().Lookup("group-vars-dir"))
	viper.BindPFlag("ejson-privkey", rootCmd.PersistentFlags().Lookup("ejson-privkey"))
	viper.BindPFlag("ejson-key-dir", rootCmd.PersistentFlags().Lookup("ejson-key-dir"))
	viper.BindPFlag("skip-eval", rootCmd.PersistentFlags().Lookup("skip-eval"))
	viper.BindPFlag("inventory", rootCmd.PersistentFlags().Lookup("inventory"))
	viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("json-log", rootCmd.PersistentFlags().Lookup("json-log"))
	viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))
	viper.BindPFlag("limit", rootCmd.PersistentFlags().Lookup("limit"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//func main() {
//	execute()
//}
