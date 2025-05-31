package service

import (
	"context"
	"fmt"
	pb "grpc-server/grpc-server/grpc"
	"sync"
	"sync/atomic"
)

type ProfileService struct {
	pb.UnimplementedProfileServiceServer
	profiles map[int32]*pb.ProfileResponse
	mu       sync.RWMutex
	nextID   int32
}

func NewProfileService() *ProfileService {
	return &ProfileService{
		profiles: make(map[int32]*pb.ProfileResponse),
		nextID:   1,
	}
}

func (s *ProfileService) GetProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	profile, exists := s.profiles[req.Id]
	if !exists {
		return nil, fmt.Errorf("profile with id %d not found", req.Id)
	}

	return profile, nil
}

func (s *ProfileService) GetAllProfiles(ctx context.Context, req *pb.EmptyRequest) (*pb.ProfileListResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	response := &pb.ProfileListResponse{
		Profiles: make([]*pb.ProfileResponse, 0, len(s.profiles)),
	}

	for _, profile := range s.profiles {
		response.Profiles = append(response.Profiles, profile)
	}

	return response, nil
}

func (s *ProfileService) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.ProfileResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := atomic.AddInt32(&s.nextID, 1)
	profile := &pb.ProfileResponse{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	s.profiles[id] = profile
	return profile, nil
}

func (s *ProfileService) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.ProfileResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.profiles[req.Id]; !exists {
		return nil, fmt.Errorf("profile with id %d not found", req.Id)
	}

	profile := &pb.ProfileResponse{
		Id:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	}

	s.profiles[req.Id] = profile
	return profile, nil
}

func (s *ProfileService) DeleteProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.DeleteProfileResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.profiles[req.Id]; !exists {
		return &pb.DeleteProfileResponse{
			Success: false,
			Message: fmt.Sprintf("profile with id %d not found", req.Id),
		}, nil
	}

	delete(s.profiles, req.Id)
	return &pb.DeleteProfileResponse{
		Success: true,
		Message: fmt.Sprintf("profile with id %d successfully deleted", req.Id),
	}, nil
}

// Вспомогательный метод для тестирования
func (s *ProfileService) AddProfile(id int32, name, email string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.profiles[id] = &pb.ProfileResponse{
		Id:    id,
		Name:  name,
		Email: email,
	}
}
