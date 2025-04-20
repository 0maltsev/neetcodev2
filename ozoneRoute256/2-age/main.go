package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Relation struct {
	person1, person2 string
	diff          int
}

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var testCases int
	fmt.Fscanln(input, &testCases)

	targetName := regexp.MustCompile(`How old is (\w+)\?`)

	for test := 0; test < testCases; test++ {
		query, _ := input.ReadString('\n')
		query = strings.TrimSpace(query)

		match := targetName.FindStringSubmatch(query)
		targetName := match[1]

		ages := make(map[string]int)
		relations := make([]Relation, 0)

		for i := 0; i < 3; i++ {
			line, _ := input.ReadString('\n')
			line = strings.TrimSpace(line)
			sentence := strings.Split(line, " ")

			name1 := sentence[0]

			if sentence[2] == "the" && sentence[3] == "same" {
				name2 := sentence[6]
				relations = append(relations, Relation{name1, name2, 0})
			} else if sentence[3] == "years" && sentence[4] == "older" {
				ageDifference, _ := strconv.Atoi(sentence[2])
				name2 := sentence[6]
				relations = append(relations, Relation{name1, name2, ageDifference})
			} else if sentence[3] == "years" && sentence[4] == "younger" {
				ageDifference, _ := strconv.Atoi(sentence[2])
				name2 := sentence[6]
				relations = append(relations, Relation{name1, name2, -ageDifference})
			} else {
				age, _ := strconv.Atoi(sentence[2])
				ages[name1] = age
			}
		}

		updated := true
		for updated {
			updated = false
			for _, relation := range relations {
				person1, person2, ageDiff := relation.person1, relation.person2, relation.diff
				if ageOfPerson2, exists := ages[person2]; exists {
					if _, found := ages[person1]; !found {
						ages[person1] = ageOfPerson2 + ageDiff
						updated = true
					}
				} else if ageOfPerson1, exists := ages[person1]; exists {
					if _, found := ages[person2]; !found {
						ages[person2] = ageOfPerson1 - ageDiff
						updated = true
					}
				}
			}
		}

		fmt.Fprintln(output, ages[targetName])
	}
}
