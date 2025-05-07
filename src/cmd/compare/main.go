package main

import (
	"Go_day01-1/nydiamig/internal/parser"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var version = "dev"
var buildTime = "unknown"

func main() {
	versionFlag := flag.Bool("version", false, "Print version information and exit")
	oldFile := flag.String("old", "", "Path to the old database file (JSON or XML)")
	newFile := flag.String("new", "", "Path to the new database file (JSON or XML)")
	formatFlag := flag.String("format", "", "Optional format override (json or xml)")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("Version: %s\nBuild Time: %s\n", version, buildTime)
		return
	}
	if *oldFile == "" || *newFile == "" {
		log.Fatal("Please provide both --old and --new file paths")
	}

	oldRecipe, err := loadRecipe(*oldFile, *formatFlag)
	if err != nil {
		log.Fatalf("Error loading old database: %v", err)
	}

	newRecipe, err := loadRecipe(*newFile, *formatFlag)
	if err != nil {
		log.Fatalf("Error loading new database: %v", err)
	}

	diff := parser.CompareRecipes(oldRecipe, newRecipe)

	output, err := json.MarshalIndent(diff, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal diff to JSON: %v", err)
	}

	fmt.Println(string(output))
	PrintHumanReadableDiff(diff)
}

func loadRecipe(filePath, formatOverride string) (parser.Recipe, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return parser.Recipe{}, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	format := strings.ToLower(formatOverride)
	if format == "" {
		format = parser.DetectFormatFromExtension(filePath)
	}
	if !parser.IsSupportedFormat(format) {
		return parser.Recipe{}, fmt.Errorf("unsupported or unknown file format: %s", format)
	}

	recipe, err := parser.ParseData(format, data)
	if err != nil {
		return parser.Recipe{}, fmt.Errorf("failed to parse %s file: %w", format, err)
	}

	return recipe, nil
}

func PrintHumanReadableDiff(diff parser.DiffResult) {
	for _, cakeDiff := range diff.CakeDiffs {
		switch {
		case cakeDiff.MissingInFirst:
			fmt.Printf("ADDED cake \"%s\"\n", cakeDiff.CakeName)
		case cakeDiff.MissingInSecond:
			fmt.Printf("REMOVED cake \"%s\"\n", cakeDiff.CakeName)
		default:
			if cakeDiff.TimeDifference != "" {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - %s\n", cakeDiff.CakeName, cakeDiff.TimeDifference)
			}
			for _, ingredientDiff := range cakeDiff.IngredientDiffs {
				switch {
				case ingredientDiff.MissingInFirst:
					fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", ingredientDiff.Name, cakeDiff.CakeName)
				case ingredientDiff.MissingInSecond:
					fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", ingredientDiff.Name, cakeDiff.CakeName)
				default:
					if ingredientDiff.UnitDifference != "" {
						fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - %s\n", ingredientDiff.Name, cakeDiff.CakeName, ingredientDiff.UnitDifference)
					}
					if ingredientDiff.CountDifference != "" {
						fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - %s\n", ingredientDiff.Name, cakeDiff.CakeName, ingredientDiff.CountDifference)
					}
				}
			}
		}
	}
}
