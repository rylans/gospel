package gospel

import (
  "fmt"
  "errors"
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
  _, exists := corr.dictionaryWords[strings.ToLower(str)]
  return exists
}

func (corr *Corrector) Correct(str string) (string, error) {
  if corr.Contains(str) { 
    return str, nil
  } 

  for _, word := range edits(str) {
    if corr.Contains(word) {
      return word, nil
    }
  }

  return "", errors.New("Unable to correct '" + str + "'")
}

// Corrector builders
func ForEnglish() Corrector {
  dictWords := makeDictOfWords(readWords())
  numWords := strconv.Itoa(len(dictWords))
  desc := "English spelling corrector (words loaded: " + numWords + ")"
  return Corrector{descriptor: desc, dictionaryWords: dictWords}
}

func OfWords(words []string) Corrector {
  dictWords := makeDictOfWords(words)
  numWords := strconv.Itoa(len(dictWords))
  desc := "Spelling corrector (words loaded: " + numWords + ")"
  return Corrector{descriptor: desc, dictionaryWords: dictWords}
}

func makeDictOfWords(words []string) map[string]int {
  dictWords := make(map[string]int)
  for _, w := range words {
    dictWords[strings.ToLower(w)] = 1
  }
  return dictWords
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

func edits(word string) []string {
  splits := []string{}

  // deletions
  for i, _ := range word {
    splits =  append(splits, word[:i] + word[i+1:])
  }

  return splits
}


func main(){
  x := ForEnglish()
  fmt.Println(x)
  fmt.Println(x.Contains("scolytus"))
  fmt.Println(x.Contains("skolytus"))

  fmt.Println(edits("knight"))
}

