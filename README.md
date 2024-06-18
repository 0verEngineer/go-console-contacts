## CLI contact book
This is a CLI contact book created in go with cobra for learning purposes.
It is the first Program i created in GO.

- The books are created as text files, every line is one contact
- The books are created at the same path as the executable because this was only a learning project

### Usage
#### Help:

`./console-contacts help`

Output:
```
Usage:
console-contacts [command]

Available Commands:
add         Add a new contact to the specified contact book
clear       Deletes all saved contacts from specified contact book
completion  Generate the autocompletion script for the specified shell
create      Creates a new contact book
delete      Deletes the specified contact from the specified contact book
delete-book Deletes the specified contact book
help        Help about any command
list        Lists all contacts in the given book
version     Print the version number of the application.

Flags:
-h, --help   help for console-contacts
```

---
#### Creating a book:

`./console-contacts create book1`

Output:
```
Created new book at: /home/username/Documents/test/book2.txt
```

---
#### Add a contact

`./console-contacts add book1 Max Mustermann max@mustermann.io +49123456789`

---
#### List contacts

`./console-contacts list book1`

Output:
```

ID First name Last name  E-Mail            Phone number
1  Max1       Mustermann max@mustermann.io +49123456789
2  Max2       Mustermann max@mustermann.io +49123456789
3  Max3       Mustermann max@mustermann.io +49123456789
```

---
#### Delete a contact

`./console-contacts delete book1 3`

- This deletes the contact with the ID 3 from the book

