![go](https://img.shields.io/static/v1?label=GO&labelColor=50bfff&message=API&color=000000&logo=go&logoColor=ffffff&style=flat-square)

# HTTP RACOONS

**Just like HTTP Cats, but not so good and with racoons instead of cats.**

This is my first Go project and I'm using it to try Go and see what I can do with it.

The language itself is not bad, is actually pretty interesting. The thing that bothers me the most is the habit of using a single letter for variable names.

> I know that this code is not well written, but I'm happy for the results.

## Running the project locally:

- Clone the repo:

```bash
git clone https://github.com/devlulcas/http-racoons.git
```

- Build it:

```bash
cd http-racoons

go build cmd/http-racoons/main.go
```

- See the project in your [localhost:3000/200](http://localhost:3000/200)

## Routes:

- /{http-code}

> Returns a JSON with the code and the url for the image for that code

```json
{
  "code": 200,
  "image": "/images/200"
}
```

- /images/{http-code}

> Returns a image with racoon for your code

`/images/403`

![forbidden](https://raw.githubusercontent.com/devlulcas/http-racoons/main/static/403.png)
