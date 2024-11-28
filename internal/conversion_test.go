package internal_test

import (
	"asciiartweb/nyeltay/algaliyev/internal"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "",
			want:  "test-1.txt",
		},
		{
			input: "hello",
			want:  "test-2.txt",
		},
		{
			input: "HELLO",
			want:  "test-3.txt",
		},
		{
			input: "HeLlo HuMaN",
			want:  "test-4.txt",
		},
		{
			input: "1Hello 2There",
			want:  "test-5.txt",
		},
		{
			input: `Hello\nThere`,
			want:  "test-6.txt",
		},
		{
			input: "{Hello & There #}",
			want:  "test-7.txt",
		},
		{
			input: "hello There 1 to 2!",
			want:  "test-8.txt",
		},
		{
			input: "MaD3IrA&LiSboN",
			want:  "test-9.txt",
		},
		{
			input: "1a\"#FdwHywR&/()=",
			want:  "test-10.txt",
		},
		{
			input: "{|}~",
			want:  "test-11.txt",
		},
		{
			input: `[\]^_ 'a`,
			want:  "test-12.txt",
		},
		{
			input: "RGB",
			want:  "test-13.txt",
		},
		{
			input: ":;<=>?@",
			want:  "test-14.txt",
		},
		{
			input: `\!" #$%&'()*+,-./`,
			want:  "test-15.txt",
		},
		{
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			want:  "test-16.txt",
		},
		{
			input: "abcdefghijklmnopqrstuvwxyz",
			want:  "test-17.txt",
		},
	}

	for _, testCase := range testCases {
		got, _ := internal.Convert(testCase.input, "standard.txt")
		want, _ := os.ReadFile("./testcases/" + testCase.want)

		if got != string(want) {
			t.Fatalf("got:\n%s\nwant\n%s\n", got, string(want))
		}
	}
}

func TestValidInput(t *testing.T) {
	testCases := []struct {
		input  string
		banner string
	}{
		{
			input:  "фывфыв",
			banner: "standard.txt",
		},
	}

	for _, testCase := range testCases {
		_, err := internal.ValidInput(testCase.input)

		if err == nil {
			t.Fatalf("\n%s\nexpected error\n", testCase.input)
		}
	}
}

func TestBanner(t *testing.T) {
	testCases := []struct {
		input  string
		banner string
	}{
		{
			input:  "фывфыв",
			banner: "std.txt",
		},
	}

	for _, testCase := range testCases {
		_, err := internal.Convert(testCase.input, testCase.banner)

		if err == nil {
			t.Fatalf("\n%s\nexpected error\n", testCase.input)
		}
	}
}
