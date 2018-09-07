package canvas

import (
	"time"
)

type (
	
	attachment struct {
		ID int `json:"id"`
		UUID string `json:"uuid"`
		FolderID int `json:"folder_id"`
		DisplayName string `json:"display_name"`
		Filename string `json:"filename"`
		ContentType string `json:"content-type"`
		Url string `json:"url"`
		Size int `json:"size"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		UnlockAt string `json:"unlock_at"`
		Locked bool `json:"locked"`
		Hidden bool `json:"hidden"`
		LockAt string `json:"lock_at"`
		HiddenForUser bool `json:"hidden_for_user"`
		ThumbnailUrl string `json:"thumbnail_url"`
		ModifiedAt time.Time `json:"modified_at"`
		MimeClass string `json:"mime_class"`
		MediaEntryID int `json:"media_entry_id"`
		LockedForUser bool `json:"locked_for_user"`
		PreviewUrl string `json:"preview_url"`
	}
	
	submission struct {
		AssignmentID int `json:"assignment_id"`
		// Assignment string `json:"assignment"`
		// Course string `json:"course"`
		ID int `json:"id"`
		CachedDueDate time.Time `json:"cached_due_date"`
		Attempt int `json:"attempt"`
		Body string `json:"body"`
		Grade string `json:"grade"`
		EnteredGrade string `json:"entered_grade"`
		GradeMatchesCurrentSubmission bool `json:"grade_matches_current_submission"`
		//HtmlUrl string `json:"html_url"`
		PreivewUrl string `json:"preview_url"`
		Score float64 `json:"score"`
		EnteredScore float64 `json:"entered_scored"`
		//SubmissionComments []*submissionComment `json:"submission_comments"`
		SubmissionType string `json:"submission_type"`
		SubmittedAt time.Time `json:"submitted_at"`
		Url string `json:"url"`
		UserID int `json:"user_id"`
		GraderID int `json:"grader_id"`
		GradedAt time.Time `json:"graded_at"`
		GradingPeriodID string `json:"grading_period_id"`
		Late bool `json:"late"`
		//AssignmentVisible bool `json:"assignment_visible"`
		Excused bool `json:"excused"`
		Missing bool `json:"missing"`
		LatePolicyStatus string `json:"late_policy_status"`
		PointsDeducted float64 `json:"points_deducted"`
		SecondsLate int `json:"seconds_late"`
		WorkflowState string `json:"workflow_state"`
		Attachments []attachment `json:"attachments"`
	}

	assignmentSubmission struct {
		UserID int
		SecondsLate int
		MostRecentSubmission string
		Missing bool
		GradeUrl string
	}

	mediaComment struct {
		ContentType string `json:"content-type"`
		DisplayName string `json:"display_name"`
		MediaID string `json:"media_id"`
		MediaType string `json:"media_type"`
		Url string `json:"url"`
	}
	
	submissionComment struct {
		ID int `json:"id"`
		AuthorID int `json:"author_id"`
		AuthorName string `json:"author_name"`
		//Author *userDisplay `json:"author"`
		Comment string `json:"comment"`
		CreatedAt string `json:"created_at"`
		EditedAt string `json:"edited_at"`
		MediaComment *mediaComment `json:"media_comment"`
	}
) 
