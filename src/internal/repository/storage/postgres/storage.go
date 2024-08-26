package postgres

import(
	"myApi/internal/domain/user"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"time"
)

const (
    host     = "localhost"
    port     = 5432
    userName = "slavik"
    password = "1234"
    dbname   = "BDservice"
)

type UserRepository struct {
	db *sql.DB
}

func New() (*UserRepository, error) {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, userName, password, dbname)
	
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal()
	}
	resRepo := &UserRepository{
		db : db,
	}
	return resRepo, err
}
 
func (s *UserRepository) CreateEmployee(newUser *user.User) error{
	err := s.db.Ping()
    if err != nil {
        log.Fatal("Ping error:", err)
    }
	_, err = s.db.Exec(
		"INSERT INTO users (first_name, last_name, date_of_birth) VALUES ($1, $2, $3)",
		newUser.Name, newUser.SurName, newUser.BirthDay)
	return err
}

func (s *UserRepository)GetAllEmployees() (*user.Employees, error) {
	err := s.db.Ping()
    if err != nil {
        log.Fatal("Ping error:", err)
    }
	rows, err := s.db.Query("SELECT id, first_name, last_name, date_of_birth FROM users")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

	employees := user.NewEmployees()
	for rows.Next() {
        var id int
        var name, surName string
		var birthDay time.Time
        if err := rows.Scan(&id, &name, &surName, &birthDay); err != nil {
            log.Println("Error query")
			return nil, err
        }
		employee := user.New(id, name, surName, birthDay)
		employees.AddEmployee(*employee)
    }
	return employees, nil
}

// func (s *UserRepository) ReadUser(id int) user.User{
// 	return nil
// }

// func (s *UserRepository) UpdateUser(updatedUser user.User) {
// 	return nil
// }

// func (s *UserRepository) DeleteUser(id int) {
// 	fmt.Println(id)
// }
