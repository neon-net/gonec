package gonec

import (
	"errors"
	"io"
	"strings"

	"github.com/covrom/gonec/gonecscan"
)

func (i *interpreter) Lexer(r io.Reader, w io.Writer) (tokens []token, err error) {

	//лексический анализ
	tokens = []token{}

	var s gonecscan.Scanner

	s.Error = func(s *gonecscan.Scanner, msg string) {
		err = errors.New(msg)
	}

	s.Init(r)

	var tok rune

	for tok != gonecscan.EOF {
		tok = s.Scan()
		if err != nil {
			return
		}

		//fmt.Println(s.Pos(), ":", gonecscan.TokenString(tok), ":", s.TokenText())

		nt := token{literal: s.TokenText(),
			srcline: s.Line,
			srccol:  s.Column,
		}

		ntlit := strings.ToLower(nt.literal)
		var ok bool
		switch tok {
		case gonecscan.Ident:
			nt.toktype, ok = keywordMap[ntlit]
			if !ok {
				nt.category = defIdentifier
			} else {
				nt.category = defKeyword
				if breaksMap[nt.toktype] == true {
					//вставляем разделитель ;
					tokens = append(tokens, token{literal: ";",
						srcline: s.Line,
						srccol:  s.Column,
						category: defDelimiter,
						toktype: oSemi,
					})
					//сбрасываем признак предыдущего присвоения
					s.Preassign = false
				}
			}
		case gonecscan.Assignator:
			nt.category = defAssignator
		case gonecscan.String:
			// строки возвращаются без переносов и комментариев
			nt.category = defValueString
		case gonecscan.Int:
			nt.category = defValueInt
		case gonecscan.Float:
			nt.category = defValueFloat
		case gonecscan.Date:
			nt.category = defValueDate
		case gonecscan.EOF:
			nt.category = defEOF
		default:
			nt.toktype, ok = operMap[ntlit]
			if !ok {
				nt.toktype, ok = delimMap[ntlit]
				if !ok {
					nt.toktype, ok = pointMap[ntlit]
					if !ok {
						nt.category = defUnknown
					} else {
						nt.category = defPoint
					}
				} else {
					nt.category = defDelimiter
				}
			} else {
				nt.category = defOperator
			}
		}
		tokens = append(tokens, nt)
	}

	tokens = append(tokens, token{category: defEOF}) //гарантированное окончание даже для пустого файла

	return
}