package managers

import (
	"database/sql"
	"farmish/models"
	"fmt"
	"time"
)

type ProvisionRepo struct {
	Conn *sql.DB
}

func NewProvisionRepo(db *sql.DB) *ProvisionRepo {
	return &ProvisionRepo{Conn: db}
}

func (p *ProvisionRepo) GetAllIds() ([]int, error) {
	var ids []int

	query := "SELECT id FROM provision"
	fmt.Println(p)
	fmt.Println(p.Conn)
	rows, err := p.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.
			Scan(
				&id,
			)

		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (p *ProvisionRepo) CreateProvision(provision models.BodyProvision, id int) error {

	query := "insert into provision (id, type, animal_type, quantity) values ($1, $2, $3, $4) returning *"
	_, err := p.Conn.Exec(query, id, provision.Type, provision.AnimalType, provision.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProvisionRepo) GetAllProvision(filter models.Filter, typ, animal_type string, quantity float64) ([]models.GetProvision, error) {
	var provisions []models.GetProvision

	query := "SELECT id, type, quantity FROM provision"
	var args []interface{}
	var conditions []string

	if typ != "" {
		args = append(args, typ)
		conditions = append(conditions, fmt.Sprintf("type = $%d", len(args)))
	}
	if animal_type != "" {
		args = append(args, animal_type)
		conditions = append(conditions, fmt.Sprintf("animal_type = $%d", len(args)))
	}
	if quantity != 0 {
		args = append(args, quantity)
		conditions = append(conditions, fmt.Sprintf("quantity >= $%d", len(args)))
	}

	if len(conditions) > 0 {
		query += " WHERE " + conditions[0]
		for i := 1; i < len(conditions); i++ {
			query += " AND " + conditions[i]
		}
	}

	var limit int

	_ = p.Conn.QueryRow("SELECT COUNT(1) FROM provision WHERE deleted_at=0").Scan(&limit)

	if filter.Limit == 0 {
		filter.Limit = limit
	}
	args = append(args, filter.Limit, filter.OFFSET)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := p.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pr models.GetProvision
		err = rows.Scan(&pr.ID, &pr.Type, &pr.Quantity)
		if err != nil {
			return nil, err
		}
		provisions = append(provisions, pr)
	}

	return provisions, nil
}
func (p *ProvisionRepo) GetProvisionById(id int) (*models.GetProvision, error) {
	res := models.GetProvision{}
	query := `select id, type, quantity from provision where id=$1`

	err := p.Conn.QueryRow(query, id).Scan(
		&res.ID,
		&res.Type,
		&res.Quantity,
	)

	return &res, err
}

func (p *ProvisionRepo) UpdateProvision(provision *models.UpdateProvision, id int) error {
	query := "UPDATE provision SET type = $1, quantity = $2, updated_at = now() WHERE id = $3"

	_, err := p.Conn.Exec(query, provision.Type, provision.Quantity, id)
	return err
}

func (p *ProvisionRepo) DeleteProvision(id int) error {
	now := time.Now().Unix()

	query := "update provision set deleted_at=$1 where id = $2"

	_, err := p.Conn.Exec(query, now, id)

	return err
}

func (p *ProvisionRepo) ProvisionData(animal_type string) (bool, error) {
	var provision models.Provision
	var quantity float64
	query := "select sum(avg_consumption) from animals where animal_type = $1"

	err := p.Conn.QueryRow(query, animal_type).Scan(
		&quantity,
	)
	if err != nil {
		return false, err
	}

	query = "select quantity from provision where animal_type=$1"

	err = p.Conn.QueryRow(query, animal_type).Scan(
		&provision.Quantity,
	)
	if err != nil {
		return false, err
	}

	if quantity*18 > provision.Quantity {
		return false, nil
	}

	return true, nil
}
