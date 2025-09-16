package renderer

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Render(t *testing.T) {
	src, expected := ReadTestCase("./testdata/Link.txt")
	html, err := Render(t.Context(), src)
	assert.NoError(t, err)
	assert.Equal(t, expected, html)
}

func Test_Render_List(t *testing.T) {
	src, expected := ReadTestCase("./testdata/List.txt")
	html, err := Render(t.Context(), src)
	assert.NoError(t, err)
	assert.Equal(t, expected, html)
}

func Test_Render_Heading(t *testing.T) {
	src, expected := ReadTestCase("./testdata/Heading.txt")
	html, err := Render(t.Context(), src)
	assert.NoError(t, err)
	assert.Equal(t, expected, html)
}

func Test_Render_Collapse(t *testing.T) {
	src, expected := ReadTestCase("./testdata/Collapse.txt")
	html, err := Render(t.Context(), src)
	assert.NoError(t, err)
	assert.Equal(t, expected, html)
}

// テストケースと期待される出力をtestdataのファイルから取得する
func ReadTestCase(path string) (string, string) {
	f, err := os.Open(path)
	if err != nil {
		return "", err.Error()
	}
	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()

	// ファイルの内容を読み込む
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	if err != nil {
		return "", err.Error()
	}

	// `\n//\n`で分割して、各行をテストケースとして扱う
	testCases := strings.Split(buf.String(), "\n//\n")
	if len(testCases) < 2 {
		return "", "Test case or expected output not found"
	}
	src := testCases[0]      // 最初の行をソースとして使用
	expected := testCases[1] // 2行目を期待される出力として使用

	return src, expected
}
