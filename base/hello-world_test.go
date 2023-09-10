package base_test

import (
	"testing"

	"github.com/continuous1024/examples/base"
)

func TestHelloWorld(t *testing.T) {
	t.Run("First test", func(t *testing.T) {
		base.HelloWorld()
	})
}
