package lexer

type TextScanner struct {
	input  []rune
	offset int
	char   rune
	eof    bool
	line   int
	column int
	stack  []SourcePosition
}

type SourcePosition struct {
	Offset int
	Line   int
	Column int
}

func NewTextScanner(source string) *TextScanner {
	s := TextScanner{input: []rune(source)}
	if len(s.input) > 0 {
		s.char = s.input[0]
		s.eof = false
	}
	s.stack = make([]SourcePosition, 0)
	return &s
}

// 当前字符
func (s *TextScanner) Current() rune {
	return s.char
}

func (s *TextScanner) safeGetChar(offset int) rune {
	max := len(s.input)
	if offset < max {
		return s.input[offset]
	}
	return 0
}

// 移到下一个字符
func (s *TextScanner) Next() rune {
	s.offset++
	s.char = s.safeGetChar(s.offset)
	if s.char == 0 {
		s.eof = true
	}
	s.column++
	if s.char == '\n' {
		s.line++
		s.column = 0
	}
	return s.char
}

// 偷看下一个字符
func (s *TextScanner) Peek() rune {
	return s.safeGetChar(s.offset + 1)
}

// 当前位置
func (s *TextScanner) Offset() int {
	return s.offset
}

// 当前位置信息
func (s *TextScanner) Position() SourcePosition {
	return SourcePosition{Offset: s.offset, Line: s.line, Column: s.column}
}

// 是否到末尾
func (s *TextScanner) EOF() bool {
	return s.eof
}

// 将当前位置压栈
func (s *TextScanner) StackPush() {
	p := SourcePosition{Offset: s.offset, Line: s.line, Column: s.column}
	s.stack = append(s.stack, p)
	// log.Printf("%d '%s' StackPush: %+v\n", s.offset, string(s.char), p)
}

// 从栈中弹出位置
// reset 如果为true则根据栈顶的信息重置当前位置
func (s *TextScanner) StackPop(reset bool) {
	i := len(s.stack) - 1
	if i < 0 {
		panic("TextScanner.StackPop() failed")
	}
	p := s.stack[i]
	s.stack = s.stack[0 : len(s.stack)-1]
	if reset {
		s.offset = p.Offset
		s.line = p.Line
		s.column = p.Column
		s.char = s.input[p.Offset]
	}
	// log.Printf("%d '%s' StackPop(%v): %+v\n", s.offset, string(s.char), reset, p)
}
