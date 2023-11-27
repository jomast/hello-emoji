package main

import (
	"fmt"
	"html"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		output := ""
		for i := 0; i < count(r); i++ {
			output += randomEmoji()
		}
		fmt.Fprintf(w, output)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	ports := portToListenOn()
	wg := sync.WaitGroup{}
	wg.Add(len(ports))
	for _, port := range ports {
		fmt.Printf("listening on port %d...\n", port)
		go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}
	wg.Wait()
}

func randomEmoji() string {
	// http://apps.timwhitlock.info/emoji/tables/unicode
	emoji := [][]int{
		{128513, 128591}, // Emoticons icons
		{128640, 128704}, // Transport and map symbols
	}
	r := emoji[rand.Int()%len(emoji)]
	min := r[0]
	max := r[1]
	n := rand.Intn(max-min+1) + min
	return html.UnescapeString("&#" + strconv.Itoa(n) + ";")
}

func count(r *http.Request) (count int) {
	// user requested count overrides all
	rawCount := r.FormValue("count")
	if rawCount != "" {
		val, err := strconv.Atoi(rawCount)
		if err == nil {
			count = val
		}
		return
	}

	// environment supplied count
	if rawCount, ok := os.LookupEnv("COUNT"); ok {
		val, err := strconv.Atoi(rawCount)
		if err == nil {
			count = val
			return
		}
	}

	// default
	return 1
}

func portToListenOn() []int {
	// accepts comma separated numbers
	if portEnv, ok := os.LookupEnv("PORTS"); ok {
		rawPorts := strings.Split(portEnv, ",")
		output := make([]int, len(rawPorts))

		for i := range rawPorts {
			val, err := strconv.Atoi(strings.TrimSpace(rawPorts[i]))
			if err == nil {
				output[i] = val
			}
		}

		if len(output) > 0 {
			return output
		}
	}

	return []int{80}
}
