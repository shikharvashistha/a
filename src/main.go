//Certain parts of code have to be blocked while we are using the code.
//If the execution time of main function is 1 second and we call a function which takes
//more than 1 second, then the code won't be executed ie the function, so here we need a wait group.
//Wait groups are bascially used to wait for the completion of a set of go routines.
//it is a counter which is incremented/decremented when a go routine is started and decremented when the go routine is finished.
/*
package main

import (
    "fmt"
    "sync"
)

func myFunc(waitgroup *sync.WaitGroup) {
    fmt.Println("Inside my goroutine")
    waitgroup.Done()
}

func main() {
    fmt.Println("Hello World")

    var waitgroup sync.WaitGroup
    waitgroup.Add(1)
    go myFunc(&waitgroup)
    waitgroup.Wait()//to ensure that our main function dosn't pass the point untill our internal counter has been decremented to 0.
	//or we can say that it blocks untill the counter is 0.

    fmt.Println("Finished Execution")
}
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"http://www.google.com",
	"http://www.facebook.com",
	"http://www.twitter.com",
	"http://www.yahoo.com",
	"http://www.linkedin.com",
	"http://www.github.com",
	"http://www.stackoverflow.com",
	"http://www.reddit.com",
	"http://www.youtube.com",
	"http://www.amazon.com",
	"http://www.microsoft.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request){
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			} else {
				fmt.Fprintf(w, "%+v\n", resp)//resp.Status to return the status code.
			}
			wg.Done()
		}(url)
	}
	wg.Wait()
	//To block the execution of our fetchStatus function untill all the go routines are completed.
}
func main(){
	fmt.Println(" HELLO ")
	http.HandleFunc("/", fetchStatus)// / is the base path
	//The request is coming to our server and then it is mapped against / endpoint 
	//and triggerd the fetchStatus function. Where we've created a go routine for each url.
	//but our fetchStatus function is not returning anything as it has terminated before 
	//the http responses have been recived.
	log.Fatal(http.ListenAndServe(":8080", nil))
	//visit http://localhost:8080/ for responses.
}