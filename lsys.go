package lsys

import (
  "bytes"
  "math/rand"
)

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

type Choice struct {
  probability float64
  successor string
}

type StochasticGen struct {
  axiom string
  rules map[string][]Choice
  randomFunc func()float64
}

func NewStochastic(axiom string, rules map[string][]Choice) *StochasticGen {
  return &StochasticGen{
    axiom: axiom,
    rules: rules,
    randomFunc: func() float64 {
      return rand.Float64()
    },
  }
}

func NewStochasticFunc(axiom string,
  rules map[string][]Choice,
  randomFunc func()float64) *StochasticGen {
  return &StochasticGen{
    axiom: axiom,
    rules: rules,
    randomFunc: randomFunc,
  }
}

func (g *StochasticGen) Generate(n int) string {
  current := g.axiom
  //var buf bytes.Buffer

  //for i := 0; i < n; i++ {
  // for _, c := range current {
  //   if choices, ok := g.rules[string(c)]; ok {
  //
  //     buf.WriteString(v)
  //   } else {
  //     buf.WriteRune(c)
  //   }
  // }
  // current = buf.String()
  // buf.Reset()
  //}

  return current
}