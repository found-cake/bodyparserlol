package errors

import (
	"fmt"

	"github.com/found-cake/bodyparserlol/config"
)

type CTFError struct {
	error
}

func (CTFError) Error() string {
	return fmt.Sprintf("flag is %s", config.FLAG)
}
