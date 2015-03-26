package shell

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const testScript = `#!/usr/bin/env sh
test $FRESH_PRINCE = "humperdink"
`

func TestExec(t *testing.T) {
	err := os.Setenv("FRESH_PRINCE", "humperdink")
	assert.Nil(t, err)
	err = Exec(testScript)
	assert.Nil(t, err)
}
