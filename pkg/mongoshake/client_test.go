package mongoshake

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestClient_GetRepl(t *testing.T) {
	c := NewClient("http://localhost:54116")
	repl, err := c.GetRepl()
	if err != nil {
		t.Errorf("Client.GetRepl() error = %v", err)
	}
	spew.Dump(repl)
}
