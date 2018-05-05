### Testkit

Testkit provides tools for testing Go code.

```go
func TestSample(t *testing.T){
	testkit.AssertEqual(t, 1,1)
	testkit.AssertIsNil(t, nil)
}
```
