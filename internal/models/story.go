package models

import (
	"encoding/json"
	"log"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JsonStory(logger *log.Logger) *Story {
	f, err := os.Open("internal/data/story.json")
	if err != nil {
		logger.Fatalf("failed to open story.json: %v", err)
	}
	defer f.Close()

	var story Story
	if err := json.NewDecoder(f).Decode(&story); err != nil {
		logger.Fatalf("failed to decode story.json: %v", err)
	}

	return &story
}

func (s Story) GetChapterByArc(arc string) (Chapter, bool) {
	ch, ok := s[arc]

	return ch, ok
}
