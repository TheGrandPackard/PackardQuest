package musicplayer

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"
)

type MusicPlayer interface {
	GetCurrentSongName() string
	GetSongs() []string

	Play() error
	Stop() error
	Interrupt(string) error
	Next() error
}

type musicPlayer struct {
	currentSong   *int
	songs         []string
	interruptSong *string

	streamer beep.StreamSeekCloser
	ctrl     *beep.Ctrl
	doneChan chan bool
}

func NewMusicPlayer(songs []string) MusicPlayer {
	p := &musicPlayer{
		songs:    songs,
		doneChan: make(chan bool),
	}

	// handle song done to play next
	go p.doneHandler()

	return p
}

func (p *musicPlayer) doneHandler() {
	for {
		<-p.doneChan
		// Reset interrupt song
		if p.interruptSong != nil {
			p.interruptSong = nil
		}

		if err := p.Next(); err != nil {
			log.Println("Error playing next:", err.Error())
			continue
		}
		log.Println("Playing next song: ", p.GetCurrentSongName())
	}
}

func (p *musicPlayer) GetCurrentSongName() string {
	if p.interruptSong != nil {
		return *p.interruptSong
	}
	if p.currentSong == nil {
		return "No current song selected"
	}
	return p.songs[*p.currentSong]
}

func (p *musicPlayer) GetSongs() []string {
	return p.songs
}

func (p *musicPlayer) Play() error {
	if p.currentSong == nil {
		if err := p.getNextSong(); err != nil {
			return err
		}
	}

	return p.play()
}

func (p *musicPlayer) Stop() error {
	return p.stop()
}

func (p *musicPlayer) Interrupt(fileName string) error {
	p.interruptSong = &fileName

	return p.play()
}

func (p *musicPlayer) Next() error {
	// Reset interrupt song
	if p.interruptSong != nil {
		p.interruptSong = nil
	}
	if err := p.getNextSong(); err != nil {
		return err
	}

	return p.play()
}

func (p *musicPlayer) getNextSong() error {
	if p.currentSong == nil {
		if len(p.songs) == 0 {
			return errors.New("no songs in player")
		}

		idx := 0
		p.currentSong = &idx
	} else {
		if *p.currentSong+1 >= len(p.songs) {
			idx := 0
			p.currentSong = &idx
		} else {
			idx := *p.currentSong + 1
			p.currentSong = &idx
		}
	}

	return nil
}

func (p *musicPlayer) play() error {
	fileName := p.songs[*p.currentSong]

	// Override with interrupt song
	if p.interruptSong != nil {
		fileName = *p.interruptSong
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var format beep.Format

	if strings.HasSuffix(fileName, ".mp3") {
		p.streamer, format, err = mp3.Decode(f)
	} else if strings.HasSuffix(fileName, ".ogg") {
		p.streamer, format, err = vorbis.Decode(f)
	} else {
		return errors.New("unsupported format for file:" + fileName)
	}
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	p.ctrl = &beep.Ctrl{Streamer: beep.Loop(1, p.streamer), Paused: false}

	speaker.Play(beep.Seq(p.ctrl, beep.Callback(func() {
		p.doneChan <- true
	})))

	return nil
}

func (p *musicPlayer) stop() error {
	if p.currentSong != nil {
		speaker.Lock()
		p.ctrl.Paused = true
		speaker.Unlock()
	}

	return nil
}
