# Authentication system with golang
## Introduction
The auth package allows you to easily create and store usernames and password without a database. It uses tokens so the password hasn't to be stored on the client side and uses hash with salt so it remains safe even with leaks.

## Installation
To install go-auth, run:
```
go get gopkg.in/yaml.v2
go get github.com/Neirpyc/Go-auth
```
This package requires [go-yaml](https://github.com/go-yaml/yaml) so it can read the config file and store the accounts a a yaml file.
## Setup
Once you're in your project file, just create these files (you can rename them as you want, but you'll have to modify some stuff later):
 - logs.log
 - settings.yml
 - accounts.yml
##
In *settings.yml* you have to put this (modify the paths if you changed the files names):
```yaml
randomlevel: 1
debuglevel: 1
checkip: false
accountsfile: "accounts.yml"
logsfile: "logs.log"
tokenvalidity: 604800
```




### Configuration
 - **randomlevel** should be *0* or *1*.
   - *0* means using [math.rand](https://golang.org/pkg/math/rand/) for random numbers generation, which is unsafe!
   - *1* means using [crypto.rand](https://golang.org/pkg/crypto/rand/) which is cryptographically secure.  
 - **debuglevel** should be *0* or *1*
   - *0* means only log when there is a fatal error
   - *1* means log when a main function is called or when there is a fatal error
- **checkip**: should be *true* or *false*
   - *true* means a token will be invalid when used with an incorrect ip
   - *false* means a token will be valid whatever ip is used
 - **accountsfile**: should be the path to the file in which the accounts are stored. Make sure this file exists and is readable.
 - **logsfile**: should be the path to the file in which the logs are stored. Make sure this file exists and is readable
  - **tokenvalidity**: should be the time in second during which a token is valid

## License
The auth package is licensed under the GNU GPL license. Please see the LICENSE file for details.

## Example
- First you have to import the settings file:
  ```go
  package main

  import goauth "github.com/Neirpyc/Go-auth"

  func main(){
	  goauth.SetSettingsFromFile("settings.yml")
  }
  ```
- Then you can register your first user:
  ```go
  err := goauth.Register("Username", "Password") //this is not a good
                            //practise (read the Good Practices part to learn more)
  if err != nil{
	  panic(err)
	}
  ```
- One this is done you can login  this user:
  ```go
  token, err := goauth.Login("Username", "Password", "0.0.0.0")
  if err != nil{
	  panic(err)
	}
  fmt.Println(token.Token) //prints the 512 base64 characters of the token
  ```
 - When the user is already logged in, and you want to verify it's him, just ask him his token, and do the following:
   ```go
   valid, err := goauth.ValidateToken("Username", "0.0.0.0", token)
   if err != nil{
	   panic(err)
	}
   if valid{
	   fmt.Println("Valid!")
	else{
	  fmt.Println("Unvalid :(")
	}
	```

## Good practices
- Never store the password of the user
- Never send the password from the client to the server and prefer to get it hashed 500-5000 times before it reaches your server with a secure algorithm such as sha512
- Use the token system (do not use Login() every time!)
