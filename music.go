package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dhowden/tag"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

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
		m.IdNum = m.IdNum + 1
	}
	m.SaveMusicListToFile()
	return string(musicListToJson(*m))
}

func (m *MusicList) deleteDuplicateMusic() {
	uniqueMusic := []Music{}
	uniqueMusicMap := make(map[string]string)
	for _, music := range m.MusicList {
		if music.Title == "" {
			if _, value := uniqueMusicMap[music.Path]; !value {
				uniqueMusicMap[music.FileName] = music.FileName
				uniqueMusic = append(uniqueMusic, music)
			}
		} else {
			if _, value := uniqueMusicMap[music.Title]; !value {
				uniqueMusicMap[music.Title] = music.Artist
				uniqueMusic = append(uniqueMusic, music)
			}
		}
	}
	m.MusicList = uniqueMusic
}

func (m *MusicList) SaveMusicListToFile() {
	m.deleteDuplicateMusic()
	json := musicListToJson(*m)
	err := WriteFile(musicFile, json)
	if err != nil {
		fmt.Println(err)
	}
}

//play music function
//TODO: add way to stop music
//TODO: add way to play next music
//TODO: handle changes for switching music in between

func (m *MusicList) Play(id int) string {
  if id >= len(m.MusicList) || id < 0 {
    return "Invalid id"
  }
  musicPath := m.MusicList[id].Path
  return player(musicPath)
}

func player(musicPath string) string {
  	file, err := os.Open(musicPath)
	if err != nil {
		return "Failed to play" + musicPath
	}

	//Decode mp3 file
	decodedMp3, err := mp3.NewDecoder(file)
	if err != nil {
		return "Failed to play" + musicPath
	}

	op := &oto.NewContextOptions{}

	op.SampleRate = 48000

	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
	// 1 is mono sound, and 2 is stereo (most speakers are stereo).
	op.ChannelCount = 2

	// Format of the source. go-mp3's format is signed 16bit integers.
	op.Format = oto.FormatSignedInt16LE

	// Remember that you should **not** create more than one context
	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	// Create a new 'player' that will handle our sound. Paused by default.
	player := otoCtx.NewPlayer(decodedMp3)

	// Play starts playing the sound and returns without waiting for it (Play() is async).
	player.Play()

  for player.IsPlaying() {
    time.Sleep(time.Millisecond)
  }

  // Now that the sound finished playing, we can restart from the beginning (or go to any location in the sound) using seek
    // newPos, err := player.(io.Seeker).Seek(0, io.SeekStart)
    // if err != nil{
    //     panic("player.Seek failed: " + err.Error())
    // }
    // println("Player is now at position:", newPos)
    // player.Play()

    // If you don't want the player/sound anymore simply close
    err = player.Close()
    if err != nil {
        panic("player.Close failed: " + err.Error())
    }

  return "Playing " + musicPath
}
