# Password Generator

A simple and customizable password generator written in Go. This program generates secure, random passwords with options to include numbers, special characters, and uppercase letters.

## Features

- Generate random passwords with a customizable length.
- Option to include:
  - Numbers
  - Special characters
  - Uppercase letters
- Secure randomization using `crypto/rand`.

## Requirements

- Go 1.16 or later.

## Usage

1. Clone the repository or save the code as a `.go` file.
2. Build or run the code using `go run` or `go build`.

### Command-line Flags

The program accepts the following flags:

- `-length`: Specify the length of the password (default: `12`).
- `-numbers`: Include numbers in the password (default: `true`).
- `-specials`: Include special characters in the password (default: `true`).
- `-uppercase`: Include uppercase letters in the password (default: `true`).

### Example

Generate a password with the default settings (12 characters, including numbers, special characters, and uppercase letters):

```bash
go run main.go
```

Output:

```bash
Generated Password: a8$B7c!KpT6q
```

Generate a password with a custom length of 20 and exclude special characters:

```bash
go run main.go -length 20 -specials=false
```

Output:

```bash
Generated Password: aQ8z2X9tN4LpK7dMfRg5
```
### Code Overview

Main Functionality
* The program reads user input through command-line flags to determine password length and character set preferences.
* The generatePassword function creates a random password based on the specified settings.

Randomization
* The program uses the crypto/rand package for secure random number generation.
* Password characters are selected from a customizable charset built based on user preferences.
