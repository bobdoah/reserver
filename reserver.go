package main

import "github.com/bndr/gojenkins"
import "fmt"
import "os"

func main() {

    jenkins := gojenkins.CreateJenkins(nil, "<url>", "<username>", "<password>")
    _, err := jenkins.Init()
    if err != nil {
        fmt.Printf("failed to initalize\n")
    }

    plugin, err := jenkins.HasPlugin("lockable-resources"); 
    if err != nil{
        fmt.Printf("failed to get plugins\n")
    }

    if plugin == nil {
        fmt.Printf("Lockable resources not installed\n")
        os.Exit(1) 
    }
    fmt.Printf("Lockable resources installed\n")
    os.Exit(0)
}
