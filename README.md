# Newton's Method

```go
// find an approximate solution to `x^2=612`
var (
    ƒ     = func(x float64) float64 { return x * x }
    dƒ    = func(x float64) float64 { return 2 * x }
    value = 612.0
)

// iterate 5 times, start with approximation
// of 10
approx := 10.0
iterate := 10

soln := Apply(value, ƒ, dƒ, iterate, approx)

fmt.Printf("sqrt(%.1f)=%.4f", value, truncateDigits(soln, 4))
// Output:
// sqrt(612.0)=24.7386
```
