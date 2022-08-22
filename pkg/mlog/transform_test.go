package mlog_test

import (
	"testing"

	"github.com/solohin/mindustry-compiler/pkg/mlog"
	"github.com/stretchr/testify/require"
)

func TestCompile(t *testing.T) {
	original := `set Hello "1"
jump SIX equal found "six"
set Hello "2"
# comment 1
THREE: set Hello "3"
set Hello "4"
FIVE: set Hello "5"
# comment 3


jump THREE equal found "three"
SIX: set Hello "6"
# more comments
STARTLOOP gNum 3
	set Hello "My name is George number $gNum"
	jump LABEL_$gNum equal 1 1 
	set Hello "Never executed"
	LABEL_$gNum: set Hello "3"
	# comment
ENDLOOP gNum
end
`

	expected := `set Hello "1"
jump 7 equal found "six"
set Hello "2"
set Hello "3"
set Hello "4"
set Hello "5"
jump 3 equal found "three"
set Hello "6"
set Hello "My name is George number 0"
jump 11 equal 1 1
set Hello "Never executed"
set Hello "3"
set Hello "My name is George number 1"
jump 15 equal 1 1
set Hello "Never executed"
set Hello "3"
set Hello "My name is George number 2"
jump 19 equal 1 1
set Hello "Never executed"
set Hello "3"
end`

	compiled, err := mlog.Compile(original)
	require.NoError(t, err)
	require.Equal(t, expected, compiled)
}
