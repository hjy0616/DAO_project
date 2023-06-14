package main

import "fmt"

type Token struct {
	balances map[string]int
	total    int
	allowed  map[string]map[string]int
}

func NewToken() *Token {
	return &Token{
		balances: make(map[string]int),
		allowed:  make(map[string]map[string]int),
	}
}

func (t *Token) BalanceOf(address string) int {
	return t.balances[address]
}

func (t *Token) TotalSupply() int {
	return t.total
}

func (t *Token) Transfer(from, to string, tokens int) bool {
	if t.balances[from] < tokens {
		return false
	}

	t.balances[from] -= tokens
	t.balances[to] += tokens

	return true
}

func (t *Token) Approve(owner, spender string, tokens int) bool {
	if t.balances[owner] < tokens {
		return false
	}

	if t.allowed[owner] == nil {
		t.allowed[owner] = make(map[string]int)
	}

	t.allowed[owner][spender] = tokens
	return true
}

func (t *Token) TransferFrom(from, to, spender string, tokens int) bool {
	if t.allowed[from][spender] < tokens {
		return false
	}

	t.balances[from] -= tokens
	t.balances[to] += tokens
	t.allowed[from][spender] -= tokens

	return true
}

func (t *Token) Allowance(owner, spender string) int {
	return t.allowed[owner][spender]
}

func (t *Token) Mint(to string, tokens int) {
	t.total += tokens
	t.balances[to] += tokens
}

func main() {
	token := NewToken()

	token.Mint("0x0", 1000)
	fmt.Println(token.TotalSupply()) // Output: 1000

	token.Transfer("0x0", "0x1", 500)
	fmt.Println(token.BalanceOf("0x1")) // Output: 500
	fmt.Println(token.BalanceOf("0x0")) // Output: 500

	token.Approve("0x1", "0x2", 200)
	fmt.Println(token.Allowance("0x1", "0x2")) // Output: 200

	token.TransferFrom("0x1", "0x0", "0x2", 200)
	fmt.Println(token.BalanceOf("0x1")) // Output: 300
	fmt.Println(token.BalanceOf("0x0")) // Output: 700
}
