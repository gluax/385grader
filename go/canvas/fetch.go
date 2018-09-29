package canvas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"385grader/utils"
)

func FetchAllAssignmentUrls(courseID, assignmentID, token string) (subs []*assignmentSubmission) {
	page, page_num := true, 1
	url := fmt.Sprintf(FETCH_ALL_ASSIGNMENTS, CANVAS_API_DOMAIN, API_VERSION, courseID, assignmentID, token)

	for page {
		cur_url := fmt.Sprintf("%s&page=%d", url, page_num)

		resp, err := http.Get(cur_url)
		utils.HandleError(err, "Get request failed. Did API change or network failure or is there service down.", true)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		utils.HandleError(err, "Could not parse request data, did API schema change?", true)

		var submissions []submission

		err = json.Unmarshal(body, &submissions)
		utils.HandleError(err, "Failed to unpack values to struct.", true)

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

func FetchOne(courseID, assignmentID, userID, token string) *assignmentSubmission {
	url := fmt.Sprintf(FETCH_ONE_ASSIGNMENT, CANVAS_API_DOMAIN, API_VERSION, courseID, assignmentID, userID, token)

	resp, err := http.Get(url)
	utils.HandleError(err, "Get request failed. Did API change or network failure or is there service down.", true)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	utils.HandleError(err, "Could not parse request data, did API schema change?", true)

	var sub submission

	err = json.Unmarshal(body, &sub)
	utils.HandleError(err, "Failed to unpack values to struct.", true)

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
