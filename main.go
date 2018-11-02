package main
import (
	"net/http"
	"github.com/go-chi/chi"
        "io/ioutil"
         "regexp"
        "fmt"
)

var acceptString = regexp.MustCompile(`^[a-zA-Z0-9]{1,10}$`)

//Ewentualne badanie wielkosci (1MB) naokolo. Na razie nie znalazlem lepszej metody
//        fi, err := os.Stat("tresc.txt")
//	if err != nil {
//	}
//	fmt.Println(fi.Size())




func main() {

var id_structure = []byte("osoby")
  
	r := chi.NewRouter()
	r.Get("/api/objects", func(w http.ResponseWriter, r *http.Request) {
            getAllKeys(id_structure)
	})

	r.Get("/api/objects/{id}", func(w http.ResponseWriter, r *http.Request) {
            key := []byte(chi.URLParam(r, "id"))
            getData(id_structure,key)
           
	})

        r.Put("/api/objects/{id}", func(w http.ResponseWriter, r *http.Request) {
            body, err := ioutil.ReadAll(r.Body)
            if err != nil {
                    panic(err)
                }
                key := []byte(chi.URLParam(r, "id"))
                value := []byte(string(body))
                if acceptString.MatchString(string(key)) {
                        fmt.Println("data ok")
                        setData(id_structure,key,value)
                //respondwithJSON(w, 200s, "Ok")
                 } else {
                        fmt.Println("data error")
                //respondwithJSON(w, 400, "Bad Request")
                 }
	})

        r.Delete("/api/objects/{id}", func(w http.ResponseWriter, r *http.Request) {
                key := []byte(chi.URLParam(r, "id"))
                if acceptString.MatchString(string(key)) {
                        fmt.Println("data ok")
                        deleteData(id_structure,key)
                //respondwithJSON(w, 200s, "Ok")
                 } else {
                        fmt.Println("data error")
                //respondwithJSON(w, 400, "Bad Request")
                 }
	})
 
	http.ListenAndServe(":8080", r)
}
