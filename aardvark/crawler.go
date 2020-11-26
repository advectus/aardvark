package aardvark

import (
  "fmt"
  "net/http"
  "crypto/md5"
  "strings"
  "encoding/json"
  "io/ioutil"
)


func Crawl(config string) string {
  fmt.Println("Aardvark config: ", config)

  dat, err := ioutil.ReadFile(config)
  check(err)

  var data AardvarkConfig
  err = json.Unmarshal(dat, &data)
  check(err)

  for _,asset := range data.Aardvark.Assets {
    //fmt.Println(asset)
    exec(data.Aardvark.Envs, asset)
  }
  return config
}

func exec(envs []string, path string) {

  asset := Asset{
    Name: path,
    Status: "success",
    Envs: []Env{},
  }

  for _,env := range envs {
    fmt.Println(env)
    href := env + "/" + path
    fmt.Println(href)
    response, err := http.Get(href)
    check(err)

    body, err := ioutil.ReadAll(response.Body)
    check(err)

    bodyString := string(body)
    data := []byte(bodyString)
    md5 := md5.Sum(data)
    //fmt.Printf("%x",md5)
    md5String := fmt.Sprintf("%x", md5)

    asset.Envs = append(asset.Envs, Env{
      Href: href,
      StatusCode: response.StatusCode,
      Md5: string(md5String),
    })
  }

  // examine md5s for all of the assets
  temp := ""
  for _,env := range asset.Envs {
    if (temp=="") {
      temp = env.Md5
    }
    if (temp!=env.Md5||env.StatusCode!=200) {
      asset.Status = "failure"
    }
  }

  b, err := json.Marshal(asset)
  check(err)

  fmt.Printf("[%s] %s\n", asset.Status, asset.Name)

  pathSub := strings.Replace(path, "/", "", -1)
  fileName := strings.Replace(pathSub, ".", "", -1)
  ioutil.WriteFile("data/results/" + fileName + ".json", b, 0644)

}
