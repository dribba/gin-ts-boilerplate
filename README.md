# gin-typescript-react-boilerplate

- [About](#about)
- [Requirements](#requirements)
- [Setup](#setup)
- [License](#license)

## About

This project is a boilerplate for [Gin](https://gin-gonic.github.io/gin/) using [React](https://facebook.github.io/react/) and [TypeScript](https://github.com/Microsoft/TypeScript) based off [wadahiro/gin-react-boilerplate](https://github.com/wadahiro/gin-react-boilerplate).

## Requirements

- [Golang](http://golang.org/)
- [Node.js](https://nodejs.org/)

## Setup

1. Install some golang tools by `go get`

```bash
go get -u github.com/GeertJohan/go.rice
go get -u github.com/GeertJohan/go.rice/rice
go get -u github.com/gin-gonic/gin
go get -u github.com/gin-gonic/contrib/static
go get -u github.com/pilu/fresh
```

2. Install JavaScript dependencies

```bash
npm install
```

## Run with development mode

Start webpack and gin with watch mode.

```bash
npm run dev & fresh
```

## Release Build

Run webpack with production mode. All you have to do is run `npm run build`.

```bash
npm run build
```

The built artifact is created at `./dist/server`.

## License

Licensed under the [MIT](/LICENSE.txt) license.
