package uu

//Err contains information about the failure
type Err string

//Error returns printable string of the failure
func (meErr Err) Error() string {
	return string(meErr)
}

const (
	//ErrInvalidLineLen is invalid line length
	ErrInvalidLineLen = Err("Invalid line length!")
	//ErrInvalidDataLen is invalid data length
	ErrInvalidDataLen = Err("Invalid data length!")
	//ErrInvalidData is invalid data
	ErrInvalidData = Err("Invalid data!")
)
