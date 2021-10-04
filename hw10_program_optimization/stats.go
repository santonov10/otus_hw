package hw10programoptimization

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/mailru/easyjson/jlexer"
)

//easyjson:json
type User struct {
	// ID       int
	// Name     string
	// Username string
	Email string
	// Phone    string
	// Password string
	// Address  string
}

func (u *User) getEmailDomain() string {
	return strings.ToLower(after(u.Email, "@"))
}

func after(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:]
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	suf := "." + domain
	bufReader := bufio.NewReader(r)
	var user User
	for {
		// line, err := bufReader.ReadString('\n') // не проходит по памяти...
		line, _, err := bufReader.ReadLine() // наверное, нужно больше логики с получением второго параметра и его проверкой
		lexer := jlexer.Lexer{
			Data:              line,
			UseMultipleErrors: false,
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return result, err
		}

		user.UnmarshalEasyJSON(&lexer)
		if lexer.Error() != nil {
			return result, lexer.Error()
		}

		if strings.HasSuffix(user.Email, suf) {
			result[user.getEmailDomain()]++
		}
	}
	return result, nil
}
