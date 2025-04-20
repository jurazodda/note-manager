package main

import (
	"bufio"
	"fmt"
	"os"
)

type Note struct {
	ID      int
	Title   string
	Content string
}

var notes []Note
var nextID int = 1

func main() {
	fmt.Println("__ Console Notes manager __")

	for {
		fmt.Println("1 - Create a note.")
		fmt.Println("2 - List all notes.")
		fmt.Println("3 - Get a note by ID.")
		fmt.Println("4 - Update note.")
		fmt.Println("5 - Delete note.")
		fmt.Println("6 - Exit.")
		fmt.Println()
		fmt.Print("Enter your choice: ")

		var userChoice int
		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			createNote()
		case 2:
			getNotes()
		case 3:
			getNoteByID()
		case 4:
			updateNote()
		case 5:
			deleteNote()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid command.")
			fmt.Println()
			continue
		}
	}
}

func createNote() {
	fmt.Print("Enter the title of the note: ")
	reader := bufio.NewReader(os.Stdin)
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter the content of the note: ")
	reader = bufio.NewReader(os.Stdin)
	content, _ := reader.ReadString('\n')

	notes = append(notes, Note{
		ID:      nextID,
		Title:   title,
		Content: content,
	})
	nextID++

	fmt.Println("Note crated.")
	fmt.Println()
}

func getNotes() {
	fmt.Println("Your notes:")
	for i := range notes {
		fmt.Printf("ID: %d, Title: %s, Content: %s\n", notes[i].ID, notes[i].Title, notes[i].Content)
	}

	if len(notes) == 0 {
		fmt.Println("You have no notes.")
	}
	fmt.Println()
}

func getNoteByID() {
	fmt.Print("Enter the ID of the note: ")
	var noteID int
	fmt.Scan(&noteID)
	for i := range notes {
		if notes[i].ID == noteID {
			fmt.Printf("ID: %d, Title: %s, Content: %s\n", notes[i].ID, notes[i].Title, notes[i].Content)
			fmt.Println()
			return
		}
	}

	fmt.Println("Note not found.")
	fmt.Println()
}

func updateNote() {
	fmt.Print("Enter the ID of the note: ")
	var noteID int
	fmt.Scan(&noteID)

	for i := range notes {
		if notes[i].ID == noteID {
			fmt.Print("Enter the title of the note: ")
			reader := bufio.NewReader(os.Stdin)
			title, _ := reader.ReadString('\n')

			fmt.Print("Enter the content of the note: ")
			reader = bufio.NewReader(os.Stdin)
			content, _ := reader.ReadString('\n')

			notes[i] = Note{
				ID:      nextID,
				Title:   title,
				Content: content,
			}
			nextID++
			fmt.Println("Note updated")
			fmt.Println()
			return
		}
	}

	fmt.Println("Note not found.")
	fmt.Println()
}

func deleteNote() {
	fmt.Println("Enter the ID of the note: ")
	var noteID int
	fmt.Scan(&noteID)

	for i := range notes {
		if notes[i].ID == noteID {
			notes = append(notes[:i], notes[i+1:]...)
			fmt.Println("Note deleted.")
			fmt.Println()
			return
		}
	}

	fmt.Println("Note not found.")
	fmt.Println()
}
