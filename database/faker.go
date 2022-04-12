package databaseTools

import (
	crud "inprinte/backend/crud/insert"
	"inprinte/backend/utils"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

func Faker() {
	db := utils.DbConnect()
	stateCommandLine := []string{"pending", "printing", "printed"}
	stateCommand := []string{"payed", "in progress", "sent", "done"}
	categories := []string{"Transport", "Air", "Terre", "Automobile", "Train", "Espace", "Mer", "Technologie", "Relations", "Amour", "Sexualité", "Médecine", "Travail", "Argent", "Militaire"}

	crud.InsertIntoRole(db, "admin")
	crud.InsertIntoRole(db, "member")

	for _, values := range categories {
		crud.InsertIntoCategory(db, values, gofakeit.Bool())
	}

	for i := 0; i < 100; i++ {
		crud.InsertIntoAddress(db, gofakeit.Address().Street, gofakeit.Address().City, gofakeit.Address().State, gofakeit.Address().Country, gofakeit.Address().Zip)
	}

	for i := 0; i < 100; i++ {
		crud.InsertIntoUser(db, gofakeit.FirstName(), gofakeit.LastName(), gofakeit.Email(), gofakeit.Word(), gofakeit.Phone(), gofakeit.Bool(), utils.Random(1, 100), utils.Random(1, 3))
	}

	crud.InsertIntoPicture(db, gofakeit.ImageURL(400, 400), false)

	for i := 0; i < 100; i++ {
		crud.InsertIntoPicture(db, gofakeit.ImageURL(400, 400), true)
		crud.InsertIntoPicture(db, gofakeit.ImageURL(400, 400), false)
	}

	for i := 0; i < 100; i++ {
		crud.InsertIntoProduct(db, gofakeit.Word(), gofakeit.Sentence(20), utils.Random(1, 100), utils.Random(1, len(categories)), utils.Random(1, 100), gofakeit.Bool(), gofakeit.Bool())

	}

	for i := 2; i < 200; i += 2 {
		crud.InsertIntoRate(db, utils.Random(1, 100), utils.Random(1, 100), utils.Random(1, 5))
		crud.InsertIntoProductPicture(db, i/2, i)
		crud.InsertIntoProductPicture(db, i/2, i+1)
		crud.InsertIntoProductFile(db, utils.Random(1, 100), "url du modèle 3D")
	}

	for i := 0; i < 100; i++ {
		crud.InsertIntoFavoriteRequest(db, utils.Random(1, 100), utils.Random(1, 100))
		crud.InsertIntoCommand(db, utils.Random(1, 100), strconv.Itoa(utils.Random(11111111, 99999999)), stateCommand[utils.Random(1, len(stateCommand))])
	}

	for i := 0; i < 100; i++ {
		crud.InsertIntoCommandLine(db, utils.Random(1, 100), utils.Random(1, 100), utils.Random(1, 10), stateCommandLine[utils.Random(1, len(stateCommandLine))])
	}

}
