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
	MusicList  []Music `json:"musicList"`
	IdNum      int     `json:"idNum"`
	fileLoaded bool
}

func musicMetadata(musicPath string) Music {
	musicfile, err := os.Open(musicPath)
	defer musicfile.Close()
	if err != nil {
		fmt.Println(err)
	}
	meta, err := tag.ReadFrom(musicfile)
	filename := ""
	if err != nil {
		fmt.Println(err)
	}
	filenameslice := strings.Split(musicPath, "\\")
	filename = filenameslice[len(filenameslice)-1]

	title := ""
	if meta.Title() != "" {
		title = meta.Title()
	} else {
		title = filename
	}

	artist := ""
	if meta.Artist() != "" {
		artist = meta.Artist()
	} else {
		artist = "Unknown"
	}

	year := ""
	if fmt.Sprint(meta.Year()) != "" {
		year = fmt.Sprint(meta.Year())
	} else {
		year = "Unknown"
	}

	music := Music{
		Path:     musicPath,
		FileName: filename,
		Title:    title,
		Artist:   artist,
		Year:     year,
	}
	return music
}

func musicListToJson(music MusicList) []byte {
	json, err := json.Marshal(music)
	if err != nil {
		fmt.Println(err)
	}
	return json
}

func jsonToMusicList(musicListStr []byte) MusicList {
	var musicList MusicList
	err := json.Unmarshal(musicListStr, &musicList)
	if err != nil {
		fmt.Println(err)
	}
	return musicList
}

func (m *MusicList) GetMusicList() string {
	if !m.fileLoaded {
		*m = LoadMusicListFromFile()
		m.fileLoaded = true
	}
	musicListStr := musicListToJson(*m)
	return string(musicListStr)
}

func LoadMusicListFromFile() MusicList {
	musicListStr, err := LoadFile(musicFile)
	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		}
	}
	musicList := jsonToMusicList(musicListStr) //Kush: fix unnecessary conversion from json to struct and back to json
	return musicList
}

func (m *MusicList) GetMusicListFromLocalFiles(musicListPaths []string) string {
	for _, musicPath := range musicListPaths {
		music := musicMetadata(musicPath)
		m.MusicList = append(m.MusicList, music)
	}
	return string(musicListToJson(*m))
}

func (m *MusicList) SaveMusicListToFile() {
	json := musicListToJson(*m)
	err := WriteFile(musicFile, json)
	if err != nil {
		fmt.Println(err)
	}
}
