package canvas

const (
	CANVAS_API_DOMAIN = "sit.instructure.com"
	API_VERSION = "v1"
	FETCH_ALL_ASSIGNMENTS = "https://%s/api/%s/courses/%s/assignments/%s/submissions?zip=1&access_token=%s&per_page=1000"
	FETCH_ONE_ASSIGNMENT = ""
	GRADE_ONE = "https://%s/api/%s/courses/%s/assignments/%s/submissions/%d?access_token=%s"
	LOOK_UP_STUDENT_NAME = ""	
)
