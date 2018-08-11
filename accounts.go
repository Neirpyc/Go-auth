package auth

import (
	"errors"
	"io/ioutil"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
)

//used to store multiple accounts in one place
type Accounts struct {
	Accounts []Account
}

//converts a []byte to an Accounts
func parseAccounts(accs []byte) Accounts {
	a := Accounts{}
	err := yaml.Unmarshal(accs, &a)
	if err != nil {
		log.Fatal(err)
	}
	return a
}

//adds an Account to an Accounts
func (accounts *Accounts) addAccount(account Account) error {
	for _, element := range accounts.Accounts {
		if element.Username == account.Username {
			return errors.New("Username " + account.Username + " already taken!")
		}
	}
	accounts.Accounts = append(accounts.Accounts, account)
	return nil
}

//writes an Accounts to the specified path
func (accounts Accounts) writeToFile(path string) {
	a, err := yaml.Marshal(accounts)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(path, a, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

//Changes the value of an Account in an Accounts
func (accounts *Accounts) setAccount(account Account) error {
	for i, element := range accounts.Accounts {
		if element.Username == account.Username {
			accounts.Accounts[i] = account
			return nil
		}
	}
	return errors.New("There is no account width username " + account.Username + "!")
}

//returns an Accounts with the specified username from an Accounts
func (accounts Accounts) getAccount(username string) (*Account, error) {
	for _, element := range accounts.Accounts {
		if element.Username == username {
			return &element, nil
		}
	}
	return &Account{}, errors.New("There is no account width username " + username + "!")
}

//used to store the data of an user
type Account struct {
	Username string
	Password string
	Salt     string
	Tokens   []Token
}

//returns a valid token for an account.
//if there is none it creates none
//if there is an expired one it destroys it
func (account *Account) getToken(ip string) Token {
	for i, element := range account.Tokens {
		if element.Expire < time.Now().Unix() {
			account.Tokens = append(account.Tokens[:i], account.Tokens[i+1:]...)
			i--
		} else {
			if Settings.CheckIp {
				if element.Ip == ip {
					return element
				}
			} else {
				return element
			}
		}
	}
	return account.addToken(ip)
}

//adds a token with specified ip to an account
func (account *Account) addToken(ip string) Token {
	tok := Token{}
	tok.Token = getRandomString(512)
	tok.Ip = ip
	tok.Expire = time.Now().Unix() + Settings.TokenValidity
	account.Tokens = append(account.Tokens, tok)
	return tok
}

//returns true of false depending on wether the password is valid or not
func (account Account) checkPassword(password string) bool {
	return account.Password == mySha512(password+account.Salt)
}
