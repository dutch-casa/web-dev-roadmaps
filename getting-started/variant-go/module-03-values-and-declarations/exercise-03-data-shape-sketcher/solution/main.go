package main

import "fmt"

// Scenario 1: Playlist

type Song struct {
	Title    string
	Artist   string
	Duration int // seconds
}

type Playlist struct {
	Name     string
	Songs    []Song
	Duration int // total seconds, sum of all songs
}

// Scenario 2: Recipe book

type Ingredient struct {
	Name   string
	Amount float64
	Unit   string
}

type Recipe struct {
	Name        string
	Ingredients []Ingredient
	Steps       []string
	PrepMinutes int
}

// Scenario 3: Class schedule

type MeetingTime struct {
	Day       string // "Monday", "Tuesday", etc.
	StartHour int    // 24-hour format: 14 means 2:00 PM
}

type Course struct {
	Name       string
	Instructor string
	Room       string
	Meetings   []MeetingTime
}

func main() {
	playlist := Playlist{
		Name: "Study Mix",
		Songs: []Song{
			{Title: "Intro", Artist: "The xx", Duration: 128},
			{Title: "Svefn-g-englar", Artist: "Sigur Ros", Duration: 610},
			{Title: "Midnight City", Artist: "M83", Duration: 243},
		},
		Duration: 981,
	}
	fmt.Printf("Playlist: %s (%d songs, %d seconds)\n",
		playlist.Name, len(playlist.Songs), playlist.Duration)

	recipe := Recipe{
		Name: "Guacamole",
		Ingredients: []Ingredient{
			{Name: "avocado", Amount: 3, Unit: "whole"},
			{Name: "lime juice", Amount: 2, Unit: "tablespoons"},
			{Name: "salt", Amount: 0.5, Unit: "teaspoons"},
			{Name: "cilantro", Amount: 0.25, Unit: "cups"},
		},
		Steps: []string{
			"Halve and pit the avocados.",
			"Scoop into a bowl and mash with a fork.",
			"Mix in lime juice, salt, and cilantro.",
		},
		PrepMinutes: 10,
	}
	fmt.Printf("Recipe: %s (%d ingredients, %d min)\n",
		recipe.Name, len(recipe.Ingredients), recipe.PrepMinutes)

	schedule := []Course{
		{
			Name:       "Intro to CS",
			Instructor: "Dr. Smith",
			Room:       "Shelby 1103",
			Meetings: []MeetingTime{
				{Day: "Monday", StartHour: 9},
				{Day: "Wednesday", StartHour: 9},
				{Day: "Friday", StartHour: 9},
			},
		},
		{
			Name:       "Calculus II",
			Instructor: "Dr. Chen",
			Room:       "Parker 224",
			Meetings: []MeetingTime{
				{Day: "Tuesday", StartHour: 14},
				{Day: "Thursday", StartHour: 14},
			},
		},
	}
	for _, c := range schedule {
		fmt.Printf("Course: %s with %s in %s (%d meetings/week)\n",
			c.Name, c.Instructor, c.Room, len(c.Meetings))
	}
}
