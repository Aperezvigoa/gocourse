package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ErrPokemonYaCapturado struct {
	message string
}

func (err ErrPokemonYaCapturado) Error() string {
	return err.message
}

type ErrTipo struct {
	message string
}

func (err ErrTipo) Error() string {
	return err.message
}

func ValidarTipo(tipo string) error {
	types := regexp.MustCompile(`^(Agua|Fuego|Planta|Electrico)$`)
	if !types.MatchString(tipo) {
		return ErrTipo{"Error: Tipo invalido."}
	}
	return nil
}

var Legendarios = map[string]bool{
	"Mewtwo":   true,
	"Lugia":    true,
	"Ho-oh":    true,
	"Dialga":   true,
	"Palkia":   true,
	"Giratina": true,
}

var Tipos = [3]string{
	"Fuego", "Agua", "Electrico",
}

type Pokemon struct {
	Nombre string
	Tipo   string
	Nivel  int
}

func (p *Pokemon) SubirNivel() {
	if p.Nivel < 100 {
		p.Nivel++
	} else {
		fmt.Println("El Pokemon ya es nivel maximo")
	}
}

func (p Pokemon) EsLegendario() {
	for i, _ := range Legendarios {
		if p.Nombre == i {
			fmt.Println("\nES LEGENDARIO!!!")
			return
		}
	}
}

func CapturarPokemon(p Pokemon, pokedex *[]Pokemon) error {
	for _, pk := range *pokedex {
		if p.Nombre == pk.Nombre {
			return ErrPokemonYaCapturado{"El Pokemon ya esta registrado!"}
		}
	}
	*pokedex = append(*pokedex, p)
	fmt.Printf("%s capturado!", p.Nombre)
	return nil
}

var pokemonTmpl = template.Must(template.New("pokemon").Parse("\n---------------------------------\nNombre: {{.Nombre}}\nTipo: {{.Tipo}}\nNivel: {{.Nivel}}"))

func MostrarPokedex(pokedex *[]Pokemon) {
	if len(*pokedex) == 0 {
		fmt.Println("No tienes ningun Pokemon")
	} else {
		for _, p := range *pokedex {
			pokemonTmpl.Execute(os.Stdout, p)
			p.EsLegendario()
		}
	}
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	var pokedex []Pokemon
	fmt.Println("=== POKEDEX ===")

	for {
		fmt.Println()
		fmt.Println("\n1. Capturar Pokemon")
		fmt.Println("2. Mostrar Pokedex")
		fmt.Println("3. Subir nivel")
		fmt.Println("4. Salir")
		fmt.Println()

		fmt.Print("Selecciona una opcion: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			nombre, tipo, nivel := requestPokemonData()
			errTipo := ValidarTipo(tipo)
			if errTipo != nil {
				fmt.Println(errTipo)
				continue
			}
			CapturarPokemon(Pokemon{Nombre: nombre, Tipo: tipo, Nivel: nivel}, &pokedex)
		case "2":
			MostrarPokedex(&pokedex)
		case "3":
			fmt.Print("Nombre del Pokemon: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)
			for i := range pokedex {
				if pokedex[i].Nombre == nombre {
					pokedex[i].SubirNivel()
				}
			}
		case "4":
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func requestPokemonData() (string, string, int) {
	fmt.Print("Nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)
	fmt.Print("Tipo: ")
	tipo, _ := reader.ReadString('\n')
	tipo = strings.TrimSpace(tipo)
	fmt.Print("Nivel: ")
	nivelstr, _ := reader.ReadString('\n')
	nivelstr = strings.TrimSpace(nivelstr)

	nivel, _ := strconv.Atoi(nivelstr)

	return nombre, tipo, nivel
}
