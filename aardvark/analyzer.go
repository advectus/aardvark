package aardvark

import (
  "fmt"
  "encoding/json"
  "os"
  "io/ioutil"
  "path/filepath"
)

func Analyze(config string) string {
  filepath.Walk(config, func(path string, info os.FileInfo, err error) error {
    if !info.IsDir() {
      fileName := info.Name()

      dat, err := ioutil.ReadFile(config + "/" + fileName)
      check(err)

      var data ResultsConfig
      err = json.Unmarshal(dat, &data)
      check(err)

      fmt.Printf("[%s] %s\n", data.Status, data.Asset)

      for _,assetEnv := range data.Envs {
        fmt.Printf("\t%s %d\n", assetEnv.Href, assetEnv.StatusCode)
      }
    }
    return nil
  })
  return config
}
