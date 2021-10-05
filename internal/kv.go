package internal

import (
	"strings"
)

func Get(repo, token, scope, key string) ([]byte, error) {
	res, err := githubGetFile(repo, token, generateKey(scope, key))
	if err != nil {
		return nil, err
	}
	return res.RawContent, nil
}

func Set(repo, token, scope, key string, val []byte) error {
	key = generateKey(scope, key)
	return githubUploadFile(repo, token, key, val)
}

func generateKey(scope, key string) string {
	keys := splitBySizeChat(getMd5([]byte(key)), 4)
	keys = append([]string{"kv", scope}, keys...)
	keys = append(keys, key)
	return strings.Join(keys, "/")
}
