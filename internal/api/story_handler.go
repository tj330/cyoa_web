package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/cyoa_web/internal/models"
	"github.com/cyoa_web/internal/utils"
)

type StoryHandler struct {
	template *template.Template
	logger   *log.Logger
	story    *models.Story
}

func NewStoryHandler(template *template.Template, logger *log.Logger, story *models.Story) *StoryHandler {
	return &StoryHandler{
		template: template,
		logger:   logger,
		story:    story,
	}
}

func (sh *StoryHandler) HanldeGetRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "cyoa/story/intro", http.StatusSeeOther)
}

func (sh *StoryHandler) HandleGetStoryByArc(w http.ResponseWriter, r *http.Request) {
	arc, err := utils.ReadArcParam(r)
	if err != nil {
		sh.logger.Printf("ERROR: readArcParam: %v", err)
	}

	ch, ok := sh.story.GetChapterByArc(*arc)
	if ok {
		sh.template.Execute(w, ch)
	}
}
