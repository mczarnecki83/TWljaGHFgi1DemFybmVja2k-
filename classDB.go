package main

import (
    "fmt"
    "log"
    "github.com/boltdb/bolt"
)

func deleteData(id_structure []byte,key []byte){
db, err := bolt.Open("osoby.db", 0644, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
 
    err = db.Update(func(tx *bolt.Tx) error {
        bucket, err := tx.CreateBucketIfNotExists(id_structure)
        if err != nil {
            return err
        }

        err = bucket.Delete(key)
        if err != nil {
            return err
        }
        //respondwithJSON(w, http.StatusOK, "saveOK")
        return nil
    })

    if err != nil {
        log.Fatal(err)
    }

}

func setData(id_structure []byte,key []byte,value []byte){
    db, err := bolt.Open("osoby.db", 0644, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()


    err = db.Update(func(tx *bolt.Tx) error {
        bucket, err := tx.CreateBucketIfNotExists(id_structure)
        if err != nil {
            return err
        }

        err = bucket.Put(key, value)
        if err != nil {
            return err
        }
        //respondwithJSON(w, http.StatusOK, "saveOK")
        return nil
    })

    if err != nil {
        log.Fatal(err)
    }

}



func getData(id_structure []byte,key []byte) {

    db, err := bolt.Open("osoby.db", 0644, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    err = db.View(func(tx *bolt.Tx) error {
 
        bucket := tx.Bucket(id_structure)
        if bucket == nil {
           
        }

        val := bucket.Get(key)
        fmt.Println(string(val))

        return  nil
        //return string(val)
    })

   if err != nil {
        log.Fatal(err)
    }
     
}


func getAllKeys(id_structure []byte){
    db, err := bolt.Open("osoby.db", 0644, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    db.View(func(tx *bolt.Tx) error {
	b := tx.Bucket(id_structure)
	c := b.Cursor()
	for k, v := c.First(); k != nil; k, v = c.Next() {
            if v == nil {
            }
		fmt.Printf("%s \n", k)
                
	}
        //respondwithJSON
	return nil
})
}



//func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
//	response, _ := json.Marshal(payload)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(code)
//	w.Write(response)
//}
