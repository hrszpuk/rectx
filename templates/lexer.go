package templates

func (tp *TemplateParser) Lex() {
	for tp.index < len(tp.content) {
		if tp.current() == "#" {
			for tp.index < len(tp.content) {
				tp.increment()
				if tp.current() == "\n" {
					tp.column = 0
					tp.line += 1
					break
				}
			}
			tp.increment()
			continue
		} else if tp.current() == "\"" {
			buffer := ""
			tp.increment()
			for tp.index < len(tp.content) && tp.current() != "\"" {
				buffer += tp.current()
				tp.increment()
				if tp.current() == "\n" {
					tp.line++
					tp.column = 1
				}
			}
			tp.increment()
			tp.tokens = append(tp.tokens, NewToken(buffer, STRING_TKN, tp.line, tp.column))
		} else if tp.current() == "{" && tp.index+1 < len(tp.content) && tp.content[tp.index+1] == '%' {
			buffer := ""
			tp.index += 2
			tp.column += 2
			contentColumn := tp.column
			contentLine := tp.line
			for tp.index < len(tp.content) {
				buffer += tp.current()
				tp.increment()
				if tp.current() == "%" && tp.index+1 < len(tp.content) && tp.content[tp.index+1] == '}' {
					break
				} else if tp.current() == "\n" {
					tp.column = 1
					tp.index++
					tp.line++
				}
			}
			tp.index += 2
			tp.column += 2
			tp.tokens = append(tp.tokens, NewToken(buffer, CONTENT_TKN, contentLine, contentColumn))
		} else if tp.current() == " " || tp.current() == "\t" || tp.current() == "\n" {
			if tp.current() == "\n" {
				tp.column = 1
				tp.index++
				tp.line++
			} else {
				tp.increment()
			}
		} else {
			buffer := ""
			for tp.index < len(tp.content) && tp.current() != " " &&
				tp.current() != "\t" && tp.current() != "\n" {
				buffer += tp.current()
				tp.increment()
			}
			for _, keyword := range KEYWORDS {
				if buffer == keyword {
					tp.tokens = append(tp.tokens, NewToken(buffer, KEYWORD_TKN, tp.line, tp.column))
				}
			}
			if tp.current() == "\n" {
				tp.column = 1
				tp.index++
				tp.line++
			} else {
				tp.increment()
			}
		}
	}
}

func (tp *TemplateParser) current() string {
	return string(tp.content[tp.index])
}

func (tp *TemplateParser) increment() {
	tp.index++
	tp.column++
}
