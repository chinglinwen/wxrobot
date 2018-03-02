package all

import (
	"fmt"
	"testing"
)

func TestRequest(t *testing.T) {
	reply, err := request(backendurl, bodytype, []byte("hello"))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(reply)
}
