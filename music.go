package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dhowden/tag"
)

//TODO: convert from musicList struct to json func
//read file function and return musicList struct
//save musicList struct to file

type Music struct {
	Path     string `json:"path"`
	FileName string `json:"fileName"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Year     string `json:"year"`
}

type MusicList struct {
	MusicList []Music `json:"musicList"`
	IdNum     int     `json:"idNum"`
}

func musicMetadata(musicPath string) Music {
	musicfile, err := os.Open(musicPath)
	if err != nil {
		fmt.Println(err)
	}
	meta, err := tag.ReadFrom(musicfile)
	filename := ""
	if err != nil {
		fmt.Println(err)
	}
	filenameslice := strings.Split(musicPath, "/")
	filename = filenameslice[len(filename)-1]
	music := Music{
		Path:     musicPath,
		FileName: filename,
		Title:    meta.Title(),
		Artist:   meta.Artist(),
		Year:     fmt.Sprint(meta.Year()),
	}
	return music
}

func musicToJson(music MusicList) []byte {
	json, err := json.Marshal(music)
	if err != nil {
		fmt.Println(err)
	}
	return json
}
