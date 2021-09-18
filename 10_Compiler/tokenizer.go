package compiler

import (
	"fmt"
	"html"
	"os"
	"regexp"
)

var (
	reComment1 = regexp.MustCompile("(?m)//.*$")
	reComment2 = regexp.MustCompile(`/\*.*\*/`)
	reComment3 = regexp.MustCompile(`(?s)/\*\*.*\*/`)

	reInt = regexp.MustCompile("^\\d*$")
)

type ITokenizer interface {
	HasMoreToken() bool
	Advance()
	TokenType() TokenType
	KeyWord() KeyWordType
	Symbol() string
	Identifier() string
	IntVal() int
	StringVal() string
}

type Tokenizer struct {
	tokens []Token
	jack   string
}

func NewTokenizer(jack string) (*Tokenizer, error) {
	tk := Tokenizer{jack: jack}
	tk.parse()
	return &tk, nil
}

func (tk *Tokenizer) parse() error {
	buf, err := os.ReadFile(tk.jack)
	if err != nil {
		return err
	}

	s := string(buf)
	s = reComment1.ReplaceAllString(s, "")
	s = reComment2.ReplaceAllString(s, "")
	s = reComment3.ReplaceAllString(s, "")

	stringConstant := false
	for _, r := range s {
		if stringConstant && r == '"' {
			stringConstant = false
			tk.tokens = append(tk.tokens, Token{})
			continue
		} else if stringConstant {
			token := tk.lastToken()
			token.value += string(r)
			continue
		} else if !stringConstant && r == '"' {
			stringConstant = true
			tk.tokens = append(tk.tokens, Token{tokenType: TkStringConst})
			continue
		}

		switch r {
		case '\r', '\n', ' ', '	':
			if !tk.lastTokenIsEmpty() {
				tk.tokens = append(tk.tokens, Token{})
			}
		case '{', '}', '(', ')', '[', ']', '.', ',', ';',
			'+', '-', '*', '/', '&', '|', '<', '>', '=', '~':
			if tk.lastTokenIsEmpty() {
				lastToken := tk.lastToken()
				lastToken.tokenType = TkSymbol
				lastToken.value = string([]rune{r})
			} else {
				tk.tokens = append(tk.tokens, Token{
					tokenType: TkSymbol,
					value:     string([]rune{r}),
				})
			}
			tk.tokens = append(tk.tokens, Token{})
		default:
			lastToken := tk.lastToken()
			if lastToken == nil {
				tk.tokens = append(tk.tokens, Token{value: string(r)})
				lastToken = tk.lastToken()
			} else {
				lastToken.value += string(r)
			}
		}
	}
	tk.finalizeTokens()
	return nil
}

func (tk *Tokenizer) finalizeTokens() {
	for i := len(tk.tokens) - 1; i >= 0; i-- {
		token := &tk.tokens[i]
		if token.value == "" {
			tk.tokens = append(tk.tokens[0:i], tk.tokens[i+1:]...)
			continue
		} else if token.tokenType != "" {
			continue
		}

		if reInt.MatchString(token.value) {
			token.tokenType = TkIntConst
			continue
		}

		switch token.value {
		case string(KwClass),
			string(KwMethod),
			string(KwFunction),
			string(KwConstructor),
			string(KwInt),
			string(KwBoolean),
			string(KwChar),
			string(KwVoid),
			string(KwVar),
			string(KwStatic),
			string(KwField),
			string(KwLet),
			string(KwDo),
			string(KwIf),
			string(KwElse),
			string(KwWhile),
			string(KwReturn),
			string(KwTrue),
			string(KwFalse),
			string(KwNull),
			string(KwThis):
			token.tokenType = TkKeyWord
			continue
		default:
			token.tokenType = TkIdentifier
		}
	}
}

func (tk *Tokenizer) lastToken() *Token {
	if len(tk.tokens) == 0 {
		return nil
	} else {
		return &tk.tokens[len(tk.tokens)-1]
	}
}

func (tk *Tokenizer) lastTokenIsEmpty() bool {
	token := tk.lastToken()
	if token == nil {
		return false
	} else {
		return token.value == ""
	}
}

func (tk *Tokenizer) ToXml() string {
	s := ""
	for _, t := range tk.tokens {
		s += t.ToString() + "\n"
	}
	return "<tokens>\n" + s + "</tokens>\n"
}

type Token struct {
	tokenType TokenType
	value     string
}

func (t *Token) ToString() string {
	return fmt.Sprintf(
		"<%s> %s </%s>",
		t.tokenType,
		html.EscapeString(t.value),
		t.tokenType,
	)
}
