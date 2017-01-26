package main

import (
    "os"
    "fmt"
    "flag"
    "path/filepath"
    "encoding/json"
    "io"
)

var (
    config  Config
    backup  bool
    pips    int
)

// Config struct
// Holds all of our settings and the games
type Config struct {
    SourcePath  string
    TargetPath  string
    BackupPath  string
    Week        int
    Games       []Game
}

// Game struct
// Holds game info and the plays
type Game struct {
    Grade string
    Title string
    Plays []Play
}

// Play struct
// Holds info about each play
type Play struct {
    Name string
    Path string
    Size int
}

// gets config from json file
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

// makes directories for games
func makeTargetDirs() {
    if len(config.Games) == 0 {
        fmt.Println("No games found in config")
    }

    for i := range config.Games {
        game := config.Games[i]
        name := game.Grade + " - " + game.Title

        e := os.Mkdir(config.TargetPath + string(filepath.Separator) + name, 0777)

        if e == nil {
            fmt.Println(name)
        } else {
            fmt.Println(e)
        }

    }
}

// does something maybe
func copyPlays() {

}

// copies file from source to destination
func copyPlay(src, dst string) error {
    srcFile, err := os.Open(src)
    if err != nil { return err }

    defer srcFile.Close()

    desFile, err := os.Create(dst)
    if err != nil { return err }

    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    copyErr := dstFile.Close()

    if err != nil { return err }

    return copyErr
}

func main() {
    flag.BoolVar(&backup, "backup", true, "Choose whether to copy files to backup path")
    flag.Parse()

    config = getConfig()

    makeTargetDirs()
    copyPlays()
}
