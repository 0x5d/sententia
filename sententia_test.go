package sententia_test

import (
	"strings"
	"testing"

	"github.com/castillobgr/sententia"
	"github.com/stretchr/testify/assert"
)

func TestNouns(t *testing.T) {
	nouns := []string{"potato"}
	sententia.SetNouns(nouns)
	sentence, err := sententia.Make("{{ noun }}")
	assert.NoError(t, err)
	assert.Equal(t, sentence, nouns[0])
}

func TestAdjectives(t *testing.T) {
	adjectives := []string{"silly"}
	sententia.SetAdjectives(adjectives)
	sentence, err := sententia.Make("{{ adjective }}")
	assert.NoError(t, err)
	assert.Equal(t, sentence, adjectives[0])
}

func TestAddNouns(t *testing.T) {
	sententia.SetNouns([]string{})
	customNouns := []string{"rocket"}
	sententia.AddNouns(customNouns)
	sentence, err := sententia.Make("{{ noun }}")
	assert.NoError(t, err)
	assert.Equal(t, sentence, customNouns[0])
}

func TestAddAdjectives(t *testing.T) {
	sententia.SetAdjectives([]string{})
	customAdjectives := []string{"fast"}
	sententia.AddAdjectives(customAdjectives)
	sentence, err := sententia.Make("{{ adjective }}")
	assert.NoError(t, err)
	assert.Equal(t, sentence, customAdjectives[0])
}

func TestAddActions(t *testing.T) {
	customActions := map[string]interface{}{
		"capitalize": strings.Title,
	}
	sententia.AddActions(customActions)
	sententia.SetNouns([]string{})
	customNouns := []string{"rocket"}
	sententia.AddNouns(customNouns)
	sentence, err := sententia.Make("{{ capitalize noun }}")
	assert.NoError(t, err)
	assert.Equal(t, sentence, "Rocket")
}
