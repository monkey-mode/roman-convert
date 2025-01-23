# Roman Numerals Converter

This project provides functions to convert between Roman numerals and Arabic numbers in Go.

## Functions

### `romanToArabic`

Converts a Roman numeral string to an Arabic number.

#### Example

```go
fmt.Println(romanToArabic("X"))     // 10
fmt.Println(romanToArabic("XXI"))   // 21
fmt.Println(romanToArabic("IV"))    // 4
fmt.Println(romanToArabic("XXXIX")) // 39
```

### `arabicToRoman`

Converts a Roman numeral string to an Arabic number.

#### Example

```go
fmt.Println(arabicToRoman(10))  // X
fmt.Println(arabicToRoman(21))  // XXI
fmt.Println(arabicToRoman(4))   // IV
fmt.Println(arabicToRoman(39))  // XXXIX
```

## Running Tests

To run the tests for these functions, use the following command:

```bash
go test -v
```

## Example Usage

Here's an example of how to use these functions in your main.go file:


```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(romanToArabic("X"))     // 10
    fmt.Println(romanToArabic("XXI"))   // 21
    fmt.Println(romanToArabic("IV"))    // 4
    fmt.Println(romanToArabic("XXXIX")) // 39

    fmt.Println(arabicToRoman(10))  // X
    fmt.Println(arabicToRoman(21))  // XXI
    fmt.Println(arabicToRoman(4))   // IV
    fmt.Println(arabicToRoman(39))  // XXXIX
}
```