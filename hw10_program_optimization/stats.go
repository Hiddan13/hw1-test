package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	return countDomains(r, domain)
}

func countDomains(r io.Reader, domain string) (DomainStat, error) {
	var str strings.Builder
	str.WriteRune('.')
	str.WriteString(domain)

	fileScanner := bufio.NewScanner(r)
	var userrr User
	result := make(DomainStat)
	dm := str.String()
	for fileScanner.Scan() {
		if err := easyjson.Unmarshal(fileScanner.Bytes(), &userrr); err != nil {
			return nil, err
		}
		if strings.Contains(userrr.Email, dm) {
			result[strings.ToLower(strings.SplitN(userrr.Email, "@", 2)[1])]++
		}
	}
	return result, nil
}
