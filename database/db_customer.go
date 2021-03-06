package database

import (
	"github.com/tranhuy-dev/IStockGolang/models"
	"github.com/tranhuy-dev/IStockGolang/models/response_models"
	"github.com/tranhuy-dev/IStockGolang/core/mathematic"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"errors"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo/options"
	"crypto/sha256"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LoginAccount(loginBody models.LoginBody) (interface{} , error) {
	customerCollection := Client.Database(DatabaseName).Collection("customer")
	// fmt.Println(loginBody)
	filter := bson.D{
		{"email",loginBody.Email},
		{"password" , loginBody.Password},
	}

	var customer models.Customer
	err := customerCollection.FindOne(context.TODO(),filter).Decode(&customer)
	if err != nil {
		return nil,errors.New("Login fail")
	}

	token := mathematic.GetHash(customer.Email)
	_ , errSession := CreateSessionToken(token , customer.Email)
	if errSession != nil {
		return nil , errSession
	}
	dataResponse := map[string]interface{}{}
	dataResponse["customer"] = customer
	dataResponse["token"] = token
	return dataResponse ,nil
}

func InsertCustomer(req models.CustomerReq) interface{} {
	newCustomer := models.Customer{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Phone: req.Phone,
		Address: req.Address,
		Age: req.Age,
		Status:1,
		Email:req.Email,
		Password:req.Password}
	customerCollection := Client.Database(DatabaseName).Collection("customer")
	_, errorQueryInsert := customerCollection.InsertOne(context.TODO(), newCustomer)
	if errorQueryInsert != nil {
		log.Fatal(errorQueryInsert)
	}
	sendmail , err := SendMail()
	if err != nil {
		return err
	}
	hashToken := sha256.Sum256([]byte(newCustomer.Email))
	responseBody := map[string]interface{}{}
	responseBody["token"] = hashToken[:]
	responseBody["email"] = sendmail
	return responseBody
}

func RetrieveAllCustomer() interface{} {
	var customer []*models.Customer
	customerCollection := Client.Database(DatabaseName).Collection("customer")
	findOptions := options.Find()
	findOptions.SetLimit(100)
	cur, err := customerCollection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Customer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		customer = append(customer, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	// Syncnorize()
	responseBody := map[string]interface{}{}
	responseBody["customer"] = customer
	responseBody["size"] = len(customer)
	return responseBody
}

func UpdateCustomer(req models.CustomerReq , token string) (*models.Customer , error){
	customerCollection := Client.Database(DatabaseName).Collection("customer")
	dataSession, errSession := CheckToken(token)
	if errSession != nil {
		return nil , errSession
	}
	filter := bson.D{{"email",dataSession.Customer}}

	updateBody := bson.D{
		{"$set" , bson.D{
			{"first_name" , req.FirstName},
			{"last_name",req.LastName},
			{"age",req.Age},
			{"phone",req.Phone},
			{"address",req.Address},
		}},
	}
	var customer models.Customer
	err := customerCollection.FindOneAndUpdate(context.TODO() , filter , updateBody).Decode(&customer)
	if err != nil {
		return nil , errors.New(constant.MessageUserNotFound)
	}
	return &customer, nil
}

func DeleteCustomer(token string) (*models.Customer , error) {
	customerCollection := Client.Database(DatabaseName).Collection("customer")
	dataSession, errSession := CheckToken(token)
	if errSession != nil {
		return nil , errSession
	}
	filter := bson.D{{"email" , dataSession.Customer}}
	updateBody := bson.D{
		{"$set", bson.D{
			{"status",0},
		}},
	}
	var customer models.Customer
	err := customerCollection.FindOneAndUpdate(context.TODO() , filter , updateBody).Decode(&customer)
	if err != nil {
		return nil,errors.New(constant.MessageUserNotFound)
	}
	return &customer,nil
}

func FindUserByEmail(token , email string) (interface{}, error) {
	var customer models.Customer
	customerCollection := Client.Database(DatabaseName).Collection("customer")
	err := customerCollection.FindOne(context.TODO() , bson.D{
		{"email" , email},
	}).Decode(&customer)
	if err != nil {
		return nil, errors.New(constant.MessageUserNotFound)
	}
	dataStock , errResultStock := RetrieveStockUser(token , email)
	if errResultStock != nil {
		return nil , errResultStock
	}

	customerResponseModel := response_models.CustomerResponse{
		FirstName: customer.FirstName,
		LastName: customer.LastName,
		Age: customer.Age,
		Address: customer.Address,
		Email: customer.Email,
		Stock: dataStock,
	}

	return customerResponseModel,nil
}
func IncID() interface{} {
	var sequenceID models.SequenceID
	seCollection := Client.Database(DatabaseName).Collection("sequence")
	filter := bson.D{{"sequence_type","sequence_id"}}
	updateBody := bson.D{
		{"$inc", bson.D{
			{"count",1},
		}},
	}
	seCollection.UpdateOne(context.TODO() , filter , updateBody)

	err := seCollection.FindOne(context.TODO() , filter).Decode(&sequenceID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("mysequence %+v",sequenceID)
	return sequenceID.Count
}

func Syncnorize() {
	customerCollection := Client.Database(DatabaseName).Collection("customer")
	updateBody := bson.D{
		{"$set",bson.D{
			{"role",constant.ROLE_USER},
		}},
	}
	customerCollection.UpdateMany(context.TODO(),bson.D{{}} , updateBody)
}

func ChangePassword(passwordReq models.ChangePasswordReq , token string) (*models.Customer , error) {
	customerCollection := Client.Database(DatabaseName).Collection("customer")
	dataSession, errSession := CheckToken(token)
	if errSession != nil {
		return nil , errSession
	}
	filter := bson.D{
		{"password",passwordReq.OldPassword},
		{"email", dataSession.Customer},
	}

	updateBody := bson.D{
		{"$set", bson.D{
			{"password",passwordReq.NewPassword},
		}},
	}

	var customer models.Customer
	err := customerCollection.FindOneAndUpdate(context.TODO() , filter , updateBody).Decode(&customer)
	if err != nil {
		return nil,errors.New("Change password fail")
	}
	return &customer,nil
}

func RetrieveCustomerByFilter(filterBody models.FilterUser) ([]*models.Customer, error) {
	filter := bson.M{}

	if filterBody.Age != 0 {
		filter["age"] = filterBody.Age
	}

	if filterBody.Address != "" {
		filter["address"] = filterBody.Address
	}

	if filterBody.Name != "" {
		filter["$or"] = bson.A{
			bson.M{"first_name" : primitive.Regex{Pattern: filterBody.Name, Options:"i"}},
			bson.M{"last_name" : primitive.Regex{Pattern: filterBody.Name, Options:"i"}},
		}

	}

	customerCollection := Client.Database(DatabaseName).Collection("customer")
	var customers []*models.Customer

	findOption := options.Find()
	findOption.SetLimit(100)

	cur,err := customerCollection.Find(context.TODO() , filter , findOption)
	if err != nil {
		return nil , errors.New(constant.MessageUnexpectedError)
	}
	for cur.Next(context.TODO()) {
		var element models.Customer
		err := cur.Decode(&element)
		if err != nil {
			return nil , errors.New(constant.MessageUnexpectedError)
		}

		customers = append(customers , &element)
	}

	if err := cur.Err(); err != nil {
		return nil , errors.New(constant.MessageUnexpectedError)
	}

	cur.Close(context.TODO())

	return customers , nil
}

func SendMail() (interface{}, error) {
	from := mail.NewEmail("Example user", "tranquochuy15091996@gmail.com")
	subject := "Test sendgrid"
	to := mail.NewEmail("Example user", "tqhuy1996.developer@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(SendgridApiKey)
	response, err := client.Send(message)
	if err != nil {
		return nil, errors.New("Send mail fail")
	}

	return response , nil
}