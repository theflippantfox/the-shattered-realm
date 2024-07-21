package story

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Option struct {
	Text        string `yaml:"text"`
	NextScene   int    `yaml:"nextScene,omitempty"`
	NextChapter int    `yaml:"nextChapter,omitempty"`
}

type Scene struct {
	Text    string   `yaml:"text"`
	Options []Option `yaml:"options"`
}

type Chapter struct {
	Title  string  `yaml:"title"`
	Scenes []Scene `yaml:"scenes"`
}

type Story struct {
	Chapters []Chapter `yaml:"chapters"`
}

type StoryManager struct {
	Story          Story
	CurrentChapter int
	CurrentScene   int
}

func NewStoryManager(filename string) (*StoryManager, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var story Story
	err = yaml.Unmarshal(file, &story)
	if err != nil {
		return nil, err
	}

	return &StoryManager{
		Story:          story,
		CurrentChapter: 0,
		CurrentScene:   0,
	}, nil
}
func (sm *StoryManager) CurrentChapterTitle() string {
	return sm.Story.Chapters[sm.CurrentChapter].Title
}

func (sm *StoryManager) CurrentText() string {
	return sm.Story.Chapters[sm.CurrentChapter].Scenes[sm.CurrentScene].Text
}

func (sm *StoryManager) CurrentOptions() []Option {
	return sm.Story.Chapters[sm.CurrentChapter].Scenes[sm.CurrentScene].Options
}

func (sm *StoryManager) ChooseOption(optionIndex int) {
	option := sm.CurrentOptions()[optionIndex]
	if option.NextScene != 0 {
		sm.CurrentScene = option.NextScene - 1 // Adjust for 0-based index
	} else if option.NextChapter != 0 {
		sm.CurrentChapter = option.NextChapter - 1 // Adjust for 0-based index
		sm.CurrentScene = 0
	} else {
		fmt.Println("No valid next scene or chapter specified")
	}
}

func (sm *StoryManager) IsEnd() bool {
	return sm.CurrentChapter >= len(sm.Story.Chapters) ||
		sm.CurrentScene >= len(sm.Story.Chapters[sm.CurrentChapter].Scenes)
}
