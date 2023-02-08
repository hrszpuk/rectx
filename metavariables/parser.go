package metavariables

import "fmt"

type Parser struct {
	content   string
	output    string
	index     int
	variables map[string]string
}

func NewParser(content string, variables map[string]string) *Parser {
	fmt.Println(variables)

	return &Parser{
		content:   content,
		output:    "",
		index:     0,
		variables: variables,
	}
}

func (p *Parser) Parse() string {
	for p.index < len(p.content) {
		if p.current() == "%" {
			// Write %...% to a buffer
			buffer := ""
			for p.index < len(p.content) {
				buffer += p.current()
				if p.current() == "\n" || p.current() == " " ||
					p.current() == "\t" {
					break
				}
				p.index++
				if p.current() == "%" {
					buffer += p.current()
					fmt.Printf("BUFFER EXIT: %s\n", buffer)
					break
				}
			}

			if buffer[len(buffer)-1] != '%' {
				p.output += buffer
				p.index++
				continue
			}

			// If found, replace with metavariable value
			if metaContent, found := p.variables[buffer]; found {
				p.output += metaContent
				fmt.Printf("Replaced %s with %s\n", buffer, metaContent)
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
