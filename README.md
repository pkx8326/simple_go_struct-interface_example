# Simple Go Struct-Interface Example

### Overview
This Go program demonstrates that structs of different types can be passed into common functions with the help of the interface type variable. This program is an improvement based on a [simpler program](https://github.com/pkx8326/simple_go_json_note-taking) in which no interface is used.

### Program manual
When run, the program asks the user to input the following information in the following order:

- Note title
- Note content
- Todo

Then, the program will show messages containing the input information and, if there's no error, will notify the user that the note and todo are successfully saved (in json file format).

There is no input validation for this program because every piece of information are in the form of free text. However, the program is designed to catch error when saving the files. The user will be notified if there's any error while saving each file. In case of error, the program will stops after displaying the error message.

### Code structure
 The codes of this projects are organized into the code for the main program (stored in the ```main.go``` file) and the codes for creating structs and functions to get, display, and save user inputs stored in the ops package folder which contains ```ops.go``` file.

In this program the display and save functions' arguments are of ```interface``` type. This means that any type of variable can be passed into the same function and those functions can be used as common functions to perform operations on variables of different types.

### Program flow
1. The user inputs the note title as a string
2. The user inputs the note content as a string
3. The program takes those inputs to create a struct which stores the note title, content, and the timestamp at its creation
4. The user inputs the todo text
5. The program takes the todo text input to create another struct which stores the todo text (without any title)
6. The program displays messages to confirm the note's title and its content from the inputs
7. The program displays a message to confirm the todo
8. The program displays a message to notify the user that it's saving the note
9. The program attempts to save the inputs as a json file with json field names according to the struct tags given in the code
10. The program displays the message that it saves the file successfully
11. The program repeat the same process from 8. for the todo (the todo's file name is already hard-coded in the program and can't be changed)
