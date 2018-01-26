package gospel

import "fmt"

func IsCorrect(str string) (bool, error) {
  if "correct" == str {
    return true, nil
  } else {
    return false, nil
  }
}

func main(){
  fmt.Println("hello")

}

