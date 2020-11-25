## Aardvark CLI

A CLI script used to help track static assets that have been deployed to multiple servers. Use the aardvark script to compare the md5 of each HTTP response for each asset on every server in your infrastructure. The results for each asset will be saved to data/results/assetName, and will specify the md5 obtained when polling that particular asset on each of the specified environments.

## Config

### aardvark.json config

```
{
  "Aardvark": {
    "assets": [
      "js/custom.js",
      "css/custom.css"
    ],
    "envs": [
      "https://dev.company.com",
      "https://staging.company.com",
      "https://www.company.com"
    ]
  }
}
```

### specify assets to track
```
"assets": [
  "js/custom.js",
  "css/custom.css"
]
```
Specify the assets that you are looking to examine

### specify environments to fetch each asset from
```
"envs": [
  "https://dev.company.com",
  "https://staging.company.com",
  "https://www.company.com"
]
```
Specify the environments that you wish to fetch the assets from


## Usage
```
wget -O aardvark https://raw.githubusercontent.com/advectus/aardvark-cli/master/bin/aardvark-osx && chmod +x aardvark
./aardvark crawl aardvark.json
./aardvark analyze /data/results
```

## Development

### Install golang

https://golang.org/doc/install

```
$ go version
go version go1.15.5 darwin/amd64
```

### manually install dependencies
```
$ go get github.com/edhere/aardvark-cli/aardvark
$ go get github.com/urfave/cli
```

### Build 

```
$ go build -o aardvark-osx cmd/aardvark/main.go
$ mv aardvark-osx bin/aardvark-osx
```

### Setup aardvark.json

At project root, create an aardvark.json using the following schema:

```
{
  "Aardvark": {
    "assets": [],
    "envs": []
  }
}
```

### test interactive with program

```
$ ./bin/aardvark-osx crawl ../aardvark.json
```