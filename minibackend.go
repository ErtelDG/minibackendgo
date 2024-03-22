package main

import (
	"encoding/json"
	"fmt"
	"io"
	"minibackend/structures"
	"net/http"
	"os"
	"strconv"
	"time"
)

// The main function sets up routes and handlers for a web server in Go and starts the server on port
// 8080.
func main() {
	mux := http.NewServeMux()

	// The `routes` variable in the Go code snippet is a map that associates specific URL paths with
	// corresponding handler functions. Each key-value pair in the map represents a route path and the
	// handler function that should be executed when a request is made to that path.
	routes := map[string]http.HandlerFunc{
		"/contacts":       contacts,
		"/categories":     categories,
		"/tasks":          tasks,
		"/add_contact":    add_contact,
		"/remove_contact": removeContact,
		"/add_task":       add_task,
		"/del_task":       deleteTask,
		"/update_task":    updateTask,
	}

	// The code snippet `for route, handler := range routes { mux.HandleFunc(route, handler) }` is
	// iterating over the `routes` map, where each key represents a specific URL path and the corresponding
	// value is a handler function.
	for route, handler := range routes {
		mux.HandleFunc(route, handler)
	}

	// The code snippet `err := http.ListenAndServe(":8080", mux)` is starting a web server in Go on port
	// 8080. If there is an error during the server startup, the subsequent code checks if the error is
	// related to a TLS handshake error. If there is an error, it is printed to the console with the
	// message "TLS Handshake Error:" followed by the error message, and then the program panics with the
	// error. This means that the program will stop execution and display the error message if there is an
	// issue with starting the server.
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("TLS Handshake Error:", err)
		panic(err)
	}
}

