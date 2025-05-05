package main

import (
	"Go_day01-1/nydiamig/internal/parser"
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
	filePath := flag.String("f", "", "Path to the input file (JSON or XML)")
	formatFlag := flag.String("format", "", "Optional format override (json or xml)")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("Version: %s\nBuild Time: %s\n", version, buildTime)
		return
	}
	if *filePath == "" {
		log.Fatal("Please provide a file path using -f flag")
	}

	if err := processAndDisplayDatabase(*filePath, *formatFlag); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// processAndDisplayDatabase читает, парсит и красиво выводит базу
func processAndDisplayDatabase(filePath, formatOverride string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	format := strings.ToLower(formatOverride)
	if format == "" {
		format = parser.DetectFormatFromExtension(filePath)
	}
	if !parser.IsSupportedFormat(format) {
		return fmt.Errorf("unsupported or unknown file format: %s", format)
	}

	recipe, err := parser.ParseData(format, data)
	if err != nil {
		return fmt.Errorf("failed to parse %s file: %w", format, err)
	}

	// Печатаем результат
	fmt.Println("Parsed Data:")
	for _, cake := range recipe.Cakes {
		fmt.Printf("Cake: %s, Time: %s\n", cake.Name, cake.Time)
		for _, ingredient := range cake.Ingredients {
			if ingredient.Unit != "" {
				fmt.Printf("  Ingredient: %s, Count: %s, Unit: %s\n", ingredient.Name, ingredient.Count, ingredient.Unit)
			} else {
				fmt.Printf("  Ingredient: %s, Count: %s (unit not specified)\n", ingredient.Name, ingredient.Count)
			}
		}
	}

	return nil
}
