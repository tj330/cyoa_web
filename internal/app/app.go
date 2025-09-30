package app

import (
	"html/template"
	"log"
	"os"

	"github.com/cyoa_web/internal/api"
	"github.com/cyoa_web/internal/models"
)

type Application struct {
	Story        *models.Story
	StoryHandler *api.StoryHandler
	Template     *template.Template
	Logger       *log.Logger
}

func NewApplication() (*Application, error) {

	template := template.Must(template.ParseFiles("internal/templates/index.html"))
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	story := models.JsonStory(logger)
	storyHandler := api.NewStoryHandler(template, logger, story)
	app := &Application{
		Logger:       logger,
		Story:        story,
		StoryHandler: storyHandler,
		Template:     template,
	}

	return app, nil
}
