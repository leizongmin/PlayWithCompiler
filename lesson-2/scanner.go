package main

type Scanner struct {
	input  []rune
	offset int
	char   rune
	eof    bool
	line   int
	column int
	stack  []positionInfo
}

type positionInfo struct {
	offset int
	line   int
	column int
}

func NewScanner(source string) *Scanner {
	s := Scanner{input: []rune(source)}
	if len(s.input) > 0 {
		s.char = s.input[0]
		s.eof = false
	}
	s.stack = make([]positionInfo, 0)
	return &s
}

// 当前字符
func (s *Scanner) Current() rune {
	return s.char
}

func (s *Scanner) safeGetChar(offset int) rune {
	max := len(s.input)
	if offset < max {
		return s.input[offset]
	}
	return 0
}

// 移到下一个字符
func (s *Scanner) Next() rune {
	s.offset++
	s.char = s.safeGetChar(s.offset)
	if s.char == 0 {
		s.eof = true
	}
	return s.char
}

// 偷看下一个字符
func (s *Scanner) Peek() rune {
	return s.safeGetChar(s.offset + 1)
}

// 当前位置
func (s *Scanner) Offset() int {
	return s.offset
}

// 是否到末尾
func (s *Scanner) EOF() bool {
	return s.eof
}

// 将当前位置压栈
func (s *Scanner) StackPush() {
	s.stack = append(s.stack, positionInfo{offset: s.offset, line: s.line, column: s.column})
}

// 从栈中弹出位置
// reset 如果为true则根据栈顶的信息重置当前位置
func (s *Scanner) StackPop(reset bool) {
	i := len(s.stack) - 1
	if i < 0 {
		panic("Scanner.StackPop() failed")
	}
	p := s.stack[i]
	s.stack = s.stack[0 : len(s.stack)-1]
	if reset {
		s.offset = p.offset
		s.line = p.line
		s.column = p.column
		s.char = s.input[p.offset]
	}
}
