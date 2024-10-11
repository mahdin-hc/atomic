package elements

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"github.com/mattn/go-colorable"
)

// Element represents a chemical element and its properties
type Element struct {
	Symbol   string
	Category string
	Number   int
	Group    int
	Amu      float64
	Fact     string
	Period   int
	Phase    string
	Name     string
	Colour   string
	Charge int
}

func (el Element) ToString() string {
	if el.Charge > 1 {
		return fmt.Sprintf("%s+%d", el.Symbol, el.Charge)
	} else if el.Charge < -1 {
		return fmt.Sprintf("%s%d", el.Symbol, el.Charge)
	} else if el.Charge == -1 {
		return el.Symbol + "-"
	} else if el.Charge == 1 {
		return el.Symbol + "+"
	}
	return el.Symbol
}



func (el Element) GetElectronConfiguration() ElectronConfiguration {
	return GenerateElectronConfiguration(el.Number)
}

func (el Element) GetQuantumNumbers(i int) (QuantumNumbers, error) {
	return GetQuantumNumbers(el.Number, i)
}

// Molecule represents a parsed chemical formula
type Molecule struct {
	Elements []Element
	Charge   int
	Name     string
	State    string
}

func (m Molecule) ToCompound() Compound {
	return Compound{Molecules: []Molecule{m}, Charge: m.Charge, State: m.State}
}

func (m Molecule) Simplify() string {
	var str string
	for _, el := range m.Elements {
		str += el.Symbol
	}
	return str
}

func (m Molecule) ToString() string {
	var str string
	count := 1
	prevSym := ""

	for i, el := range m.Elements {
		// If this is the first element or the symbol is the same as the previous one
		if i == 0 {
			prevSym = el.ToString()
			continue
		}

		if el.ToString() == prevSym {
			count++ // Same symbol, increase the count
		} else {
			// Append the previous element and its count
			str += prevSym
			if count > 1 {
				str += fmt.Sprintf("%d", count)
			}
			// Reset for the new element
			prevSym = el.ToString()
			count = 1
		}
	}

	// Append the last element and its count
	str += prevSym
	if count > 1 {
		str += fmt.Sprintf("%d", count)
	}

	// Add charge representation
	if m.Charge == 1 {
		str += "+"
	} else if m.Charge == -1 {
		str += "-"
	} else if m.Charge > 1 {
		str += fmt.Sprintf("+%d", m.Charge)
	} else if m.Charge < -1 {
		str += fmt.Sprintf("%d", m.Charge)
	}

	return str
}



func (m Molecule) GetMass() float64 {
	var sum float64
	for _, el := range m.Elements {
		sum += el.Amu
	}
	return sum
}

var ElementTable = map[string]Element{}
var CompoundTable = map[string]Compound{}

// LoadElements loads the CSV data of elements into the ElementTable map
func LoadElements(data string) error {
	r := csv.NewReader(strings.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %v", err)
	}

	for _, record := range records[1:] { // Skip header
		number, _ := strconv.Atoi(record[2])
		group, _ := strconv.Atoi(record[3])
		amu, _ := strconv.ParseFloat(record[4], 64)
		period, _ := strconv.Atoi(record[6])

		element := Element{
			Symbol:   record[0],
			Category: record[1],
			Number:   number,
			Group:    group,
			Amu:      amu,
			Fact:     record[5],
			Period:   period,
			Phase:    record[7],
			Name:     record[8],
			Colour:   record[9],
		}

		ElementTable[element.Symbol] = element
	}

	return nil
}


// LoadMolecules loads the CSV data into a slice of Molecules
func LoadMolecules(data string) error {
	r := csv.NewReader(strings.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %v", err)
	}

	for ri, record := range records[1:] { // Skip header
		if len(record) < 3 {
			return errors.New("invalid record: insufficient columns")
		}

		formula := record[0]
		name := record[1]
		state := record[2]

		// Parse the formula to get elements
		p := NewParser(formula)
		compound, err := p.ParseCompound()
		if err != nil {
			return fmt.Errorf("%s:%d, %s\n", formula, ri+2, err.Error())
		}
		
		// Set the Name and Charge
		compound.Name = name
		compound.State = state

		CompoundTable[compound.ToString()] = compound
	}

	return nil
}


// Compound represents a chemical compound made up of multiple molecules or ions
type Compound struct {
	Molecules []Molecule
	Name      string
	State     string
	Charge    int
}

