package service

import (
  "github.com/spf13/viper"
  "log"
  "fmt"
  "path"
  "net/http"
  "html"
  "os"
)

// Handle args from command line if needed
func Run(args []string) error {

  logger := log.New(os.Stderr, "{{AppName}}", log.Lshortfile)
  err := setupConfig()
  logger.Println("service:started")
  startApi()
  return err
}

func setupConfig()  error {
  viper.AutomaticEnv()

  filename := os.Args[0]
  configpath := path.Join(path.Dir(filename), "../", "config")
  viper.SetConfigName("config") // name of config file (without extension)
  viper.SetConfigType("yaml")
  viper.AddConfigPath("$HOME/.{{AppName}}") // call multiple times to add many search paths
  viper.AddConfigPath(configpath)

  err := viper.ReadInConfig() // Find and read the config file
  if err != nil {             // Handle errors reading the config file
    log.Println("Fatal error config file: %s \n", err)
    panic(err)
  }
  return err
}

func startApi() {
  // TODO more fun here
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":"+viper.GetString("API_PORT"), nil))
}