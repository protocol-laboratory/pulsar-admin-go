package padmin

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func dummyHttpServer() {
	http.HandleFunc("/timeout", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Millisecond * 500)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func Test_newHttpClient(t *testing.T) {
	// run a dummy server
	go dummyHttpServer()

	cli := &http.Client{Timeout: time.Second}
	req, err := http.NewRequest("HEAD", "http://localhost:8080/timeout", nil)
	if !assert.NoError(t, err) {
		return
	}
	resp, err := cli.Do(req)
	if !assert.NoError(t, err) {
		return
	}
	all, err := ReadAll(resp.Body)
	if !assert.NoError(t, err) {
		return
	}
	fmt.Println(all)
}

func TestMarshal(t *testing.T) {
	var a interface{} = 1
	bb, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte{49}, bb)
	var b interface{} = "1"
	bb, err = json.Marshal(b)
	assert.NoError(t, err)
	t.Log(bb)
}
