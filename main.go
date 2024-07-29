package main

import (
  "fmt"
  gha "github.com/sethvargo/go-githubactions"
  "regexp"
  "os"
)

func main() {

  // Obtain the github context containing the payload
  context, err := gha.Context()

  if err != nil {
    fmt.Print(err.Error())
  }

  // Event contains the pull_request information
  // along with other details
  pr := context.Event["pull_request"].(map[string]interface {})
  pr_body := pr["body"].(string)


  // An empty checkbox means task is incomplete
  // Therefore the check must fail
  empty_checkbox_pattern := `[-*]\s*\[ \]`

  re, err := regexp.Compile(empty_checkbox_pattern)

  if err != nil {
    fmt.Print(err.Error())
  }

  matches := re.FindAllString(pr_body, -1)


  if len(matches) > 0 {
    fmt.Printf("Check Failed !! Please complete all the tasks\n")
    os.Exit(1)
  }


  username := gha.GetInput("username")
  message := fmt.Sprintf("Congrats %s for completing all the tasks !! Well done :)\n", username)
  gha.SetOutput("status", message) 

}
