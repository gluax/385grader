package canvas

import (
	"time"
)

type (
	attachment struct {
		ID            int       `json:"id"`
		UUID          string    `json:"uuid"`
		FolderID      int       `json:"folder_id"`
		DisplayName   string    `json:"display_name"`
		Filename      string    `json:"filename"`
		ContentType   string    `json:"content-type"`
		Url           string    `json:"url"`
		Size          int       `json:"size"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		UnlockAt      string    `json:"unlock_at"`
		Locked        bool      `json:"locked"`
		Hidden        bool      `json:"hidden"`
		LockAt        string    `json:"lock_at"`
		HiddenForUser bool      `json:"hidden_for_user"`
		ThumbnailUrl  string    `json:"thumbnail_url"`
		ModifiedAt    time.Time `json:"modified_at"`
		MimeClass     string    `json:"mime_class"`
		MediaEntryID  int       `json:"media_entry_id"`
		LockedForUser bool      `json:"locked_for_user"`
		PreviewUrl    string    `json:"preview_url"`
	}

	submission struct {
		AssignmentID                  int          `json:"assignment_id"`
		ID                            int          `json:"id"`
		CachedDueDate                 time.Time    `json:"cached_due_date"`
		Attempt                       int          `json:"attempt"`
		Body                          string       `json:"body"`
		Grade                         string       `json:"grade"`
		EnteredGrade                  string       `json:"entered_grade"`
		GradeMatchesCurrentSubmission bool         `json:"grade_matches_current_submission"`
		PreivewUrl                    string       `json:"preview_url"`
		Score                         float64      `json:"score"`
		EnteredScore                  float64      `json:"entered_scored"`
		SubmissionType                string       `json:"submission_type"`
		SubmittedAt                   time.Time    `json:"submitted_at"`
		Url                           string       `json:"url"`
		UserID                        int          `json:"user_id"`
		GraderID                      int          `json:"grader_id"`
		GradedAt                      time.Time    `json:"graded_at"`
		GradingPeriodID               string       `json:"grading_period_id"`
		Late                          bool         `json:"late"`
		Excused                       bool         `json:"excused"`
		Missing                       bool         `json:"missing"`
		LatePolicyStatus              string       `json:"late_policy_status"`
		PointsDeducted                float64      `json:"points_deducted"`
		SecondsLate                   int          `json:"seconds_late"`
		WorkflowState                 string       `json:"workflow_state"`
		Attachments                   []attachment `json:"attachments"`
	}

	user struct {
		UserID       int    `json:"id"`
		Name         string `json:"name"`
		SortableName string `json:"sortable_name"`
		ShortName    string `json:"short_name"`
	}

	com struct {
		TextComment string `json:"text_comment"`
	}

	sub struct {
		PostedGrade string `json:"posted_grade"`
	}

	grade struct {
		Comment    com `json:"comment"`
		Submission sub `json:"submission"`
	}

	assignmentSubmission struct {
		UserID               int
		SecondsLate          int
		MostRecentSubmission string
		Missing              bool
		GradeUrl             string
		NameUrl              string
	}
)
