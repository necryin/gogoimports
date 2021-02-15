# gogoimports collapse imports to two groups: stdlib and other.

Install: `go get github.com/necryin/gogoimports`

**For example:**
```go
import (
    "github.com/prometheus/client_golang/prometheus"
    "flag"
    
    "os"
    "os/signal"
    "syscall"

    "go.uber.org/zap"
)
```
**became** 
```go
import (
    "flag"
    "os"
    "os/signal"
    "syscall"

    "github.com/prometheus/client_golang/prometheus"
    "go.uber.org/zap"
)
```

**NB! Should be used together with _goimports_ and called after it**
