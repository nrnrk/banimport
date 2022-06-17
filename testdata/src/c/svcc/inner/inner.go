package inner

import (
	innerb "c/svcb/inner" // want `import of "c/svcb/inner" prohibited`
)

func DoNothing() {
	innerb.DoNothing()
}
