package templates

import (
	"strings"
)

var KEYWORDS = []string{
	"folder", "file", "command",
}

type TemplateParser struct {
	content string
	tokens  []*Token
	index   int
	errors  []string
	token   *Token
}

func NewTemplateParser(content string) *TemplateParser {
	return &TemplateParser{
		content: content,
		tokens:  []*Token{},
		index:   0,
	}
}

func (tp *TemplateParser) Parse() {
	tp.Lex()
	/* Patterns:
	FOLDER STRING
	FOLDER STRING STRING
	FILE STRING
	FILE STRING STRING
	FILE STRING STRING BLOCK
	FILE STRING BLOCK
	COMMAND STRING
	*/
	tp.index = 0
	tp.token = tp.tokens[tp.index]

	for tp.index < len(tp.tokens) {
		if tp.token.Kind == KEYWORD_TKN {
			if strings.ToLower(tp.token.Value) == "folder" {
				tp.ParseFolder()
			} else if strings.ToLower(tp.token.Value) == "file" {
				tp.ParseFile()
			} else if strings.ToLower(tp.token.Value) == "command" {
				tp.ParseCommand()
			}
		} else {
			// ERROR
		}
	}
}

func (tp *TemplateParser) ParseFolder() {
	tp.index++
	var folderName string
	if tp.tokens[tp.index].Kind == STRING_TKN {
		folderName = tp.tokens[tp.index].Value
		tp.index++
	} else {
		// ERROR
	}

	var sourceFolder string
	if tp.tokens[tp.index].Kind == STRING_TKN {
		sourceFolder = tp.tokens[tp.index].Value
		tp.index++
	} else {
		// PATTERN ENDS... EXIT (check if keyword though)
	}

	// BUILD FOLDER STATEMENT
}

func (tp *TemplateParser) ParseFile() {

}

func (tp *TemplateParser) ParseCommand() {

}
