package main

import (
	"go-mongo/delivery"
)

// const uri = "mongodb://localhost:27017"

func main() {
	// credential := options.Credential{
	// 	Username: "jack",
	// 	Password: "12345678",
	// }

	// clientOptions := options.Client()
	// clientOptions.ApplyURI(uri).SetAuth(credential)

	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// // coll := client.Database("db_enigma").Collection("students")
	// collProd := client.Database("db_enigma").Collection("products")
	// // const layout = "2006-01-02"
	// // dt, _ := time.Parse(layout, "2022-01-05")
	// // newStudent := entity.Student{
	// // 	Id:       primitive.NewObjectID(),
	// // 	Name:     "Ari",
	// // 	Gender:   "M",
	// // 	Age:      25,
	// // 	JoinDate: dt,
	// // 	IdCard:   "304",
	// // 	Senior:   false,
	// // }
	// // newId, err := coll.InsertOne(ctx, newStudent)
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// // fmt.Println("id", newId)
	// // menu.InsertOneStudent(ctx, coll, newStudent)
	// // menu.FindAllStudent(ctx, coll)
	// // menu.FindStudentByGenderAndAge(ctx, coll, "M", 25)
	// // menu.CountStudent(ctx, coll)
	// // menu.CountStudentsByAge(ctx, coll, 25)
	// // menu.FindAllStudentByStructGenderAndAge(ctx, coll, "M", 25)
	// menu.CountProductByCategory(ctx, collProd, "food")

	//mdb, err := db.InitResource()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//ctx, cancel := utils.InitContext()
	//defer cancel()
	//defer func() {
	//	if err = mdb.DB.Client().Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	//
	//if err := mdb.DB.Client().Ping(ctx, readpref.Primary()); err != nil {
	//	panic(err)
	//}
	//fmt.Println("Successfully connected and Pinged")
	//stdRepo := repository.NewStudentRepository(mdb)
	////const layout = "2006-01-02"
	////dt, _ := time.Parse(layout, "2022-01-05")
	////newStudent := model.Student{
	////	Id:       primitive.NewObjectID(),
	////	Name:     "Saifool",
	////	Gender:   "M",
	////	Age:      24,
	////	JoinDate: dt,
	////	IdCard:   "304",
	////	Senior:   false,
	////}
	////std, err := stdRepo.CreateOne(newStudent)
	////if err != nil {
	////	fmt.Print(err)
	////} else {
	////	fmt.Println(std.Id)
	////	fmt.Println(std.Name)
	////}
	//stdRepo.GetWithPage(3, 1)
	var server delivery.Routes
	server.StartGin()
}
