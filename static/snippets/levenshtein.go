package main
import ("fmt"; "io/ioutil"; "log"; "os"; "strconv"; "strings"; "time")

func min(a, b int) int { if a <= b { return a }; return b }
func max(a, b int) int { if a >= b { return a }; return b }
func failOn(err error) { if err != nil { log.Fatal(err) } }

func distance(left, right string) int {
  result := max(len(left), len(right))
  rightMap := mapLetters(right)
  for x := 0; x < len(left); x++ {
    letter := string(left[x])
    if rightMap[letter] != nil {
      for _, y := range rightMap[letter] {
        beforeMatch := max(x, y)
        if beforeMatch < result {
          xOut := x
          yOut := y
          for xOut < len(left) && yOut < len(right) &&
              left[xOut] == right[yOut] {
            xOut++
            yOut++
          }
          afterMatch := distance(left[xOut:], right[yOut:])
          result = min(result, beforeMatch + afterMatch)
        }
      }
    }
  }
  return result
}

func mapLetters(text string) map[string][]int {
  result := make(map[string][]int)
  for offset := 0; offset < len(text); offset++ {
    letter := string(text[offset])
    if result[letter] == nil {
      result[letter] = make([]int, 0, len(text))
    }
    result[letter] = append(result[letter], offset)
  }
  return result
}

// command-line interface

func main() {
  if len(os.Args) < 3 {
    fmt.Println("USAGE: go run levenshtein.go /path/to/dict numberOfWords")
    os.Exit(1)
  }

  source := "/usr/share/dict/american-english"
  text,  err := ioutil.ReadFile(source);  failOn(err)
  limit, err := strconv.Atoi(os.Args[1]); failOn(err)
  words := strings.Split(string(text), "\n")
  count := 0
  start := time.Now()
  
  for i, left := range words[:limit - 1] {
    for _, right := range words[i + 1:limit] {
      fmt.Printf("%s vs. %s: %d\n", left, right, distance(left, right))
      count++
    }
  }
  
  fmt.Printf("Performed %d calculations in %d ms\n", count, time.Now().Sub(start) / 1e6)
}
