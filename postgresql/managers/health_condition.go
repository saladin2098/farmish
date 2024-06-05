package managers

import "database/sql"

type HealthConditionRepo struct {
	Conn *sql.DB
}

func NewHealthConditionRepo(conn *sql.DB) *HealthConditionRepo {
	return &HealthConditionRepo{Conn: conn}
}

func (hr *HealthConditionRepo) GetAllHealthConditionIds() ([]int, error) {
	query := "SELECT id FROM health_conditions"
	rows, err := hr.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}
