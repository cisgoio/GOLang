//ref: https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)
    fmt.Println(exPath)
}


//===Older way:
/*
package main

import (
    "fmt"
    "os"
)

func main() {
    pwd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(pwd)
}
*/
