package goverisure

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerisureOverview(t *testing.T) {
	v := New()
	assert.Nil(t, v.Login())

	installations, err := v.Installations()
	assert.Nil(t, err)

	overview, err := v.Overview(installations[0].Giid)
	assert.Nil(t, err)

	log.Printf("overview: %+v", overview)
}
