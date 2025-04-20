package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testAmount int
	fmt.Fscan(in, &testAmount)

	for test := 0; test < testAmount; test++ {
		var n int
		fmt.Fscan(in, &n)
		in.ReadString('\n')

		resultMap := make(map[string]int)
		var sound string

		for i := 0; i < n; i++ {
			speaker, sentence := parseInput(in)
            _, ok := resultMap[speaker]
            if !ok {
                resultMap[speaker] = 0
            }

			if sentence[0] == "I" {
				if sentence[1] == "am" {
                    if sentence[2] == "not"{
                        x := sentence[3]
                        sound = x
                        resultMap[speaker] -= 1
                    } else {
                        x := sentence[2]
                        sound = x
                        resultMap[speaker] += 2
                    } 
                } else if sentence[1] == "is" {
                        if sentence[2] == "not"{
                            x := sentence[3]
                            sound = x
                            resultMap["I"] -= 1
                        } else {
                            x := sentence[2]
                            sound = x
                            resultMap["I"] += 1
                        }
                    }
			} else {
				target := sentence[0]
				if sentence[1] == "is"{
                    if sentence[2] == "not"{
                        x := sentence[3]
                        sound = x
                        resultMap[target] -= 1
                    } else {
                        x := sentence[2]
                        sound = x
                        resultMap[target] += 1
                    }
                }
			}
		}

		maxResult := -2147483648
		for _, value := range resultMap {
			if value > maxResult {
				maxResult = value
			}
		}

		var speaker []string
		for name, result := range resultMap {
			if result == maxResult {
				speaker = append(speaker, name)
			}
		}

		sort.Strings(speaker)
		for _, name := range speaker {
			fmt.Fprintf(out, "%s is %s.\n", name, sound)
		}
	}
}

func parseInput(in *bufio.Reader) (string, []string){
    line, _ := in.ReadString('\n')
    line = strings.TrimSpace(line)

    parts := strings.SplitN(line, ":", 2)
    speaker := strings.TrimSpace(parts[0])
    statement := strings.TrimSpace(parts[1])

    sentence := strings.Fields(statement)
    // sentence = sentence[:len(sentence)-1]
    last := len(sentence) - 1
    sentence[last] = strings.TrimSuffix(sentence[last], "!")
    return speaker, sentence
}
