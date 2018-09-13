package canvas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"385grader/utils"
)

func FetchAllAssignmentUrls(courseID, assignmentID, token string) (subs []*assignmentSubmission) {
	url := fmt.Sprintf(FETCH_ALL_ASSIGNMENTS, CANVAS_API_DOMAIN, API_VERSION, courseID, assignmentID, token)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var submissions []submission

	if err = json.Unmarshal(body, &submissions); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(submissions)

	for _, submission := range submissions {
		var latest attachment

		for _, attachment := range submission.Attachments {
			if attachment.CreatedAt.After(latest.CreatedAt) {
				latest = attachment
			}
		}

		subs = append(subs, &assignmentSubmission{
			UserID:               submission.UserID,
			SecondsLate:          submission.SecondsLate,
			MostRecentSubmission: latest.Url,
			Missing:              submission.Missing,
			GradeUrl:             fmt.Sprintf(GRADE_ONE, CANVAS_API_DOMAIN, API_VERSION, courseID, assignmentID, submission.UserID, token),
			NameUrl:              fmt.Sprintf(LOOK_UP_STUDENT, CANVAS_API_DOMAIN, API_VERSION, courseID, submission.UserID, token),
		})
	}

	return
}

func GradeAllSubmissions(testscript string, subs []*assignmentSubmission) {
	cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	temp_dir := filepath.Join(cwd, "temp")

	err = os.Mkdir(temp_dir, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//utils.Cp(testscript, temp_dir)

	var fp string

	for _, sub := range subs {
		fmt.Println(sub.UserID)
		// if sub.UserID != 19628 {
		// 	continue
		// }
		
		fp = filepath.Join(temp_dir, fmt.Sprintf("%d.zip", sub.UserID))

		utils.DownloadFileFromUrl(sub.MostRecentSubmission, fp)
		sub.GradeAndComment(temp_dir, fp, testscript)

		// if sub.UserID == 19762 {
		// 	os.Exit(1)
		// }

		utils.Rm(filepath.Join(temp_dir))

		err = os.Mkdir(temp_dir, 0777)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//utils.Cp(testscript, temp_dir)
		
	}

	os.RemoveAll(filepath.Join(cwd, "temp"))
}
