package internal

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var ErrNotFound = fmt.Errorf("not found")

func githubUploadFile(repo, token, path string, content []byte) error {
	sha := getSHA(repo, token, path)
	var body io.Reader
	if sha == "" {
		body = bytes.NewReader([]byte(fmt.Sprintf(`{
  "message": "upload %s",
  "content": "%s"
}`, path, base64.StdEncoding.EncodeToString(content))))
	} else {
		body = bytes.NewReader([]byte(fmt.Sprintf(`{
  "message": "update %s",
  "content": "%s",
  "sha": "%s"
}`, path, base64.StdEncoding.EncodeToString(content), sha)))
	}

	uri := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repo, path)
	_, err := githubDoRequest(http.MethodPut, uri, body, token)
	return err
}

func githubGetFile(repo, token, path string) (*githubGetFileResp, error) {
	uri := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repo, path)
	res, err := githubDoRequest(http.MethodGet, uri, nil, token)
	if err != nil {
		return nil, err
	}
	resp := new(githubGetFileResp)
	if err = json.Unmarshal(res, resp); err != nil {
		return nil, err
	}
	res, err = base64.StdEncoding.DecodeString(resp.Content)
	if err != nil {
		return nil, err
	}
	resp.RawContent = res
	return resp, nil
}

func githubDoRequest(method, url string, body io.Reader, token string) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{Timeout: time.Second * 3}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		var res struct {
			Message string `json:"message"`
		}
		if _ = json.Unmarshal(bs, &res); res.Message != "" {
			if resp.StatusCode == 404 && res.Message == "Not Found" {
				return nil, ErrNotFound
			}
			return nil, fmt.Errorf("[github] %s %s fail, code=%d, res=%s", method, url, resp.StatusCode, res.Message)
		}
		return nil, fmt.Errorf("[github] %s %s fail, code=%d, res=%s", method, url, resp.StatusCode, bs)
	}

	return bs, nil
}

func getSHA(repo, token, path string) string {
	uri := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repo, path)
	res, err := githubDoRequest(http.MethodGet, uri, nil, token)
	if err != nil {
		return ""
	}

	var sha struct {
		SHA string `json:"sha"`
	}
	if err = json.Unmarshal(res, &sha); err != nil {
		return ""
	}

	return sha.SHA
}

type githubGetFileResp struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HtmlURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
	RawContent  []byte `json:"-"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}
