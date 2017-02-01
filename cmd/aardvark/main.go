package main

import (
  "fmt"
  "os"
  "runtime"
  "github.com/urfave/cli"
  "github.com/advectus/aardvark-cli/aardvark"
)

func main() {
  fmt.Printf("Runtime %s/%s\n", runtime.GOOS, runtime.GOARCH)

  if _, err := os.Stat("data"); os.IsNotExist(err) {
    os.Mkdir("data", os.ModePerm)
  }

  if _, err := os.Stat("data/results"); os.IsNotExist(err) {
    os.Mkdir("data/results", os.ModePerm)
  }

  app := cli.NewApp()
  app.Name = "Aardvark"
  app.Version = "0.0.1"
  app.Commands = []cli.Command{
    {
      Name: "crawl",
      Aliases: []string{"c"},
      Usage: "crawl assets",
      Action: func(c *cli.Context) error {
        config := c.Args().First()
        if (config!="") {
          aardvark.Crawl(config)
        }
        return nil
      },
    },
    {
      Name: "analyze",
      Aliases: []string{"a"},
      Usage: "analyze results",
      Action: func(c *cli.Context) error {
        config := c.Args().First()
        if (config!="") {
          aardvark.Analyze(config)
        }
        return nil
      },
    },
    {
      Name: "env",
      Aliases: []string{"e"},
      Usage: "options for envs",
      Subcommands: []cli.Command{
        {
          Name: "add",
          Usage: "add a new environment",
          Action: func(c *cli.Context) error {
            fmt.Println("add environment: ", c.Args().First())
            return nil
          },
        },
        {
          Name: "remove",
          Usage: "remove an environment",
          Action: func(c *cli.Context) error {
            fmt.Println("remove environment: ", c.Args().First())
            return nil
          },
        },
      },
    },
    {
      Name: "asset",
      Aliases: []string{"a"},
      Usage: "options for assets",
      Subcommands: []cli.Command{
        {
          Name: "add",
          Usage: "add a new asset",
          Action: func(c *cli.Context) error {
            fmt.Println("add asset: ", c.Args().First())
            return nil
          },
        },
        {
          Name: "remove",
          Usage: "remove an asset",
          Action: func(c *cli.Context) error {
            fmt.Println("remove asset: ", c.Args().First())
            return nil
          },
        },
      },
    },
  }

  app.Run(os.Args)
}
