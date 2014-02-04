package genmai_test

import (
	"fmt"
	"log"

	"github.com/naoina/genmai"
)

func ExampleDB_Select_all() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT * FROM "test_model";
	if err := db.Select(&results); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_where() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT * FROM "test_model" WHERE "id" = 1;
	if err := db.Select(&results, db.Where("id", "=", 1)); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_whereAnd() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT * FROM "test_model" WHERE "id" = 1 AND "name" = "alice";
	if err := db.Select(&results, db.Where("id", "=", 1).And("name", "=", "alice")); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_whereNested() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT * FROM "test_model" WHERE "id" = 1 OR ("name" = "alice" AND "addr" != "Tokyo");
	if err := db.Select(&results, db.Where("id", "=", 1).Or(db.Where("name", "=", "alice").And("addr", "!=", "Tokyo"))); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_in() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT * FROM "test_model" WHERE "id" IN (1, 3, 5);
	if err := db.Select(&results, db.Where("id").In(1, 3, 5)); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_like() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT * FROM "test_model" WHERE "name" LIKE "alice%";
	if err := db.Select(&results, db.Where("name").Like("alice%")); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_between() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT * FROM "test_model" WHERE "id" BETWEEN 3 AND 5;
	if err := db.Select(&results, db.Where("id").Between(3, 5)); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_orderBy() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT * FROM "test_model" ORDER BY "name" DESC;
	if err := db.Select(&results, db.OrderBy("name", genmai.DESC)); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_limit() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT * FROM "test_model" LIMIT 3;
	if err := db.Select(&results, db.Limit(3)); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_offset() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT * FROM "test_model" OFFSET 10;
	if err := db.Select(&results, db.Offset(10)); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_distinct() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT DISTINCT "name" FROM "test_model";
	if err := db.Select(&results, db.Distinct("name")); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_count() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var result int64
	// SELECT COUNT(*) FROM "test_model";
	if err := db.Select(&result, db.Count(), db.From(TestModel{})); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func ExampleDB_Select_countDistinct() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var result int64
	// SELECT COUNT(DISTINCT "name") FROM "test_model";
	if err := db.Select(&result, db.Count(db.Distinct("name")), db.From(TestModel{})); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func ExampleDB_Select_columns() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT "id", "name" FROM "test_model";
	if err := db.Select(&results, []string{"id", "name"}); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func ExampleDB_Select_complex() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	var results []TestModel
	// SELECT "name" FROM "test_model"
	//   WHERE "name" LIKE "%alice%" OR ("id" > 100 AND "id" < 200) OR ("id" BETWEEN 700 AND 1000)
	//   ORDER BY "id" ASC LIMIT 2 OFFSET 5
	if err := db.Select(&results, "name", db.Where("name").
		Like("%alice%").
		Or(db.Where("id", ">", 100).And("id", "<", 200)).
		Or(db.Where("id").Between(700, 1000)).
		Limit(2).Offset(5).OrderBy("id", genmai.ASC),
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}