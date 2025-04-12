package models

type Section struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CourseId    int    `json:"course_id"`
}
