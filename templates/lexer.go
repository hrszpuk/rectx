package templates

func (tp *TemplateParser) Lex() {
	for tp.index < len(tp.content) {
		if tp.current() == "#" {
			for tp.index < len(tp.content) {
				tp.index++
				tp.column++
				if tp.current() == "\n" {
					tp.column = 0
					tp.line += 1
					break
				}
			}
			tp.index++
			tp.column++
			continue
		} else if tp.current() == "\"" {
			buffer := ""
			tp.index++
			tp.column++
			for tp.index < len(tp.content) && tp.current() != "\"" {
				buffer += tp.current()
				tp.index++
				tp.column++
			}
			tp.index++
			tp.column++
			tp.tokens = append(tp.tokens, NewToken(buffer, STRING_TKN, tp.line, tp.column))
		} else if tp.current() == "{" && tp.index+1 < len(tp.content) && tp.content[tp.index+1] == '%' {
			buffer := ""
			tp.index += 2
			tp.column += 2
			for tp.index < len(tp.content) {
				buffer += tp.current()
				tp.index++
				tp.column++
				if tp.current() == "%" && tp.index+1 < len(tp.content) && tp.content[tp.index+1] == '}' {
					break
				} else if tp.current() == "\n" {
					tp.column = 0
					tp.line++
				}
			}
			tp.tokens = append(tp.tokens, NewToken(buffer, CONTENT_TKN, tp.line, tp.column))
		} else if tp.current() == " " || tp.current() == "\t" || tp.current() == "\n" {
			tp.index++
			tp.column++
		} else {
			buffer := ""
			for tp.index < len(tp.content) && tp.current() != " " &&
				tp.current() != "\t" && tp.current() != "\n" {
				buffer += tp.current()
				tp.index++
			}
			for _, keyword := range KEYWORDS {
				if buffer == keyword {
					tp.tokens = append(tp.tokens, NewToken(buffer, KEYWORD_TKN, tp.line, tp.column))
				}
			}
			tp.index++
		}
	}
}

func (tp *TemplateParser) current() string {
	return string(tp.content[tp.index])
}
