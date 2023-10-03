package product

// func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request){

// 	uh.mu.Lock()
// 	defer uh.mu.Unlock()

// 	query := "SELECT id, user FROM Users"

// 	rows, err := database.DB.Query(query)
// 	if err != nil {
// 		response.ResponseWithError(w, 500, "Error querying Database")
// 		return
//     }
// 	defer rows.Close()

// 	for rows.Next(){
// 		var user User

// 		if err:= rows.Scan(&user.ID, &user.User); err != nil{
// 			response.ResponseWithError(w, 500, "Error iterating data!")
// 			return
// 		}

// 		uh.users = append(uh.users, user)
// 	}

// 	response.ResponseWithJSON(w, http.StatusOK, uh.users)
// 	uh.users = nil
// }