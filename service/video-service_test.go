package service

import (
	"gintoki/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TITLE = "Sample Title"
var DESCRIPTION = "Sample description"
var URL = "sample_url.com"

func getVideo() entity.Video {
	return entity.Video{
		Title:       TITLE,
		Description: DESCRIPTION,
		URL:         URL,
	}
}

func TestFindAll(t *testing.T) {
	service := New()
	service.Save(getVideo())
	videos := service.FindAll()

	firstVideo := videos[0]
	assert.NotNil(t, videos)
	assert.Equal(t, TITLE, firstVideo.Title)
	assert.Equal(t, DESCRIPTION, firstVideo.Description)
	assert.Equal(t, URL, firstVideo.URL)
}
