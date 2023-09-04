package controllers

import (
	"net/http"
	"ukzuhdi/configs"
	"ukzuhdi/models"

	"github.com/labstack/echo/v4"
)

func GetProductController(c echo.Context) error {
	var data []models.Product
	result := configs.DB.Find(&data)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Gagal get data dari product",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil get data dari product",
		Data:    data,
	})
}

func AddProductController(c echo.Context) error {
	var productRequest models.Product
	c.Bind(&productRequest)

	result := configs.DB.Create(&productRequest)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed Insert data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "Success insert data to database",
		Data:    productRequest,
	})
}

func DeleteProductController(c echo.Context) error {
	productID := c.Param("id")

	result := configs.DB.Delete(&models.Product{}, productID)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Gagal Hapus Data Dari Database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Produk Berhasil Dihapus",
		Data:    nil,
	})
}

func UpdateProductController(c echo.Context) error {
	productID := c.Param("id")

	var updatedProduct models.Product
	if err := c.Bind(&updatedProduct); err != nil {
		return err
	}

	// Update produk dalam database berdasarkan ID
	result := configs.DB.Model(&models.Product{}).Where("id = ?", productID).Updates(updatedProduct)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to update product in the database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Product updated successfully",
		Data:    updatedProduct,
	})
}
