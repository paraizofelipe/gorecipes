# gorecipes

This project aims to develop a Rest API to control phone call records and generate bills.

# Getting Started

These instructions will provide you a copy of the project that can be run on your local machine for development and testing purposes.
Consult [deployment](#deployment) item for notes on how to deploy the project on a live system.

# Prerequisites

This package is built with go1.15, and all you need is provide with the go standard library.

# Installing

This is what you need to install the application from the source code:

```shell script
    git clone https://github.com/paraizofelipe/gorecipes.go
    go install
```

## Caution!

Before executing any command create your file with the environment variables by copying from lcoal.env

```shell script
    cat local.env > .env
``` 

To build the docker version you can use the `Makefile`:

```shell script
    make dk-build 
```

# Running the tests

Until I finish this README there is not so much Unit tests written.

But I will try to coverage unless 70% of unit tests for this code as soon as possible.

You can run tests like this:

```shell script
    make test
```

To run this code locally for test purposes use:

```shell script
    GIPHY_TOKEN=<TOKEN> HOST=0.0.0.0 PORT=8989 DEBUG=true make start
```

# Deployment

This codebase is cloud-native by design so you can use lots of environments to make this run anywhere you want.

But to make this even easier to you the codebase also provides a Dockerfile.

There is also a Makefile to make all this even easier.

Run with Makefile:

```shell script
    make dk-start
```

# API

## POST - /api/recipes

### Body of response

**Status**: 200

```json
{
  "keywords": [
    "tomato",
    "onion",
    "orange"
  ],
  "recipes": [
    {
      "title": "Tomato & Orange Cottage Cheese Salad \r\n\t\t\n",
      "ingredients": [
        "balsamic vinaigrette",
        "basil",
        "cottage cheese",
        "orange",
        "red onions",
        "tomato"
      ],
      "link": "http://www.kraftfoods.com/kf/recipes/tomato-orange-cottage-cheese-54326.aspx",
      "gif": "https://giphy.com/gifs/deadsetonlife-surprise-salad-l4FGxA2WdhH8K9qEg"
    },
    ....
  ]
}
```

**Status**: 500

```json
{
    "error": "error message"
}
```

# Authors

Felipe Paraizo - Initial work - [paraizo](http://paraizo.dev)
