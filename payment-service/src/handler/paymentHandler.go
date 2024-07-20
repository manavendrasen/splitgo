package handler

import (
	"context"
	"payment-service/src/repository"

	pb "common"

	"google.golang.org/grpc"
)

type handler struct {
	pb.UnimplementedPaymentServiceServer
}

func NewHandler(grpcServer *grpc.Server) {
	pb.RegisterPaymentServiceServer(grpcServer, &handler{})
}

func (h *handler) GetPayment(c context.Context, req *pb.GetPaymentRequest) (*pb.PaymentList, error) {
	payments, err := repository.ViewPayment(uint(req.From))

	if err != nil {
		return nil, err
	}

	pbPayments := &pb.PaymentList{}

	for _, v := range payments {
		pbPayments.Payments = append(pbPayments.Payments, &pb.Payment{
			ID:          int32(v.ID),
			From:        uint64(v.FromUserId),
			To:          v.To,
			Amount:      v.Amount,
			Description: v.Description,
			UpdatedAt:   v.UpdatedAt.String(),
			CreatedAt:   v.CreatedAt.String(),
		})
	}
	return pbPayments, nil
}

func (h *handler) AddPayment(c context.Context, req *pb.AddPaymentRequest) (*pb.Payment, error) {
	payment, err := repository.AddPayment(uint(req.From), req.To, req.Amount, req.Description)

	if err != nil {
		return nil, err
	}

	return &pb.Payment{
		ID: int32(payment.ID),
	}, nil
}

func (h *handler) UpdatePayment(c context.Context, req *pb.UpdatePaymentRequest) (*pb.Payment, error) {
	payment, err := repository.UpdatePayment(uint(req.From), uint(req.PaymentId), req.Amount, req.To, req.Description)

	if err != nil {
		return nil, err
	}

	return &pb.Payment{
		ID:          int32(payment.ID),
		To:          payment.To,
		Amount:      payment.Amount,
		Description: payment.Description,
		UpdatedAt:   payment.UpdatedAt.String(),
		CreatedAt:   payment.CreatedAt.String(),
	}, nil
}

func (h *handler) DeletePayment(c context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error) {
	payment, err := repository.DeletePayment(uint(req.UserId), uint(req.PaymentId))

	if err != nil {
		return nil, err
	}

	result := make(map[string]int)
	result["rowsAffected"] = int(payment)
	return &pb.DeletePaymentResponse{
		RowsAffected: int32(payment),
	}, nil
}
