package Repository

import (
	"notihub.panos.codes/Database"
	"notihub.panos.codes/Github"
)


func GetPulls() []Pull {
	db := Database.Get()
	q := "select * from pulls"

	rows, err := db.Query(q)
	if err != nil {
		panic(err)
	}

	var pulls []Pull
	for rows.Next() {
		var p Pull

		rows.Scan(&p.ID, &p.RemoteId, &p.Url, &p.Title, &p.Author, &p.CreatedAt)

		pulls = append(pulls, p)
	}

	return pulls
}

func BulkInsertPullsIfNotExist(pulls []Github.Pull) {
	db := Database.Get()

	q := `insert into pulls (remote_id, url, title, author)
	      values (?,?,?,?)`
	statement, _ := db.Prepare(q)

	for _, pull := range pulls {
		statement.Exec(pull.ID, pull.Url, pull.Title, pull.Author.Handle)
	}

	defer statement.Close()
}
