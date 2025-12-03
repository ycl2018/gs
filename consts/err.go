package consts

import "fmt"

type CrashError struct {
	VmStack   []byte
	CodeTrace string
}

func (c *CrashError) Error() string {
	return fmt.Sprintf("code trace:\n%s\nvm stack:\n%s\n", c.CodeTrace, c.VmStack)
}
