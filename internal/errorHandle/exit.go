package errorHandle

import (
	"fmt"
	"os"
)

func ExitWithError(errorMessage string) {
	fmt.Println("Error:", errorMessage)
	os.Exit(1)
}
