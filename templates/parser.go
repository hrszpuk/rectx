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
			// TODO ERROR
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
		// TODO ERROR
	}

	var sourceFolder string
	if tp.tokens[tp.index].Kind == STRING_TKN {
		sourceFolder = tp.tokens[tp.index].Value
		tp.index++
	} else {
		// TODO PATTERN ENDS... EXIT
	}

	// TODO BUILD FOLDER STATEMENT
}

func (tp *TemplateParser) ParseFile() {
	tp.index++
	var fileName string
	if tp.tokens[tp.index].Kind == STRING_TKN {
		fileName = tp.tokens[tp.index].Value
		tp.index++
	} else {
		// TODO ERROR
	}

	var sourceFolder string
	if tp.tokens[tp.index].Kind == STRING_TKN {
		sourceFolder = tp.tokens[tp.index].Value
		tp.index++
	} else {
		// TODO PATTERN ENDS... EXIT
	}

	var contentBlock string
	if tp.tokens[tp.index].Kind == CONTENT_TKN {
		contentBlock = tp.tokens[tp.index].Value
		tp.index++
	} else {
		// TODO PATTERN ENDS... EXIT (check if keyword next though)
	}

	// TODO BUILD FILE STATEMENT
}

func (tp *TemplateParser) ParseCommand() {
	tp.index++
	var command string
	if tp.tokens[tp.index].Kind == STRING_TKN {
		command = tp.tokens[tp.index].Value
		tp.index++
	} else {
		// TODO ERROR
	}

	// TODO: BUILD COMMAND STATEMENT
}
