package db

import (
	"blog/models"
)

func (db *BlogDBImpl) CreateUserIdDB(userData *models.UserData) (*models.UserDataResponse, error) {
	// userIdResp := models.NewUserDataResponse()
	// tx := db.dbConn.MustBegin()
	// _, err := tx.NamedQuery(`INSERT INTO userdatabase(user_id,name,email,phone_number)VALUES(:user_id,:name,:email,:phone_number)`, userData)
	// if err != nil {
	// 	return nil, err
	// }
	// err = tx.Commit()
	// // if err != nil {
	// // 	return nil, db_error.NewInternalServerError(err.Error())
	// // }
	// err = db.dbConn.Get(userIdResp, `SELECT * FROM userdatabase WHERE user_id=?`, *userData.UserId)
	// // if err != nil {
	// // 	return nil, db_error.NewInternalServerError(err.Error())
	// // }
	return nil, nil
}
