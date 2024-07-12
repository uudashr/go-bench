package bench_test

import (
	"bytes"
	_ "embed"
	"io"
	"io/fs"
	"os"
	"strings"
	"testing"
	"testing/fstest"
)

func BenchmarkReadReader_DirFS_readAllJSONString(b *testing.B) {
	afs := dirFS()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, err := afs.Open("error.json")
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		out, err := io.ReadAll(f)
		if err != nil {
			b.Fatal(err)
		}

		s := string(out)
		_ = s

		b.StopTimer()
		f.Close()
		b.StartTimer()
	}
}

func BenchmarkReadReader_DirFS_stringsBuilderJSON(b *testing.B) {
	afs := dirFS()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, err := afs.Open("error.json")
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		var sb strings.Builder
		_, err = io.Copy(&sb, f)
		if err != nil {
			b.Fatal(err)
		}

		s := sb.String()
		_ = s

		b.StopTimer()
		f.Close()
		b.StartTimer()
	}
}

func BenchmarkReadReader_DirFS_bytesBufferJSONString(b *testing.B) {
	afs := dirFS()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, err := afs.Open("error.json")
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		var buf bytes.Buffer
		_, err = buf.ReadFrom(f)
		if err != nil {
			b.Fatal(err)
		}

		s := buf.String()
		_ = s

		b.StopTimer()
		f.Close()
		b.StartTimer()
	}
}

func BenchmarkReadReader_MapFS_readAllJSONString(b *testing.B) {
	afs := mapFS()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, err := afs.Open("error.json")
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		out, err := io.ReadAll(f)
		if err != nil {
			b.Fatal(err)
		}

		s := string(out)
		_ = s

		b.StopTimer()
		f.Close()
		b.StartTimer()
	}
}

func BenchmarkReadReader_MapFS_stringsBuilderJSON(b *testing.B) {
	afs := mapFS()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, err := afs.Open("error.json")
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		var sb strings.Builder
		_, err = io.Copy(&sb, f)
		if err != nil {
			b.Fatal(err)
		}

		s := sb.String()
		_ = s

		b.StopTimer()
		f.Close()
		b.StartTimer()
	}
}

func BenchmarkReadReader_MapFS_bytesBufferJSONString(b *testing.B) {
	afs := mapFS()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, err := afs.Open("error.json")
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		var buf bytes.Buffer
		_, err = buf.ReadFrom(f)
		if err != nil {
			b.Fatal(err)
		}

		s := buf.String()
		_ = s

		b.StopTimer()
		f.Close()
		b.StartTimer()
	}
}

//go:embed testdata/json/error.json
var jsonString string

func mapFS() fstest.MapFS {
	return fstest.MapFS{
		"error.json": {
			Data: []byte(jsonString),
		},
	}
}

func dirFS() fs.FS {
	return os.DirFS("testdata/json")
}
