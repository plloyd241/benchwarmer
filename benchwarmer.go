package main

import (
    "os"
    "fmt"
    "errors"
    "flag"
    "encoding/json"
    "io/ioutil"
)

var (
    config  Config
    backup  bool
    pips    int
)

type Config struct {
    SourcePath  string
    TargetPath  string
    BackupPath  string
    Week        int
    Games       []Game
}

type Game struct {
    Grade string
    Title string
}

func getConfig() Config {
    file, e := ioutil.ReadFile("./config.json")

    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    var c Config
    json.Unmarshal(file, &c)

    return c
}

func makeTargetDirs() error {
    if len(config.Games) == 0 {
        return errors.New("No games found in config")
    }

    for i := range config.Games {
        game := config.Games[i]

        fmt.Printf("%s - %s\n", game.Grade, game.Title)
    }

    return nil
}

func copyPlays() error {
    return nil
}

func main() {
    flag.BoolVar(&backup, "backup", true, "Choose whether to copy files to backup path")
    flag.Parse()

    config = getConfig()

    dirErr := makeTargetDirs();
    if dirErr != nil {
        fmt.Printf("Error creating directories: %v\n", dirErr)
    }

    copyErr := copyPlays();
    if copyErr != nil {
        fmt.Printf("Error copying files: %v\n", copyErr)
    }
}
