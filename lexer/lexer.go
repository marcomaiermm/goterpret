package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // have we reached the end of our position in the input string?
		l.ch = 0 // ASCII code for "NUL" -> "we havent read anything yet" or "end of file"
	} else {
		l.ch = l.input[l.readPosition] // set l.ch to the next character
	}
	l.position = l.readPosition // always point to the position where we last read
	l.readPosition += 1         // always point to the next position
}
