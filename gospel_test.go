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
  corrected := c.Correct("Gxsbyzz")

  assert.Equal(t, "Gxsbyzz", corrected, "should have returned input word")
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

func TestEnglishCorrectorCorrectsCommonWordThe(t *testing.T){
  c := ForEnglish()
  assert.Equal(t, "the", c.Correct("the"))
  assert.Equal(t, "the", c.Correct("thi"))
  assert.Equal(t, "the", c.Correct("teh"))
}

func TestEnglishCorrectorCorrectsCommonWordThis(t *testing.T){
  c := ForEnglish()
  assert.Equal(t, "this", c.Correct("this"))
  assert.Equal(t, "this", c.Correct("thiss"))
}

func TestEnglishCorrectorCorrectsCommonWordWhich(t *testing.T){
  c := ForEnglish()
  assert.Equal(t, "which", c.Correct("which"))
  assert.Equal(t, "which", c.Correct("whch"))
  assert.Equal(t, "which", c.Correct("whih"))
}

func TestWordsCorrectorCorrectsTransposition(t *testing.T){
  c := OfWords([]string{"knight"})

  corrected := c.Correct("tnighk")
  assert.Equal(t, "knight", corrected, "should have corrected to word 'knight'")
}

func TestWordsCorrectorCorrectsComplexEdit(t *testing.T){
  c := OfWords([]string{"zygomycetes"})

  corrected := c.Correct("zgomyteces")
  assert.Equal(t, "zygomycetes", corrected, "should have corrected to word 'zygomycetes'")
}

func TestEnglishCorrectorSplitsTheDifference(t *testing.T){
  c := ForEnglish()

  corrected := c.Correct("thedifference")
  assert.Equal(t, "the difference", corrected, "should have corrected word to 'the difference'")
}

func TestEnglishCorrectorSplitsFreeGift(t *testing.T){
  c := ForEnglish()

  corrected := c.Correct("freegift")
  assert.Equal(t, "free gift", corrected, "should have corrected word to 'free gift'")
}

func TestEnglishCorrectorDoesNotSplitDoorknob(t *testing.T){
  c := ForEnglish()

  cx := c.Correct("doorknob")
  assert.Equal(t, "doorknob", cx, "No change expected for word 'doorknob'")
}

func TestEnglishCorrectorDoesNotSplitButterfly(t *testing.T){
  c := ForEnglish()

  cx := c.Correct("butterfly")
  assert.Equal(t, "butterfly", cx, "No change expected for word 'butterfly'")
}
