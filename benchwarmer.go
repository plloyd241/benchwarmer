package main

import (
    "os"
    "fmt"
    "flag"
    "encoding/json"
    "io/ioutil"
)

var (
    backup bool
)

type Config struct {
    week        string
    sourcePath  string
    targetPath  string
    backupPath  string
    games       []Game
}

type Game struct {
    grade string
    teams string
}

func getConfig() jsonobject {
    file, e := ioutil.ReadFile("./config.json")
    d := json.NewDecoder()

    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
}

func main() {
    config := getConfig()

    fmt.Printf("Results: %v\n", config)

    backup := flag.Bool("backup", true, "Choose whether to copy files to backup path")
    flag.Parse()

    fmt.Println("Hello,", flagVal);
}
