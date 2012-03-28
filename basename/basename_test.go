package main

import (
	. "github.com/manveru/gobdd"
	"testing"
)

func TestEverthing(*testing.T) {}

func init() {
	defer PrintSpecReport()

	Describe("basename", func() {
		It("does `basneame foobar`", func() {
			Expect(basename("foobar", ""), ToEqual, "foobar")
		})

		It("does `basename foobar.txt`", func() {
			Expect(basename("foobar.txt", ""), ToEqual, "foobar.txt")
		})

		It("does `basename foobar.txt .txt`", func() {
			Expect(basename("foobar.txt", ".txt"), ToEqual, "foobar")
		})

		It("does `basename foobar.txt.txt .txt`", func() {
			Expect(basename("foobar.txt.txt", ".txt"), ToEqual, "foobar.txt")
		})

		It("does `basename foobar.txt.txt .txt`", func() {
			Expect(basename("foobar.txt.txt", ".txt"), ToEqual, "foobar.txt")
		})

		It("does `basename /some/thing/in/foo`", func() {
			Expect(basename("/some/thing/in/foo", ""), ToEqual, "foo")
		})

		It("does `basename /some/thing/in/foo.bar .bar`", func() {
			Expect(basename("/some/thing/in/foo.bar", ".bar"), ToEqual, "foo")
		})

		It("does `basename x x`", func() {
			Expect(basename("x", "x"), ToEqual, "")
		})

		It("does `basename ax x`", func() {
			Expect(basename("ax", "x"), ToEqual, "a")
		})

		It("does `basename x xxx`", func() {
			Expect(basename("x", "xxx"), ToEqual, "x")
		})
	})

	Describe("dirname", func() {
		It("does `basename -d foobar`", func() {
			Expect(dirname("foobar"), ToEqual, ".")
		})

		It("does `basename -d /`", func() {
			Expect(dirname("/"), ToEqual, "")
		})

		It("does `basename -d //`", func() {
			Expect(dirname("//"), ToEqual, "/")
		})

		It("does `basename -d /.`", func() {
			Expect(dirname("/./"), ToEqual, "/.")
		})

		It("does `basename -d /a/b/c/d`", func() {
			Expect(dirname("/a/b/c/d"), ToEqual, "/a/b/c")
		})

		It("does `basename -d /a/b/c/d/`", func() {
			Expect(dirname("/a/b/c/d/"), ToEqual, "/a/b/c/d")
		})
	})
}
