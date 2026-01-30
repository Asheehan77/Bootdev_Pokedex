package main
import(
	"github.com/Asheehan77/Bootdev_Pokedex/internal"
	"time"
)

func main(){
	pokeclient := internal.NewClient(5*time.Second)
	cfg := &config{
		pokeapiClient: pokeclient,
	}
	runRepl(cfg)
}