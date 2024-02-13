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

func main() {
	http.HandleFunc("/contacts", contacts)
	http.HandleFunc("/categories", categories)
	http.HandleFunc("/tasks", tasks)
	http.HandleFunc("/add_contact", add_contact)
	http.HandleFunc("/remove_contact", removeContact)
	http.HandleFunc("/add_task", add_task)
	http.HandleFunc("/del_task", deleteTask)
	http.HandleFunc("/update_task", updateTask)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("TLS Handshake Error:", err)
		panic(err)
	}
}

func contacts(w http.ResponseWriter, r *http.Request) {
	read_file, err := os.ReadFile("./data/contacts.json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = w.Write(read_file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func add_contact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Read existing contacts from file
	contactsFile := "./data/contacts.json"
	existingContacts, err := readContactsFromFile(contactsFile)
	if err != nil {
		http.Error(w, "Error reading existing contacts", http.StatusInternalServerError)
		return
	}

	// Create a new contact with a unique ID
	newContactID := "cont" + strconv.FormatInt(time.Now().Unix(), 10)
	newContact := &structures.Contact{ID_contact: newContactID}
	err = json.NewDecoder(r.Body).Decode(newContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the new contact to the existing contacts
	existingContacts[newContactID] = newContact

	// Write the updated contacts back to the file
	err = writeContactsToFile(existingContacts, contactsFile)
	if err != nil {
		http.Error(w, "Error writing updated contacts", http.StatusInternalServerError)
		return
	}

	// Respond with the updated contacts
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(existingContacts)
}

func categories(w http.ResponseWriter, r *http.Request) {
	read_file, err := os.ReadFile("./data/categories.json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = w.Write(read_file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func tasks(w http.ResponseWriter, r *http.Request) {
	read_file, err := os.ReadFile("./data/tasks.json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = w.Write(read_file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func add_task(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Read existing tasks from file
	tasksFile := "./data/tasks.json"
	existingTasks, err := readTasksFromFile(tasksFile)
	if err != nil {
		http.Error(w, "Error reading existing tasks", http.StatusInternalServerError)
		return
	}

	// Create a new task with a unique ID
	newTaskID := "tk" + strconv.FormatInt(time.Now().Unix(), 10)
	newTask := &structures.Task{ID_task: newTaskID}
	err = json.NewDecoder(r.Body).Decode(newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the new contact to the existing tsks
	existingTasks[newTaskID] = newTask

	// Write the updated tsks back to the file
	err = writeTasksToFile(existingTasks, tasksFile)
	if err != nil {
		http.Error(w, "Error writing updated tasks", http.StatusInternalServerError)
		return
	}

	// Respond with the updated contacts
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(existingTasks)
}

// Handler-Funktion für den Endpunkt "/update_task"
func updateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Lese den Request-Body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Dekodiere den JSON-Body in einen Task
	var task structures.Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Lese vorhandene Tasks aus der Datei
	tasksFile := "./data/tasks.json"
	existingTasks, err := readTasksFromFile(tasksFile)
	if err != nil {
		http.Error(w, "Error reading existing tasks", http.StatusInternalServerError)
		return
	}

	// Aktualisiere den Task oder füge ihn hinzu
	existingTasks[task.ID_task] = &task

	// Schreibe die aktualisierten Tasks zurück in die Datei
	if err := writeTasksToFile(existingTasks, tasksFile); err != nil {
		http.Error(w, "Error writing updated tasks", http.StatusInternalServerError)
		return
	}

	// Bestätige die erfolgreiche Aktualisierung
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Task updated successfully")
}

// Helper function to read existing contacts from a file
func readContactsFromFile(filePath string) (map[string]*structures.Contact, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	contacts := make(map[string]*structures.Contact)
	err = json.Unmarshal(data, &contacts)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

// Helper function to read existing tasks from a file
func readTasksFromFile(filePath string) (map[string]*structures.Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	tasks := make(map[string]*structures.Task)
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Helper function to write contacts to a file
func writeContactsToFile(contacts map[string]*structures.Contact, filePath string) error {
	data, err := json.MarshalIndent(contacts, "", "   ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Helper function to write tasks to a file
func writeTasksToFile(tasks map[string]*structures.Task, filePath string) error {
	data, err := json.MarshalIndent(tasks, "", "   ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Lese den Request-Body, um die Task-ID zu extrahieren
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Dekodiere den JSON-Body, um die Task-ID zu erhalten
	var requestData struct {
		TaskID string `json:"task_id"`
	}
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Überprüfe, ob die Task-ID bereitgestellt wurde
	if requestData.TaskID == "" {
		http.Error(w, "Task ID not provided", http.StatusBadRequest)
		return
	}

	// Lese die vorhandenen Aufgaben aus der Datei
	tasksFile := "./data/tasks.json"
	existingTasks, err := readTasksFromFile(tasksFile)
	if err != nil {
		http.Error(w, "Error reading existing tasks", http.StatusInternalServerError)
		return
	}

	// Entferne die Aufgabe mit der angegebenen Task-ID
	delete(existingTasks, requestData.TaskID)

	// Schreibe die aktualisierten Aufgaben zurück in die Datei
	err = writeTasksToFile(existingTasks, tasksFile)
	if err != nil {
		http.Error(w, "Error writing updated tasks", http.StatusInternalServerError)
		return
	}

	// Bestätige, dass die Aufgabe erfolgreich gelöscht wurde
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task with ID %s deleted successfully", requestData.TaskID)
}

func removeContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Lese den Request-Body, um die Contact-ID zu extrahieren
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Dekodiere den JSON-Body, um die Contact-ID zu erhalten
	var requestData struct {
		ID_contact string `json:"ID_contact"`
	}
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Überprüfe, ob die Contact-ID bereitgestellt wurde
	if requestData.ID_contact == "" {
		http.Error(w, "Contact ID not provided", http.StatusBadRequest)
		return
	}

	// Lese die vorhandenen Aufgaben aus der Datei
	contactsFile := "./data/contacts.json"
	existingContact, err := readContactsFromFile(contactsFile)
	if err != nil {
		http.Error(w, "Error reading existing tasks", http.StatusInternalServerError)
		return
	}

	// Entferne die Aufgabe mit der angegebenen Contact-ID
	delete(existingContact, requestData.ID_contact)

	// Schreibe die aktualisierten Aufgaben zurück in die Datei
	err = writeContactsToFile(existingContact, contactsFile)
	if err != nil {
		http.Error(w, "Error writing updated tasks", http.StatusInternalServerError)
		return
	}

	// Bestätige, dass die Aufgabe erfolgreich gelöscht wurde
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task with ID %s deleted successfully", requestData.ID_contact)
}
