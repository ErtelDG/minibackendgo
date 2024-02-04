package main

import (
	"encoding/json"
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
	http.ListenAndServe(":8080", nil)
}

type Contact struct {
	ID_contact string
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

func contacts(w http.ResponseWriter, r *http.Request) {
	read_file, err := os.ReadFile("./data/contacts.json")
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
	newContact := &Contact{ID_contact: newContactID}
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(existingContacts)
}

func categories(w http.ResponseWriter, r *http.Request) {
	read_file, err := os.ReadFile("./data/categories.json")
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

// Helper function to read existing contacts from a file
func readContactsFromFile(filePath string) (map[string]*Contact, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	contacts := make(map[string]*Contact)
	err = json.Unmarshal(data, &contacts)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

// Helper function to write contacts to a file
func writeContactsToFile(contacts map[string]*Contact, filePath string) error {
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
