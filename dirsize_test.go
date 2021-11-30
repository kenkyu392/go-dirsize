package dirsize

import (
	"errors"
	"io/fs"
	"testing"
)

func TestCalc(t *testing.T) {
	const root = "./testdata"
	t.Run("case=B", func(t *testing.T) {
		got, err := Calc(root, B)
		if err != nil {
			t.Fatal(err)
		}
		if want := 11264.00; got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("case=KB", func(t *testing.T) {
		got, err := Calc(root, KB)
		if err != nil {
			t.Fatal(err)
		}
		if want := 11.00; got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("case=MB", func(t *testing.T) {
		got, err := Calc(root, MB)
		if err != nil {
			t.Fatal(err)
		}
		if want := 0.01; got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("case=error", func(t *testing.T) {
		_, err := Calc("./unknown", B)
		if err == nil {
			t.Errorf("must be an error")
		}
	})
}

func TestOptionFunc(t *testing.T) {
	const root = "./testdata"
	t.Run("case=IgnoreFile/IgnoreDir", func(t *testing.T) {
		got, err := Calc(root, B, IgnoreFile("ignore_*"), IgnoreDir("ignore"))
		if err != nil {
			t.Fatal(err)
		}
		if want := 3072.00; got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("case=Ignore", func(t *testing.T) {
		got, err := Calc(root, B, Ignore("ignore*"))
		if err != nil {
			t.Fatal(err)
		}
		if want := 3072.00; got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("case=error", func(t *testing.T) {
		want := errors.New("unknown")
		_, err := Calc(root, B, func(_ string, _ fs.FileInfo, _ error) error {
			return want
		})
		if !errors.Is(err, want) {
			t.Errorf("got: %v, want: %v", err, want)
		}
	})
}
