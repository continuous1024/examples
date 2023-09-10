package base_test

import (
	"testing"

	"github.com/continuous1024/examples/base"
)

func TestMaxHeapify(t *testing.T) {
	t.Run("Max heapify", func(t *testing.T) {
		data := []string{"1", "2", "3"}
		result := base.MaxHeapify(data, 1)
		t.Log(result)
	})
}
