package http

import (
	"errors"
	"io"

	"github.com/author_name/project_name/configs"
)

func HttpGet(api string) (string, error) {
	if len(api) == 0 {
		return "", errors.New("api is empty")
	}

	respRaw, err := configs.HttpClient.Get(api)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(respRaw.Body)
	if err != nil {
		return "", err
	}

	if respRaw.StatusCode != 200 {
		return "", errors.New("server error")
	}

	bodyBytes, err := io.ReadAll(respRaw.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
