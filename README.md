# Stopwatch

**Stopwatch** is a Go package intended for **simple stepwise benchmarking** of code. It was written as a personal project of mine to familiarize with the language, feel free to contribute.

## Installation

```
$ go get github.com/mpragliola/stopwatch
```

## Usage

The usage is very simple:

* `import` the package in your `.go` file

* instantiate a stopwatch with `stopwatch.New()` (multiple instances will allow for independent measurements)

* mark checkpoints with `.Mark(label)`

* get the measurement **directly as text dump** (`.Dump`) or as `struct` (`.Data()`). In either case you will be able to read a table with a checkpoint's **label**, **timestamp** and **delta timestamp** from the previous one.

```go
package foo

import "github.com/mpragliola/stopwatch"

// ... code

w := stopwatch.NewStart()
// code
w.Mark("Aggregated")
// code
w.Mark("Transformed")
// code
w.Dump()
```

Will output something similar to: 

```
(Start)          |       0ms |       0ms |
Aggregated       |      35ms |     +35ms |      
Transformed      |      85ms |     +40ms |
(End)            |     113ms |     +28ms | 
```

#### As structured data

`w.Dump()` is a convenience method to output results immediately; we can also use them programmatically with `w.Data()`.
It will return a `[]stopwatch.DataRow` slice, where each `DataRow` item mirrors the above data in the `.Label`, `.At` and `.Delta` fields.


  ```go
  data := w.GetData()

  fmt.Println(data[1].At) // time.Millisecond * 35
  fmt.Println(data[2].Label) // "Transformed"
  ```

> The time fields will be an instance of `time.Duration` rounded to `time.Millisecond`.

#### Halted stopwatch

`stopwatch.NewStart()` creates an "active" stopwatch, where the start time is already marked as a first entry and the clock is "ticking"; use `stopwatch.New()` instead to create a "halted" stopwatch; it will begin to count not from creation time but from the time you `.Mark()` the first time stamp:

```go
w := stopwatch.New()
// code
w.Mark("Aggregated")
// code
w.Mark("Transformed")
// code
w.Dump()
```

Will output something similar to (notice the absence of `Start`): 

```
Aggregated       |       0ms |      +0ms |      
Transformed      |      40ms |     +40ms |
(End)            |      68ms |     +28ms | 
```

#### Json

You can also get a JSON representation of the stopwatch with `w.Json()`.
