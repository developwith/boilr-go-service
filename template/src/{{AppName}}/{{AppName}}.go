package {{AppName}}

import (
  "github.com/spf13/viper"
  "log"
  "fmt"

  "os"
  "runtime"
  "util/round"
)

func main() {

  logger := log.New(os.Stderr, "{{AppName}}", log.Lshortfile)
  viper.AutomaticEnv()
  _, filename, _, _ := runtime.Caller(1)
  configpath := path.Join(path.Dir(filename), "config")
  viper.SetConfigName("config") // name of config file (without extension)
  viper.SetConfigType("yaml")
  viper.AddConfigPath("$HOME/.{{AppName}}") // call multiple times to add many search paths
  viper.AddConfigPath(configpath)

  err := viper.ReadInConfig() // Find and read the config file
  if err != nil {             // Handle errors reading the config file
    logger.Println("Fatal error config file: %s \n", err)
    panic(err)
  }
  logger.Println("service:starting")

  startMaxProcs()
  startApi()

}

func setupMaxProcs() {
  numCPU := util.Round(float64(runtime.NumCPU()/2), 1, 0)
  runtime.GOMAXPROCS(int(numCPU))
}

func startApi() {
  // TODO more fun here
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":"+viper.GetString("API_PORT"), nil))
}