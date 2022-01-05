package menu

import (
	"context"
	"fmt"
	"go-mongo/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOneStudent(ctx context.Context, coll *mongo.Collection, newStudent model.Student) {
	newId, err := coll.InsertOne(ctx, newStudent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID baru: ", (*newId).InsertedID)
}

func FindAllStudent(ctx context.Context, coll *mongo.Collection) {
	var results []bson.M
	allDocumentStudentCursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer allDocumentStudentCursor.Close(ctx)
	err = allDocumentStudentCursor.All(ctx, &results)
	if err != nil {
		log.Fatal(err)
	}
	for _, doc := range results {
		fmt.Printf("_id: %v, name: %v, age: %v \n", doc["_id"], doc["name"], doc["age"])
	}
}

func FindStudentByGenderAndAge(ctx context.Context, coll *mongo.Collection, gender string, age int) {
	filterGenderAndAge := bson.D{
		{
			"$and", bson.A{
				bson.D{
					{"gender", gender},
					{"age", age},
				},
			},
		},
	}
	projection := bson.D{
		{"_id", 0},
		{"name", 1},
	}

	findOpts := options.Find().SetProjection(projection)
	var result []bson.M
	resultCursor, err := coll.Find(ctx, filterGenderAndAge, findOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer resultCursor.Close(ctx)

	err = resultCursor.All(ctx, &result)
	if err != nil {
		log.Fatal(err)
	}
	for _, doc := range result {
		fmt.Printf("_id: %v, name: %v, age: %v \n", doc["_id"], doc["name"], doc["age"])
	}
}

func CountStudent(ctx context.Context, coll *mongo.Collection) {
	doc := bson.D{}
	number, err := coll.CountDocuments(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("number of students", number)
}

func CountStudentsByAge(ctx context.Context, coll *mongo.Collection, age int) {
	doc := bson.D{
		{"age", age},
	}
	number, err := coll.CountDocuments(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("number of students by age", number)
}

func FindAllStudentByStructGenderAndAge(ctx context.Context, coll *mongo.Collection, gender string, age int) {
	filterGenderAndAge := bson.D{
		{
			"$and", bson.A{
				bson.D{
					{"gender", gender},
					{"age", age},
				},
			},
		},
	}
	projection := bson.D{
		{"_id", 0},
		{"name", 1},
	}

	findOpts := options.Find().SetProjection(projection)
	result := make([]*model.Student, 0)
	resultCursor, err := coll.Find(ctx, filterGenderAndAge, findOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer resultCursor.Close(ctx)

	for resultCursor.Next(ctx) {
		var row model.Student
		err := resultCursor.Decode(&row)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &row)
	}

	for _, doc := range result {
		fmt.Printf("name: %v \n", doc.Name)
	}
}

func CountProductByCategory(ctx context.Context, productColl *mongo.Collection, category string) {
	matchStage := bson.D{{"$match", bson.D{{"category", category}}}}
	groupStage := bson.D{{"$group", bson.D{
		{"_id", "$category"},
		{"total", bson.D{{"$sum", 1}}},
	}}}

	aggCursor, err := productColl.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		log.Fatal(err)
	}
	defer aggCursor.Close(ctx)

	var results []bson.M
	if err = aggCursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	for _, doc := range results {
		fmt.Println()
		fmt.Printf("Group: %v, Total: %v \n", doc["_id"], doc["total"])
	}
}
