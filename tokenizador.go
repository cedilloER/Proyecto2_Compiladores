package main

import (
	"io"
	"sort"
	"sync"
)

const newLine = '\n'

type TokenKey int

const (
	TokenUnknown        TokenKey = -6
	TokenStringFragment TokenKey = -5
	TokenString         TokenKey = -4
	TokenFloat          TokenKey = -3
	TokenInteger        TokenKey = -2
	TokenKeyword        TokenKey = -1
	TokenUndef          TokenKey = 0
)

const (
	fStopOnUnknown          uint16 = 0b1
	fAllowKeywordUnderscore uint16 = 0b10
	fAllowNumberUnderscore  uint16 = 0b100
	fAllowNumberInKeyword   uint16 = 0b1000
)

const BackSlash = '\\' //comentario

var defaultWhiteSpaces = []byte{' ', '\t', '\n', '\r'}

// Delimitadores
var DefaultStringEscapes = map[byte]byte{
	'n':  '\n',
	'r':  '\r',
	't':  '\t',
	'\\': '\\',
}

type tokenRef struct {
	Key   TokenKey
	Token []byte
}

type QuoteInjectSettings struct {
	StartKey TokenKey
	EndKey   TokenKey
}

type StringSettings struct {
	Key          TokenKey
	StartToken   []byte
	EndToken     []byte
	EscapeSymbol byte
	SpecSymbols  map[byte]byte
	Injects      []QuoteInjectSettings
}

func (q *StringSettings) AddInjection(startTokenKey, endTokenKey TokenKey) *StringSettings {
	q.Injects = append(q.Injects, QuoteInjectSettings{StartKey: startTokenKey, EndKey: endTokenKey})
	return q
}

func (q *StringSettings) SetEscapeSymbol(symbol byte) *StringSettings {
	q.EscapeSymbol = symbol
	return q
}

func (q *StringSettings) SetSpecialSymbols(special map[byte]byte) *StringSettings {
	q.SpecSymbols = special
	return q
}

type Tokenizer struct {
	flags   uint16
	tokens  map[TokenKey][]*tokenRef
	index   map[byte][]*tokenRef
	quotes  []*StringSettings
	wSpaces []byte
	pool    sync.Pool
}

func New() *Tokenizer {
	t := Tokenizer{
		flags:   0,
		tokens:  map[TokenKey][]*tokenRef{},
		index:   map[byte][]*tokenRef{},
		quotes:  []*StringSettings{},
		wSpaces: defaultWhiteSpaces,
	}
	t.pool.New = func() interface{} {
		return new(Token)
	}
	return &t
}

func (t *Tokenizer) SetWhiteSpaces(ws []byte) *Tokenizer {
	t.wSpaces = ws
	return t
}

func (t *Tokenizer) StopOnUndefinedToken() *Tokenizer {
	t.flags |= fStopOnUnknown
	return t
}

func (t *Tokenizer) AllowKeywordUnderscore() *Tokenizer {
	t.flags |= fAllowKeywordUnderscore
	return t
}

func (t *Tokenizer) AllowNumbersInKeyword() *Tokenizer {
	t.flags |= fAllowNumberInKeyword
	return t
}

func (t *Tokenizer) DefineTokens(key TokenKey, tokens []string) *Tokenizer {
	var tks []*tokenRef
	if key < 1 {
		return t
	}
	for _, token := range tokens {
		ref := tokenRef{
			Key:   key,
			Token: s2b(token),
		}
		head := ref.Token[0]
		tks = append(tks, &ref)
		if t.index[head] == nil {
			t.index[head] = []*tokenRef{}
		}
		t.index[head] = append(t.index[head], &ref)
		sort.Slice(t.index[head], func(i, j int) bool {
			return len(t.index[head][i].Token) > len(t.index[head][j].Token)
		})
	}
	t.tokens[key] = tks

	return t
}
func (t *Tokenizer) DefineStringToken(key TokenKey, startToken, endToken string) *StringSettings {
	q := &StringSettings{
		Key:        key,
		StartToken: s2b(startToken),
		EndToken:   s2b(endToken),
	}
	if q.StartToken == nil {
		return q
	}
	t.quotes = append(t.quotes, q)

	return q
}

func (t *Tokenizer) allocToken() *Token {
	return t.pool.Get().(*Token)
}

func (t *Tokenizer) freeToken(token *Token) {
	token.next = nil
	token.prev = nil
	token.value = nil
	token.indent = nil
	token.offset = 0
	token.line = 0
	token.id = 0
	token.key = 0
	token.string = nil
	t.pool.Put(token)
}

func (t *Tokenizer) ParseString(str string) *Stream {
	return t.ParseBytes(s2b(str))
}

func (t *Tokenizer) ParseBytes(str []byte) *Stream {
	p := newParser(t, str)
	p.parse()
	return NewStream(p)
}

func (t *Tokenizer) ParseStream(r io.Reader, bufferSize uint) *Stream {
	p := newInfParser(t, r, bufferSize)
	p.preload()
	p.parse()
	return NewInfStream(p)
}

// Prueba de tokenizador
/*
func TestTokenize(t *testing.T) {
	type item struct {
		str   string
		token Token
	}
	tokenizer := New()


	for _, v := range data {
		stream := tokenizer.ParseBytes([]byte(v.str))
		expected := &v.token
		expected.value = []byte(v.str)
		actual := &Token{
			value:  stream.current.value,
			key:    stream.current.key,
			string: stream.current.string,
		}
		require.Equalf(t, expected, actual, "parse %s: %s", v.str, stream.current)
	}
}
*/
