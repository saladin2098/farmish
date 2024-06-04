package managers

import (
	"database/sql"
	"farmish/models"
	"strconv"
	"strings"
	"time"
)

type ProvisionRepo struct {
	db *sql.DB
}

func NewProvisionRepo(db *sql.DB) *ProvisionRepo {
	return &ProvisionRepo{db}
}

func (p *ProvisionRepo) GetAllIds() ([]int, error) {
	var ids []int

	query := "SELECT id FROM provision"
	rows, err := p.db.Query(query)
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

func (p *ProvisionRepo) CreateProvision(provision models.CreateProvision) (*models.CreateProvision, error) {

	query := "insert into provision (id, type, animal_type, quantity) values ($1, $2, $3, $4) returning *"
	_, err := p.db.Exec(query, provision.ID, provision.Type, provision.AnimalType, provision.Quantity)
	if err != nil {
		return nil, err
	}

	var createdPrv models.CreateProvision
	err = p.db.QueryRow(query,
		provision.ID,
		provision.Type,
		provision.AnimalType,
		provision.Quantity).Scan(
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

	query := "select id, type, quantity from provision LIMIT $1 OFFSET $2"
	rows, err := p.db.Query(query, filter.Limit, filter.OFFSET)
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
	err := p.db.QueryRow(query, args...).
		Scan(
			&provision.ID,
			&provision.Type,
			&provision.Quantity,
		)

	return &provision, err
}

func (p *ProvisionRepo) UpdateProvision(provision *models.UpdateProvision) error {
	query := "update provision set type = $1, quantity = $2,  updated_at=now() where id = $3"

	_, err := p.db.Exec(query, provision.Type, provision.Quantity, provision.ID)

	return err
}

func (p *ProvisionRepo) DeleteProvision(id int) error {
	now := time.Now().Unix()

	query := "update provision set deleted_at=$1 where id = $2"

	_, err := p.db.Exec(query, now, id)

	return err
}
