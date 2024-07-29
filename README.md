# Go Structs Example

This project provides a basic example of defining and using structs in Go. It demonstrates how to create structs, embed one struct within another, and define methods to interact with struct fields.

## Overview

The program illustrates the following concepts:
1. Defining structs with fields.
2. Embedding one struct within another.
3. Defining methods on structs.
4. Creating and manipulating struct instances.

## Usage

To run the program, ensure you have Go installed on your system. Then, follow these steps:

1. Save the code to a file, for example, `main.go`.
2. Open a terminal and navigate to the directory containing `main.go`.
3. Run the following command to execute the program:

```bash
go run main.go
```

## Code Explanation

The main components of the code are:

1. **Defining Structs:**
   ```go
   type contactInfo struct {
       email    string
       zipCode  int
   }

   type person struct {
       firstName string
       lastName  string
       contactInfo
   }
   ```

   - `contactInfo` is a struct that holds contact details, such as `email` and `zipCode`.
   - `person` is a struct that includes personal information (`firstName` and `lastName`) and embeds the `contactInfo` struct.

2. **Creating and Initializing Struct Instances:**
   ```go
   func main() {
       jim := person{
           firstName: "jim",
           lastName:  "party",
           contactInfo: contactInfo{
               email:   "ali@j.com",
               zipCode: 12345,
           },
       }
       jim.update("sooz")
       jim.print()
   }
   ```

   - An instance of `person` named `jim` is created and initialized with values for `firstName`, `lastName`, and `contactInfo`.
   - The `update` method is called to change `firstName` to "sooz".
   - The `print` method is called to display the `person` struct's details.

3. **Defining Methods on Structs:**
   ```go
   func (p person) print() {
       fmt.Printf("%+v", p)
   }

   func (p *person) update(name string) {
       (*p).firstName = name
   }
   ```

   - `print` is a method that takes a `person` value receiver and prints the struct's details.
   - `update` is a method that takes a pointer receiver (`*person`) and updates the `firstName` field.

## Notes

- Structs are used to group related data together.
- Method receivers allow you to define methods that operate on structs. A value receiver (e.g., `person`) copies the struct, while a pointer receiver (e.g., `*person`) allows you to modify the original struct.
- Embedded structs allow you to compose structs and reuse fields, improving code organization.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
