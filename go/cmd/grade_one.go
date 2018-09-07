package cmd

import (
	"errors"
	"fmt"
	
	"github.com/spf13/cobra"
)

var gradeOneCmd = &cobra.Command{
	Use: "gradeOne",
	Short: "Grades a user's submitted assignment.",
	Long: "Grades the latest submission of a specific person who has submitted, generates comments and grades automatically on canvas.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flag("ctoken").Value.String() == "" {
			return errors.New("The argument ctoken is required.")
		}
		if cmd.Flag("cid").Value.String() == "" {
			return errors.New("The argument cid is required.")
		}

		if cmd.Flag("aid").Value.String() == "" {
			return errors.New("The argument aid is required.")
		}

		if cmd.Flag("uid").Value.String() == "" {
			return errors.New("The argument uid is required.")
		}
		
		return nil
	},
	Run: func(cmd *cobra.Command, args []string){
		fmt.Printf("%s, %s, %s", cmd.Flag("ctoken").Value.String(), cmd.Flag("cid").Value.String(), cmd.Flag("aid").Value.String())
	},
}

func init() {
	gradeOneCmd.PersistentFlags().String("uid", "", "Canvas User ID.")
		
	rootCmd.AddCommand(gradeOneCmd)
}
