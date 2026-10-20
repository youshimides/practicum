package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	euro float64 = 94.04
	usd  float64 = 83.92
)

func rubToUsd(rub float64) float64  { return rub / usd }
func rubToEuro(rub float64) float64 { return rub / euro }
func usdToEuro(d float64) float64   { return rubToEuro(d * usd) }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>Hello, World!</h1>")

		if r.Method == "POST" {
			r.ParseForm()

			if r.FormValue("action") == "random" {
				fmt.Fprintln(w, "<h2>Random number:", rand.Intn(100)+1, "</h2>")
			}

			if r.FormValue("action") == "convert" {
				rub, err := strconv.ParseFloat(r.FormValue("rub"), 64)

				if err == nil {
					fmt.Fprintf(w, "<h2>USD: %.2f</h2>", rubToUsd(rub))
					fmt.Fprintf(w, "<h2>EUR: %.2f</h2>", rubToEuro(rub))
				} else {
					fmt.Fprintln(w, "<p>Введите правильную сумму</p>")
				}
			}
		}

		fmt.Fprintln(w, `
			<h2>Random generator</h2>
			<form method="POST">
				<button name="action" value="random">Generate number</button>
			</form>

			<h2>Currency converter</h2>
			<form method="POST">
				<input type="number" name="rub" step="any" placeholder="Сумма в рублях" required>
				<button name="action" value="convert">Конвертировать</button>
			</form>
		`)
	})

	http.ListenAndServe(":8080", nil)
}