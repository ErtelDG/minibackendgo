// The code defines Go structs for Contact, Task, and Subtask with specified fields and JSON tags.
// @property {string} ID_contact - The `ID_contact` field in the `Contact` struct represents the unique
// identifier for a contact.
// @property {string} First_Name - The `First_Name` field is a string type in the `Contact` struct with
// the JSON tag `first_name`. It is used to store the first name of a contact in a contact management
// system.
// @property {string} Last_Name - Last_Name is a field in the Contact struct that represents the last
// name of a contact. It is of type string and is tagged with `json:"last_name"` for JSON marshaling
// and unmarshaling.
// @property {string} Email - The `Email` property is a field in the `Contact` struct. It represents
// the email address of a contact.
// @property {string} Phone - The `Phone` property is a field in the `Contact` struct. It represents
// the phone number of a contact.
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