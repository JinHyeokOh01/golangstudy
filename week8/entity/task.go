package entity

import "time"

type TaskID int64
type TaskStatus string

const (
	TaskStatusTodo TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone TaskStatus = "done"
)

type Task struct{
	ID			TaskID		`json:"id"`
	UserID		UserID		`json:"user_id" db:"user_id"`
	Title 		string		`json:"title"`
	Status 		TaskStatus	`json:"status"`
	Created 	time.Time	`json:"created"`
	Modified 	time.Time	`json:"modified" db:"modified"`
}

type Tasks []*Task