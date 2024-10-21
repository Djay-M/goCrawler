# Web Crawler Service
Web crawler Service built using Go Lang.

Why Go-lang and not python, because go lang gives us the prower of go routines, which are very lightWieght and fast compared to python threads.

<!-- TABLE OF CONTENTS -->
## Table of Contents
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
        </li>
      </ul>
    </li>
    <li><a href="#contact">Contact</a></li>
  </ol>

<!-- ABOUT THE PROJECT -->

## About The Project

Developed a backend service, which provides APIs for fetching product URL for requested websites.

- Backend: GO
- Database: N/A, but there is scope of adding mongo DB

## API Details

All the apis are version controlled.()

- #### Web Crawler Apis:

  - Api to crawl web(POST).
    - path: /api/v1/crawler
    - response: (mock)
    ```
    [{
        "website": "website request to crawl",
        "productLinks": [ "product urls (array of string)"]
    }]
    ```

- #### Example CURL:

```
curl --location 'localhost:4000/api/v1/crawler' \
--header 'Content-Type: application/json' \
--data '{
    "websites": [
        "https://www.meesho.com/accessories-men/pl/3tp"
    ]
}'
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

The following commands starts the server.
```sh
go run .
```

## Contact

Dhananjaya.M - [@my_website](https://elements.getpostman.com/redirect?entityId=33121059-44362fbd-84ab-46e9-ba80-ad911b366bf7&entityType=collection)

Project Link: [https://github.com/Djay-M/goCrawler](https://github.com/Djay-M/goCrawler)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
