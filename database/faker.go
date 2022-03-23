package databaseTools

import (
	"inprinte/backend/utils"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

func Faker() {
	db := DbConnect()

	// var
	stateCommandLine := []string{"pending", "printing", "printed"}
	stateCommand := []string{"payed", "in progress", "sent", "done"}
	categories := []string{"Transport", "Air", "Terre", "Automobile", "Train", "Espace", "Mer", "Technologie", "Relations", "Amour", "Sexualité", "Médecine", "Travail", "Argent", "Militaire"}

	// Add fake roles
	InsertIntoRole(db, "admin")
	InsertIntoRole(db, "member")

	// Add fake categories
	for _, values := range categories {
		InsertIntoCategory(db, values)
	}

	// Add fake address
	for i := 0; i < 100; i++ {
		InsertIntoAddress(db, gofakeit.Address().Street, gofakeit.Address().City, gofakeit.Address().State, gofakeit.Address().Country, gofakeit.Address().Zip)
	}

	//add fake users
	for i := 0; i < 100; i++ {
		InsertIntoUser(db, gofakeit.FirstName(), gofakeit.LastName(), gofakeit.Email(), gofakeit.Word(), gofakeit.Phone(), gofakeit.Bool(), utils.Random(1, 100), utils.Random(1, 3))
	}

	//add fake picture
	for i := 0; i < 100; i++ {
		InsertIntoPicture(db, gofakeit.ImageURL(400, 400))
	}

	for i := 0; i < 100; i++ {
		// Add fake products
		InsertIntoProduct(db, gofakeit.Word(), gofakeit.Sentence(20), utils.Random(1, 100), utils.Random(1, len(categories)), utils.Random(1, 100), gofakeit.Bool(), gofakeit.Bool())

	}

	for i := 0; i < 100; i++ {
		// Add fake rates
		InsertIntoRate(db, utils.Random(1, 100), utils.Random(1, 100), utils.Random(1, 5))
		// Add fake product pictures
		InsertIntoProductPicture(db, utils.Random(1, 100), utils.Random(1, 100))
		// Add fake product files
		InsertIntoProductFile(db, utils.Random(1, 100), "url du modèle 3D")
	}

	for i := 0; i < 100; i++ {
		// Add fake favorites
		InsertIntoFavorite(db, utils.Random(1, 100), utils.Random(1, 100))
		InsertIntoCommand(db, utils.Random(1, 100), strconv.Itoa(utils.Random(11111111, 99999999)), stateCommand[utils.Random(1, len(stateCommand))])
	}

	for i := 0; i < 100; i++ {
		InsertIntoCommandLine(db, utils.Random(1, 100), utils.Random(1, 100), utils.Random(1, 10), stateCommandLine[utils.Random(1, len(stateCommandLine))])
	}

}
