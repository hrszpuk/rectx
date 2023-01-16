package templates

import (
	"strings"
)

var KEYWORDS = []string{
	"folder", "file", "command",
}

type TemplateParser struct {
	content    string
	tokens     []*Token
	index      int
	errors     []string
	token      *Token
	statements []Statement
}

func NewTemplateParser(content string) *TemplateParser {
	return &TemplateParser{
		content: content,
		tokens:  []*Token{},
		index:   0,
	}
}

func (tp *TemplateParser) Parse() []*Statement {
	tp.Lex()
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

	stmt := NewFolderStatement(folderName, sourceFolder)
	tp.statements = append(tp.statements, stmt)
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

	stmt := NewFileStatement(fileName, sourceFolder, contentBlock)
	tp.statements = append(tp.statements, stmt)
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

	stmt := NewCommandStatement(command)
	tp.statements = append(tp.statements, stmt)
}
