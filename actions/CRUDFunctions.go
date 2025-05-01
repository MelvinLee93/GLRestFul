package actions

import( 
"net/http"
"time"
"Prototype/config"
"Prototype/models"
"github.com/labstack/echo/v4"
)

func AddEmployee(c echo.Context) error {
	b := new(models.Users)
	db := config.DB()
	b.Created_At = time.Now()
    b.Updated_At = time.Now()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	users := &models.Users{
		Username: b.Username,
		Password: b.Password,
		Email: b.Email,
		Created_At: b.Created_At,
		Updated_At: b.Updated_At,
	}

	if err := db.Create(&users).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": b,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateEmployee(c echo.Context) error {
	id := c.Param("id")
	b := new(models.Users)
	b.Updated_At = time.Now()
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existing_Users := new(models.Users)

	if err := db.First(&existing_Users, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existing_Users.Username = b.Username
	existing_Users.Password = b.Password
	existing_Users.Updated_At = b.Updated_At
	if err := db.Save(&existing_Users).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existing_Users,
	}

	return c.JSON(http.StatusOK, response)
}

func FindEmployee(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var users []*models.Users

	if res := db.Find(&users, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": users[0],
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteEmployee(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	users := new(models.Users)

	err := db.Delete(&users, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "An employee has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}
