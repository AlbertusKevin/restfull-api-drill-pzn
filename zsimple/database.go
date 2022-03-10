package zsimple

// ini sebagai objek yang dibutuhkan class db repository
type Database struct {
	Name string
}

// // untuk database, tersedia 2 provider yang keduanya akan digunakan db repository. Kedua provider ini sama-sama return tipe struct Database
// func NewDatabasePostgre() *Database {
// 	return &Database{Name: "PostgrSQL"}
// }
// func NewDatabaseMongoDB() *Database {
// 	return &Database{Name: "MongoDB"}
// }

// // class ini yang membutuhkan dependency dari Database, membutuhkan 2 data yang sama-sama bertipe Database
// type DatabaseRepository struct {
// 	DBPostgre *Database
// 	DBMongoDB *Database
// }

// // provider dari repository, menerima 2 parameter, namun keduanya bertipe Database. Ini akan membuat bingung google wire ketika harus memasukkan value parameter ke variable yang mana.
// func NewDatabaseRepository(postgreSQL *Database, mongoDB *Database) *DatabaseRepository {
// 	return &DatabaseRepository{DBPostgre: postgreSQL, DBMongoDB: mongoDB}
// }

// buat type aliasnya, untuk membedakan tipe data meski keduanya sama-sama Database
type DatabaseMongoDB Database
type DatabasePostgre Database

// return datanya sesuai dengan type aliasnya
func NewDatabasePostgre() *DatabasePostgre {
	return &DatabasePostgre{Name: "PostgreSQL"}
}
func NewDatabaseMongoDB() *DatabaseMongoDB {
	return &DatabaseMongoDB{Name: "MongoDB"}
}

type DatabaseRepository struct {
	DBPostgre *DatabasePostgre
	DBMongoDB *DatabaseMongoDB
}

func NewDatabaseRepository(
	postgreSQL *DatabasePostgre, 
	mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DBPostgre: postgreSQL, DBMongoDB: mongoDB}
}