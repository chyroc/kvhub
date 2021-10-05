package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_github(t *testing.T) {
	as := assert.New(t)
	token := os.Getenv("GITHUB_TOKEN")
	repo := "chyroc/data-store" // os.Getenv("REPO")

	t.Run("", func(t *testing.T) {
		res, err := githubGetFile(repo, token, "fake/1.txt")
		as.Nil(err)
		as.Equal("1\n\n2\n\n3", string(res.RawContent))
	})

	t.Run("", func(t *testing.T) {
		res, err := githubGetFile(repo, token, "fake/a/2.json")
		as.Nil(err)
		as.Equal("{\n\t\"a\": \"b\"\n}", string(res.RawContent))
	})

	t.Run("", func(t *testing.T) {
		err := githubUploadFile(repo, token, "fake/3.txt", []byte("# hi"))
		as.Nil(err)
		res, err := githubGetFile(repo, token, "fake/3.txt")
		as.Nil(err)
		as.Equal("# hi", string(res.RawContent))
	})
}
