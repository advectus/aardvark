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
Specify each asset that you are looking to examine

### specify environments to fetch each asset from
```
"envs": [
  "https://dev.company.com",
  "https://staging.company.com",
  "https://www.company.com"
]
```
Specify the environment that you wish to fetch the assets from


## Usage
```
wget -O aardvark https://github.com/advectus/aardvark-cli/blob/master/bin/aardvark-osx
./aardvark crawl aardvark.json
./aardvark analyze /data/results
```
