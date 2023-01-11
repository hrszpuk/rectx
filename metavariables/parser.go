package metavariables

type Parser struct {
	content       string
	output        string
	index         int
	metavariables map[string]string
}

func NewParser(content string, metavariables map[string]string) *Parser {
	return &Parser{
		content:       content,
		output:        "",
		index:         0,
		metavariables: metavariables,
	}
}

func (p *Parser) parse() string {
	for p.index < len(p.content) {
		if p.current() == "%" {
			// Write %...% to a buffer
			buffer := ""
			for p.index < len(p.content) {
				buffer += p.current()
				if p.current() == "\n" || p.current() == " " ||
					p.current() == "\t" || p.current() == "%" {
					break
				}
				p.index++
			}

			if buffer[len(buffer)-1] != '%' {
				p.output += buffer
				p.index++
				continue
			}

			// Lookup %...% in metavariables
			metacontent, ok := p.metavariables[buffer]

			// If found, replace with metavariable value
			if ok {
				p.output = metacontent
				p.index++
			} else {
				// Otherwise, write buffer to output
				p.output += buffer
				p.index++
			}
		} else {
			p.output += p.current()
			p.index++
		}
	}
	return p.output
}

func (p *Parser) current() string {
	return string(p.content[p.index])
}
