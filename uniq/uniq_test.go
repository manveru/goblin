package main

import (
	"bytes"
	. "github.com/manveru/gobdd"
	"testing"
)

func TestEverything(t *testing.T) {}

func init() {
	defer PrintSpecReport()

	out := bytes.NewBufferString("")
	reset := func() { out.Reset() }

	Describe("uniq -u", func() {
		BeforeEach(reset)
		flagMode = 'u'

		It("prints a single line", func() {
			in := bytes.NewBufferString("a\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\n")
		})

		It("prints two distinct lines", func() {
			in := bytes.NewBufferString("a\nb\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\nb\n")
		})

		It("prints three distinct lines", func() {
			in := bytes.NewBufferString("a\nb\nc\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\nb\nc\n")
		})

		It("prints four distinct lines", func() {
			in := bytes.NewBufferString("a\nb\nc\nd\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\nb\nc\nd\n")
		})

		It("prints none of two identical lines", func() {
			in := bytes.NewBufferString("a\na\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "")
		})

		It("prints none of three identical lines", func() {
			in := bytes.NewBufferString("a\na\na\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "")
		})

		It("prints none of four identical lines", func() {
			in := bytes.NewBufferString("a\na\na\na\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "")
		})

		It("handles more complex examples", func() {
			in := bytes.NewBufferString("a\na\nb\nc\nc\nd\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "b\nd\n")
		})
	})

	Describe("uniq -d prints one copy of duplicate lines", func() {
		BeforeEach(reset)
		flagMode = 'd'

		It("doesn't print a single line", func() {
			in := bytes.NewBufferString("a\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "")
		})

		It("prints one of two duplicate lines", func() {
			in := bytes.NewBufferString("a\na\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\n")
		})

		It("prints one of three duplicate lines", func() {
			in := bytes.NewBufferString("a\na\na\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\n")
		})

		It("prints one of four duplicate lines", func() {
			in := bytes.NewBufferString("a\na\na\na\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\n")
		})

		It("handles more complex cases", func() {
			in := bytes.NewBufferString("a\na\na\nb\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\nb\n")
		})

		It("handles more complex cases", func() {
			in := bytes.NewBufferString("a\na\na\nb\nb\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\nb\n")
		})
	})

	Describe("uniq -c prints number of duplicates with each line", func() {
		BeforeEach(reset)
		flagMode = 'c'

		It("handles a single line", func() {
			in := bytes.NewBufferString("a\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "   1 a\n")
		})

		It("handles two duplicate lines", func() {
			in := bytes.NewBufferString("a\na\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "   2 a\n")
		})

		It("handles three duplicate lines", func() {
			in := bytes.NewBufferString("a\na\na\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "   3 a\n")
		})

		It("various lines", func() {
			in := bytes.NewBufferString("a\nb\nb\nc\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "   1 a\n   2 b\n   1 c\n")
		})
	})

	Describe("uniq prints lines without duplicates", func() {
		BeforeEach(reset)
		flagMode = ' '

		It("handles a single line", func() {
			in := bytes.NewBufferString("a\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\n")
		})

		It("handles three lines", func() {
			in := bytes.NewBufferString("a\na\nb\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\nb\n")
		})

		It("various lines", func() {
			in := bytes.NewBufferString("a\nb\nc\nc\nc\nd\nd\ne\ne\n")
			uniq(in, out)
			Expect(out.String(), ToEqual, "a\nb\nc\nd\ne\n")
		})
	})

	Describe("skip", func() {
		It("skips 1 field", func() {
			Expect(skip("", 1, 0), ToEqual, "")
			Expect(skip("foo bar", 1, 0), ToEqual, "bar")
			Expect(skip("foo bar  baz", 1, 0), ToEqual, "bar  baz")
		})

		It("skips chars", func() {
			Expect(skip("", 0, 1), ToEqual, "")
			Expect(skip("foo", 0, 1), ToEqual, "oo")
			Expect(skip("foobar", 0, 3), ToEqual, "bar")
		})

		It("skips fields then chars", func() {
			Expect(skip("", 1, 1), ToEqual, "")
			Expect(skip("foo bar", 1, 1), ToEqual, "ar")
			Expect(skip("foobar baz hoge", 2, 1), ToEqual, "oge")
		})
	})
}
