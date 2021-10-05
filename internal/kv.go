package internal

import (
	"strings"
)

func Get(repo, token, key string) ([]byte, error) {
	res, err := githubGetFile(repo, token, generateKey(key))
	if err != nil {
		return nil, err
	}
	return res.RawContent, nil
}

func Set(repo, token, key string, val []byte) error {
	key = generateKey(key)
	return githubUploadFile(repo, token, key, val)
}

func generateKey(key string) string {
	keys := splitBySizeChat(getMd5([]byte(key)), 4)
	keys = append([]string{"kv"}, keys...)
	keys = append(keys, key)
	return strings.Join(keys, "/")
}
