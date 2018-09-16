package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var timeOut int
var post bool

var rootCmd = &cobra.Command{
	Use:   "385grader",
	Short: "385 programming quick grader.",
	Long:  "A cli executable that leverages docker to help grade Stevens 385 programming assignments easily and automatically by leveraging the canvas api. Source available at https://www.gitlab.com/gluaxspeed/385grader.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func defaultFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("token", "t", "", "Required: Canvas API Token.")
	cmd.PersistentFlags().StringP("cid", "c", "", "Required: Canvas Course ID.")
	cmd.PersistentFlags().StringP("aid", "a", "", "Required: Canvas Assignment ID.")
	cmd.PersistentFlags().StringP("test_script", "s", "", "Required: Path to shell test script.")
	cmd.PersistentFlags().StringP("entrypoint", "e", "", "Required: The entrypoint cpp file where they write their name and pledge.")
	cmd.PersistentFlags().IntVarP(&timeOut, "timeout", "o", 30, "Amount of time in seconds to run test script to completion.")
	cmd.PersistentFlags().BoolVarP(&post, "post", "p", false, "Whether or not to post the grade and comments to canvas.")
	cmd.PersistentFlags().StringP("valgrind", "g", "", "Path to valgrind input text file.")
}

func defaultArgCheck(cmd *cobra.Command, args []string) error {
	if cmd.Flag("token").Value.String() == "" {
		return errors.New("The argument ctoken is required")
	}
	if cmd.Flag("cid").Value.String() == "" {
		return errors.New("The argument cid is required")
	}

	if cmd.Flag("aid").Value.String() == "" {
		return errors.New("The argument aid is required")
	}

	if cmd.Flag("test_script").Value.String() == "" {
		return errors.New("The argument test_script is required")
	}

	if cmd.Flag("entrypoint").Value.String() == "" {
		return errors.New("The argument entrypoint is required")
	}

	return nil
}
