package Data

import (
	"taskAPI/Model"
	"time"
)

var Tasks = []Model.Task{
	{
		ID:          1,
		Title:       "Finish Report",
		Description: "Work on the detailed report for the project.",
		Deadline:    time.Now().AddDate(0, 0, 3),
		Status:      Model.Discontinued,
	},
	{
		ID:          2,
		Title:       "Code Review",
		Description: "Review and refactor existing codebase.",
		Deadline:    time.Now().AddDate(0, 0, 5),
		Status:      Model.Pending,
	},
	{
		ID:          3,
		Title:       "Bug Fixing",
		Description: "Identify and fix bugs reported by QA.",
		Deadline:    time.Now().AddDate(0, 0, 2),
		Status:      Model.Pending,
	},
	{
		ID:          4,
		Title:       "Meeting Prep",
		Description: "Prepare agenda and materials for the team meeting.",
		Deadline:    time.Now().AddDate(0, 0, 4),
		Status:      Model.Finished,
	},
	{
		ID:          5,
		Title:       "Deploy New Feature",
		Description: "Implement and deploy new feature according to specifications.",
		Deadline:    time.Now().AddDate(0, 0, 7),
		Status:      Model.NotStarted,
	},
}
