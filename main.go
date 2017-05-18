package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	errorEmptyString = errors.New("Unwilling to print an emty string")
)

type Geometry interface {
	Area() int
}

type Triangle struct {
	Base   int
	Height int
}

func (t Triangle) Area() int {
	return (t.Base * t.Height) / 2
}

func PrintArea(g Geometry) {
	area := g.Area()
	fmt.Println(area)
}

func makeRandoms(c chan int) {
	for {
		c <- rand.Intn(1000)
	}
}

func makeID(c chan int) {
	var id int
	id = 0
	for {
		c <- id
		id++
	}
}

func emit(wordChannel chan string, done chan bool) {
	defer close(wordChannel)
	words := []string{"The", "quick", "brown", "fox"}
	i := 0
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}

		case <-done:
			fmt.Printf("GOT DONE\n")
			close(done)
			return

		case <-t.C:
			return
		}
	}
}

func getPage(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	return len(body), nil
}

func getter(url string, size chan string) {
	length, err := getPage(url)
	if err == nil {
		size <- fmt.Sprintf("%s has legnt %d\n", url, length)
	}
}

var (
	running int64 = 0
)

func work() {
	atomic.AddInt64(&running, 1)
	fmt.Printf("[%d ", running)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	atomic.AddInt64(&running, -1)
	fmt.Printf("]")
}

func worker(ch chan bool) {
	<-ch
	work()
	ch <- true
}

type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}
	defer resp.Body.Close()
	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}
}

func (w *webPage) isOK() bool {
	return w.err == nil
}

type shuffler interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	/* INTERFACES */
	is := intSlice{1, 2, 3, 4, 5, 6}
	shuffle(is)
	fmt.Printf("%q\n", is)
	/* TYPES */
	// w := &webPage{url: "http://www.google.com/"}
	// w.get()
	// if w.isOK() {
	// 	fmt.Printf("URL: %s Error: %s Legth: %d\n", w.url, w.err, len(w.body))
	// } else {
	// 	fmt.Printf("Something went wrong")
	// }

	/* BUFFERED CHANNLES */
	//ch := make(chan bool, 20)
	//for i := 0; i < 1000; i++ {
	//	go worker(ch)
	//}
	//for i := 0; i < cap(ch); i++ {
	//	ch <- true
	//}
	//time.Sleep(30 * time.Second)

	/* READER & WRITERS */
	//urls := []string{"http://www.google.com/", "http://www.yahoo.com", "https://portal.selcukluholding.com.tr", "http://siparis.ilko.com.tr"}
	//size := make(chan string)
	//for _, url := range urls {
	//	go getter(url, size)
	//}
	//for i := 0; i < len(urls); i++ {
	//	fmt.Printf(<-size)
	//}

	/* CHANNEL USAGE-2 */
	//wordCh := make(chan string)
	//doneCh := make(chan bool)
	//go emit(wordCh, doneCh)
	//for word := range wordCh {
	//	fmt.Printf("%s ", word)
	//}

	/* CHANNEL USAGE-1 */
	//idChan := make(chan int)
	//go makeID(idChan)
	//fmt.Printf("%d\n", <-idChan)
	//fmt.Printf("%d\n", <-idChan)

	/* RANDOM NUMBER GENERATOR */
	//randoms := make(chan int)
	//go makeRandoms(randoms)
	//for n := range randoms {
	//	fmt.Printf("%d ", n)
	//}

	// STRUCT GEOMETRIC CALCULATION
	//t := Triangle{0, 5}
	//PrintArea(t)

	/* SLICE USAGE */
	// newSlice := make([]float32, 5, 10) // 5 elemanli ve max 10 kapasiteli slice tanimi
	// newSlice[0] = 1.32
	// newSlice = append(newSlice, math.Pi, 2.277) // Slice'in sonuna eleman ekledi
	// fmt.Println(newSlice)

	// newSlice1 := []int{3, 5, 1, 7, 8} // 5 elemanli ve kapasiteli bir slice
	// newSlice2 := make([]int, 3)       // 3 elemanli ve kapasiteli bir slice
	// copy(newSlice2, newSlice1)
	// fmt.Println(newSlice1)
	// fmt.Println(newSlice2)

	/* MAP USAGE */
	// dayMonths := make(map[string]int)
	// dayMonths["Jan"] = 31
	// dayMonths["Feb"] = 28
	// dayMonths["Mar"] = 31
	// dayMonths["Apr"] = 30
	// dayMonths["May"] = 31
	// dayMonths["Jun"] = 30
	// dayMonths["Jul"] = 31
	// dayMonths["Aug"] = 31
	// dayMonths["Sep"] = 30
	// dayMonths["Oct"] = 31
	// dayMonths["Nov"] = 30
	// dayMonths["Dec"] = 31
	//
	// for month, day := range dayMonths {
	// 	//fmt.Printf("%d\n", days)
	// 	if month == "Feb" {
	// 		fmt.Printf("%d days in February\n", day)
	// 	}
	// }

	/* CUSTOM ERROR USAGE */
	// if err := printer(""); err != nil {
	// 	if err == errorEmptyString {
	// 		fmt.Printf("You tried to print an empty string!\n")
	// 	} else {
	// 		fmt.Printf("printer failed: %s\n", err)
	// 	}
	// 	os.Exit(1)
	// }
}
