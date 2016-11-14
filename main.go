package main

import (
	_ "database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	//"time",
	"fmt"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "31780", "localost", "testbase"))
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=solox2 sslmode=disable password=31780")
	if err != nil {
		fmt.Printf("Database opening error -->%v\n", err)
		//panic("Database error")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

}









//db.AutoMigrate(&User{})
//
//db.AutoMigrate(&User{}, &Product{}, &Order{})
//
//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
//
//db.HasTable(&User{})
//
//db.CreateTable(&User{})
//
//// will append "ENGINE=InnoDB" to the SQL statement when creating table `users`
//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})


//const DB_CONNECT_STRING =
//	"host=localhost users=postgres dbname=Solox sslmode=disable password=31780"
//
//func main() {
//
//	db, err := sql.Open("postgres", DB_CONNECT_STRING)
//	defer db.Close()
//
//	if err != nil {
//		fmt.Printf("Database opening error -->%v\n", err)
//		panic("Database error")
//	}
//
//	init_database(&db)
//	make_insertion(&db)
//}
//
//func init_database(pdb **sql.DB) {
//
//	db := *pdb
//
//	init_db_strings := []string{
//		"DROP SCHEMA IF EXISTS sb CASCADE;",
//		"CREATE SCHEMA sb;",
//		//be careful - next multiline string is quoted by backquote symbol
//		`CREATE TABLE sb.test_data(
//		 id serial,
//		 device_id integer not null,
//		 parameter_id integer not null,
//		 value varchar(100),
//		 event_ctime timestamp default current_timestamp,
//		 constraint id_pk primary key (id));`}
//
//	for _, qstr := range init_db_strings {
//		_, err := db.Exec(qstr)
//
//		if err != nil {
//			fmt.Printf("Database init error -->%v\n", err)
//			panic("Query error")
//		}
//	}
//	fmt.Println("Database rebuilded successfully")
//}
//
///*-----------------------------------------------------------------------------*/
//func make_insertion(pdb **sql.DB) {
//
//	db := *pdb
//	const TEST_NUMBER = 400000
//
//	// backquotes for next multiline string
//	const INSERT_QUERY = `insert into sb.test_data(device_id, parameter_id, value)
//                                  values ($1, $2, $3);`
//
//	insert_query, err := db.Prepare(INSERT_QUERY)
//	defer insert_query.Close()
//
//	if err != nil {
//		fmt.Printf("Query preparation error -->%v\n", err)
//		panic("Test query error")
//	}
//
//	t1 := time.Now()
//
//	for i := 0; i < TEST_NUMBER; i++ {
//
//		_, err = insert_query.Exec(i, i, "0")
//
//		if err != nil {
//			fmt.Printf("Query execution error -->%v\n", err)
//			panic("Error")
//		}
//	}
//
//	t2 := time.Since(t1)
//
//	fmt.Printf("%v queries are executed for %v seconds (%v per second)\n",
//		TEST_NUMBER, t2.Seconds(), TEST_NUMBER / t2.Seconds())
//
//	 //do not forget clean up after work )
//	_, err = db.Query("TRUNCATE sb.test_data;")
//}
