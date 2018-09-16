package canvas

const (
	CANVAS_API_DOMAIN     = "sit.instructure.com"
	API_VERSION           = "v1"
	FETCH_ALL_ASSIGNMENTS = "https://%s/api/%s/courses/%s/assignments/%s/submissions?zip=1&access_token=%s&per_page=100"
	FETCH_ONE_ASSIGNMENT  = "https://%s/api/%s/courses/%s/assignments/%s/submissions/%s?zip=1&access_token=%s"
	GRADE_ONE             = "https://%s/api/%s/courses/%s/assignments/%s/submissions/%d?access_token=%s"
	LOOK_UP_STUDENT       = "https://%s/api/%s/courses/%s/users/%d?access_token=%s"
)
