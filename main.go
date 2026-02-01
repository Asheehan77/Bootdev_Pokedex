package main
import(
	"github.com/Asheehan77/Bootdev_Pokedex/internal/pokeapi"
	"time"
)

func main(){
	pokeclient := pokeapi.NewClient(5*time.Second,10*time.Minute)
	cfg := &config{
		pokeapiClient: pokeclient,
	}
	runRepl(cfg)
}