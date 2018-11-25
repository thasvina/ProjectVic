package main
import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
    Database struct {
        Host     string `json:"host"`
        Password string `json:"password"`
    } `json:"database"`
    Host string `json:"host"`
    Port string `json:"port"`
}
func LoadConfiguration(file string) (Config, error) {
var config Config
configFile, err := os.Open(file)
defer configFile.Close()
if err != nil {
return config, err
}

jsonParser := json.NewDecoder(configFile)
err = jsonParser.Decode(&config)
return config, err
}
func main() {

fmt.Println("startting the application")
config, _ :=LoadConfiguration("config.json")
fileWriter, _ := os.Create("output.txt")
	json.NewEncoder(fileWriter).Encode(config)
fmt.Println(config)

}
