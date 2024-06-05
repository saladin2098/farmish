package managers

import (
	"database/sql"
	"farmish/models"
	"fmt"
	"strconv"
	"strings"
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

func (p *ProvisionRepo) CreateProvision(provision models.BodyProvision, id int) (*models.CreateProvision, error) {

	query := "insert into provision (id, type, animal_type, quantity) values ($1, $2, $3, $4) returning *"
	_, err := p.Conn.Exec(query, id, provision.Type, provision.AnimalType, provision.Quantity)
	if err != nil {
		return nil, err
	}
	var cProvision models.CreateProvision
	var createdPrv models.CreateProvision
	err = p.Conn.QueryRow(query,
		cProvision.ID,
		cProvision.Type,
		cProvision.AnimalType,
		cProvision.Quantity).Scan(
		&createdPrv.ID,
		&createdPrv.Type,
		&createdPrv.AnimalType,
		&createdPrv.Quantity,
	)
	if err != nil {
		return nil, err
	}

	return &createdPrv, nil
}

func (p *ProvisionRepo) GetAllProvision(filter models.Filter) (*models.GetAllProvisions, error) {
	var provisions []models.GetProvision

	var limit int

	p.Conn.QueryRow("select count(1) from provision").Scan(&limit)

	if filter.Limit == 0 {
		filter.Limit = limit
	}
	query := "select id, type, quantity from provision LIMIT $1 OFFSET $2"
	rows, err := p.Conn.Query(query, filter.Limit, filter.OFFSET)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var i int16
	for rows.Next() {
		var pr models.GetProvision
		err = rows.Scan(
			&pr.ID,
			&pr.Type,
			&pr.Quantity,
		)
		if err != nil {
			return nil, err
		}
		provisions = append(provisions, pr)
		i++
	}

	return &models.GetAllProvisions{Count: i, Provisions: &provisions}, err
}

func (p *ProvisionRepo) GetProvision(id int, typ, animal_type string, quantity float64) (*models.GetProvision, error) {
	query := `select id, type, quantity from provision where `
	var args []interface{}
	var conditions []string

	if id != 0 {
		args = append(args, id)
		conditions = append(conditions, "id = $"+strconv.Itoa(len(args)))
	}
	if typ != "" {
		args = append(args, typ)
		conditions = append(conditions, "type = $"+strconv.Itoa(len(args)))
	}
	if animal_type != "" {
		args = append(args, animal_type)
		conditions = append(conditions, "animal_type = $"+strconv.Itoa(len(args)))
	}
	if quantity != 0 {
		args = append(args, quantity)
		conditions = append(conditions, "quantity <= $"+strconv.FormatFloat(quantity, 'f', 10, 32))
	}

	if len(args) > 1 {
		query += " " + strings.Join(conditions, " and ")
	} else {
		query += " " + conditions[0]
	}

	var provision models.GetProvision
	err := p.Conn.QueryRow(query, args...).
		Scan(
			&provision.ID,
			&provision.Type,
			&provision.Quantity,
		)

	return &provision, err
}

func (p *ProvisionRepo) UpdateProvision(provision *models.UpdateProvision, id int) error {
	query := "update provision set type = $1, quantity = $2,  updated_at=now() where id = $3"

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

	err := p.Conn.QueryRow(query, animal_type).Scan(&quantity)
	if err != nil {
		return false, err
	}
	query = "select quantity from provision where animal_type=$1"
	err = p.Conn.QueryRow(query, animal_type).Scan(&provision.Quantity)
	if err != nil {
		return false, err
	}

	if quantity*18 > provision.Quantity {
		return false, nil
	}

	return true, nil
}
