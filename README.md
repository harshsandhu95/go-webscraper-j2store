# Web Scraper for j2store.net

This Go program scrapes product data from j2store.net and saves it to a JSON file.

## Requirements

- Go programming language
- Colly web scraping library (`go get github.com/gocolly/colly`)

## Usage

1.  Make sure you have Go and Colly installed.
2.  Run the code: `go run main.go`
3.  The scraped product data will be saved in `products.json`.

## Code explanation

- **`main.go`:**
    - Imports necessary packages for JSON encoding, file operations, and web scraping.
    - Defines an `item` struct to represent product data (name, price, image URL).
    - Uses Colly to:
        - Create a collector instance with allowed domain `j2store.net`.
        - Define callbacks to extract product data from HTML elements using CSS selectors.
        - Handle pagination to scrape data from multiple pages.
        - Print visited URLs for logging.
    - Encodes the collected data into JSON format.
    - Writes the JSON data to `products.json`.

## Data format

The `products.json` file contains an array of product objects, each with the following fields:

- **`name`:** Product name
- **`price`:** Product price
- **`imgurl`:** Product image URL

## Note

- This code is specifically designed for scraping data from j2store.net. You may need to modify the CSS selectors and pagination logic if you want to use it for other websites.
- Always respect the website's robots.txt and terms of service when scraping.
