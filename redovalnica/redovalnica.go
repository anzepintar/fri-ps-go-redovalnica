// Paket redovalnica omogoča upravljanje z ocenami študentov.
package redovalnica

import "fmt"

// Student predstavlja študenta z osebnimi podatki in ocenami.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// Config hrani nastavitve za upravljanje ocen.
type Config struct {
	MinOcena int // Najnižja možna ocena
	MaxOcena int // Najvišja možna ocena
	StOcen   int // Minimalno število ocen
}

// DefaultConfig vrne privzete nastavitve (ocene 1-10, min 3 ocene).
func DefaultConfig() Config {
	return Config{
		MinOcena: 1,
		MaxOcena: 10,
		StOcen:   3,
	}
}

// DodajOceno doda oceno študentu z dano vpisno številko.
// Preveri, ali je ocena veljavna glede na nastavitve.
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, cfg Config) {
	if ocena < cfg.MinOcena || ocena > cfg.MaxOcena {
		fmt.Printf("Ocena mora biti med %d in %d.\n", cfg.MinOcena, cfg.MaxOcena)
		return
	}
	if student, found := studenti[vpisnaStevilka]; found {
		student.Ocene = append(student.Ocene, ocena)
		studenti[vpisnaStevilka] = student
	} else {
		fmt.Println("Študenta z vpisno številko", vpisnaStevilka, "ni na seznamu.")
	}
}

// povprecje izračuna povprečno oceno študenta.
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	if student, found := studenti[vpisnaStevilka]; found {
		if len(student.Ocene) == 0 {
			return 0.0
		}
		sum := 0
		for _, ocena := range student.Ocene {
			sum += ocena
		}
		return float64(sum) / float64(len(student.Ocene))
	}
	fmt.Println("Študenta z vpisno številko", vpisnaStevilka, "ni na seznamu.")
	return -1.0
}

// IzpisVsehOcen izpiše seznam vseh študentov in njihovih ocen.
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpis, student := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpis, student.Ime, student.Priimek, student.Ocene)
	}
}

// IzpisiKoncniUspeh izpiše končni uspeh za vse študente.
// Uspeh se določi na podlagi povprečja in števila ocen.
func IzpisiKoncniUspeh(studenti map[string]Student, cfg Config) {
	fmt.Println("KONČNI USPEH:")
	for vpis, student := range studenti {
		if len(student.Ocene) < cfg.StOcen {
			fmt.Printf("%s %s: premalo ocen (%d/%d) -> Premalo ocen\n",
				student.Ime, student.Priimek, len(student.Ocene), cfg.StOcen)
			continue
		}

		povp := povprecje(studenti, vpis)
		uspeh := ""
		switch {
		case povp >= 9.0:
			uspeh = "Odličen študent!"
		case povp >= 6.0:
			uspeh = "Povprečen študent"
		default:
			uspeh = "Neuspešen študent"
		}
		fmt.Printf("%s %s: povprečna ocena %.1f -> %s\n", student.Ime, student.Priimek, povp, uspeh)
	}
}
