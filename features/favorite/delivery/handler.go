package delivery

import (
	"net/http"
	"warehouse/features/favorite"
	"warehouse/middlewares"
	"warehouse/utils/helper"

	"github.com/labstack/echo/v4"
)

type FavoriteDelivery struct {
	favoriteUsecase favorite.UsecaseInterface
}

func New(e *echo.Echo, usecase favorite.UsecaseInterface) {
	handler := &FavoriteDelivery{
		favoriteUsecase: usecase,
	}
	e.POST("/favorite", handler.PostFavorite, middlewares.JWTMiddleware())

}

func (delivery *FavoriteDelivery) PostFavorite(c echo.Context) error {
	token, role, errToken := middlewares.ExtractToken(c)

	if role != "penitip" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	var dataFav FavRequest

	errBind := c.Bind(&dataFav)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}
	dataFav.UserID = token

	row, err := delivery.favoriteUsecase.PostFavorite(toCore(dataFav))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
}
