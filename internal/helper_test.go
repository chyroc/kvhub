package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trimPrefix(t *testing.T) {
	as := assert.New(t)

	as.Equal("a", trimPrefixPath("./a"))
	as.Equal("a", trimPrefixPath(".//a"))
	as.Equal("a", trimPrefixPath("//a"))
	as.Equal("a", trimPrefixPath(".a"))
	as.Equal("a", trimPrefixPath("a"))
}

func Test_md5(t *testing.T) {
	assert.Equal(t, "18158894e2973a002b17a52abd05736c", getMd5([]byte("235689876543")))
}

func Test_split(t *testing.T) {
	assert.Equal(t, []string{"18", "58", "94", "29", "3a", "02", "17", "52", "bd", "57", "6c"}, splitBySizeChat("18158894e2973a002b17a52abd05736c", 2))
}
