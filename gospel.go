package gospel

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

type Corrector struct {
  descriptor string
  dictionaryWords map[string]int
}

func (corr Corrector) String() string {
  return corr.descriptor
}

func (corr *Corrector) Contains(str string) bool {
  _, exists := corr.dictionaryWords[str]
  return exists
}


// Corrector builders
func ForEnglish() Corrector {
  dictWords := make(map[string]int)

  for _, w := range readWords() {
    dictWords[strings.ToLower(w)] = 1
  }

  numWords := strconv.Itoa(len(dictWords))
  desc := "English spelling corrector (words loaded: " + numWords + ")"
  return Corrector{descriptor: desc, dictionaryWords: dictWords}

}


func readWords() []string {
  dat, err := ioutil.ReadFile("/usr/share/dict/words")
  if err != nil {
    panic(err)
  }
  return strings.Split(string(dat), "\n")
}

func IsCorrect(str string) (bool, error) {
  if "correct" == str {
    return true, nil
  } else {
    return false, nil
  }
}

func main(){

  x := ForEnglish()
  fmt.Println(x)
  fmt.Println(x.Contains("scolytus"))
  fmt.Println(x.Contains("skolytus"))
}

