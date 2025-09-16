package extension

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// extender

type collapse struct {
}

var Collapse = &collapse{}

func NewCollapse() goldmark.Extender {
	return &collapse{}
}

func (c *collapse) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(
				NewCollapseParser(),
				100, // Set priority to 100 to ensure it runs before other parsers
			),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(
				NewCollapseHTMLRenderer(),
				100, // Set priority to 100 to ensure it runs before other renderers
			),
		),
	)
}

// parser

type collapseParser struct {
}

var defaultCollapseParser = &collapseParser{}

// NewCollapseParser returns a new BlockParser that
// parses collapsible sections.
func NewCollapseParser() parser.BlockParser {
	return defaultCollapseParser
}

func (b *collapseParser) Trigger() []byte {
	return []byte{'|'}
}

func (b *collapseParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, _ := reader.PeekLine()
	w, pos := util.IndentWidth(line, reader.LineOffset())
	// '|>('で始まる行を探す
	indentSpacesNum := 3
	if w > indentSpacesNum || pos >= len(line) || line[pos] != '|' {
		return nil, parser.NoChildren
	}
	pos++
	if pos >= len(line) || line[pos] != '>' {
		return nil, parser.NoChildren
	}
	pos++
	if pos >= len(line) || (line[pos] != '(') {
		return nil, parser.NoChildren
	}
	pos++

	// <summary>部分の処理　残りの文字列を読み飛ばす
	//　同時にsummary部分を抽出する
	summaryText := line[pos:]
	if len(summaryText) > 0 && summaryText[len(summaryText)-1] == '\n' {
		summaryText = summaryText[:len(summaryText)-1]
	}
	for pos < len(line) && line[pos] != '\n' {
		pos++
	}
	reader.Advance(pos)
	// CollapseBlock ノードを作成
	// NewCollapseBlockに summary の内容を渡すと記録される
	collapseNode := NewCollapseBlock(summaryText)
	return collapseNode, parser.HasChildren
}

func (b *collapseParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, _ := reader.PeekLine()
	w, pos := util.IndentWidth(line, reader.LineOffset())

	// 空行の場合は続行
	if util.IsBlank(line) {
		return parser.Continue | parser.HasChildren
	}

	// ')'の行が来たら終了
	indentSpacesNum := 3
	if w <= indentSpacesNum && pos < len(line) && line[pos] == ')' {
		reader.AdvanceToEOL()
		return parser.Close
	}

	// その他の行は子ノードとして処理
	return parser.Continue | parser.HasChildren
}

func (b *collapseParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	// nothing to do
}

func (b *collapseParser) CanInterruptParagraph() bool {
	return true
}

func (b *collapseParser) CanAcceptIndentedLine() bool {
	return true
}

// html renderer

type collapseHTMLRenderer struct {
}

func NewCollapseHTMLRenderer() *collapseHTMLRenderer {
	return &collapseHTMLRenderer{}
}

func (r *collapseHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(kindCollapse, r.renderCollapse)
}

func (r *collapseHTMLRenderer) renderCollapse(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		collapseNode, ok := node.(*CollapseBlock)
		if !ok {
			return ast.WalkStop, nil
		}
		// <details> タグを開始
		_, err := w.WriteString("<details class=\"collapse\">\n")
		if err != nil {
			return ast.WalkStop, err
		}
		// <summary> タグを開始
		_, err = w.WriteString("<summary>")
		if err != nil {
			return ast.WalkStop, err
		}
		// <summary> 内のテキストを出力
		_, err = w.Write(collapseNode.summary)
		if err != nil {
			return ast.WalkStop, err
		}
		_, err = w.WriteString("</summary>\n")
		if err != nil {
			return ast.WalkStop, err
		}
	} else {
		// </details> タグを終了
		_, err := w.WriteString("</details>\n")
		if err != nil {
			return ast.WalkStop, err
		}
	}
	return ast.WalkContinue, nil
}

// CollapseBlock
type CollapseBlock struct {
	ast.BaseBlock

	// summary
	// 折りたたみ(Detail)のsummary部分を保持する
	summary []byte
}

var kindCollapse = ast.NewNodeKind("Collapse")

func (n *CollapseBlock) Kind() ast.NodeKind {
	return kindCollapse
}

func (n *CollapseBlock) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

func NewCollapseBlock(summary []byte) *CollapseBlock {
	return &CollapseBlock{
		BaseBlock: ast.BaseBlock{},
		summary:   summary,
	}
}
