package ops

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

//////////////////////////////////////////////////////

type note struct {
	Title     string    `json:"note_title"`
	Content   string    `json:"note_content"`
	CreatedAt time.Time `json:"created_at"`
}

type todo struct {
	Todo string `json:"todo"`
}

//////////////////////////////////////////////////////

func Getinfo(prompt string) (info string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		info, _ = reader.ReadString('\n')
		info = strings.TrimSpace(info) // This line is to remove the blank line createed from the /n and to remove preceeding and trailing spaces
		if info == "" {
			fmt.Println("This information can't be left blank. Please try again.")
			continue
		} else {
			return
		}
	}
}

func NoteConstructor(notetitle, notecontent string) (newnote note) {
	newnote = note{
		Title:     notetitle,
		Content:   notecontent,
		CreatedAt: time.Now(),
	}
	return
}

func TodoConstructor(todotext string) (newtodo todo) {
	newtodo = todo{
		Todo: todotext,
	}
	return

}

func Displaydata(data interface{}) {
	switch dtype := data.(type) {
	case note:
		fmt.Println("Your note title is :", dtype.Title)
		fmt.Println("Your note text is :", dtype.Content)
	case todo:
		fmt.Println("Your todo is :", dtype.Todo)
	}
}

func Savejson(data interface{}) (err error) {

	switch dtype := data.(type) {
	case note:
		filename := strings.ReplaceAll(dtype.Title, " ", "_")
		filename = strings.ToLower(filename) + ".json"
		json, saveError := json.Marshal(dtype) //The .Marshal() method works only if fields in the struct are publicly accessible
		if saveError != nil {
			err = fmt.Errorf("error saving note file")
			return
		} else {
			fmt.Println("Saving the note to json...")
			return os.WriteFile(filename, json, 0644)
		}

	case todo:
		filename := "todo.json"
		json, saveError := json.Marshal(dtype)
		if saveError != nil {
			err = fmt.Errorf("error saving todo file")
			return
		} else {
			fmt.Println("Saving todo file...")
			return os.WriteFile(filename, json, 0644)
		}
	default:
		err = fmt.Errorf("error! Unsupported data type: %T", dtype)
		return
	}
}
