package padmin

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func StatusNok(code int) bool {
	return code < 200 || code >= 300
}

func StatusOk(code int) bool {
	return code >= 200 && code < 300
}

func HttpCheck(response *http.Response) error {
	defer response.Body.Close()
	if StatusNok(response.StatusCode) {
		str, err := ReadAll(response.Body)
		if err != nil {
			return err
		}
		return errors.New(str)
	}
	return nil
}

func HttpCheckReadBytes(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	if StatusNok(response.StatusCode) {
		str, err := ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(str)
	}
	return io.ReadAll(response.Body)
}

func ReadAll(r io.Reader) (string, error) {
	bytes, err := io.ReadAll(r)
	return string(bytes), err
}

func EasyReader(resp *http.Response, ptr interface{}) error {
	body, err := HttpCheckReadBytes(resp)
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return nil
	}
	return json.Unmarshal(body, ptr)
}
