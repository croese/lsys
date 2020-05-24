package lsys

import "testing"

func TestDetGen(t *testing.T) {
  tests := []struct {
    axiom string
    rules map[string]string
    cases []struct {
      n int
      expected string
    }
  }{
    {
      "A",
      map[string]string{
        "A": "AB",
        "B": "A",
      },
      []struct {
        n        int
        expected string
      }{
      {0, "A"},
      {3, "ABAAB"},
      {7, "ABAABABAABAABABAABABAABAABABAABAAB"},
      },
    },
    {
      "0",
      map[string]string{
        "1": "11",
        "0": "1[0]0",
      },
      []struct {
        n        int
        expected string
      }{
        {0, "0"},
        {1, "1[0]0"},
        {2, "11[1[0]0]1[0]0"},
        {3, "1111[11[1[0]0]1[0]0]11[1[0]0]1[0]0"},
      },
    },
  }

  for i, tt := range tests {
    g := NewDeterministic(tt.axiom, tt.rules)
    for j, ex := range tt.cases {
      actual := g.Generate(ex.n)

      if actual != ex.expected {
        t.Fatalf("tests[%d, %d] - incorrect output. expected=%q. got=%q",
          i, j, ex.expected, actual)
      }
    }
  }
}
