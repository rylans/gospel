package gospel

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestWordsCorrectorContainsWord(t *testing.T){
  c := OfWords([]string{"hello", "there"})
  assert.Equal(t, true, c.Contains("There"), "should contain word 'there'")
}

func TestWordsCorrectorNotContainsWord(t *testing.T){
  c := OfWords([]string{"hello", "there"})
  assert.Equal(t, false, c.Contains("seven"), "should not contain word 'seven'")
}

func TestEnglishCorrectorNotContainsWord(t *testing.T){
  c := ForEnglish()
  assert.Equal(t, false, c.Contains("FOOBAR"), "should not contain non-existant word")
}

func TestEnglishCorrectorContainsWord(t *testing.T){
  c := ForEnglish()
  assert.Equal(t, true, c.Contains("Gospel"), "should contain word 'gospel'")
}

func TestEnglishCorrectorCorrectsWord(t *testing.T){
  c := ForEnglish()
  corrected, _ := c.Correct("Gospel")

  assert.Equal(t, "Gospel", corrected, "should have corrected 'gospel'")
}

func TestEnglishCorrectorFailsToCorrectWord(t *testing.T){
  c := ForEnglish()
  corrected, err := c.Correct("Gxsby")

  assert.Equal(t, "", corrected, "should have returned empty string")
  assert.NotNil(t, err, "should have returned an error")
}

func TestWordsCorrectorCorrectsDeletionInitial(t *testing.T){
  c := OfWords([]string{"knight"})

  corrected, _ := c.Correct("zknight")
  assert.Equal(t, "knight", corrected, "should have corrected to word 'knight'")
}

func TestWordsCorrectorCorrectsDeletionFinal(t *testing.T){
  c := OfWords([]string{"knight"})

  corrected, _ := c.Correct("knightx")
  assert.Equal(t, "knight", corrected, "should have corrected to word 'knight'")
}

func TestWordsCorrectorCorrectsDeletionMedial(t *testing.T){
  c := OfWords([]string{"knight"})

  corrected, _ := c.Correct("knieght")
  assert.Equal(t, "knight", corrected, "should have corrected to word 'knight'")
}

