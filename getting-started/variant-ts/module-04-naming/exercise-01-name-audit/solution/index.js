const filterAvailable = (books) => {
  const available = [];
  for (const book of books) {
    if (book.available) {
      available.push(book);
    }
  }
  return available;
};

const averagePrice = (books) => {
  if (books.length === 0) {
    return 0;
  }
  let total = 0;
  for (const book of books) {
    total += book.price;
  }
  return total / books.length;
};

const groupByYear = (books) => {
  const byYear = {};
  for (const book of books) {
    if (!byYear[book.year]) {
      byYear[book.year] = [];
    }
    byYear[book.year].push(book);
  }
  return byYear;
};

const formatBook = (book) => {
  const parts = [];
  parts.push(book.title);
  parts.push(`(year ${book.year})`);
  parts.push(`$${book.price.toFixed(2)}`);
  if (book.available) {
    parts.push("[available]");
  } else {
    parts.push("[sold]");
  }
  return parts.join(" ");
};

const catalog = [
  { title: "JS Programming", year: 2023, price: 49.99, available: true },
  { title: "The Art of SQL", year: 2021, price: 39.95, available: false },
  { title: "Systems Design", year: 2024, price: 54.99, available: true },
  { title: "Network Protocols", year: 2022, price: 44.50, available: true },
  { title: "Data Structures", year: 2023, price: 42.00, available: false },
];

console.log("All:");
for (const book of catalog) {
  console.log(" ", formatBook(book));
}

const available = filterAvailable(catalog);
console.log(`\nAvailable: ${available.length}`);
for (const book of available) {
  console.log(" ", formatBook(book));
}

console.log(`\nAverage price: $${averagePrice(catalog).toFixed(2)}`);

const byYear = groupByYear(catalog);
console.log("\nBy year:");
for (const year of Object.keys(byYear)) {
  console.log(`  ${year}: ${byYear[year].length} items`);
}
