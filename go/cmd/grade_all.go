package cmd

import (
	"github.com/spf13/cobra"

	"385grader/canvas"
)

var gradeAllCmd = &cobra.Command{
	Use:     "gradeAll",
	Short:   "Grades all submitted assignments.",
	Long:    "Grades the latest submission of each person who has submitted, generates comments and grades automatically on canvas.",
	PreRunE: defaultArgCheck,
	RunE: func(cmd *cobra.Command, args []string) error {
		token := cmd.Flag("token").Value.String()
		cid := cmd.Flag("cid").Value.String()
		aid := cmd.Flag("aid").Value.String()
		testScript := cmd.Flag("test_script").Value.String()
		entrypoint := cmd.Flag("entrypoint").Value.String()
		executable := cmd.Flag("executable").Value.String()
		valgrind := cmd.Flag("valgrind").Value.String()

		subs := canvas.FetchAllAssignmentUrls(cid, aid, token)
		canvas.GradeAllSubmissions(entrypoint, testScript, executable, valgrind, subs, timeOut, post, view)

		return nil
	},
}

func init() {
	defaultFlags(gradeAllCmd)
	rootCmd.AddCommand(gradeAllCmd)
}
