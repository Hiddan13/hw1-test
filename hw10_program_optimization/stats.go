package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
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

func main() {
	data := `{"Id":1,"Name":"Howard Mendoza","Username":"0Oliver","Email":"aliquid_qui_ea@Browsedrive.gov","Phone":"6-866-899-36-79","Password":"InAQJvsq","Address":"Blackbird Place 25"}
{"Id":2,"Name":"Jesse Vasquez","Username":"qRichardson","Email":"mLynch@broWsecat.com","Phone":"9-373-949-64-00","Password":"SiZLeNSGn","Address":"Fulton Hill 80"}
{"Id":3,"Name":"Clarence Olson","Username":"RachelAdams","Email":"RoseSmith@Browsecat.com","Phone":"988-48-97","Password":"71kuz3gA5w","Address":"Monterey Park 39"}
{"Id":4,"Name":"Gregory Reid","Username":"tButler","Email":"5Moore@Teklist.net","Phone":"520-04-16","Password":"r639qLNu","Address":"Sunfield Park 20"}
{"Id":5,"Name":"Janice Rose","Username":"KeithHart","Email":"nulla@Linktype.com","Phone":"146-91-01","Password":"acSBF5","Address":"Russell Trail 61"}`
	result, err := GetDomainStat(bytes.NewBufferString(data), "com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [10]User

func getUsers(r io.Reader) (result users, err error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		go func(ii int, lline string) {
			var user User
			if err = json.Unmarshal([]byte(lline), &user); err != nil {
				return
			}
			result[ii] = user
			fmt.Println(user.Email)
		}(i, line)
	}
	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		//matched, err := regexp.Match("."+domain, []byte(user.Email))
		matched := strings.Contains("."+domain, user.Email)
		fmt.Println(domain, user.Email)
		fmt.Println(len(u))

		fmt.Println(matched)
		// if err != nil {
		// 	return nil, err
		// }

		// if matched {
		// 	//num [i] = user
		// 	num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
		// 	num++
		// 	result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		// }
	}
	return result, nil
}
