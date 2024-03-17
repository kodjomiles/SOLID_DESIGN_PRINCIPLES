# Journal Management System

The Journal Management System is a simple, lightweight tool designed for creating, managing, and persisting personal or professional journal entries. It is built using Go and emphasizes the principle of Single Responsibility by separating concerns between journal entry management and data persistence.

## Features

- **Add Entries**: Easily add new entries to your journal with automatic numbering for easy tracking.
- **Print Entries**: Print all entries in the journal to the console for a quick review.
- **Remove Entries**: Functionality to remove entries by index (Implementation to be added).
- **Persistence**: Save your journal entries to a file for permanent storage and load them back into the system when needed.
- **Extensibility**: Designed with the Single Responsibility Principle (SRP) in mind, allowing for easy extension and modification without overcomplicating the journal's core functionalities.

## Components

### Journal

The `Journal` struct is the core of the system, responsible for managing journal entries. It allows adding new entries, printing all entries, and (with future implementation) removing specific entries.

### Persistence

To adhere to the SRP, journal persistence functions are separated from the `Journal` struct. Persistence is managed by:

- A standalone function `SaveToFile`, for a functional approach.
- A `Persistence` struct, offering a more structured approach with the possibility to customize the line separator used in saved files.

## Usage

1. **Creating a Journal**: Instantiate a `Journal` object to start adding entries.
2. **Adding Entries**: Use `AddEntry` to add new journal entries. Each entry is automatically prefixed with an incrementing number.
3. **Printing Entries**: Call `Print` to display all current journal entries in the console.
4. **Saving Entries**: Use either the standalone `SaveToFile` function or an instance of the `Persistence` struct to save journal entries to a file. This ensures that journal management and data persistence are kept as separate responsibilities.

### Example

```go
j := Journal{}
j.AddEntry("I ate banku today")
j.AddEntry("I think the new keyword works too")
j.Print()

// Saving to file using a functional approach
SaveToFile(&j, "functional_save.txt")

// Saving to file using the Persistence struct
p := Persistence{"\n"}
p.SaveToFile(&j, "package_save.txt")
```

## Extensibility

The system is designed for easy modification and extension. For instance, additional functionalities like entry removal, loading entries from a web source, and more sophisticated data persistence strategies can be seamlessly integrated without interfering with the journal's primary responsibilities.

## Conclusion

The Journal Management System offers a straightforward yet flexible way to manage and persist journal entries, adhering to the principles of clean code and single responsibility. Whether for personal use or as a building block for more complex systems, it provides a solid foundation with room for growth and customization.