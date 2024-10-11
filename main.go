package main

import (
	_ "embed"
	"flag"
	"fmt"
	"atomic/elements"
)

//go:embed data/elements.csv
var elementsCSV string

//go:embed data/molecules.csv
var moleculesCSV string

func main() {
	ptCmd := flag.Bool("pt", false, "Draw periodic table")
	eCmd := flag.Bool("e", false, "Show electron configurations")
	flag.Parse()
	args := flag.Args()

	// Ensure a chemical formula is provided
	if len(args) == 0 {
		return
	}
	
	formula := args[0] // First non-flag argument is the formula

	// Load elements data from CSV
	err := elements.LoadElements(elementsCSV)
	if err != nil {
		fmt.Printf("Error loading elements: %v\n", err)
		return
	}

	// Load molecules data from CSV
	err = elements.LoadMolecules(moleculesCSV)
	if err != nil {
		fmt.Println("Error loading molecules:", err)
		return
	}

	// Parse the chemical formula
	compound, err := elements.ParseFormula(formula)
	if err != nil {
		fmt.Println(err)
		return
	}

	molecule := compound.ToMolecule()

	fmt.Println()

	// Draw the periodic table if the draw command is used
	if *ptCmd {
		elements.DrawPeriodicTable(molecule)
		fmt.Println()
	}
	
	// fmt.Printf("%+v\n\n",compound)
	// for i, m := range compound.Molecules {
		// fmt.Printf("%d %+v\n", i, m.ToString())
	// }
	// fmt.Println("")

	// Output chemical information
	name := compound.GetName()
	el, exists := elements.ElementTable[formula]
	if exists {
		name = el.Name
	}
	fmt.Printf("  Molecule : %s\n", compound.ToString())
	fmt.Printf("  Simplify : %s\n", compound.ToMolecule().Simplify())
	fmt.Printf("  Name     : %s\n", name)
	if compound.State != "" {
		fmt.Printf("  State    : %s\n", compound.State)
	}
	fmt.Printf("  Mass     : %f\n", compound.GetMass())
	fmt.Printf("  Charge   : %d\n", compound.GetCharge())
	if exists {
		fmt.Printf("  Number   : %d\n", el.Number)
		fmt.Printf("  Category : %s\n", el.Category)
		fmt.Printf("  Fact     : %s\n", el.Fact)
	}
	fmt.Println()

	// Show electron configurations if -e is passed
	if *eCmd {
		printed := make(map[string]bool)
		for _, el := range molecule.Elements {
			if !printed[el.Symbol] {
				fmt.Printf("  %s(%d): %s\n", el.Symbol, el.Number, el.GetElectronConfiguration().ToString())
				printed[el.Symbol] = true
			}
		}
	}
}
