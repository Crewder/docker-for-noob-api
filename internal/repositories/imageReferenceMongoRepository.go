package repositories

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/pkg/goDotEnv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strings"
)

type imageReferenceRepository struct {
	client *mongo.Client
}

func NewImageReferenceRepository() *imageReferenceRepository {
	mongoUri := goDotEnv.GetEnvVariable("MONGO_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	return &imageReferenceRepository{
		client: client,
	}
}

func (repository *imageReferenceRepository) Read(imageName string) (domain.ImageReference, error) {
	coll := repository.client.Database("docker-for-noob").Collection("reference")

	var result bson.M

	err := coll.FindOne(context.TODO(), bson.D{{"Name", imageName}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return domain.ImageReference{}, err
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return domain.ImageReference{}, err
	}

	var response domain.ImageReference

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return domain.ImageReference{}, err
	}

	return response, nil
}

func (repository *imageReferenceRepository) Add(imageReference domain.ImageReference) error {

	coll := repository.client.Database("docker-for-noob").Collection("reference")
	doc := bson.D{{"Name", imageReference.Name}, {"Port", imageReference.Port}, {"workdir", imageReference.Workdir}, {"Env", imageReference.Env}}
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}

	return nil
}

func (repository *imageReferenceRepository)  AddAllTagReferenceFromApi() error {

	coll := repository.client.Database("docker-for-noob").Collection("reference")

	pathToInputData := goDotEnv.GetEnvVariable("BATCH_REFERENTIEL_BUFFER")

	if _, err := os.Stat(pathToInputData);
		err != nil {
		return err
	}

	f, err := os.Open(pathToInputData)
	if err != nil {
		return err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	var allReferenceToAdd []interface{}

	for _, element := range csvData {
		allReferenceToAdd = append(allReferenceToAdd, mapCsvResultToTagReferenceStruct(element))
	}

	_, err = coll.InsertMany(context.TODO(), allReferenceToAdd)
	if err != nil {
		return err
	}

	return nil
}

func mapCsvResultToTagReferenceStruct(csvLine []string) domain.ImageReference {

	tagReference := domain.ImageReference{}
	tagReference.Name = csvLine[0]
	tagReference.Port = strings.Fields(csvLine[2])
	tagReference.Workdir = strings.Fields(csvLine[3])
	return tagReference
}