// The `contacts` function reads a JSON file containing contacts data and serves it as a response with
// appropriate headers in a Go HTTP server.
func contacts(w http.ResponseWriter, r *http.Request) {

	// The above code snippet in Go is attempting to read the contents of a file named "contacts.json"
	// located in the "./data" directory. It then sets the necessary headers for allowing cross-origin
	// requests and specifying the content type as JSON. If an error occurs during the file reading
	// process, it will return a 404 Not Found status along with the error message.
	read_file, err := os.ReadFile("./data/contacts.json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// The above code is attempting to write the contents of a file (`read_file`) to the response writer
	// (`w`) in a Go web application. If an error occurs during the write operation, it will return a 500
	// Internal Server Error response with the error message.
	_, err = w.Write(read_file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// The function `add_contact` in Go handles POST requests to add a new contact by reading, decoding,
// updating, and writing contact data to a JSON file, setting appropriate HTTP headers, and returning
// the updated contacts as a JSON response.
func add_contact(w http.ResponseWriter, r *http.Request) {

	// The above code is checking if the HTTP request method is POST. If the method is not POST, it returns
	// a "Method not allowed" error with status code 405 (Method Not Allowed).
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//  It is attempting to read contacts data
	// from a JSON file located at "./data/contacts.json". It first tries to read the existing contacts
	// data from the file using the `readContactsFromFile` function. If there is an error during the
	// reading process, it will return an HTTP 500 Internal Server Error response with the message "Error
	// reading existing contacts".
	contactsFile := "./data/contacts.json"
	existingContacts, err := readContactsFromFile(contactsFile)
	if err != nil {
		http.Error(w, "Error reading existing contacts", http.StatusInternalServerError)
		return
	}

	// The above code snippet in Go is creating a new contact ID by concatenating the string "cont" with
	// the current Unix timestamp converted to a string. It then creates a new Contact struct instance with
	// the generated ID. The code then decodes the JSON data from the request body into the newContact
	// struct. If there is an error during the decoding process, it will return a Bad Request HTTP error
	// with the error message.
	newContactID := "cont" + strconv.FormatInt(time.Now().Unix(), 10)
	newContact := &structures.Contact{ID_contact: newContactID}
	err = json.NewDecoder(r.Body).Decode(newContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the new contact to the existing contacts
	existingContacts[newContactID] = newContact

	// The above code snippet is attempting to write the existing contacts to a file specified by
	// `contactsFile`. If an error occurs during the write operation, it will return an HTTP 500 Internal
	// Server Error response with the message "Error writing updated contacts".
	err = writeContactsToFile(existingContacts, contactsFile)
	if err != nil {
		http.Error(w, "Error writing updated contacts", http.StatusInternalServerError)
		return
	}

	//  It is setting HTTP headers for CORS
	// (Cross-Origin Resource Sharing) and content type in the response. Here's a breakdown of what each
	// line is doing:
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(existingContacts)
}

// The function `categories` reads a JSON file containing categories data and serves it over HTTP with
// appropriate headers.
func categories(w http.ResponseWriter, r *http.Request) {

	// The code snippet provided is written in Go programming language. Here's a breakdown of what the code
	// is doing:
	read_file, err := os.ReadFile("./data/categories.json")

	// The above code snippet is written in Go and it is handling an error condition. If the `err` variable
	// is not `nil`, it will return a HTTP 404 Not Found error with the error message as the response body.
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// The above code is writing the contents of a file (`read_file`) to the response writer (`w`) in a Go
	// web application. If an error occurs during the writing process, it will return an internal server
	// error with the error message.
	_, err = w.Write(read_file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// The function `tasks` reads a JSON file containing tasks and sends it as a response with appropriate
// headers in a Go HTTP server.
func tasks(w http.ResponseWriter, r *http.Request) {
	read_file, err := os.ReadFile("./data/tasks.json")

	// The above code snippet is handling an error in a Go program. If the `err` variable is not `nil`, it
	// will return a 404 Not Found HTTP status code along with the error message in the response.
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// The above code is writing the contents of a file (`read_file`) to the response writer (`w`) in a Go
	// web application. If an error occurs during the writing process, it will return an internal server
	// error with the error message.
	_, err = w.Write(read_file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// The `add_task` function handles adding a new task to a list of existing tasks stored in a JSON file
// in a Go web application.
func add_task(w http.ResponseWriter, r *http.Request) {

	// The above code is checking if the HTTP request method is POST. If the method is not POST, it
	// returns a "Method not allowed" error with status code 405 (Method Not Allowed).
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// The above code is reading tasks from a JSON file named "tasks.json" using the `readTasksFromFile`
	// function. If there is an error reading the tasks from the file, it will return an internal server
	// error response with the message "Error reading existing tasks".
	tasksFile := "./data/tasks.json"
	existingTasks, err := readTasksFromFile(tasksFile)
	if err != nil {
		http.Error(w, "Error reading existing tasks", http.StatusInternalServerError)
		return
	}

	// The above code snippet in Go is generating a new task ID by concatenating "tk" with the current Unix
	// timestamp using `time.Now().Unix()`. It then creates a new task object with the generated ID. The
	// code then decodes the JSON data from the request body into the new task object using
	// `json.NewDecoder(r.Body).Decode(newTask)`. If there is an error during decoding, it will return a
	// Bad Request response with the error message.
	newTaskID := "tk" + strconv.FormatInt(time.Now().Unix(), 10)
	newTask := &structures.Task{ID_task: newTaskID}
	err = json.NewDecoder(r.Body).Decode(newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// The above code snippet is adding a new task to an existing list of tasks and then writing the
	// updated tasks to a file. If there is an error while writing the tasks to the file, it will return a
	// 500 Internal Server Error response with the message "Error writing updated tasks".
	existingTasks[newTaskID] = newTask
	err = writeTasksToFile(existingTasks, tasksFile)
	if err != nil {
		http.Error(w, "Error writing updated tasks", http.StatusInternalServerError)
		return
	}

	// The above code is written in Go programming language. It is setting the response headers for CORS
	// (Cross-Origin Resource Sharing) to allow requests from any origin ("*") and specifying the allowed
	// headers as "Content-Type". It then sets the Content-Type of the response to "application/json", sets
	// the HTTP status code to 201 (Created), and encodes an existingTasks variable as JSON and writes it
	// to the response writer (w).
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(existingTasks)
}

// The `updateTask` function in Go handles updating a task by reading the request body, parsing JSON
// data, updating the task in a tasks file, and sending a success response.
func updateTask(w http.ResponseWriter, r *http.Request) {

	// The above code is checking if the HTTP request method is POST. If the method is not POST, it returns
	// a "Method not allowed" error with a status code of 405 (Method Not Allowed).
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// The above code is reading the request body from an HTTP request in Go. It uses the `io.ReadAll`
	// function to read the request body and stores the result in the `body` variable. If there is an error
	// while reading the request body, it will return an HTTP error response with a status code of 500
	// (Internal Server Error) and a message "Error reading request body".
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// The above code is attempting to unmarshal JSON data from the `body` variable into a `task` struct in
	// Go. If there is an error during the unmarshalling process, it will return a "Invalid JSON format"
	// error with a status code of 400 (Bad Request).
	var task structures.Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// The above code is reading tasks from a JSON file located at "./data/tasks.json". It first attempts
	// to read the tasks from the file using the `readTasksFromFile` function. If there is an error
	// reading the tasks, it returns an HTTP 500 Internal Server Error response with the message "Error
	// reading existing tasks".
	tasksFile := "./data/tasks.json"
	existingTasks, err := readTasksFromFile(tasksFile)
	if err != nil {
		http.Error(w, "Error reading existing tasks", http.StatusInternalServerError)
		return
	}

	// The above code snippet is written in Go programming language. It is updating an existing task in a
	// map called `existingTasks` by assigning a reference to the task with the key `task.ID_task`. After
	// updating the task in the map, it then calls a function `writeTasksToFile` to write the updated tasks
	// to a file specified by `tasksFile`. If there is an error during the writing process, it returns an
	// HTTP error response with status code 500 (Internal Server Error).
	existingTasks[task.ID_task] = &task
	if err := writeTasksToFile(existingTasks, tasksFile); err != nil {
		http.Error(w, "Error writing updated tasks", http.StatusInternalServerError)
		return
	}

	// The above code is written in Go programming language and it is setting the response header
	// "Content-Type" to "application/json", setting the HTTP status code to 200 (OK), and writing the
	// message "Task updated successfully" to the response body. This code is typically used in a web
	// server to send a JSON response indicating that a task has been updated successfully.
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Task updated successfully")
}

// The function reads contacts data from a file in JSON format and returns a map of contacts.
func readContactsFromFile(filePath string) (map[string]*structures.Contact, error) {

	// The above code in Go is reading the contents of a file specified by the `filePath` variable using
	// the `os.ReadFile` function. If there is an error during the file reading process, it will return
	// `nil` and the error.
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// The above code is decoding JSON data into a map of string keys and pointers to Contact structures in
	// the Go programming language. It uses the `json.Unmarshal` function to unmarshal the JSON data into
	// the `contacts` map. If there is an error during the unmarshalling process, it returns `nil` and the
	// error.
	contacts := make(map[string]*structures.Contact)
	err = json.Unmarshal(data, &contacts)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

// The function `readTasksFromFile` reads tasks from a file in JSON format and returns them as a map of
// task IDs to task objects.
func readTasksFromFile(filePath string) (map[string]*structures.Task, error) {

	// The above code snippet is reading the contents of a file located at the `filePath` using the
	// `os.ReadFile` function in Go. If there is an error during the file reading process, it will return
	// `nil` and the error.
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// The above code is unmarshalling JSON data into a map of tasks in Go. It creates a map called `tasks`
	// with string keys and pointers to `structures.Task` values. It then uses `json.Unmarshal` to decode
	// the JSON data into the `tasks` map. If there is an error during unmarshalling, it returns `nil` and
	// the error. Otherwise, it returns the `tasks` map and a `nil` error.
	tasks := make(map[string]*structures.Task)
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Helper function to write contacts to a file
func writeContactsToFile(contacts map[string]*structures.Contact, filePath string) error {

	// The above code is marshaling the `contacts` variable into a JSON format with an indented structure
	// using the `json.MarshalIndent` function in Go. The resulting JSON data is stored in the `data`
	// variable. If there is an error during the marshaling process, the error is returned.
	data, err := json.MarshalIndent(contacts, "", "   ")
	if err != nil {
		return err
	}

	// The above code is writing data to a file specified by the `filePath` variable using the
	// `os.WriteFile` function in Go. It sets the file permissions to 0644. If there is an error during the
	// write operation, it returns the error. If the write operation is successful, it returns nil.
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// The `writeTasksToFile` function writes tasks stored in a map to a JSON file in Go.
func writeTasksToFile(tasks map[string]*structures.Task, filePath string) error {

	// The above code is marshaling the `tasks` variable into a JSON format with an indented structure
	// using the `json.MarshalIndent` function in Go. If there is an error during the marshaling process,
	// it will return the error.
	data, err := json.MarshalIndent(tasks, "", "   ")
	if err != nil {
		return err
	}

	// The above code is writing data to a file specified by the `filePath` variable using the
	// `os.WriteFile` function in Go. It sets the file permissions to 0644. If there is an error during the
	// write operation, it returns the error. If the write operation is successful, it returns nil.
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// The `deleteTask` function handles deleting a task based on the provided task ID from a JSON file in
// a Go HTTP server.
func deleteTask(w http.ResponseWriter, r *http.Request) {

	// The above code is checking if the HTTP request method is not DELETE. If the method is not DELETE, it
	// returns a "Method not allowed" error with status code 405 (Method Not Allowed).
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// The above code is reading the request body from an HTTP request in a Go program. It uses the
	// `io.ReadAll` function to read the entire body of the request and stores it in the `body` variable.
	// If there is an error while reading the request body, it will return an HTTP error response with a
	// status code of 500 (Internal Server Error) and the message "Error reading request body".
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// The above code in Go is defining a struct named `requestData` with a single field `TaskID` of type
	// string. The `json:"task_id"` tag is used to specify the JSON key for this field when marshaling and
	// unmarshaling JSON data.
	var requestData struct {
		TaskID string `json:"task_id"`
	}

	// The above code is attempting to unmarshal a JSON request body into a struct variable named
	// `requestData`. If there is an error during the unmarshaling process, it will return a 400 Bad
	// Request response with the message "Invalid JSON format".
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// The above code is written in Go programming language. It checks if the `TaskID` field in the
	// `requestData` object is empty. If it is empty, it returns a HTTP 400 Bad Request error with the
	// message "Task ID not provided". This code snippet is used to validate the presence of a Task ID in
	// the incoming request data.
	if requestData.TaskID == "" {
		http.Error(w, "Task ID not provided", http.StatusBadRequest)
		return
	}

	// The above code is reading tasks from a JSON file located at "./data/tasks.json" using the
	// `readTasksFromFile` function. If there is an error reading the tasks from the file, it will return
	// a 500 Internal Server Error response with the message "Error reading existing tasks".
	tasksFile := "./data/tasks.json"
	existingTasks, err := readTasksFromFile(tasksFile)
	if err != nil {
		http.Error(w, "Error reading existing tasks", http.StatusInternalServerError)
		return
	}

	// The code is deleting an entry from a map called `existingTasks` using the key `requestData.TaskID`.
	delete(existingTasks, requestData.TaskID)

	// The above code is attempting to write the existing tasks to a file named `tasksFile` using the
	// `writeTasksToFile` function. If an error occurs during the writing process, it will return a 500
	// Internal Server Error response with the message "Error writing updated tasks".
	err = writeTasksToFile(existingTasks, tasksFile)
	if err != nil {
		http.Error(w, "Error writing updated tasks", http.StatusInternalServerError)
		return
	}

	// The above code is written in Go programming language and it is setting HTTP headers for CORS
	// (Cross-Origin Resource Sharing) and content type. It allows requests from a specific origin
	// (*), specifies allowed headers (Content-Type), sets the content type to JSON,
	// and then writes a response to the client with a message indicating that a task with a specific ID
	// has been deleted successfully.
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task with ID %s deleted successfully", requestData.TaskID)
}

// The `removeContact` function handles deleting a contact from a JSON file based on the provided
// contact ID in a Go HTTP server.
func removeContact(w http.ResponseWriter, r *http.Request) {

	// The above code is checking if the HTTP request method is not DELETE. If the method is not DELETE, it
	// returns a "Method not allowed" error with status code 405 (Method Not Allowed).
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// The above code is reading the request body from an HTTP request in a Go program. It uses the
	// `io.ReadAll` function to read the entire request body into a byte slice named `body`. If there is an
	// error while reading the request body, it will return an internal server error response with the
	// message "Error reading request body".
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// The above code is written in Go and it is attempting to unmarshal JSON data from the `body` variable
	// into a struct named `requestData`. The struct has a single field `ID_contact` with a tag specifying
	// the JSON key to map to. If there is an error during the unmarshalling process, it will return a
	// "Invalid JSON format" error with a status code of 400 (Bad Request).
	var requestData struct {
		ID_contact string `json:"ID_contact"`
	}
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// The above code is checking if the `ID_contact` field in the `requestData` is empty. If it is empty,
	// it will return a HTTP 400 Bad Request error with the message "Contact ID not provided".
	if requestData.ID_contact == "" {
		http.Error(w, "Contact ID not provided", http.StatusBadRequest)
		return
	}

	// The above code snippet is written in Go programming language. It is attempting to read contacts data
	// from a JSON file located at "./data/contacts.json". It first tries to read the existing contacts
	// data from the file using the `readContactsFromFile` function. If there is an error during the
	// reading process, it will return an HTTP 500 Internal Server Error response with the message "Error
	// reading existing tasks".
	contactsFile := "./data/contacts.json"
	existingContact, err := readContactsFromFile(contactsFile)
	if err != nil {
		http.Error(w, "Error reading existing tasks", http.StatusInternalServerError)
		return
	}

	// The code is deleting a contact with the ID specified in the `requestData.ID_contact` from the
	// `existingContact` data structure.
	delete(existingContact, requestData.ID_contact)

	// The above code snippet is attempting to write the existing contact information to a file named
	// `contactsFile`. If an error occurs during the write operation, it will return an HTTP error response
	// with the message "Error writing updated tasks" and a status code of 500 (Internal Server Error).
	err = writeContactsToFile(existingContact, contactsFile)
	if err != nil {
		http.Error(w, "Error writing updated tasks", http.StatusInternalServerError)
		return
	}

	// The above code is written in Go programming language and is setting HTTP headers for CORS
	// (Cross-Origin Resource Sharing) and content type. It allows requests from a specific origin
	// (https://join.denniscodeworld.de), specifies allowed headers (Content-Type), sets the content type to JSON,
	// and sends a response with a success message indicating that a task with a specific ID has been
	// deleted successfully.

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task with ID %s deleted successfully", requestData.ID_contact)
}
