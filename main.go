//go:generate swagger generate spec

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func parseOption(r *http.Request, key string) (string, error) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return "", err
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", err
	}

	if len(m[key]) > 0 {
		return m[key][0], nil
	}

	return "", nil
}

func main() {
	fmt.Println("starting the server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Alive")
	})

	rand.Seed(time.Now().UTC().UnixNano())

	// swagger:route GET /ping ping
	//
	// Pings the service to check if it's alive.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I really would prefer to play darts..."))
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	})

	http.HandleFunc("/fail", func(w http.ResponseWriter, r *http.Request) {
		panic("panicing")
	})

	http.HandleFunc("/terminate", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(1)
	})

	http.HandleFunc("/code", func(w http.ResponseWriter, r *http.Request) {
		code, err := parseOption(r, "code")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not parse optional paramter"))
		}

		retCode := 200
		if code != "" {
			retCode, _ = strconv.Atoi(code)
		}

		fmt.Println("returning status code of ", retCode)
		w.WriteHeader(retCode)
		w.Write([]byte("Playing a little bit of roulette."))
	})

	http.HandleFunc("/random/error", func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(2) == 1 {
			intCode := 500 + rand.Intn(11)
			fmt.Println("using random code of ", intCode)

			w.WriteHeader(intCode)
			w.Write([]byte("I feel like this should have worked"))
		} else {
			intCode := 400 + rand.Intn(51)
			fmt.Println("using random code of ", intCode)

			w.WriteHeader(intCode)
			w.Write([]byte("I feel like this should have worked"))
		}
	})

	http.HandleFunc("/maybe/error", func(w http.ResponseWriter, r *http.Request) {
		rate, err := parseOption(r, "rate")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not parse optional paramter"))
		}

		intRate := 50
		if rate != "" {
			intRate, _ = strconv.Atoi(rate)
		}

		fmt.Println("using random rate of ", intRate, "%")
		if rand.Intn(100) < intRate {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("No where did I put that function...?"))
		}

		w.Write([]byte("Oh, right - THERE's the return statement!"))
	})

	http.HandleFunc("/maybe/fail", func(w http.ResponseWriter, r *http.Request) {
		rate, err := parseOption(r, "rate")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not parse optional paramter"))
		}

		intRate := 50
		if rate != "" {
			intRate, _ = strconv.Atoi(rate)
		}

		fmt.Println("using random rate of ", intRate, "%")
		if rand.Intn(100) < intRate {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("BEES!!!!"))
		}

		w.Write([]byte("I've got my towel, all is well."))
	})

	http.HandleFunc("/maybe/terminate", func(w http.ResponseWriter, r *http.Request) {
		rate, err := parseOption(r, "rate")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not parse optional paramter"))
		}

		intRate := 1
		if rate != "" {
			intRate, _ = strconv.Atoi(rate)
		}

		fmt.Sprintf("using random rate of %d%%", intRate)
		if rand.Intn(100) < intRate {
			w.WriteHeader(http.StatusInternalServerError)
			os.Exit(1)
		} else {
			w.Write([]byte("I'll be back. (Puts sunglasses on dramatically)"))
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
