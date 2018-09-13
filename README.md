# primelist

This library uses sieve-of-eratosthenes algorithm to genenrate list of prime numbers within given size.

```go

package main

import (
  "fmt"

  "github.com/iAziz786/primelist"
)

func main() {
  primes := primelist.Generate(200)
  fmt.Println(primes)
  // [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]
}
```