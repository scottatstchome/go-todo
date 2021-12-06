// database/connect.go
package database

// import (
//     "fmt"
//     "strconv"

//     "github.com/scottatstchome/go-todo/config"

//     "github.com/jinzhu/gorm"
//     _ "github.com/lib/pq"
// )

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
)

func ConnectDB() {
    var err error
    // p := config.Config("DB_PORT")
    // p := "5535"
    // port, err := strconv.ParseUint(p, 10, 32)

    // configData := fmt.Sprintf(
    //     "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    //     config.Config("DB_HOST"),
    //     port,
    //     config.Config("DB_USER"),
    //     config.Config("DB_PASSWORD"),
    //     config.Config("DB_NAME"),
    // )

    configData := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        "go-todo-db",
        "5535",
        "postgres",
        "root",
        "go-todo-db",
    )

    DB, err = gorm.Open(
        "postgres",
        configData,
    )

    if err != nil {
        fmt.Println(
            err.Error(),
        )
        panic("failed to connect database")
    }

    fmt.Println("Connection Opened to Database")
}