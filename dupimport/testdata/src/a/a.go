package a

import fmt1 "fmt"
import fmt2 "fmt" // want "duplicated import: fmt"

func f() {
	fmt1.Println("test")
	fmt2.Println("test")
}
