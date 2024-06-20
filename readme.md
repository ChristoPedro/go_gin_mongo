# Study of Gonlang with Gin and MongoDB

This is a study of query mongoDB and exposes it though restAPI using Gin.

<!-- TOC -->

- [Study of Gonlang with Gin and MongoDB](#study-of-gonlang-with-gin-and-mongodb)
  - [Routs](#routs)
  - [How to run](#how-to-run)

<!-- /TOC -->

## Routs

```/movies```

Return all movies from MongoDB sample collection *movies*.

```/movies/{movie name}```

Return all movies that mach parcially or with the full string.

```/movie/id/{id}```

Return the mobie by given MongoDB *_id*.

## How to run

1. Export Enviroment Variables

```bash
export MONGO_STRING={Your mongo atlas string}
```

2.Build

```bash
go build .
```

3.Run

Just run the created executable file.
