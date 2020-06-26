package common

import (
	"errors"
	"fmt"

	"github.com/rollbar/rollbar-go"
)

func Rescue() {
	if r := recover(); r != nil {
		err := errors.New(fmt.Sprintf("%+v", r))
		fmt.Println(err)
		rollbar.Critical(err)
		rollbar.Wait()
	}
}
