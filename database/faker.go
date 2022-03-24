package databaseTools

import (
	CRUD "inprinte/backend/CRUD/insert"
	"inprinte/backend/utils"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

func Faker() {
	db := DbConnect()
	stateCommandLine := []string{"pending", "printing", "printed"}
	stateCommand := []string{"payed", "in progress", "sent", "done"}
	categories := []string{"Transport", "Air", "Terre", "Automobile", "Train", "Espace", "Mer", "Technologie", "Relations", "Amour", "Sexualité", "Médecine", "Travail", "Argent", "Militaire"}

	CRUD.InsertIntoRole(db, "admin")
	CRUD.InsertIntoRole(db, "member")

	for _, values := range categories {
		CRUD.InsertIntoCategory(db, values)
	}

	for i := 0; i < 100; i++ {
		CRUD.InsertIntoAddress(db, gofakeit.Address().Street, gofakeit.Address().City, gofakeit.Address().State, gofakeit.Address().Country, gofakeit.Address().Zip)
	}

	for i := 0; i < 100; i++ {
		CRUD.InsertIntoUser(db, gofakeit.FirstName(), gofakeit.LastName(), gofakeit.Email(), gofakeit.Word(), gofakeit.Phone(), gofakeit.Bool(), utils.Random(1, 100), utils.Random(1, 3))
	}

	CRUD.InsertIntoPicture(db, gofakeit.ImageURL(400, 400), false, true)

	for i := 0; i < 100; i++ {
		CRUD.InsertIntoPicture(db, gofakeit.ImageURL(400, 400), true, false)
		CRUD.InsertIntoPicture(db, gofakeit.ImageURL(400, 400), false, false)
	}

	for i := 0; i < 100; i++ {
		CRUD.InsertIntoProduct(db, gofakeit.Word(), gofakeit.Sentence(20), utils.Random(1, 100), utils.Random(1, len(categories)), utils.Random(1, 100), gofakeit.Bool(), gofakeit.Bool())

	}

	for i := 2; i < 200; i += 2 {
		CRUD.InsertIntoRate(db, utils.Random(1, 100), utils.Random(1, 100), utils.Random(1, 5))
		CRUD.InsertIntoProductPicture(db, i/2, i)
		CRUD.InsertIntoProductPicture(db, i/2, i+1)
		CRUD.InsertIntoProductFile(db, utils.Random(1, 100), "url du modèle 3D")
	}

	for i := 0; i < 100; i++ {
		CRUD.InsertIntoFavorite(db, utils.Random(1, 100), utils.Random(1, 100))
		CRUD.InsertIntoCommand(db, utils.Random(1, 100), strconv.Itoa(utils.Random(11111111, 99999999)), stateCommand[utils.Random(1, len(stateCommand))])
	}

	for i := 0; i < 100; i++ {
		CRUD.InsertIntoCommandLine(db, utils.Random(1, 100), utils.Random(1, 100), utils.Random(1, 10), stateCommandLine[utils.Random(1, len(stateCommandLine))])
	}

}
