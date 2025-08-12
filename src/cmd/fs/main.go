package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"Go_day01-1/nydiamig/internal/parser"
)

// ParsePathsAsRecipe — превращает список путей в Recipe
func ParsePathsAsRecipe(filePath string) (parser.Recipe, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return parser.Recipe{}, err
	}
	defer file.Close()

	var r parser.Recipe
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := scanner.Text()
		r.Cakes = append(r.Cakes, parser.Cake{
			Name: path, // путь будет "именем торта"
		})
	}
	if err := scanner.Err(); err != nil {
		return parser.Recipe{}, err
	}

	return r, nil
}

func main() {
	if len(os.Args) != 5 {
		log.Fatalf("Usage: %s --old snapshot1.txt --new snapshot2.txt", os.Args[0])
	}

	var oldPath, newPath string
	for i := 1; i < len(os.Args); i += 2 {
		switch os.Args[i] {
		case "--old":
			oldPath = os.Args[i+1]
		case "--new":
			newPath = os.Args[i+1]
		default:
			log.Fatalf("Unknown argument: %s", os.Args[i])
		}
	}

	oldRecipe, err := ParsePathsAsRecipe(oldPath)
	if err != nil {
		log.Fatal(err)
	}

	newRecipe, err := ParsePathsAsRecipe(newPath)
	if err != nil {
		log.Fatal(err)
	}

	diff := parser.CompareRecipes(oldRecipe, newRecipe)

	var added, removed []string

	for _, c := range diff.CakeDiffs {
		if c.MissingInSecond {
			removed = append(removed, c.CakeName)
		}
		if c.MissingInFirst {
			added = append(added, c.CakeName)
		}
	}

	sort.Strings(added)
	sort.Strings(removed)

	for _, path := range added {
		fmt.Println("ADDED", path)
	}
	for _, path := range removed {
		fmt.Println("REMOVED", path)
	}
}
