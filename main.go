package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anzepintar/fri-ps-go-redovalnica/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Aplikacija za upravljanje z ocenami študentov",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "stOcen",
				Value:   3,
				Usage:   "Najmanjše število ocen potrebnih za pozitivno oceno",
				Aliases: []string{"n"},
			},
			&cli.IntFlag{
				Name:    "minOcena",
				Value:   1,
				Usage:   "Najmanjša možna ocena",
				Aliases: []string{"min"},
			},
			&cli.IntFlag{
				Name:    "maxOcena",
				Value:   10,
				Usage:   "Največja možna ocena",
				Aliases: []string{"max"},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Preberemo vrednosti stikal
			cfg := redovalnica.Config{
				MinOcena: int(cmd.Int("minOcena")),
				MaxOcena: int(cmd.Int("maxOcena")),
				StOcen:   int(cmd.Int("stOcen")),
			}

			// Preverimo veljavnost konfiguracije
			if cfg.MinOcena >= cfg.MaxOcena {
				return fmt.Errorf("minOcena mora biti manjša od maxOcena")
			}
			if cfg.StOcen < 0 {
				return fmt.Errorf("stOcen mora biti pozitivno število")
			}

			// Ustvarimo študente
			studenti := make(map[string]redovalnica.Student)
			studenti["63210001"] = redovalnica.Student{
				Ime:     "Ana",
				Priimek: "Novak",
				Ocene:   []int{10, 9, 8},
			}
			studenti["63210002"] = redovalnica.Student{
				Ime:     "Boris",
				Priimek: "Kralj",
				Ocene:   []int{6, 7, 5, 8},
			}
			studenti["63210003"] = redovalnica.Student{
				Ime:     "Janez",
				Priimek: "Novak",
				Ocene:   []int{4, 5, 3, 5},
			}

			// Dodajamo ocene
			redovalnica.DodajOceno(studenti, "63210001", 10, cfg)
			redovalnica.DodajOceno(studenti, "63210001", 10, cfg)
			redovalnica.DodajOceno(studenti, "63210001", 10, cfg)

			// Izpišemo redovalnico
			redovalnica.IzpisVsehOcen(studenti)
			fmt.Println()

			// Izpišemo končni uspeh
			redovalnica.IzpisiKoncniUspeh(studenti, cfg)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Napaka: %v\n", err)
		os.Exit(1)
	}
}
