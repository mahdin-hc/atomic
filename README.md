# ATOMIC

A command-line tool to parse chemical formulas, draw the periodic table, and display electron configurations.


![icon](./icon/icon_small.png)

## Features

- Parse chemical formulas and display relevant information, such as molecular mass, charge, and state.
- Draw the periodic table.
- Display electron configurations of the elements in the formula.



## Usage

```bash
atomic [options] <formula>
```

### Options

- `-pt` : Draw the periodic table with the elements involved in the provided formula.
- `-e`  : Show electron configurations of the elements in the provided formula.

### Examples

1. **Parse a formula** and show details:

    ```bash
    atomic NaCl
    ```

    Output:
    ```
	Molecule : NaCl
	Simplify : NaCl
	Name     : Sodium Chloride
	Mass     : 58.443000
	Charge   : 0
    ```

2. **Draw the periodic table** and show the electron configurations:

    ```bash
    atomic -pt -e NaCl
    ```

    Output:
    ```
	H                                                                  He
	Li  Be                                           B   C   N   O   F  Ne
	Na  Mg                                          Al  Si   P   S  Cl  Ar
	K  Ca  Sc  Ti   V  Cr  Mn  Fe  Co  Ni  Cu  Zn  Ga  Ge  As  Se  Br  Kr
	Rb  Sr   Y  Zr  Nb  Mo  Tc  Ru  Rh  Pd  Ag  Cd  In  Sn  Sb  Te   I  Xe
	Cs  Ba      Hf  Ta   W  Re  Os  Ir  Pt  Au  Hg  Tl  Pb  Bi  Po  At  Rn
	Fr  Ra      Rf  Db  Sg  Bh  Hs  Mt  Ds  Rg  Cn  Nh  Fl  Mc  Lv  Ts  Og

		  La  Ce  Pr  Nd  Pm  Sm  Eu  Gd  Tb  Dy  Ho  Er  Tm  Yb  Lu
		  Ac  Th  Pa   U  Np  Pu  Am  Cm  Bk  Cf  Es  Fm  Md  No  Lr

	Na(11): 1s2 2s2 2p6 3s1
	Cl(17): 1s2 2s2 2p6 3s2 3p5
    ```

## Data

- Elements data is loaded from `data/elements.csv`.
- Molecules data is loaded from `data/molecules.csv`.

## Requirements

- Go 1.16 or later

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/mahdin-hc/atomic.git
    ```

2. Build the executable:

    ```bash
    go build -o atomic.exe
    ```

3. Run the tool:

    ```bash
    atomic.exe -e -pt NaCl
    ```