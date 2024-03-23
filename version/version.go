package version

import (
	"fmt"
	r "runtime"
)

func Version() {
	fmt.Println(r.Version())
}
