// Scenario 1: Playlist

const playlist = {
  name: "Study Mix",
  songs: [
    { title: "Intro", artist: "The xx", duration: 128 },
    { title: "Svefn-g-englar", artist: "Sigur Ros", duration: 610 },
    { title: "Midnight City", artist: "M83", duration: 243 },
  ],
  duration: 981, // total seconds, sum of all songs
};

console.log(`Playlist: ${playlist.name} (${playlist.songs.length} songs, ${playlist.duration} seconds)`);

// Scenario 2: Recipe book

const recipe = {
  name: "Guacamole",
  ingredients: [
    { name: "avocado", amount: 3, unit: "whole" },
    { name: "lime juice", amount: 2, unit: "tablespoons" },
    { name: "salt", amount: 0.5, unit: "teaspoons" },
    { name: "cilantro", amount: 0.25, unit: "cups" },
  ],
  steps: [
    "Halve and pit the avocados.",
    "Scoop into a bowl and mash with a fork.",
    "Mix in lime juice, salt, and cilantro.",
  ],
  prepMinutes: 10,
};

console.log(`Recipe: ${recipe.name} (${recipe.ingredients.length} ingredients, ${recipe.prepMinutes} min)`);

// Scenario 3: Class schedule

const schedule = [
  {
    name: "Intro to CS",
    instructor: "Dr. Smith",
    room: "Shelby 1103",
    meetings: [
      { day: "Monday", startHour: 9 },
      { day: "Wednesday", startHour: 9 },
      { day: "Friday", startHour: 9 },
    ],
  },
  {
    name: "Calculus II",
    instructor: "Dr. Chen",
    room: "Parker 224",
    meetings: [
      { day: "Tuesday", startHour: 14 },
      { day: "Thursday", startHour: 14 },
    ],
  },
];

for (const course of schedule) {
  console.log(`Course: ${course.name} with ${course.instructor} in ${course.room} (${course.meetings.length} meetings/week)`);
}
