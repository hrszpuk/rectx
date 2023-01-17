package templates

func (tp *TemplateParser) Lex() {
	for tp.index < len(tp.content) {
		if tp.current() == "#" {
			for tp.index < len(tp.content) {
				tp.index++
				if tp.current() == "\n" {
					break
				}
			}
			tp.index++
			continue
		} else if tp.current() == "\"" {
			buffer := ""
			tp.index++
			for tp.index < len(tp.content) && tp.current() != "\"" {
				buffer += tp.current()
				tp.index++
			}
			tp.index++
			tp.tokens = append(tp.tokens, NewToken(buffer, STRING_TKN))
		} else if tp.current() == "{" && tp.index+1 < len(tp.content) && tp.content[tp.index+1] == '%' {
			buffer := ""
			tp.index += 2
			for tp.index < len(tp.content) {
				buffer += tp.current()
				tp.index++
				if tp.current() == "%" && tp.index+1 < len(tp.content) && tp.content[tp.index+1] == '}' {
					break
				}
			}
			tp.tokens = append(tp.tokens, NewToken(buffer, CONTENT_TKN))
		} else if tp.current() == " " || tp.current() == "\t" || tp.current() == "\n" {
			tp.index++
		} else {
			buffer := ""
			for tp.index < len(tp.content) && tp.current() != " " &&
				tp.current() != "\t" && tp.current() != "\n" {
				buffer += tp.current()
				tp.index++
			}
			for _, keyword := range KEYWORDS {
				if buffer == keyword {
					tp.tokens = append(tp.tokens, NewToken(buffer, KEYWORD_TKN))
				}
			}
			tp.index++
		}
	}
}

func (tp *TemplateParser) current() string {
	return string(tp.content[tp.index])
}
