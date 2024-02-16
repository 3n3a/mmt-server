package main

import (
    "os"
    "log"
    "net"

    "github.com/urfave/cli/v2"
    "github.com/gofiber/fiber/v2"
)

type Config struct {
    Name string `json:"name"`
    Description string `json:"description"`
    Version string `json:"version"`
}

func main() {
    // cli
    var port string

    conf := Config{
        Name: "server",
        Description: "server that servers",
        Version: "v1.1",
    }

    cli := &cli.App{
        Name: conf.Name,
        Usage: conf.Description,
        Version: conf.Version,
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:  "port",
                Value: "3000",
                Usage: "port for webserver",
                Destination: &port,
            },
        },
    }

    if err := cli.Run(os.Args); err != nil {
        log.Fatal(err)
    }

    // server
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(conf)
    })

    app.Listen(net.JoinHostPort("", port))
}
