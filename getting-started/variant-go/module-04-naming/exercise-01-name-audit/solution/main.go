package main

import (
	"fmt"
	"strings"
)

type Book struct {
	Title     string
	Year      int
	Price     float64
	Available bool
}

func filterAvailable(books []Book) []Book {
	var available []Book
	for _, b := range books {
		if b.Available {
			available = append(available, b)
		}
	}
	return available
}

func averagePrice(books []Book) float64 {
	total := 0.0
	for _, b := range books {
		total += b.Price
	}
	if len(books) == 0 {
		return 0
	}
	return total / float64(len(books))
}

func groupByYear(books []Book) map[int][]Book {
	byYear := make(map[int][]Book)
	for _, b := range books {
		byYear[b.Year] = append(byYear[b.Year], b)
	}
	return byYear
}

func formatBook(book Book) string {
	var parts []string
	parts = append(parts, book.Title)
	parts = append(parts, fmt.Sprintf("(year %d)", book.Year))
	parts = append(parts, fmt.Sprintf("$%.2f", book.Price))
	if book.Available {
		parts = append(parts, "[available]")
	} else {
		parts = append(parts, "[sold]")
	}
	return strings.Join(parts, " ")
}

func main() {
	catalog := []Book{
		{Title: "Go Programming", Year: 2023, Price: 49.99, Available: true},
		{Title: "The Art of SQL", Year: 2021, Price: 39.95, Available: false},
		{Title: "Systems Design", Year: 2024, Price: 54.99, Available: true},
		{Title: "Network Protocols", Year: 2022, Price: 44.50, Available: true},
		{Title: "Data Structures", Year: 2023, Price: 42.00, Available: false},
	}

	fmt.Println("All:")
	for _, book := range catalog {
		fmt.Println(" ", formatBook(book))
	}

	available := filterAvailable(catalog)
	fmt.Printf("\nAvailable: %d\n", len(available))
	for _, book := range available {
		fmt.Println(" ", formatBook(book))
	}

	fmt.Printf("\nAverage price: $%.2f\n", averagePrice(catalog))

	byYear := groupByYear(catalog)
	fmt.Println("\nBy year:")
	for year, books := range byYear {
		fmt.Printf("  %d: %d items\n", year, len(books))
	}
}
