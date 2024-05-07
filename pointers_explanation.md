**Deep-dive on Pointers material**


At a high level, this code translates a JSON string into an array of strings. The JSON input - `"["John", "Hafsah", "Sandy"]"` - is converted into a this format in Go: `["John", "Hafsah", "Sandy"]`. This enables us to operate on the array in future code (iterating over it, indexing etc.)

```go
package main
import(
  "fmt"
  "encoding/json"
)
```

These lines set up the Go file.

`Package main` declares that the file belongs to the main package (a compulsory part of a Go project, containing the code for the executable application).

The subsequent lines import the required external packacges needed to run this code: `Fmt` (used for printing and manipulating strings), and `Encoding/json` (used here for converting json data into a format Go can work with).

```go
func main(){
```

This line sets up the main function (ie. anything that we want to execute in our program).

```go
    jsonString := `["John", "Hafsah", "Sandy"]`
    jsonBytes := []byte(jsonString)
```

These lines declare and initialise two variables: `jsonString` (assigned to a JSON array of strings), and `jsonBytes`.

The jsonBytes variable is assigned to a byte slice, ie. a slice containing each character of the JSON data as individual 'bytes'. Each byte corresponds to a UTF-8 number (similar to the ASCII table system where a character has a corresponding number). Using UTF-8 encoding, Go can start to read these characters in their new 'byte' form.

```go
    var namesList []string
```

This line declares a new variable, `namesList`, which is a currently empty slice, designed to contain strings.

```go
    err := json.Unmarshal(jsonBytes, &namesList)
```

This line uses a function imported from the external `encoding/json` package, Unmarshal. Unmarshal is used to convert the UTF-8-encoded data it receives from `jsonBytes` into alphabetical form, and to store the resulting data in the namesList slice created in the previous line.

Crucially, a namesList argument given to the Unmarshal function needs to be a **pointer** for the code to work. The pointer directs (or 'points') the code to the original namesList and is represented in the '&' before 'namesList'. This ensures that the namesList array is changed in-place, instead of a copy of namesList being created and leaving the original namesList as an unchanged empty string. 

Without the pointer, a new copy of namesList (ie. an empty array) would be created, Unmarshal will recognise that the namesList variable is a copy rather than a pointer to the original, and it will not be able to pass its decoded data to `namesList`. Consequently, an error message will be printed to the terminal instead of our intended result. 

Aside from Including pointers make Go code more memory-efficient (ie. saving memory space but not creating new copies), and more scalable (eg. if our jsonString was a much larger list of names, creating a copy of the whole structure would be inefficient and can affect performance)

```go
  if err != nil {
        fmt.Println(err)
    }
```

These two lines, along with the previous line, set up error handling for the Unmarshal function: if our attempt to unmarshal (decode) the JSON data fails, an error message should be printed to the terminal, using the Println function from the imported `fmt` package.

If no error is encountered during the Unmarshal operation (ie. the data is successfully converted from jsonByte form to Go string/array form, and err == nil), no error message is printed and the next line of code is executed...

```go
  fmt.Println(namesList)
}
```

This line uses the imported `fmt` package to prints the original `namesList` (thanks to the pointer included earlier); a Go array of strings converted from the JSON data.
