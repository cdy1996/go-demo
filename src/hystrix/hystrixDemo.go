package hystrix

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
)

func main() {

	_ = hystrix.Do("my_command", func() error {
		return nil
	}, func(err error) error {
		return nil
	})

	output := make(chan bool, 1)

	errors := hystrix.Go("my_command", func() error {
		output <- true
		return nil
	}, func(err error) error {
		return nil
	})
	select {
	case out := <-output:
		fmt.Printf("%v", out)
	//
	case err := <-errors:
		fmt.Printf("%v", err)

	}

}
