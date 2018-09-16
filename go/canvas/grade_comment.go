package canvas

import (
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

func countNumTests(filepath string) int {
	reg := regexp.MustCompile("run_test")
	matches := reg.FindAllStringIndex(utils.Cat(filepath), -1)

	return len(matches)
}

func (s *assignmentSubmission) getName() []string {
	resp, err := http.Get(s.NameUrl)
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

	var person user
	if err = json.Unmarshal(body, &person); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
		comment.WriteString(fmt.Sprintf("%s or %s not found in header comments -5.", name[0], name[1]))
	}

	reg_pledge := regexp.MustCompile("[^a-z1-9]?i pledge my honor that i have abided by the steven(')?s honor system[^a-z1-9]?")
	matches = reg_pledge.FindAllStringIndex(head, -1)

	if len(matches) == 0 {
		score += 5
		comment.WriteString("Pledge not found/spelt wrong in header comments -5.")
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
			comment.WriteString(fmt.Sprintf("%s %.1f\n", line, float64(-1.0)/float64(num_tests)*100))
			total_score -= float64(1) / float64(num_tests) * 100
		}
	}

	hours := secondsToHours(s.SecondsLate)
	if hours*2 >= total_score {
		total_score = 0.0
		comment.WriteString(fmt.Sprintf("%.1f Hours Late -%.1f", hours, hours*2))
	} else if hours > 0 {
		total_score -= hours * 2
		comment.WriteString(fmt.Sprintf("%.1f Hours Late -%.1f", hours, hours*2))
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
	fmt.Println(gradeUrl)
	data := grade{
		Comment: com{
			TextComment: comment,
		},
		Submission: sub{
			PostedGrade: fmt.Sprintf("%d", int(score)),
		},
	}
	js, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, gradeUrl, bytes.NewBuffer(js))
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
	//fmt.Println(resp)
}

func (s *assignmentSubmission) GradeAndComment(fp, zippath, testpath, entrypoint string, timeout int, post bool) {
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
			comment = "Incorrect folder structure -5."
		} else {
			comment += "\nIncorrect folder structure -5."
		}
		score -= incorrectStructure
	}

	if post {
		postGrade(s.GradeUrl, comment, score)
	}
	//fmt.Println(score, comment)

	utils.Cd("..")
}
