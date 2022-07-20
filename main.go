package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	/*config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load  :", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start Server:", err)
	}*/

	fmt.Println(HashPassword("asd123"))

}
func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %w", err)
	}
	fmt.Println(hashPassword)
	return string(hashPassword), nil
}
