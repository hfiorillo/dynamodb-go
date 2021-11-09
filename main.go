package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"strconv"
)

var dynamodb *dynamodb.DynamoDB

type Person struct {
	Id            int
	Name, Website string
}

const (
	REGION = "eu-west-1"
	TABLE_NAME = "exampletable"
)

func init() {
	dynamo = connectDynamo()
}

func connectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(sesion.Must(Sesion.NewSession(&as.Config{
		Region: aws.String(REGION)
	})))
}

func createTable() {
	result, err := dynamodb.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("Name")
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType: aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(TABLE_NAME),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
}

func PutItem(person Person) {
	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue {
			"Id" : {
				N: aws.String(strconv.Itoa(person.Id)),
			},
			"Name": {
				s: aws.String(strconv.Itoa(person.Name)),
			},
			"Website": {
				s: aws.String(strconv.Itoa(person.Website)),
			},
		},
		TableName: aws.String(TABLE_NAME),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
}

func UpdateItem(person Person) {
	_, err := dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#N": aws.String("Name"),
			"#W": aws.String("Website"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Name": {
				S: aws.String(person.Name),
			},
			":Website": {
				S: aws.String(person.Website),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Id":{
				N: aws.String(strconv.Itoa(person.Id))
			},
		},
		TableName: aws.String(TABLE_NAME),
		UpdateExpression: aws.String("SET #N = :Name, #W = :Website"),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
}

func DeleteItem(id int) {
	_, err := dynamo.DeleteItem(&dynamo.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
		TableName: aws.String(TABLE_NAME),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
}

func GetItem(id int)(person Person) {
	_, err := dynamo.GetItem(&dynamo.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
		TableName: aws.String(TABLE_NAME),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
	err = dynamodbattribute.UnmarshalMap(result.Item, &person)
	if err != nil{
		panic(err)
	}

	return person

}


func main() {
	CreateTable()

	var person Person = Person {
		Id: 1,
		Name: "Harry Example",
		Website: "fiorillo.co.uk",
	}
	fmt.Println(GetItem(1))

	PutItem(person)
	fmt.Println(GetItem(1))

	person.Name = "Harry Fiorillo"
	UpdateItem(person)
	fmt.Println(GetItem(1))

	DeleteItem(1)
	fmt.Println(GetItem(1))
}

