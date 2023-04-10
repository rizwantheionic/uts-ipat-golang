package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Mahasiswa struct {
	NPM     uint   `json:"npm"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func main() {

	// database connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/latihan_go")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	// database connection

	e := echo.New()

	// Enable CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Service API!")
	})

	// Mahasiswa
	e.GET("/mahasiswa", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM mahasiswa")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var mahasiswa []Mahasiswa
		for res.Next() {
			var m Mahasiswa
			_ = res.Scan(&m.NPM, &m.Name, &m.Phone, &m.Address)
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})

	e.POST("/mahasiswa", func(c echo.Context) error {
		var mahasiswa Mahasiswa
		c.Bind(&mahasiswa)

		sqlStatement := "INSERT INTO mahasiswa (npm,name, phone,address)VALUES (?,?, ?, ?)"
		res, err := db.Query(sqlStatement, mahasiswa.NPM, mahasiswa.Name, mahasiswa.Phone, mahasiswa.Address)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.PUT("/mahasiswa/:npm", func(c echo.Context) error {
		var mahasiswa Mahasiswa
		c.Bind(&mahasiswa)

		sqlStatement := "UPDATE mahasiswa SET name=?,phone=?,address=? WHERE npm=?"
		res, err := db.Query(sqlStatement, mahasiswa.Name, mahasiswa.Phone, mahasiswa.Address, c.Param("npm"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.DELETE("/mahasiswa/:npm", func(c echo.Context) error {
		var mahasiswa Mahasiswa
		c.Bind(&mahasiswa)

		sqlStatement := "DELETE FROM mahasiswa WHERE npm=?"
		res, err := db.Query(sqlStatement, c.Param("npm"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})
	// Mahasiswa

	e.Logger.Fatal(e.Start(":8100"))
}
