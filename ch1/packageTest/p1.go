package packageTest

import "fmt"

type Id int

func (i Id) String() string {
	return fmt.Sprintf("%g", i)
}
