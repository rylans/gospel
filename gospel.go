package gospel

import (
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

func (corr *Corrector) ValueOf(str string) int {
  return corr.dictionaryWords[strings.ToLower(str)]
}

func (corr *Corrector) Correct(str string) string {
  if corr.Contains(str) { 
    return str
  } 

  candidate := corr.maxCandidate(edits(str))
  if candidate != "" {
    return candidate
  }

  candidate = corr.maxCandidate(editsTwo(str))
  if candidate == "" {
    return str
  } 
  return candidate
}

func (corr *Corrector) maxCandidate(candidateWords []string) string {
  maxVal := 0
  maxWord := ""
  for _, candidateWord := range candidateWords {
    if (corr.Contains(candidateWord)){
      candidateValue := corr.ValueOf(candidateWord)
      if candidateValue > maxVal {
	maxVal = candidateValue
	maxWord = candidateWord
      }
    }
  }
  return maxWord
}

// Corrector builders
func ForEnglish() Corrector {
  dictWords := makeDictOfWords(readWords())
  loadFrequentWords(dictWords)
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

func loadFrequentWords(dictWords map[string]int){
  loadWordsWithZipfDistribution(56271,
    []string{"the", "of", "and", "to", "in", "i",
	     "that", "was", "his", "he", "it", "with",
	     "is", "for", "as", "had", "you", "not",
	     "be", "her", "on", "at", "by", "which",
	     "have", "or", "from", "this", "him", "but",
	     "all", "she", "they", "were", "my", "are"},
    dictWords)
}

func loadWordsWithZipfDistribution(frequencyOfFirst int, words []string, dictWords map[string]int){
  for ix, word := range words {
    dictWords[word] = (frequencyOfFirst / (1 + ix))
  }

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
  letters := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
  splits := []string{}

  // deletions
  for i, _ := range word {
    splits =  append(splits, word[:i] + word[i+1:])
  }

  // inital + medial insertions
  for i, _ := range word {
    for j, _ := range letters {
      splits = append(splits, word[:i] + letters[j] + word[i:])
    }
  }
  
  // final insertions
  for j, _ := range letters {
    splits = append(splits, word + letters[j])
  }

  return splits
}

func editsTwo(word string) []string {
  allEdits := edits(word)
  edits1 := edits(word)

  for _, v := range edits1 {
    for _, k := range edits(v){
      allEdits = append(allEdits, k)
    }
  }
  return allEdits

}
