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

type test struct {
	Link string `json:"link"`
}

func FetchAllAssignmentUrls(courseID, assignmentID, token string) (subs []*assignmentSubmission) {
	page, page_num := true, 1
	url := fmt.Sprintf(FETCH_ALL_ASSIGNMENTS, CANVAS_API_DOMAIN, API_VERSION, courseID, assignmentID, token)

	for page {
		cur_url := fmt.Sprintf("%s&page=%d", url, page_num)

		resp, err := http.Get(cur_url)
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

		if len(submissions) < 100 {
			page = false
		}
		page_num++

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
	}

	return
}

func GradeAllSubmissions(entrypoint, testscript string, subs []*assignmentSubmission, timeout int, post bool) {
	tempDir := utils.CreateTempDir()

	var fp string

	for _, sub := range subs {
		//fmt.Println(sub.UserID)
		fp = filepath.Join(tempDir, fmt.Sprintf("%d.zip", sub.UserID))

		utils.DownloadFileFromUrl(sub.MostRecentSubmission, fp)
		sub.GradeAndComment(tempDir, fp, testscript, entrypoint, timeout, post)

		// if sub.UserID == 22855 {
		// 	os.Exit(1)
		// }

		os.RemoveAll(tempDir)

		err := os.Mkdir(tempDir, 0777)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	os.RemoveAll(tempDir)
}

func FetchOne(courseID, assignmentID, userID, token string) *assignmentSubmission {
	url := fmt.Sprintf(FETCH_ONE_ASSIGNMENT, CANVAS_API_DOMAIN, API_VERSION, courseID, assignmentID, userID, token)

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

	var sub submission

	if err = json.Unmarshal(body, &sub); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var latest attachment

	for _, attachment := range sub.Attachments {
		if attachment.CreatedAt.After(latest.CreatedAt) {
			latest = attachment
		}
	}

	return &assignmentSubmission{
		UserID:               sub.UserID,
		SecondsLate:          sub.SecondsLate,
		MostRecentSubmission: latest.Url,
		Missing:              sub.Missing,
		GradeUrl:             fmt.Sprintf(GRADE_ONE, CANVAS_API_DOMAIN, API_VERSION, courseID, assignmentID, sub.UserID, token),
		NameUrl:              fmt.Sprintf(LOOK_UP_STUDENT, CANVAS_API_DOMAIN, API_VERSION, courseID, sub.UserID, token),
	}
}

func GradeOneSubmission(entrypoint, testscript string, sub *assignmentSubmission, timeout int, post bool) {
	tempDir := utils.CreateTempDir()

	var fp string

	fp = filepath.Join(tempDir, fmt.Sprintf("%d.zip", sub.UserID))

	utils.DownloadFileFromUrl(sub.MostRecentSubmission, fp)
	sub.GradeAndComment(tempDir, fp, testscript, entrypoint, timeout, post)

	os.RemoveAll(tempDir)
}
