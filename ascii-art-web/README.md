# ASCII Art Web

## Description

ASCII Art Web is a web application written in Go that allows users to generate ASCII art from text using different banner styles. Users can enter text into a web form, select a banner style, and see the resulting ASCII art displayed in the browser.

Supported banners:

- shadow
- standard
- thinkertoy

The project demonstrates building an HTTP server with Go, handling HTML forms, and rendering dynamic content using Go templates.

## Authors

- Francis Awuor
- Phil Aturo
- Janet Odalo

## Usage

1. Clone the repository:

```bash
git clone https://learn.zone01kisumu.ke/git/jvictori/ascii-art-web
cd ascii-art-web
```

2. Run the server:
```bash
go run .
```

Open your browser and go to:

`http://localhost:8080`


Enter your text, choose a banner, and click submit to generate ASCII art.

## Implementation Details
### Algorithm Overview

1. Banner Loading

    - Load the selected banner file from the banners/ directory.
    - Parse the file into a map of rune → ASCII art lines.
    - Handle errors:
2. Text Rendering

    - Split the input text by newlines.
    - For each line:
        - For each character:
            - Look up the ASCII art in the banner map.
            - Return ErrInvalidInput if the character is not supported.
        - Combine each row of the characters to form the ASCII art line.
    - Join all rendered lines and write to writer

3. Web Server

    - Serve HTML templates for the main page and ASCII art display.
    - Accept form submissions via POST /ascii-art.
    - Retrieve input text and banner choice from the form using r.FormValue.
    - Pass the input to the generator function and display the result.
    - Return appropriate HTTP status codes:

4. Template Rendering

    - Go templates are used to dynamically inject user input, banner choice, and ASCII art results into the HTML page.
    - Errors are also displayed in the template when they occur.