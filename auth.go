package auth

import (
	"errors"
	"io/ioutil"
	"log"
)

//Creates a new accounts with the given username and password
func register(username string, password string) error {
	//Logs if necessray
	if Settings.DebugLevel >= 1 {
		log.Println("Registering", username, "width password", password)
	}
	//prepares the account
	account := Account{}
	account.Salt = getRandomString(512)
	account.Password = mySha512(password + account.Salt)
	account.Username = username

	//opens the file in which the account will be stored
	a, err := ioutil.ReadFile(Settings.AccountsFile)
	if err != nil {
		log.Println(err)
		return err
	}

	//converts the accounts file into a go struct
	accounts := parseAccounts(a)

	//adds the account to the accounts strcut
	err = accounts.addAccount(account)
	if err != nil {
		log.Println(err)
		return err
	}
	//writes the accounts to the accounst file
	accounts.writeToFile(Settings.AccountsFile)
	return nil
}

//returns a token for a user with a given username and password and ip if you
//have set CheckIp in the config file to true, otherwise use "" or "0.0.0.0" as
//ip
func login(username string, password string, ip string) (Token, error) {
	if Settings.DebugLevel >= 1 {
		log.Println("Logging", username, "width password", password, "on ip", ip)
	}
	//reads the accounts file
	a, err := ioutil.ReadFile(Settings.AccountsFile)
	if err != nil {
		log.Println(err)
		return Token{}, err
	}

	//converts the accounts file to a golang struct
	accounts := parseAccounts(a)

	//gets the accounts with the specified username
	account, err := accounts.getAccount(username)
	if err != nil {
		log.Println(err)
		return Token{}, err
	}
	//checks for the password to be valid
	if account.checkPassword(password) {
		//gets the token of the user
		tok := account.getToken(ip)

		//if there is no token, the function above creates one, so we have to write
		//the changes that could have been made
		err = accounts.setAccount(*account)
		if err != nil {
			log.Fatal(err)
		}

		accounts.writeToFile(Settings.AccountsFile)
		return tok, nil
	}
	return Token{}, errors.New("Invalid password")
}

//returns true or false depending on wether a token is valid for a user
//and it's ip if you have set CheckIp in the config file to true, otherwise use
//"" or "0.0.0.0" as ip
func validateToken(username string, ip string, token string) (bool, error) {
	if Settings.DebugLevel >= 1 {
		log.Println("Validating", token, "with username", username, "and ip", ip)
	}
	//opens the accounts file
	a, err := ioutil.ReadFile(Settings.AccountsFile)
	if err != nil {
		log.Println(err)
		return false, err
	}

	//converts the accounts file to a golang struct
	accounts := parseAccounts(a)

	//gets the requested account
	account, err := accounts.getAccount(username)
	if err != nil {
		log.Println(err)
		return false, err
	}

	//check wether it's token is valid or not
	if account.getToken(ip).Token == token {
		return true, nil
	} else {
		return false, nil
	}
}
