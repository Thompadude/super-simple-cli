# Super Simple CLI
A very simple CLI written in Go.
# Why?
For learning purposes.
# How to Use
The application has two modes
- `todo`
- `swapi`
## Non-flag Arguments
Each mode is enabled by running the application with the respective name as a non-flag argument.
## Flag Arguments
The application accepts one flag – `id`. This is the ID of the to-do or the SWAPI person to fetch.  
## Example
Running the program with arguments `-id=1 swapi` will fetch `https://swapi.co/api/people/1/`
# APIs
- https://jsonplaceholder.typicode.com/
- https://swapi.co/
# Credit
Yoda ASCII made by Shanaka Dias
https://www.asciiart.eu/movies/star-wars