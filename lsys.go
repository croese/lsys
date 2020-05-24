package lsys

import "bytes"

type DetGen struct {
  axiom string
  rules map[string]string
}

func NewDeterministic(axiom string, rules map[string]string) *DetGen {
  return &DetGen{
    axiom: axiom,
    rules: rules,
  }
}

func (g *DetGen) Generate(n int) string {
  current := g.axiom
  var buf bytes.Buffer

  for i := 0; i < n; i++ {
    for _, c := range current {
      if v, ok := g.rules[string(c)]; ok {
        buf.WriteString(v)
      } else {
        buf.WriteRune(c)
      }
    }
    current = buf.String()
    buf.Reset()
  }

  return current
}