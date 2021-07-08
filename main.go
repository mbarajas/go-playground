//brian.retterer@okta.com
package main

import "fmt"
import "net/http"
import "time"
import "encoding/json"

type User struct {
  ID              string      `json:"id"`
  Status          string      `json:"status"`
  Created         time.Time   `json:"created"`
  Activated       interface{} `json:"activated"`
  StatusChanged   time.Time   `json:"statusChanged"`
  LastLogin       time.Time   `json:"lastLogin"`
  LastUpdated     time.Time   `json:"lastUpdated"`
  PasswordChanged time.Time   `json:"passwordChanged"`
  Type            struct {
    ID string `json:"id"`
  } `json:"type"`
  Profile struct {
    FirstName   string      `json:"firstName"`
    LastName    string      `json:"lastName"`
    MobilePhone interface{} `json:"mobilePhone"`
    SecondEmail interface{} `json:"secondEmail"`
    Login       string      `json:"login"`
    Email       string      `json:"email"`
  } `json:"profile"`
  Credentials struct {
    Password struct {
    } `json:"password"`
    Emails []struct {
      Value  string `json:"value"`
      Status string `json:"status"`
      Type   string `json:"type"`
    } `json:"emails"`
    RecoveryQuestion struct {
      Question string `json:"question"`
    } `json:"recovery_question"`
    Provider struct {
      Type string `json:"type"`
      Name string `json:"name"`
    } `json:"provider"`
  } `json:"credentials"`
  Links struct {
    Self struct {
      Href string `json:"href"`
    } `json:"self"`
  } `json:"_links"`
}

func main() {
  token := "00eIjHI7TXZOqDW0bB1AYsaREHPh4LnO_RFkMnsLT2"
  url := "https://php.oktapreview.com/api/v1/users/00uaz81i7cHX3cSsg0h7"
  user := &User{}

  getUser(token, url, user)
  firstName, lastName := extractName(user)

  fmt.Println(firstName, lastName)

}

func getUser(token, url string, user *User) error {
  client := &http.Client{}
  request, _ := http.NewRequest("GET", url, nil)
  request.Header.Add("Authorization", "SSWS " + token)

  response, err := client.Do(request)
  checkError(err)
  defer response.Body.Close()

  return json.NewDecoder(response.Body).Decode(&user)
}

func extractName(user *User) (firstName, lastName string){
  firstName = user.Profile.FirstName
  lastName = user.Profile.LastName

  return firstName, lastName

}

func checkError(err error){
  if err != nil {
    panic(err)
  }
}
