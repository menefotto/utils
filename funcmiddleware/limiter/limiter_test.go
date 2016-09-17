package limiter

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLimiter(t *testing.T) {
	var f LimiterFunc = func() (interface{}, error) {
		fmt.Println("Random num: ", rand.Intn(99))
		return nil, nil
	}

	for i := 0; i < 40; i++ {
		Limiter(f, 1000)
	}
}
