package aardvark

type AardvarkConfig struct {
  Aardvark struct {
    Assets []string `json:"assets"`
    Envs []string `json:"envs"`
  } `json:"aardvark"`
}

type Asset struct {
  Name string `json:"asset"`
  Status string `json:"status"`
  Envs []Env `json:"envs"`
}

type Env struct {
  Href string `json:"href"`
  StatusCode int `json:"statusCode"`
  Md5 string `json:"md5"`
}

type ResultsConfig struct {
  Asset string `json:"asset"`
  Status string `json:"status"`
  Envs []Env `json:"envs"`
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}
