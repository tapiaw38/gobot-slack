package querys

import (
	"bot-slack/models"
	"database/sql"
	"log"
)

// Function to get all messages
func GetMessage(db *sql.DB) []models.Category {

	var category models.Category
	var categories []models.Category

	query := "SELECT * FROM core_category"

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var name string
		var slug string
		var description string
		var color string
		var order int

		if err := rows.Scan(&id, &name, &slug, &description, &color, &order); err != nil {
			log.Fatal(err)
		}

		category = models.Category{
			Id:          id,
			Name:        name,
			Slug:        slug,
			Description: description,
			Color:       color,
			Order:       order,
		}

		categories = append(categories, category)

	}

	defer rows.Close()
	defer db.Close()

	return categories

}
