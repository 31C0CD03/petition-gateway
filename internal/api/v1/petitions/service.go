package petitions

import (
	"context"
	"log"

	storagev1 "github.com/31c0cd03/petition-pkg/contracts/storage/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var petitionController *PetitionService

type PetitionService struct {
	storageService storagev1.StorageServiceClient
}

func GetPetitionController() *PetitionService {
	if petitionController == nil {
		conn, err := grpc.NewClient("dns:///petition-ms-storage-svc-1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		petitionController = &PetitionService{
			storageService: storagev1.NewStorageServiceClient(conn),
		}
	}
	return petitionController
}

func (p *PetitionService) CreatePetition(name string, description string) (*storagev1.CreatePetitionResponse, error) {
	resp, err := p.storageService.CreatePetition(context.Background(), &storagev1.CreatePetitionRequest{
		Name:        name,
		Description: description,
		AuthorId:    "69420",
	})
	return resp, err
}

func (p *PetitionService) GetPetitions() {
	resp, err := p.storageService.GetPetitions(context.Background(), &emptypb.Empty{})
	log.Printf("%+v\n%v\n", resp, err)
}
