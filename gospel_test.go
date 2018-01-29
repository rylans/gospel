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
  corrected := c.Correct("Gospel")

  assert.Equal(t, "Gospel", corrected, "should have corrected 'gospel'")
}

func TestEnglishCorrectorFailsToCorrectWord(t *testing.T){
  c := ForEnglish()
  corrected := c.Correct("Gxsby")

  assert.Equal(t, "Gxsby", corrected, "should have returned input word")
}

func TestWordsCorrectorCorrectsDeletionInitial(t *testing.T){
  c := OfWords([]string{"knight"})

  corrected := c.Correct("zknight")
  assert.Equal(t, "knight", corrected, "should have corrected to word 'knight'")
}

func TestWordsCorrectorCorrectsDeletionFinal(t *testing.T){
  c := OfWords([]string{"knight"})

  corrected := c.Correct("knightx")
  assert.Equal(t, "knight", corrected, "should have corrected to word 'knight'")
}

func TestWordsCorrectorCorrectsDeletionMedial(t *testing.T){
  c := OfWords([]string{"knight"})

  corrected := c.Correct("knieght")
  assert.Equal(t, "knight", corrected, "should have corrected to word 'knight'")
}

func TestWordsCorrectorCorrectsInsertionInitial(t *testing.T){
  c := OfWords([]string{"knight"})

  corrected := c.Correct("night")
  assert.Equal(t, "knight", corrected, "should have corrected to word 'knight'")
}

func TestWordsCorrectorCorrectsInsertionFinal(t *testing.T){
  c := OfWords([]string{"knight"})

  corrected := c.Correct("knigh")
  assert.Equal(t, "knight", corrected, "should have corrected to word 'knight'")
}

func TestWordsCorrectorCorrectsInsertionMedial(t *testing.T){
  c := OfWords([]string{"knight"})

  corrected := c.Correct("kniht")
  assert.Equal(t, "knight", corrected, "should have corrected to word 'knight'")
}

func TestWordsCorrectorCorrectsInsertionAndDeletion(t *testing.T){
  c := OfWords([]string{"knight"})

  corrected := c.Correct("knihtu")
  assert.Equal(t, "knight", corrected, "should have corrected to word 'knight'")
}

func TestWordsCorrectorCorrectsAcheive(t *testing.T){
  c := OfWords([]string{"achieve"})

  corrected := c.Correct("acheive")
  assert.Equal(t, "achieve", corrected, "should have corrected to word 'achieve'")
}

func TestEnglishCorrectorCorrectsCommonWordThis(t *testing.T){
  c := ForEnglish()
  assert.Equal(t, "this", c.Correct("this"))
  assert.Equal(t, "this", c.Correct("thi"))
  assert.Equal(t, "this", c.Correct("thiss"))
}

func TestEnglishCorrectorCorrectsCommonWordWhich(t *testing.T){
  c := ForEnglish()
  assert.Equal(t, "which", c.Correct("which"))
  assert.Equal(t, "which", c.Correct("whch"))
  assert.Equal(t, "which", c.Correct("whih"))
}
