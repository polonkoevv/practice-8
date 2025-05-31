package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-client/grpc-server/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := pb.NewProfileServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createProfile, err := client.CreateProfile(ctx, &pb.CreateProfileRequest{
		Name:  "Alice Johnson",
		Email: "alice@example.com",
	})
	if err != nil {
		log.Printf("ошибка при создании профиля: %v", err)
	} else {
		fmt.Printf("Создан профиль: %+v\n", createProfile)
	}

	profiles, err := client.GetAllProfiles(ctx, &pb.EmptyRequest{})
	if err != nil {
		log.Printf("ошибка при получении всех профилей: %v", err)
	} else {
		fmt.Println("\nСписок всех профилей:")
		for _, profile := range profiles.Profiles {
			fmt.Printf("ID: %d, Name: %s, Email: %s\n", profile.Id, profile.Name, profile.Email)
		}
	}

	getProfile, err := client.GetProfile(ctx, &pb.ProfileRequest{Id: createProfile.Id})
	if err != nil {
		log.Printf("ошибка при получении профиля: %v", err)
	} else {
		fmt.Printf("\nПолучен профиль по ID: %+v\n", getProfile)
	}

	// Обновление профиля
	updateProfile, err := client.UpdateProfile(ctx, &pb.UpdateProfileRequest{
		Id:    createProfile.Id,
		Name:  "Alice Smith",
		Email: "alice.smith@example.com",
	})
	if err != nil {
		log.Printf("ошибка при обновлении профиля: %v", err)
	} else {
		fmt.Printf("\nОбновлен профиль: %+v\n", updateProfile)
	}

	// Удаление профиля
	deleteResult, err := client.DeleteProfile(ctx, &pb.ProfileRequest{Id: createProfile.Id})
	if err != nil {
		log.Printf("ошибка при удалении профиля: %v", err)
	} else {
		fmt.Printf("\nРезультат удаления: %s\n", deleteResult.Message)
	}

	// Проверка удаления - получение всех профилей
	profiles, err = client.GetAllProfiles(ctx, &pb.EmptyRequest{})
	if err != nil {
		log.Printf("ошибка при получении всех профилей: %v", err)
	} else {
		fmt.Println("\nСписок всех профилей после удаления:")
		for _, profile := range profiles.Profiles {
			fmt.Printf("ID: %d, Name: %s, Email: %s\n", profile.Id, profile.Name, profile.Email)
		}
	}
}
