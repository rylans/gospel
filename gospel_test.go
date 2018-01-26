package gospel

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

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
