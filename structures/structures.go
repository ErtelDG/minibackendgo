package structures

type Contact struct {
	ID_contact string
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

type Task struct {
	ID_task     string
	Status      string   `json:"status"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Assigned    string   `json:"assigned"`
	Prio        string   `json:"prio"`
	Due_date    string   `json:"due_date"`
	Category    string   `json:"category"`
	Subtasks    []Subtask `json:"subtasks"`
}

type Subtask struct {
    SubtaskId int64  `json:"subtaskId"`
    Title     string `json:"title"`
    Checked   bool   `json:"checked"`
}

