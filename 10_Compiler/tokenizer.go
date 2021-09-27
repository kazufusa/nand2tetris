package compiler

import (
	"fmt"
	"html"
	"os"
	"regexp"
	"strconv"
)

var (
	reInt = regexp.MustCompile("^\\d*$")
)

type ITokenizer interface {
	HasMoreToken() bool
	Advance()
	TokenType() TokenTag
	KeyWord() KeyWordTag
	Symbol() string
	Identifier() string
	IntVal() int
	StringVal() string
}

type Tokenizer struct {
	tokens []Token
	jack   string
	index  int
}

var _ ITokenizer = (*Tokenizer)(nil)

func NewTokenizer(jack string) (*Tokenizer, error) {
	tk := Tokenizer{jack: jack}
	tk.parse()
	return &tk, nil
}

func (tk *Tokenizer) HasMoreToken() bool {
	return tk.index < len(tk.tokens)
}

func (tk *Tokenizer) Advance() {
	tk.index++
}

func (tk *Tokenizer) TokenType() TokenTag {
	return tk.tokens[tk.index].tokenType
}

func (tk *Tokenizer) KeyWord() KeyWordTag {
	return KeyWordTag(tk.tokens[tk.index].value)
}

func (tk *Tokenizer) Symbol() string {
	return tk.tokens[tk.index].value
}

func (tk *Tokenizer) Identifier() string {
	return tk.tokens[tk.index].value
}

func (tk *Tokenizer) IntVal() int {
	v, _ := strconv.Atoi(tk.tokens[tk.index].value)
	return v
}

func (tk *Tokenizer) StringVal() string {
	return tk.tokens[tk.index].value
}

func (tk *Tokenizer) parse() error {
	buf, err := os.ReadFile(tk.jack)
	if err != nil {
		return err
	}

	s := string(buf)

	comment1 := false
	comment2 := false
	stringConstant := false
	for i, r := range s {
		if comment1 && r == '\n' {
			comment1 = false
			continue
		} else if comment1 {
			continue
		} else if !comment1 && !comment2 && r == '/' && i < len(s)-2 && s[i+1] == '/' {
			comment1 = true
			continue
		}

		if comment2 && r == '/' && i > 0 && s[i-1] == '*' {
			comment2 = false
			continue
		} else if comment2 {
			continue
		} else if !comment1 && !comment2 && r == '/' && i < len(s)-2 && s[i+1] == '*' {
			comment2 = true
			continue
		}

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
	tokenType TokenTag
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
