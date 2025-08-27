package auxiliar

import (
	"fmt"

	"github.com/badoux/checkmail"
)

// "Write" is a function that writes a message to the console.
func Write() {
	fmt.Println("Writing from auxiliary package")
	write2()
	// Since it's in the same package, the "write2" function can be accessed.

	err := checkmail.ValidateFormat("teste@gmail.com")
	fmt.Println(err) // <nil> if the email is valid
	if err != nil {
		fmt.Println(err)
	}

}
