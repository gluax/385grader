package canvas

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"385grader/utils"
)

func valgrind(executable, valgrindFile string) (float64, string) {
	// if _, err := os.Stat("./makefile"); os.IsNotExist(err) {
	// 	return 0.0, ""
	// }

	// if _, err := os.Stat("./Makefile"); os.IsNotExist(err) {
	// 	return 0.0, ""
	// }

	utils.Make()

	var toRemove = 0.0
	var comment strings.Builder
	var resp string

	reg := regexp.MustCompile("definitely lost: [0-9]+ bytes.+\n")
	lines := utils.ReadValgrindFile(valgrindFile)
	for _, line := range lines {
		resp = utils.RunValgrind(executable, strings.Split(line, " "))
		matches := reg.FindAllStringIndex(resp, -1)
		start, end := matches[0][0], matches[0][1]
		lost := resp[start : end-1]
		parts := strings.Fields(lost)

		if parts[2] != "0" {
			toRemove += 5
			comment.WriteString(fmt.Sprintf("\n%s -5", lost))
		}
	}

	return toRemove, comment.String()
}

func countNumTests(filepath string) int {
	reg := regexp.MustCompile("run_test")
	matches := reg.FindAllStringIndex(utils.Cat(filepath), -1)

	return len(matches) - 1
}

func (s *assignmentSubmission) getName() []string {
	resp, err := http.Get(s.NameUrl)
	utils.HandleError(err, "Get request failed. Did API change or network failure or is there service down.", true)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	utils.HandleError(err, "Could not parse request data, did API schema change?", true)

	var person user
	err = json.Unmarshal(body, &person)
	utils.HandleError(err, "Failed to unpack values to struct.", true)

	name := strings.Split(person.Name, " ")
	name[0], name[1] = strings.ToLower(name[0]), strings.ToLower(name[1])

	return name
}

func (s *assignmentSubmission) nameAndPledge(filepath string) (float64, string) {
	head := utils.Head(filepath, 10)
	head = strings.ToLower(head)

	var score float64 = 0
	var comment strings.Builder
	name := s.getName()

	reg_name := regexp.MustCompile(fmt.Sprintf("[^a-z1-9]?(%s|%s)[^a-z1-9]?", name[0], name[1]))
	matches := reg_name.FindAllStringIndex(head, -1)

	if len(matches) == 0 {
		score += 5
		comment.WriteString(fmt.Sprintf("\n%s or %s not found in header comments -5.", name[0], name[1]))
	}

	reg_pledge := regexp.MustCompile("[^a-z1-9]?i pledge my honor that i have abided by the steven(')?s honor system[^a-z1-9]?")
	matches = reg_pledge.FindAllStringIndex(head, -1)

	if len(matches) == 0 {
		score += 5
		comment.WriteString("\nPledge not found/spelt wrong in header comments -5.")
	}

	return score, comment.String()
}

func (s *assignmentSubmission) analyzeTestResults(entrypoint, testoutput string, num_tests int) (float64, string) {
	if s.Missing {
		return 0, "No Submission."
	}

	split := strings.Split(testoutput, "\n")
	if strings.Index(split[0], "done") == -1 {
		return 0, "Compilation Failed."
	}

	var total_score float64 = 100
	var comment strings.Builder

	sub, com := s.nameAndPledge(entrypoint)
	total_score -= sub
	comment.WriteString(com)

	for _, line := range split {
		if strings.Index(line, "failure") > -1 {
			comment.WriteString(fmt.Sprintf("\n%s %.1f\n", line, float64(-1.0)/float64(num_tests)*100))
			total_score -= float64(1) / float64(num_tests) * 100
		}
	}

	hours := secondsToHours(s.SecondsLate)
	if hours*2 >= total_score {
		total_score = 0.0
		comment.WriteString(fmt.Sprintf("\n%.1f Hours Late -%.1f", hours, hours*2))
	} else if hours > 0 {
		total_score -= hours * 2
		comment.WriteString(fmt.Sprintf("\n%.1f Hours Late -%.1f", hours, hours*2))
	}

	if comment.String() == "" {
		comment.WriteString("Good Job!")
	}
	return total_score, comment.String()
}

func secondsToHours(seconds int) float64 {
	hours := math.Ceil(float64(seconds) / 3600.0)
	if hours > 50.0 {
		hours = 50.0
	}
	return hours
}

