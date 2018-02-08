package gospel

// Word list from:
// https://en.wikipedia.org/wiki/Wikipedia:Lists_of_common_misspellings/For_machines

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "strconv"
)

const BENCHMARK_CORRECT_PERCENT = 0.55


func TestBatchCorrections_Ab(t *testing.T){
  c := ForEnglish()

  misspellings := []ExpectedCorrection{
    ExpectedCorrection{wrong: "abandonned", right: "abandoned"},
    ExpectedCorrection{wrong: "aberation", right: "aberration"},
    ExpectedCorrection{wrong: "abilityes", right: "abilities"},
    ExpectedCorrection{wrong: "abilties", right: "abilities"},
    ExpectedCorrection{wrong: "abilty", right: "ability"},
    ExpectedCorrection{wrong: "abondon", right: "abandon"},
    ExpectedCorrection{wrong: "abbout", right: "about"}}

  assertCorrectBenchmark(t, batchCorrectPercent(&c, misspellings))
}

func TestBatchCorrections_Em(t *testing.T){
  c := ForEnglish()

  misspellings := []ExpectedCorrection{
    ExpectedCorrection{wrong: "embarass", right: "embarrass"},
    ExpectedCorrection{wrong: "embarassed", right: "embarrassed"},
    ExpectedCorrection{wrong: "emblamatic", right: "emblematic"},
    ExpectedCorrection{wrong: "eminate", right: "emanate"},
    ExpectedCorrection{wrong: "emision", right: "emission"},
    ExpectedCorrection{wrong: "emited", right: "emitted"},
    ExpectedCorrection{wrong: "emmediately", right: "immediately"}}

  assertCorrectBenchmark(t, batchCorrectPercent(&c, misspellings))
}

func TestBatchCorrections_Hi(t *testing.T){
  c := ForEnglish()

  misspellings := []ExpectedCorrection{
    ExpectedCorrection{wrong:"hieght", right: "height"},
    ExpectedCorrection{wrong:"hierachical", right: "hierarchical"},
    ExpectedCorrection{wrong:"hieroglph", right: "hieroglyph"},
    ExpectedCorrection{wrong:"higer", right: "higher"},
    ExpectedCorrection{wrong:"higway", right: "highway"},
    ExpectedCorrection{wrong:"himselv", right: "himself"},
    ExpectedCorrection{wrong:"hitsingles", right: "hit singles"}}

  assertCorrectBenchmark(t, batchCorrectPercent(&c, misspellings))
}

func batchCorrectPercent(c *Corrector, misspellings []ExpectedCorrection) float64 {
  rights := 0.0

  for _, ms := range misspellings {
    corrected := c.Correct(ms.wrong)
    if corrected == ms.right {
      rights++
    }
  }

  return float64(rights)/float64(len(misspellings))
}

func assertCorrectBenchmark(t *testing.T, pct float64) {
  assert.True(t, pct > BENCHMARK_CORRECT_PERCENT, "Correct percent was: " + strconv.FormatFloat(pct, 'f', -1, 64))
}

type ExpectedCorrection struct {
  wrong, right string
}
