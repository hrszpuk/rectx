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
	line       int
	column     int
}

func NewTemplateParser(content string) *TemplateParser {
	return &TemplateParser{
		content: content,
		tokens:  []*Token{},
		index:   0,
	}
}

func (tp *TemplateParser) Parse() []Statement {
	tp.Lex()
	tp.index = 0

	for tp.index < len(tp.tokens) {
		if tp.tokens[tp.index].Kind == KEYWORD_TKN {
			if strings.ToLower(tp.tokens[tp.index].Value) == "folder" {
				tp.statements = append(tp.statements, tp.ParseFolder())
			} else if strings.ToLower(tp.tokens[tp.index].Value) == "file" {
				tp.statements = append(tp.statements, tp.ParseFile())
			} else if strings.ToLower(tp.tokens[tp.index].Value) == "command" {
				tp.statements = append(tp.statements, tp.ParseCommand())
			}
		} else {
			token := tp.tokens[tp.index]
			tp.index++
			tp.statements = append(
				tp.statements,
				NewBadStatement(NewToken("", KEYWORD_TKN, token.line, token.column), token),
			)
		}
	}
	return tp.statements
}

func (tp *TemplateParser) ParseFolder() Statement {
	tp.index++
	var folderName = ""
	if tp.tokens[tp.index].Kind == STRING_TKN {
		folderName = tp.tokens[tp.index].Value
		tp.index++
	} else {
		token := tp.tokens[tp.index]
		tp.index++
		return NewBadStatement(NewToken("", STRING_TKN, token.line, token.column), token)
	}

	var sourceFolder = ""
	if tp.tokens[tp.index].Kind == STRING_TKN {
		sourceFolder = tp.tokens[tp.index].Value
		tp.index++
	}

	return NewFolderStatement(folderName, sourceFolder)
}

func (tp *TemplateParser) ParseFile() Statement {
	tp.index++
	var fileName = ""
	if tp.tokens[tp.index].Kind == STRING_TKN {
		fileName = tp.tokens[tp.index].Value
		tp.index++
	} else {
		token := tp.tokens[tp.index]
		tp.index++
		return NewBadStatement(NewToken("", STRING_TKN, token.line, token.column), token)
	}

	var sourceFolder = ""
	if tp.tokens[tp.index].Kind == STRING_TKN {
		sourceFolder = tp.tokens[tp.index].Value
		tp.index++
	}

	var contentBlock = ""
	if tp.tokens[tp.index].Kind == CONTENT_TKN {
		contentBlock = tp.tokens[tp.index].Value
		tp.index++
	}

	return NewFileStatement(fileName, sourceFolder, contentBlock)
}

func (tp *TemplateParser) ParseCommand() Statement {
	tp.index++
	var command string
	if tp.tokens[tp.index].Kind == STRING_TKN {
		command = tp.tokens[tp.index].Value
		tp.index++
	} else {
		token := tp.tokens[tp.index]
		tp.index++
		return NewBadStatement(NewToken("", STRING_TKN, token.line, token.column), token)
	}

	return NewCommandStatement(command)
}