func checkFolderStructure(fp string) float64 {
	toMove := utils.FindFolders(fp)
	toMove = utils.FilterStringSlice(toMove, func(item string) bool {
		if strings.Contains(item, "__MACOSX") {
			return false
		}
		return true
	})

	if len(toMove) == 0 {
		return 0
	}

	for _, folder := range toMove {
		reg := regexp.MustCompile(" ")
		folder = reg.ReplaceAllLiteralString(folder, "\\ ")

		utils.Mv(filepath.Join(folder, "*.cpp"), fp)
		utils.Mv(filepath.Join(folder, "*.h"), fp)
		utils.Mv(filepath.Join(folder, "makefile"), fp)
		utils.Mv(filepath.Join(folder, "Makefile"), fp)

		os.RemoveAll(folder)
	}

	return 5
}

func postGrade(gradeUrl, comment string, score float64) {
	data := grade{
		Comment: com{
			TextComment: comment,
		},
		Submission: sub{
			PostedGrade: fmt.Sprintf("%d", int(score)),
		},
	}

	js, err := json.Marshal(data)
	utils.HandleError(err, "Failed to pack values to struct.", true)
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, gradeUrl, bytes.NewBuffer(js))
	utils.HandleError(err, "Put request failed. Did API change or network failure or is there service down.", true)

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	utils.HandleError(err, "Put request failed. Did API change or network failure or is there service down.", true)

	utils.Log(resp, "Grade post response.")
}

func GradeAllSubmissions(entrypoint, testscript, executable, valgrindFile string, subs []*assignmentSubmission, timeout int, post, view bool) {
	tempDir := utils.CreateTempDir()

	var fp string

	for _, sub := range subs {
		data := map[string]interface{}{
			"id": sub.UserID,
		}
		utils.Log(data, "Current User Id")
		fp = filepath.Join(tempDir, fmt.Sprintf("%d.zip", sub.UserID))

		utils.DownloadFileFromUrl(sub.MostRecentSubmission, fp)
		sub.gradeAndComment(tempDir, fp, testscript, entrypoint, executable, valgrindFile, timeout, post, view)

		os.RemoveAll(tempDir)

		err := os.Mkdir(tempDir, 0777)
		utils.HandleError(err, "Failed to make temp directory.", true)
	}

	os.RemoveAll(tempDir)
}

func GradeOneSubmission(entrypoint, testscript, executable, valgrindFile string, sub *assignmentSubmission, timeout int, post, view bool) {
	tempDir := utils.CreateTempDir()

	var fp string

	fp = filepath.Join(tempDir, fmt.Sprintf("%d.zip", sub.UserID))
	utils.DownloadFileFromUrl(sub.MostRecentSubmission, fp)
	sub.gradeAndComment(tempDir, fp, testscript, entrypoint, executable, valgrindFile, timeout, post, view)

	//os.RemoveAll(tempDir)
}

func (s *assignmentSubmission) gradeAndComment(fp, zippath, testpath, entrypoint, executable, valgrindFile string, timeout int, post, view bool) {
	utils.Unzip(zippath, fp)
	utils.Cp(testpath, fp)
	utils.Cd(fp)

	var score float64 = 0
	var incorrectStructure = checkFolderStructure(fp)
	var comment string

	_, testName := filepath.Split(testpath)

	results := utils.RunBashScript(testName, timeout)
	entrypoint = filepath.Join(fp, entrypoint)
	score, comment = s.analyzeTestResults(entrypoint, results, countNumTests(testpath))

	if incorrectStructure > 0 {
		if score == 100 {
			comment = "\nIncorrect folder structure -5."
		} else {
			comment += "\nIncorrect folder structure -5."
		}
		score -= incorrectStructure
	}

	if valgrindFile != "" {
		execPath := filepath.Join(fp, executable)
		s, c := valgrind(execPath, valgrindFile)
		score -= s
		comment += c
	}

	if view && score > 0 {
		var sub float64
		var enter, exit string

		fmt.Println(utils.Cat(entrypoint))
		fmt.Printf("Remove points and comment(Y)?")
		fmt.Scanf("%f", &enter)
		for enter == "Y" {
			fmt.Printf("Enter amount points to take off: ")
			fmt.Scanf("%f", &sub)
			fmt.Printf("Enter comment to go with removed points: ")

			in := bufio.NewReader(os.Stdin)
			add, _ := in.ReadString('\n')

			for exit != "N" && exit != "Y" {
				fmt.Printf("Enter another comment/point removal(Y/N)?: ")
				fmt.Scan(&exit)
			}

			if score == 100 {
				comment = add
			} else {
				comment += "\n" + add
			}
			score -= sub

			if exit == "N" {
				break
			}
			exit = ""
		}
	}

	if post {
		postGrade(s.GradeUrl, comment, score)
	}
	data := map[string]interface{}{
		"id":      s.UserID,
		"score":   score,
		"comment": comment,
	}

	utils.Log(data, "scores and comment.")

	utils.Cd("..")
}
