# Champion Store Scraper

Store url: [Champion Store](https://www.championstore.com/en_gb/)

## Overview

This program scrapes example items from Champion Store website. It extracts information such as the item collection, name, price, product link, and image link, and saves it to a CSV file.

## Getting Started:

### Prerequisites

- Docker installed on your machine
- Docker Compose installed on your machine
- Internet connection

### Instalation

1. Clone the repository

```bash
$ git clone https://github.com/x4trm/champion-store-scraper.git
$ cd champion-store-scraper
```

2. docker-compose:

```bash
$ docker-compose build
$ docker-compose up
```

## Configuration

Create or edit a `config.json` file in the root directory of the project. This file should contain the configuration for the items you want to scrape. Here is an example configuration:

```json
{
  "entries": [
    {
      "item": "Hoodie",
      "baseUrl": "https://www.championstore.com/en_gb/men/clothing/hoodies?p=",
      "maxPages": 3,
      "sex": "men"
    },
    {
      "item": "t-shirt",
      "baseUrl": "https://www.championstore.com/en_gb/men/clothing/t-shirts?p=",
      "maxPages": 5,
      "sex": "men"
    }
  ]
}
```

## Data

The output data should be in the output directory in the data.csv file.
