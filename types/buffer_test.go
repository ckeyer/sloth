package types

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"testing"
	"time"

	"gopkg.in/check.v1"
)

var (
	test_data = []byte("abcd")
)

func Test(t *testing.T) { check.TestingT(t) }

type BufferSuite struct {
	buf *Buffer
}

func init() {
	check.Suite(&BufferSuite{buf: NewBuf()})
}

func (b *BufferSuite) Test(c *check.C) {
	l, err := b.buf.Write(test_data)
	c.Check(err, check.IsNil)
	c.Check(l, check.Equals, 4)

	c.Assert(b.buf.Len(), check.Equals, 4)
	b.buf.Clear()
	c.Assert(b.buf.Len(), check.Equals, 0)

}

func (b *BufferSuite) TestIO(c *check.C) {
	b.buf.Clear()
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		io.Copy(b.buf, strings.NewReader(time.Now().String()+"\n"))
	}
	fmt.Println(b.buf.String())
}

func (b *BufferSuite) TestCMD(c *check.C) {
	cmd := exec.Command("sh", "test.sh")
	b.buf.Clear()
	cmd.Stderr = b.buf
	cmd.Stdout = b.buf
	cmd.Start()

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait... err, ", err)
	} else {
		fmt.Println("wait over...")
		b.buf.Over()
		fmt.Println("input times, ", b.buf.GetInputTimes())
	}
}
