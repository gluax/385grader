package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "385grader",
	Short:   "385 programming quick grader.",
	Long:    "A cli executable that leverages docker to help grade Stevens 385 programming assignments easily and automatically by leveraging the canvas api. Source available at https://www.gitlab.com/gluaxspeed/385grader.",
	PreRunE: defaultArgCheck,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s, %s, %s", cmd.Flag("ctoken").Value.String(), cmd.Flag("cid").Value.String(), cmd.Flag("aid").Value.String())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("ctoken", "", "Required: Canvas API Token.")
	rootCmd.PersistentFlags().String("cid", "", "Required: Canvas Course ID.")
	rootCmd.PersistentFlags().String("aid", "", "Required: Canvas Assignment ID.")
	rootCmd.PersistentFlags().String("test_script", "", "Path to shell test script.")
	rootCmd.PersistentFlags().String("valgrind", "", "Path to valgrind input text file.")
}

func defaultArgCheck(cmd *cobra.Command, args []string) error {
	if cmd.Flag("ctoken").Value.String() == "" {
		return errors.New("The argument ctoken is required")
	}
	if cmd.Flag("cid").Value.String() == "" {
		return errors.New("The argument cid is required")
	}

	if cmd.Flag("aid").Value.String() == "" {
		return errors.New("The argument aid is required")
	}

	return nil
}
