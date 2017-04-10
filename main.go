package main
// +build ignore
import (
	"gopkg.in/mgo.v2"
	"fmt"
	//"gopkg.in/mgo.v2/bson"
)

type BasicDetails struct {
	coll *mgo.Collection
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := mgo.Database{}
	db.Name = "test"
	db.Session = session

	collections, _ := db.CollectionNames()

	//collection := db.C(collections[0])

	//collection.


	fmt.Println(collections)
	//mux := goji.NewMux()
	//http.ListenAndServe("google.com", )
}
