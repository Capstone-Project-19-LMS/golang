package customerRepository

import (
	"errors"
	"fmt"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/models/model"
	"golang/util"
	"log"
	"math/rand"

	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

// CreateCustomer implements CustomerRepository
func (u *customerRepository) CreateCustomer(customer dto.CostumerRegister) error {

	customerModel := model.Customer{
		ID:           customer.ID,
		Name:         customer.Name,
		Email:        customer.Email,
		Password:     customer.Password,
		ProfileImage: "https://t3.ftcdn.net/jpg/03/46/83/96/360_F_346839683_6nAPzbhpSkIpb8pmAwufkC7c5eD7wYws.jpg",
		IsActive:     false,
	}

	var getAllCustomer []model.Customer
	u.db.Find(&getAllCustomer)

	for _, dataCustomer := range getAllCustomer {
		if customer.Email == dataCustomer.Email {
			if !dataCustomer.IsActive {
				var getCustomerCode model.CustomerCode
				u.db.First(&getCustomerCode, "customer_id=?", dataCustomer.ID)
				letter := []rune("1234567890")
				b := make([]rune, 7)
				for i := range b {
					b[i] = letter[rand.Intn(len(letter))]
				}
				code := string(b)

				isiEmail := fmt.Sprintf("<p>kode verifikasi yaitu <b>%s</b></p>", code)
				mailer := gomail.NewMessage()
				mailer.SetHeader("From", util.GetConfig("SENDER_NAME"))
				mailer.SetHeader("To", customer.Email, "alimuldev@gmail.com")
				mailer.SetAddressHeader("Cc", customer.Email, "Tra Lala La")
				mailer.SetHeader("Subject", "Test mail")
				mailer.SetBody("text/html", isiEmail)
				dialer := gomail.NewDialer(
					util.GetConfig("SMTP_HOST"),
					587,
					util.GetConfig("AUTH_EMAIL"),
					util.GetConfig("AUTH_PASSWORD"),
				)

				errss := dialer.DialAndSend(mailer)
				if errss != nil {
					log.Fatal(errss.Error())
				}

				getCustomerCode.Code = code
				u.db.Save(&getCustomerCode)
			}

		}
	}

	err := u.db.Create(&customerModel).Error

	if err != nil {
		return err
	}
	letter := []rune("1234567890")
	b := make([]rune, 7)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	code := string(b)

	isiEmail := fmt.Sprintf("<p>kode verifikasi yaitu <b>%s</b></p>", code)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", util.GetConfig("SENDER_NAME"))
	mailer.SetHeader("To", customer.Email)
	mailer.SetAddressHeader("Cc", customer.Email, "GENCER")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", isiEmail)
	dialer := gomail.NewDialer(
		util.GetConfig("SMTP_HOST"),
		587,
		util.GetConfig("AUTH_EMAIL"),
		util.GetConfig("AUTH_PASSWORD"),
	)

	errss := dialer.DialAndSend(mailer)
	if errss != nil {
		log.Fatal(errss.Error())
	}
	customerCode := model.CustomerCode{
		ID:         customer.CustomerCodeID,
		CustomerID: customer.ID,
		Code:       code,
	}

	err2 := u.db.Create(&customerCode).Error
	if err2 != nil {
		return err
	}
	return nil
}

func (u *customerRepository) VerifikasiCustomer(input dto.CustomerVerif) error {
	var customerCodeModel model.CustomerCode
	var customerModel model.Customer
	err := u.db.First(&customerCodeModel, "customer_id=?", input.CustomerID)

	if err.Error != nil {
		return err.Error
	}
	u.db.First(&customerModel, "id=?", input.CustomerID)

	if input.Code == customerCodeModel.Code {
		customerModel.IsActive = true
		u.db.Save(&customerModel)
		u.db.Unscoped().Delete(&customerCodeModel)
	}
	return nil
}

// LoginCustomer implements CustomerRepository
func (u *customerRepository) LoginCustomer(customer dto.CostumerLogin) (dto.CostumerResponseGet, error) {
	var customerLogin dto.CostumerResponseGet
	err := u.db.Model(&model.Customer{}).First(&customerLogin, "email = ?", customer.Email).Error
	if err != nil {
		return dto.CostumerResponseGet{}, err
	}
	match := helper.CheckPasswordHash(customer.Password, customerLogin.Password)
	if !match {
		return dto.CostumerResponseGet{}, errors.New(constantError.ErrorEmailOrPasswordNotMatch)
	}
	var customerLoginResponse = dto.CostumerResponseGet{
		ID:           customerLogin.ID,
		Name:         customerLogin.Name,
		Email:        customerLogin.Email,
		Password:     customerLogin.Password,
		ProfileImage: customerLogin.ProfileImage,
		IsActive:     customerLogin.IsActive,
	}
	if !customerLoginResponse.IsActive {
		var getCustomerCode model.CustomerCode
		u.db.First(&getCustomerCode, "customer_id=?", customerLoginResponse.ID)
		letter := []rune("1234567890")
		b := make([]rune, 7)
		for i := range b {
			b[i] = letter[rand.Intn(len(letter))]
		}
		code := string(b)

		isiEmail := fmt.Sprintf("<p>kode verifikasi yaitu <b>%s</b></p>", code)
		mailer := gomail.NewMessage()
		mailer.SetHeader("From", util.GetConfig("SENDER_NAME"))
		mailer.SetHeader("To", customerLogin.Email, "alimuldev@gmail.com")
		mailer.SetAddressHeader("Cc", customerLogin.Email, "Tra Lala La")
		mailer.SetHeader("Subject", "Test mail")
		mailer.SetBody("text/html", isiEmail)
		dialer := gomail.NewDialer(
			util.GetConfig("SMTP_HOST"),
			587,
			util.GetConfig("AUTH_EMAIL"),
			util.GetConfig("AUTH_PASSWORD"),
		)

		errss := dialer.DialAndSend(mailer)
		if errss != nil {
			log.Fatal(err.Error())
		}

		getCustomerCode.Code = code
		u.db.Save(&getCustomerCode)
	}
	return customerLoginResponse, nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
