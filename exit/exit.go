/*Package exit implements exit stuff*/
package exit

import (
	"os"
)

/*When using os.Exit the defer func will not be respected. Use recover in a Safe Exit function*/

/*Code to return, use this wrapper if you want to raise a panic with a int*/
type Code struct{ int }

/*
Safe catch the panic and do a clean os.Exit with the given code or raise the panic
You need this to respect the defer functions of your program. Must be the very first defer in your main
*/
func Safe() {
	if e := recover(); e != nil {
		if exit, ok := e.(Code); ok == true {
			os.Exit(exit.int)
		}
		panic(e) // not an Exit, bubble up
	}
}

/*
Exit is a wrapper for a simple safe exit with os codes
*/
func Exit(e int) {
	panic(Code{e})
}
