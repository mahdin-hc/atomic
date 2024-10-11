package elements

import (
	"fmt"
	"strings"
)

// Subshell represents a subshell and its electron count (e.g., 1s2)
type Subshell struct {
	Orbital string
	Electrons int
}

// ElectronConfiguration represents a series of subshells
type ElectronConfiguration struct {
	Subshells []Subshell
}

// toString returns the string representation of the electron configuration (e.g., "1s2 2s2 2p6")
func (ec ElectronConfiguration) ToString() string {
	var sb strings.Builder
	for _, subshell := range ec.Subshells {
		sb.WriteString(fmt.Sprintf("%s%d ", subshell.Orbital, subshell.Electrons))
	}
	return strings.TrimSpace(sb.String())
}

var exceptionConfigurations = map[int]string{
	// Transition Metals
	24: "1s2 2s2 2p6 3s2 3p6 4s1 3d5",  // Chromium (Cr)
	29: "1s2 2s2 2p6 3s2 3p6 4s1 3d10", // Copper (Cu)
	39: "1s2 2s2 2p6 3s2 3p6 4s2 3d5",  // Yttrium (Y)
	40: "1s2 2s2 2p6 3s2 3p6 4s2 3d10", // Zirconium (Zr)
	41: "1s2 2s2 2p6 3s2 3p6 4s2 3d5",  // Niobium (Nb)
	42: "1s2 2s2 2p6 3s2 3p6 4s2 3d10", // Molybdenum (Mo)
	43: "1s2 2s2 2p6 3s2 3p6 4s2 3d5",  // Technetium (Tc)
	44: "1s2 2s2 2p6 3s2 3p6 4s2 3d10", // Ruthenium (Ru)
	45: "1s2 2s2 2p6 3s2 3p6 4s2 3d5",  // Rhodium (Rh)
	46: "1s2 2s2 2p6 3s2 3p6 4s2 3d10", // Palladium (Pd)
	47: "1s2 2s2 2p6 3s2 3p6 4s2 3d10", // Silver (Ag)
	48: "1s2 2s2 2p6 3s2 3p6 4s2 3d10", // Cadmium (Cd)
	57: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Lanthanum (La)
	58: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Cerium (Ce)
	59: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Praseodymium (Pr)
	60: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Neodymium (Nd)
	61: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Promethium (Pm)
	62: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Samarium (Sm)
	63: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Europium (Eu)
	64: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Gadolinium (Gd)
	65: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Terbium (Tb)
	66: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Dysprosium (Dy)
	67: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Holmium (Ho)
	68: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Erbium (Er)
	69: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Thulium (Tm)
	70: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Ytterbium (Yb)
	71: "1s2 2s2 2p6 3s2 3p6 4s2 4f14 5d10 6s2", // Lutetium (Lu)
	90: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Actinium (Ac)
	91: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Protactinium (Pa)
	92: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Uranium (U)
	93: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Neptunium (Np)
	94: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Plutonium (Pu)
	95: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Americium (Am)
	96: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Curium (Cm)
	97: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Berkelium (Bk)
	98: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Californium (Cf)
	99: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Einsteinium (Es)
	100: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Fermium (Fm)
	101: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Mendelevium (Md)
	102: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Nobelium (No)
	103: "1s2 2s2 2p6 3s2 3p6 4s2 5f14 6d10 7s2", // Lawrencium (Lr)
}


// GenerateElectronConfiguration generates the electron configuration for an element based on its atomic number
func GenerateElectronConfiguration(atomicNumber int) ElectronConfiguration {
	// Check if this element has an exception
	if exceptionConfig, found := exceptionConfigurations[atomicNumber]; found {
		return parseConfigurationString(exceptionConfig)
	}

	// The list of orbitals in the order of filling
	orbitals := []string{"1s", "2s", "2p", "3s", "3p", "4s", "3d", "4p", "5s", "4d", "5p", "6s", "4f", "5d", "6p", "7s", "5f", "6d", "7p"}
	// Maximum number of electrons each subshell can hold
	maxElectrons := map[string]int{
		"s": 2, "p": 6, "d": 10, "f": 14,
	}

	configuration := ElectronConfiguration{}
	electronsRemaining := atomicNumber

	for _, orbital := range orbitals {
		// Determine the type of subshell (s, p, d, f) and its max electrons
		subshellType := string(orbital[len(orbital)-1:])
		maxInSubshell := maxElectrons[subshellType]

		// Determine how many electrons to fill in this subshell
		electronsInSubshell := min(electronsRemaining, maxInSubshell)

		// Append the subshell to the configuration
		configuration.Subshells = append(configuration.Subshells, Subshell{
			Orbital: orbital,
			Electrons: electronsInSubshell,
		})

		// Subtract the filled electrons from the remaining electrons
		electronsRemaining -= electronsInSubshell

		// If all electrons are placed, we can break out of the loop
		if electronsRemaining == 0 {
			break
		}
	}

	return configuration
}

// Helper function to return the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// parseConfigurationString parses an electron configuration string (e.g., "1s2 2s2 2p6") into ElectronConfiguration
func parseConfigurationString(configStr string) ElectronConfiguration {
	configuration := ElectronConfiguration{}
	parts := strings.Fields(configStr) // Split the string into individual subshells (e.g., "1s2", "2s2")

	for _, part := range parts {
		// Extract the orbital (e.g., "1s") and electron count (e.g., "2")
		var orbital string
		var electrons int
		fmt.Sscanf(part, "%2s%d", &orbital, &electrons) // Parse the subshell and electron count

		configuration.Subshells = append(configuration.Subshells, Subshell{
			Orbital: orbital,
			Electrons: electrons,
		})
	}

	return configuration
}

