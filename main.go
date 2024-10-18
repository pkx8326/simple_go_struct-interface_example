package main

import (
	"fmt"

	"example.com/struct_interface/ops"
)

/////////////////////////////////////////////////////

func main() {
	notetitle := ops.Getinfo("Please input the note's title: ")
	notecontent := ops.Getinfo("Please input the note's content: ")
	newnote := ops.NoteConstructor(notetitle, notecontent)
	todotext := ops.Getinfo("Please input a todo :")
	newtodo := ops.TodoConstructor(todotext)

	ops.Displaydata(newnote)
	ops.Displaydata(newtodo)

	err := ops.Savejson(newnote)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Save note successfully.")
	}

	err = ops.Savejson(newtodo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Save todo successfully.")
	}
}
