package maindb

import (
	pioi "database/sql"
	"fmt"

	khg "cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

func Check() {
	khg.RegisterDriver("ghitnj")
	pioi.Open("pgxv4", "huhfdih")
	fmt.Println("Hello")
}
