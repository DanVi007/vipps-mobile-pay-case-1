# vipps-mobile-pay-case-1
Case 1 for vipps-mobile-pay 
This project solves option 1 from vipps-mobile-pay.

## The case
### Part 1:

Using an HTTP GET method, retrieve information from Wikipedia using a given topic. Query https://en.wikipedia.org/w/api.php?action=parse&section=0&prop=text&format=json&page=[topic] to get the topic Wikipedia article.

Print the total number of times that the string [topic] appears in the article's text field. Expose your backend code as an API.
### Part 2:

Write a frontend that calls the API that you're exposing. The frontend can be very simple - just an input field and some way to submit and display the results. Feel free to be creative.

For both parts, use any language or framework that you’d like.

## Solution
I have decided to use a go backend and a react frontend.


## To build run the project
The entire project is dockerized so to build simply type:
```bash
docker compose up
```
or
```bash
docker-compose up
```

sudo might be needed if it does not work.

The frontend will run on http://localhost:3000
The backend will run on http://localhost:8080