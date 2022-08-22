package mlog_test

import (
	"testing"

	"github.com/solohin/mindustry-compiler/pkg/mlog"
	"github.com/stretchr/testify/require"
)

func TestSimpleLoop(t *testing.T) {
	original := `
set dbg "1"
STARTLOOP gNum 3
	# comment
	set dbg "Loop iteration $gNum starts"
	set dbg "This is an iteration #$gNum"
	set dbg "Loop iteration $gNum ends"
ENDLOOP gNum
end
`

	result := `set dbg "1"
set dbg "Loop iteration 0 starts"
set dbg "This is an iteration #0"
set dbg "Loop iteration 0 ends"
set dbg "Loop iteration 1 starts"
set dbg "This is an iteration #1"
set dbg "Loop iteration 1 ends"
set dbg "Loop iteration 2 starts"
set dbg "This is an iteration #2"
set dbg "Loop iteration 2 ends"
end`

	compiled, err := mlog.Compile(original)
	require.NoError(t, err)
	require.Equal(t, result, compiled)
}

func TestNestedLoop(t *testing.T) {
	original := `
set dbg "1"
STARTLOOP i 4
	STARTLOOP k 3
	set dbg "Loop iteration $i.$k starts"
	ENDLOOP k
ENDLOOP i
end
`

	result := `set dbg "1"
set dbg "Loop iteration 0.0 starts"
set dbg "Loop iteration 0.1 starts"
set dbg "Loop iteration 0.2 starts"
set dbg "Loop iteration 1.0 starts"
set dbg "Loop iteration 1.1 starts"
set dbg "Loop iteration 1.2 starts"
set dbg "Loop iteration 2.0 starts"
set dbg "Loop iteration 2.1 starts"
set dbg "Loop iteration 2.2 starts"
set dbg "Loop iteration 3.0 starts"
set dbg "Loop iteration 3.1 starts"
set dbg "Loop iteration 3.2 starts"
end`

	compiled, err := mlog.Compile(original)
	require.NoError(t, err)
	require.Equal(t, result, compiled)
}
