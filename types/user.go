// when working with multiple files we split files into packages
// the root folder package is main, any sub folders have their own package scope

package types

type User struct {
	Username string // first capital letter means public variable, lowercase private/unaccessible out of the package
	Age      int
}
