package find

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewerCreate(t *testing.T) {
	finder := NewFinder()
	finder.Newer(Created, "test/data.csv")
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test/other/DATA.csv"}, matches)
}
