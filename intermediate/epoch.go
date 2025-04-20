package intermediate

import (
	"fmt"
	"time"
)

func main() {

	// 00:00:00 UTC on Jan 1, 1970

	now := time.Now()
	unixTime := now.Unix()

	fmt.Println("Now:", now)
	fmt.Println("Now Unix:", unixTime)

	t := time.Unix(unixTime, 0) // 0 nanoseconds
	fmt.Println(t)
	fmt.Println("Formated Time:", t.Format("02/01/06 15:04"))
}
