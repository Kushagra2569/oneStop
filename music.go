package main

import (
	"bytes"
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
	MusicList     []Music `json:"musicList"`
	IdNum         int     `json:"idNum"`
	otoCtx        *oto.Context
	currentPlayer *oto.Player
	fileLoaded    bool
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
		m.otoCtx = nil
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
//TODO: add way to play next music
//TODO: add ability to change volume as well as return current volume and currentPlaying music in a json struct
//so create a struct that will hold the current playing music and the volume

func (m *MusicList) MusicController(id int, action int) string {
	musicPath := m.MusicList[id].Path

	if id >= len(m.MusicList) || id < 0 {
		return "Invalid id"
	}
	if m.otoCtx == nil {
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
		m.otoCtx = otoCtx

		// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
		<-readyChan
	}

	if action == 0 {
		if m.currentPlayer != nil {
			if m.currentPlayer.IsPlaying() {
				m.currentPlayer.Pause()
			}
		}
		return "Paused " + m.MusicList[id].Title

	} else if action == 1 {

		if m.currentPlayer != nil {
			m.currentPlayer.Close()
		}

		fileBytes, err := os.ReadFile(musicPath)
		if err != nil {
			return "Error opening file"
		}

		// Convert the pure bytes into a reader object that can be used with the mp3 decoder
		fileBytesReader := bytes.NewReader(fileBytes)

		//Decode mp3 file
		decodedMp3, err := mp3.NewDecoder(fileBytesReader)
		if err != nil {
			return "Error decoding mp3"
		}

		// Create a new 'player' that will handle our sound. Paused by default.
		playerVar := m.otoCtx.NewPlayer(decodedMp3)
		m.currentPlayer = playerVar

		go player(m.currentPlayer)

		return "Playing " + m.MusicList[id].Title

	} else if action == 2 {
		if m.currentPlayer != nil {
			if !m.currentPlayer.IsPlaying() {
				go player(m.currentPlayer)
			}
		}
		return "Resumed " + m.MusicList[id].Title

	} else {
		return "Invalid action"

	}
}

func player(player *oto.Player) {
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
	// Kush commented this part out because on pausing the player, it closed it which made it impossible to resume
	// err := player.Close()
	// if err != nil {
	// 	fmt.Println("player.Close failed: " + err.Error())
	// }

}
