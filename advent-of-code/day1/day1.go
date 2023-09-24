package day1

import (
	"fmt"
)

func max(a, b int) int {
  if a > b {
    return a;
  }
  return b;
}

func Day1() int {
  var max_calaries int = 0;
  var current_calaries int = 0;
  for {
    var input int;
    n, err := fmt.Scanln(&input);
    if err != nil {
      break;
    }
    if n == 0 {
      max_calaries = max(max_calaries, current_calaries);
      current_calaries = 0;
    }
  }
  max_calaries = max(max_calaries, current_calaries);

  return max_calaries;
}
