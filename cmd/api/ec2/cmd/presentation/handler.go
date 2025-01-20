package ec2_presentation

import (
	ec2_domain "awsManager/api/ec2/cmd/domain"
)

type Handler struct {
	svc ec2_domain.IService
}

func NewHandler(svc ec2_domain.IService) *Handler {
	return &Handler{svc: svc}
}
