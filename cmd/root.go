/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"fmt"
	"os"

	"github.com/b-turchyn/diary/lib"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"go.uber.org/zap"
)

var cfgFile string
var databaseName string
var overrideDate string
var overrideTime string
var color bool
var logDebug bool

var logger *zap.Logger

var dateTimeArgsValidator = func(cmd *cobra.Command, args []string) error {
	if overrideDate != "" && overrideTime == "" {
		return errors.New("Time is required when a date is provided")
	}

	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: lib.VERSION,
	Use:     "diary",
	Short:   "Document your day-to-day activities",
	Long: `A lot goes on in your day-to-day life. This aims to be a simple
place to log all of that and be able to pull that information
back quickly.

This tracks what you're currently working on, what you learned,
and where you made a mistake.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Sugar().Error(err)
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.diary.yaml)")
	rootCmd.PersistentFlags().StringVar(&databaseName, "database", "~/diary.sqlite3", "Diary database (full path)")
	rootCmd.PersistentFlags().StringVarP(&overrideDate, "date", "d", "", "Use a specific date (--time is required with this)")
	rootCmd.PersistentFlags().StringVarP(&overrideTime, "time", "t", "", "Use a specific time")
	rootCmd.PersistentFlags().BoolVar(&color, "color", true, "Control colored output on TTY")
	rootCmd.PersistentFlags().BoolVar(&logDebug, "debug", false, "Write more verbose logs")

	viper.BindPFlag("database", rootCmd.PersistentFlags().Lookup("database"))
	viper.BindPFlag("color", rootCmd.PersistentFlags().Lookup("color"))
	viper.BindPFlag("log.debug", rootCmd.PersistentFlags().Lookup("debug"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".diary" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".diary")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	// Logger needs to be init'd before running logger, so it exists twice here
	if err := viper.ReadInConfig(); err == nil {
		initLogger()
		logger.Debug("Using config file:", zap.String("file", viper.ConfigFileUsed()))
	} else {
		initLogger()
	}
}

/**
 * Build logger from provided configuration
 */
func initLogger() {
	loggerConfig := zap.NewDevelopmentConfig()
	if viper.GetBool("log.debug") {
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	logger, _ = loggerConfig.Build()
	defer logger.Sync()
}
