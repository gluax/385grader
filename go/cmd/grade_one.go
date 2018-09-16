package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"385grader/canvas"
)

var gradeOneCmd = &cobra.Command{
	Use:   "gradeOne",
	Short: "Grades a user's submitted assignment.",
	Long:  "Grades the latest submission of a specific person who has submitted, generates comments and grades automatically on canvas.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := defaultArgCheck(cmd, args); err != nil {
			return err
		}

		if cmd.Flag("uid").Value.String() == "" {
			return errors.New("The argument uid is required")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		token := cmd.Flag("token").Value.String()
		cid := cmd.Flag("cid").Value.String()
		aid := cmd.Flag("aid").Value.String()
		uid := cmd.Flag("uid").Value.String()
		testScript := cmd.Flag("test_script").Value.String()
		entrypoint := cmd.Flag("entrypoint").Value.String()

		sub := canvas.FetchOne(cid, aid, uid, token)
		canvas.GradeOneSubmission(entrypoint, testScript, sub, timeOut, post)
	},
}

func init() {
	defaultFlags(gradeOneCmd)
	gradeOneCmd.PersistentFlags().StringP("uid", "u", "", "Canvas User ID.")

	rootCmd.AddCommand(gradeOneCmd)
}
