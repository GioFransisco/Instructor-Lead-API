package controller

import (
	"encoding/csv"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/driver/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/usecase"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type authController struct {
	authUC         usecase.AuthUsecase
	rg             *gin.RouterGroup
	middlewareAuth middleware.JwtMiddleware
}

func (c *authController) findByUsername(ctx *gin.Context) {
	var userDto dto.UserLoginDto

	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		common.ResponseError(ctx, errors.New(config.ErrorDescriptionForInvalidData).Error(), http.StatusBadRequest)
		return
	}

	token, err := c.authUC.FindByUsername(userDto)

	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	common.ResponseSuccess(ctx, "OK", http.StatusOK, token)
}

func (c *authController) createNewUser(ctx *gin.Context) {
	_, header, errForm := ctx.Request.FormFile("data")
	var payloadData model.User
	// var err error

	if err := ctx.ShouldBindJSON(&payloadData); err != nil {
		if errForm == nil {
			fileLocation := filepath.Join("asset/filecsv", header.Filename)
			os.MkdirAll("asset/filecsv", os.ModePerm)
			ctx.SaveUploadedFile(header, fileLocation)
			fileType := common.GetFileType(fileLocation)

			switch fileType {
			case "xlsx":
				{
					err := c.addXlsxFile(fileLocation)
					if err != nil {
						ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
						return
					}

					ctx.AbortWithStatusJSON(http.StatusCreated, gin.H{"message": "succes add user form xlsx file"})
					return
				}
			case "csv":
				{
					err := c.addCsvFile(fileLocation)
					if err != nil {
						ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
						return
					}
					ctx.AbortWithStatusJSON(http.StatusCreated, gin.H{"message": "succes add user form csv file"})
					return
				}
			default:
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file type invalid"})
				return
			}

		} else {
			errMsg := common.ValidationErrors(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errMsg})
			return
		}
	} else {
		payloadData, err = c.authUC.CreateNewUser(payloadData)

		if err != nil {
			common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
			return
		}

	}
	common.ResponseSuccess(ctx, "OK", http.StatusCreated, payloadData)
}

func (c *authController) addCsvFile(filePath string) error {
	var payloadData model.User
	fileOp, err := os.Open(filePath)

	if err != nil {
		return errors.New("failed open file")
	}

	defer fileOp.Close()

	reader := csv.NewReader(fileOp)

	i := 0
	for {
		record, err := reader.Read()

		if err != nil {
			break
		}

		if record != nil {
			dataString := strings.Split(record[0], ";")

			payloadData.Name = dataString[0]
			payloadData.Email = dataString[1]
			payloadData.PhoneNumber = dataString[2]
			payloadData.Username = dataString[3]
			payloadData.Password = dataString[4]
			payloadData.Age, _ = strconv.Atoi(dataString[5])
			payloadData.Address = dataString[6]
			payloadData.Gander = dataString[7]
			payloadData.Role = dataString[8]
		}
		_, err = c.authUC.CreateNewUser(payloadData)

		if err != nil {
			return err
		}

		i++
	}
	return nil
}

func (c *authController) addXlsxFile(filePath string) error {
	xlFile, err := xlsx.OpenFile(filePath)
	var payloadUser model.User
	if err != nil {
		return errors.New("failed open file")
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for index, cell := range row.Cells {
				value := cell.String()
				switch index {
				case 0:
					payloadUser.Name = value
				case 1:
					payloadUser.Email = value
				case 2:
					payloadUser.PhoneNumber = value
				case 3:
					payloadUser.Username = value
				case 4:
					payloadUser.Password = value
				case 5:
					payloadUser.Age, _ = strconv.Atoi(value)
				case 6:
					payloadUser.Address = value
				case 7:
					payloadUser.Gander = value
				case 8:
					payloadUser.Role = value
				default:
					return errors.New("index out length")
				}
			}

			_, err := c.authUC.CreateNewUser(payloadUser)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *authController) Router() {
	c.rg.POST(config.RegisterPath, c.middlewareAuth.AuthMiddleware("Admin"), c.createNewUser)
	c.rg.POST(config.LoginPath, c.findByUsername)
}

func NewAuthController(authUC usecase.AuthUsecase, rg *gin.RouterGroup, middleware middleware.JwtMiddleware) *authController {
	return &authController{
		authUC:         authUC,
		rg:             rg,
		middlewareAuth: middleware,
	}
}
