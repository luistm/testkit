### TestKit

TestKit provides tools for testing Go code.

```go
import "github.com/luistm/testkit"

func TestSample(t *testing.T){
	testkit.AssertEqual(t, 1,1)
	testkit.AssertIsNil(t, nil)
}
```
