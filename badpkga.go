package badpkga

import (
    "fmt"

    "github.com/snowroll/badpkg"
)

func Outa() {
    message := badpkg.Hello("badpkga v1.0.0")
    fmt.Println(message)
}


