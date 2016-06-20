package amqp

import (
	"testing"
//  "fmt"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct{}

var _ = check.Suite(&S{})

func (s *S) Testdial(c *check.C) {
	f, err := Init()
	c.Assert(err, check.IsNil)
  qu, ok := f.New("test-queue")
  body := []byte("test-data")
  qu.Publish(body)
	c.Assert(ok, check.Equals, true)
}
