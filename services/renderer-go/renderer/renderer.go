package renderer

import (
	"bytes"
	"context"

	customextension "github.com/hatena/Hatena-Intern-2025/services/renderer-go/renderer/extension"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	var buf bytes.Buffer
	// Markdown を HTML に変換するための Goldmark エンジンを設定
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			customextension.NewCollapse(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	// Markdown を HTML に変換
	err := md.Convert([]byte(src), &buf)
	if err != nil {
		return "", err
	}
	// buf.Stringの行末の改行を削除
	bufString := buf.String()
	if len(bufString) > 0 && bufString[len(bufString)-1] == '\n' {
		bufString = bufString[:len(bufString)-1]
	}
	// 変換された HTML を返す
	return bufString, nil
}
