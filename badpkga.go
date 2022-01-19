package badpkga

import (
    "fmt"

    "github.com/snowroll/badpkg"
)

func Outa() {
    message := badpkg.Hello("badpkga")
    fmt.Println(message)
}


