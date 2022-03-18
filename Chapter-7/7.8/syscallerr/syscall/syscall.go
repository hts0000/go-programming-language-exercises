package syscall

import "fmt"

type Erron uintptr

var errors = [...]string{
	1: "operation not permitted",   // EPERM
	2: "no such file or directory", // ENOENT
	3: "no such process",           // ESRCH
	// ...
}

func (e Erron) Error() string {
	if 0 <= int(e) && int(e) < len(errors) {
		return errors[e]
	}
	return fmt.Sprintf("errno %d", e)
}
