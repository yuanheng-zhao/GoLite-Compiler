package scanner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
	// "proj/golite/token"
)

func Test1(t *testing.T) {
	// s1 := "This is a test string"

	f_obj, _ := os.Open("test.golite")
	var r *bufio.Reader
	r = bufio.NewReader(f_obj)
	c, _, _ := r.ReadRune()
	c, _, _ = r.ReadRune()
	c, _, _ = r.ReadRune()
	r.UnreadRune()

	c, _, _ = r.ReadRune()
	d := string(c)
	fmt.Println(d)
}

func Test2(t *testing.T) {
	f_obj, _ := os.Open("test.golite")
	var r *bufio.Reader
	r = bufio.NewReader(f_obj)
	s := New(r)
	x, _, err := s.reader.ReadRune()
	if err != nil && err == io.EOF {
		fmt.Println("EOF FOUND!")
	} else {
		fmt.Println(x)
	}
}
