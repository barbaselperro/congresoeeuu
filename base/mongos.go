package base

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

// esta funcion sirve para realizar la conexion al servidor mongodb, recibe dos parametros
// de tipo string el nombre de la base y la collection
func conexion(base, collections string) (*mongo.Collection, error) {
	var files, _ = ioutil.ReadFile("base/mongoconex.txt")
	uri := string(files)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	coll := client.Database(base).Collection(collections)
	return coll, err
}

func ObtieneMongo(usuario string) Usuario {
	var miuser Usuario
	colls, _ := conexion("Users", "user")
	var result bson.M
	errors := colls.FindOne(context.TODO(), bson.D{{"user", usuario}}).Decode(&result)
	if errors != nil {
		panic(errors)
	}
	jsonData, _ := json.MarshalIndent(result, "", "    ")
	json.Unmarshal(jsonData, &miuser)
	Senadores()
	return miuser
}

func Senadores() []Senates {
	coll, _ := conexion("congreso", "Senate")
	cursor, _ := coll.Find(context.TODO(), bson.D{{}})
	var result bson.D
	var senadores []Senates
	for cursor.Next(context.TODO()) {
		var senate Senates
		cursor.Decode(&senate)
		senadores = append(senadores, senate)
		if err := cursor.Decode(&result); err != nil {
			panic(err)
		}
	}
	return senadores
}

func House() []Houses {
	coll, _ := conexion("congreso", "House")
	cursor, _ := coll.Find(context.TODO(), bson.D{{}})
	var result bson.D
	var houses []Houses
	for cursor.Next(context.TODO()) {
		var house Houses
		cursor.Decode(&house)
		houses = append(houses, house)
		if err := cursor.Decode(&result); err != nil {
			panic(err)
		}
	}
	return houses
}
