package elements

import (
	"fmt"
)

// QuantumNumbers represents the four quantum numbers of an electron
type QuantumNumbers struct {
	n   int     // Principal quantum number
	l   int     // Azimuthal quantum number
	m int     // Magnetic quantum number
	s float64 // Spin quantum number
}

// ToString returns a string representation of the quantum numbers
func (q QuantumNumbers) ToString() string {
	return fmt.Sprintf("n: %d, l: %d, m: %d, s: %f", q.n, q.l, q.m, q.s)
}

// GetQuantumNumbers retrieves the quantum numbers for the electron at the specified index in the configuration
func GetQuantumNumbers(atomicNumber int, electronIndex int) (QuantumNumbers, error) {
	configuration := GenerateElectronConfiguration(atomicNumber)

	if electronIndex < 0 || electronIndex >= atomicNumber {
		return QuantumNumbers{}, fmt.Errorf("electron index out of bounds")
	}

	electronCount := 0

	for _, subshell := range configuration.Subshells {
		for i := 0; i < subshell.Electrons; i++ {
			if electronCount == electronIndex {
				n := getPrincipalQuantumNumber(subshell.Orbital)
				l := getAzimuthalQuantumNumber(subshell.Orbital)
				m := getMagneticQuantumNumber(l, i)
				s := getSpinQuantumNumber(i)

				return QuantumNumbers{n, l, m, s}, nil
			}
			electronCount++
		}
	}

	return QuantumNumbers{}, fmt.Errorf("quantum numbers not found for electron index %d", electronIndex)
}


// getPrincipalQuantumNumber returns the principal quantum number for a given orbital
func getPrincipalQuantumNumber(orbital string) int {
	n := 0
	fmt.Sscanf(orbital, "%d", &n)
	return n
}

// getAzimuthalQuantumNumber returns the azimuthal quantum number for a given orbital
func getAzimuthalQuantumNumber(orbital string) int {
	switch orbital[len(orbital)-1] {
	case 's':
		return 0
	case 'p':
		return 1
	case 'd':
		return 2
	case 'f':
		return 3
	default:
		return -1 // Invalid orbital
	}
}

// getMagneticQuantumNumber returns the magnetic quantum number for a given l and the electron index
func getMagneticQuantumNumber(l int, electronIndex int) int {
	if l == 0 { // s subshell
		return 0
	}
	if l == 1 { // p subshell
		return (electronIndex % 3) - 1 // -1, 0, +1 for p subshell
	}
	if l == 2 { // d subshell
		return (electronIndex % 5) - 2 // -2, -1, 0, +1, +2 for d subshell
	}
	if l == 3 { // f subshell
		return (electronIndex % 7) - 3 // -3, -2, -1, 0, +1, +2, +3 for f subshell
	}
	return 0 // Should not reach here
}

// getSpinQuantumNumber returns the spin quantum number for an electron based on its index
func getSpinQuantumNumber(electronIndex int) float64 {
	if electronIndex%2 == 0 {
		return +0.5 // +1/2 spin for even indexed electrons
	}
	return -0.5 // -1/2 spin for odd indexed electrons
}
