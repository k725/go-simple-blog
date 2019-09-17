package markdown

import (
	"github.com/russross/blackfriday/v2"
)

func Render(text string) string {
	//HTMLFlags and Renderer
	htmlFlags := blackfriday.CommonHTMLFlags         //UseXHTML | Smartypants | SmartypantsFractions | SmartypantsDashes | SmartypantsLatexDashes
	htmlFlags |= blackfriday.FootnoteReturnLinks     //Generate a link at the end of a footnote to return to the source
	htmlFlags |= blackfriday.SmartypantsAngledQuotes //Enable angled double quotes (with Smartypants) for double quotes rendering
	htmlFlags |= blackfriday.SmartypantsQuotesNBSP   //Enable French guillemets 損 (with Smartypants)
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{Flags: htmlFlags, Title: "", CSS: ""})

	//Extensions
	extFlags := blackfriday.CommonExtensions //NoIntraEmphasis | Tables | FencedCode | Autolink | Strikethrough | SpaceHeadings | HeadingIDs | BackslashLineBreak | DefinitionLists
	extFlags |= blackfriday.Footnotes        //Pandoc-style footnotes
	extFlags |= blackfriday.HeadingIDs       //specify heading IDs  with {#id}
	extFlags |= blackfriday.Titleblock       //Titleblock ala pandoc
	extFlags |= blackfriday.DefinitionLists  //Render definition lists

	html := blackfriday.Run([]byte(text), blackfriday.WithExtensions(extFlags), blackfriday.WithRenderer(renderer))
	return string(html)
}