// ToString simplifies and returns the string representation of the compound
func (c Compound) ToString() string {
	var str string
	count := 1
	prevSym := ""

	for i, mol := range c.Molecules {
		// If this is the first element or the symbol is the same as the previous one
		if i == 0 {
			prevSym = mol.ToString()
			continue
		}

		if mol.ToString() == prevSym {
			count++ // Same symbol, increase the count
		} else {
			// Append the previous element and its count
			if count > 1 {
				str +=  "(" + prevSym + ")"
				str += fmt.Sprintf("%d", count)
				// str = fmt.Sprintf("(%s)%d", str, count)
			} else {
				str += prevSym
			}
			// Reset for the new element
			prevSym = mol.ToString()
			count = 1
		}
	}

	// Append the last element and its count
	if count > 1 {
		str +=  "(" + prevSym + ")"
		str += fmt.Sprintf("%d", count)
	} else {
		str +=  prevSym
	}

	// Add charge representation
	if c.Charge == 1 {
		str += "+"
	} else if c.Charge == -1 {
		str += "-"
	} else if c.Charge > 1 {
		str += fmt.Sprintf("+%d", c.Charge)
	} else if c.Charge < -1 {
		str += fmt.Sprintf("%d", c.Charge)
	}

	return str
}

func (c Compound) ToMolecule() Molecule {
	var mol = Molecule{}
	for _, m := range c.Molecules {
		mol.Elements = append(mol.Elements, m.Elements...)
	}
	return mol
}

// GetMass calculates the total mass of the compound
func (c Compound) GetMass() float64 {
	var sum float64
	for _, mol := range c.Molecules {
		sum += mol.GetMass()
	}
	return sum
}

func (c Compound) GetCharge() int {
	var charge = c.Charge
	for i, _ := range c.Molecules {
		charge += c.Molecules[i].Charge
		for j, _ := range c.Molecules[i].Elements {
			charge += c.Molecules[i].Elements[j].Charge
		}
	}
	return charge
}

func (c Compound) GetName() string {
	var str string
	str += c.Name
	for i, mol := range c.Molecules {
		if mol.Name == "" {
			continue
		}
		str += mol.Name
		if i < len(c.Molecules) - 1 {
			str += ", "
		}
	}
	return str
}

func ParseFormula(formula string) (Compound, error) {
	p := NewParser(formula)
	
	compound, err := p.ParseCompound()
	if err != nil {
		return compound, err
	}
	
	m, exists := CompoundTable[compound.ToString()]
	if exists {
		compound.Name = m.Name
		compound.State = m.State
	}
	
	return compound, nil
}

// DrawPeriodicTable highlights elements in the molecule
func DrawPeriodicTable(molecule Molecule) {
	// Create a 2D grid for the main periodic table (7 periods, 18 groups)
	table := make([][]string, 7)
	for i := 0; i < 7; i++ {
		table[i] = make([]string, 18)
	}

	// Create a separate grid for Lanthanides and Actinides
	lanthanides := make([]string, 15) // 57 (La) to 71 (Lu)
	actinides := make([]string, 15)   // 89 (Ac) to 103 (Lr)

	// Populate the main table with element symbols
	for _, el := range ElementTable {
		if el.Number >= 57 && el.Number <= 71 {
			lanthanides[el.Number-57] = el.Symbol
		} else if el.Number >= 89 && el.Number <= 103 {
			actinides[el.Number-89] = el.Symbol
		} else if el.Period <= 7 && el.Group <= 18 && el.Group > 0 {
			table[el.Period-1][el.Group-1] = el.Symbol
		}
	}

	// Check which elements are in the molecule for highlighting
	highlighted := make(map[string]bool)
	for _, el := range molecule.Elements {
		highlighted[el.Symbol] = true
	}

	// Print the main periodic table, highlighting elements in the molecule
	for _, row := range table {
		for _, el := range row {
			if el == "" {
				fmt.Printf("%4s", "")
			} else if highlighted[el] {
				fmt.Fprintf(colorable.NewColorableStdout(), "\x1b[0;92m%4s\x1b[0m", el)
			} else {
				fmt.Printf("%4s", el)
			}
		}
		fmt.Println()
	}

	// Print the Lanthanide and Actinide series, highlighting elements in the molecule
	fmt.Println()
	fmt.Print("        ")
	for _, el := range lanthanides {
		if highlighted[el] {
			fmt.Fprintf(colorable.NewColorableStdout(), "\x1b[0;92m%4s\x1b[0m", el)
		} else {
			fmt.Printf("%4s", el)
		}
	}
	fmt.Println()
	fmt.Print("        ")
	for _, el := range actinides {
		if highlighted[el] {
			fmt.Fprintf(colorable.NewColorableStdout(), "\x1b[0;92m%4s\x1b[0m", el)
		} else {
			fmt.Printf("%4s", el)
		}
	}
	fmt.Println()
}
