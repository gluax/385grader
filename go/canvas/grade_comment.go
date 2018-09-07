package canvas

import (
	"bytes"
	"path/filepath"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"strconv"

	"385grader/utils"
)

func analyzeTestResults(testoutput string, missed bool, late int) (float64, string) {
	if missed {
		return 0, "No Submission."
	}
	
	split := strings.Split(testoutput, "\n")
	if strings.Index(split[0], "done") == -1 {
		return 0, "Compilation Failed."
	}

	run_loc := strings.Index(testoutput, "run:")
	num_tests, err := strconv.Atoi(testoutput[run_loc + 5: run_loc + 7])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	var total_score float64
	total_score = 100
	var comment strings.Builder

	for _, line := range split {
		if strings.Index(line, "failure") > -1 {
			comment.WriteString(fmt.Sprintf("%s %.1f\n", line, float64(-1.0)/float64(num_tests)*100))
			total_score -= float64(1)/float64(num_tests)*100
		}
	}

	hours := secondsToHours(late)
	if hours*2 >= total_score {
		total_score = 0.0
		comment.WriteString(fmt.Sprintf("%.1f Hours Late -%.1f", hours, hours*2))
	} else if hours > 0 {
		total_score -= hours*2
		comment.WriteString(fmt.Sprintf("%.1f Hours Late -%.1f", hours, hours*2))
	}

	if comment.String() == "" {
		comment.WriteString("Good Job!")
	}
	return total_score, comment.String()
}

func secondsToHours(seconds int) float64 {
	hours := float64(seconds/3600)
	if hours > 50.0 {
		hours = 50.0
	}
	return hours
}

type com struct {
	Text_comment string `json:"text_comment"`
}
type sub struct {
	Posted_grade string `json:"posted_grade"`
}

type grade struct {
	Comment com `json:"comment"`
	Submission sub `json:"submission"`
}


func GradeAndComment(fp, zippath, testpath, url string, lateseconds int, missed bool) {
	utils.Unzip(zippath, fp)
	utils.Cp(testpath, fp)
	utils.Cd(fp)
	to_move := utils.FindFolders(fp)
	for _, file := range to_move {
		if _, err := os.Stat(filepath.Join(file, "makefile")); err == nil {
			utils.Mv(filepath.Join(file,"makefile"), fp)
		}

		if _, err := os.Stat(filepath.Join(file, "Makefile")); err == nil {
			utils.Mv(filepath.Join(file,"Makefile"), fp)
		}

		if _, err := os.Stat(filepath.Join(file, "gcd.cpp")); err == nil {
			utils.Mv(filepath.Join(file,"gcd.cpp"), fp)
		}
		break
		
	}
	
	if _, err := os.Stat(filepath.Join(fp, "gcd.cpp")); err == nil {
		results := utils.RunBashScript("test_gcd.sh")
		score, comment := analyzeTestResults(results, missed, lateseconds)
		fmt.Println(score, comment)
		data := map[string]interface{}{
			"comment": map[string]string{
				"text_comment": comment,
			},
			"submission": map[string]string{
				"posted_grade": fmt.Sprintf("%d", int(score)),
			},
		}
		js, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		client := &http.Client{}
		
		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(js))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(resp)
	}
	if _, err := os.Stat(filepath.Join(fp, "GCD.cpp")); err == nil {
		score := 0
		comment := "File: gcd.cpp not found"
		fmt.Printf("%d, %s", int(score), comment)
		return
		data := map[string]interface{}{
			"comment": map[string]string{
				"text_comment": comment,
			},
			"submission": map[string]string{
				"posted_grade": fmt.Sprintf("%d", int(score)),
			},
		}
		js, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(js))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(resp.StatusCode)
		
	}

	utils.Cd("..")

	

	//fmt.Println(url)
}
