package main

import (
    "os"
    "fmt"
    "flag"
    "encoding/json"
    "io/ioutil"

    "github.com/plloyd241/sideliner/game"
)

type Config struct {
    week        string
    sourcePath  string
    targetPath  string
    backupPath  string
    games       []Game
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

    doBackup := flag.Bool("backup", true, "Choose whether to copy files to backup path")
    name := flag.String("name", "test", "Enter your name")
    flag.Parse()

    fmt.Println("Hello,", flagVal);
}
