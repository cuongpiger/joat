package attempt

import (
	"fmt"
	"testing"
	"time"
)

func TestAttemptWithDelay(t *testing.T) {
	AttemptWithDelay(3, 10*time.Second, func(inter int, _ time.Duration) error {
		fmt.Println("inter is: ", inter)
		if inter > 0 && inter%2 == 0 {
			fmt.Println("success")
			return nil
		}

		return fmt.Errorf("error with inter %d", inter)
	})
}
