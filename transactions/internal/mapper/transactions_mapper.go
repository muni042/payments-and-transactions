package mapper

import (
	"strconv"
	"time"

	pb "transactions/api/transactions/v1"
	"transactions/internal/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

//
// ---------- ENUM MAPPERS ----------
//

// proto enum → DB enum
func mapTransactionType(t pb.TransactionType) string {
	switch t {
	case pb.TransactionType_TRANSACTION_TYPE_DEBIT:
		return "debit"
	case pb.TransactionType_TRANSACTION_TYPE_CREDIT:
		return "credit"
	case pb.TransactionType_TRANSACTION_TYPE_FEE:
		return "fee"
	case pb.TransactionType_TRANSACTION_TYPE_REVERSAL:
		return "reversal"
	default:
		return "debit"
	}
}

func mapTransactionStatus(s pb.TransactionStatus) string {
	switch s {
	case pb.TransactionStatus_TRANSACTION_STATUS_PENDING:
		return "pending"
	case pb.TransactionStatus_TRANSACTION_STATUS_POSTED:
		return "posted"
	case pb.TransactionStatus_TRANSACTION_STATUS_ARCHIVED:
		return "archived"
	case pb.TransactionStatus_TRANSACTION_STATUS_FAILED:
		return "failed"
	default:
		return "pending"
	}
}

// DB enum → proto enum
func mapStatusToProto(s string) pb.TransactionStatus {
	switch s {
	case "pending":
		return pb.TransactionStatus_TRANSACTION_STATUS_PENDING
	case "posted":
		return pb.TransactionStatus_TRANSACTION_STATUS_POSTED
	case "archived":
		return pb.TransactionStatus_TRANSACTION_STATUS_ARCHIVED
	case "failed":
		return pb.TransactionStatus_TRANSACTION_STATUS_FAILED
	default:
		return pb.TransactionStatus_TRANSACTION_STATUS_UNSPECIFIED
	}
}

func mapTypeToProto(t string) pb.TransactionType {
	switch t {
	case "debit":
		return pb.TransactionType_TRANSACTION_TYPE_DEBIT
	case "credit":
		return pb.TransactionType_TRANSACTION_TYPE_CREDIT
	case "fee":
		return pb.TransactionType_TRANSACTION_TYPE_FEE
	case "reversal":
		return pb.TransactionType_TRANSACTION_TYPE_REVERSAL
	default:
		return pb.TransactionType_TRANSACTION_TYPE_UNSPECIFIED
	}
}

//
// ---------- ADD ----------
//

func AddReqToBiz(req *pb.AddTransactionRequest) (*model.Transaction, error) {

	accountID, err := strconv.ParseInt(req.AccountId, 10, 64)
	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		return nil, err
	}

	var postingDate time.Time
	if req.PostingDate != "" {
		postingDate, err = time.Parse("2006-01-02", req.PostingDate)
		if err != nil {
			return nil, err
		}
	}

	tx := &model.Transaction{
		AccountID:       accountID,
		TransactionType: mapTransactionType(req.TransactionType),
		Amount:          amount,
		Currency:        req.Currency,
		Status:          mapTransactionStatus(req.Status),
		Description:     req.Description,
		PostingDate:     postingDate,
	}

	if req.TransactionDate != nil {
		tx.TransactionDate = req.TransactionDate.AsTime()
	}

	return tx, nil
}

func ModelToAddResponse(tx *model.Transaction) *pb.AddTransactionResponse {
	return &pb.AddTransactionResponse{
		TransactionId: strconv.FormatInt(tx.TransactionID, 10),
		Status:        mapStatusToProto(tx.Status),
		CreatedAt:     timestamppb.New(tx.CreatedOn),
	}
}

//
// ---------- GET ----------
//

func ModelToGetResponse(tx *model.Transaction) *pb.GetTransactionResponse {

	var postingDate string
	if !tx.PostingDate.IsZero() {
		postingDate = tx.PostingDate.Format("2006-01-02")
	}

	return &pb.GetTransactionResponse{
		TransactionId: strconv.FormatInt(tx.TransactionID, 10),
		AccountId:     strconv.FormatInt(tx.AccountID, 10),

		TransactionType: mapTypeToProto(tx.TransactionType),
		Amount:          strconv.FormatFloat(tx.Amount, 'f', 4, 64),
		Currency:        tx.Currency,

		Status:      mapStatusToProto(tx.Status),
		Description: tx.Description,

		PostingDate:     postingDate,
		TransactionDate: timestamppb.New(tx.TransactionDate),

		CreatedAt: timestamppb.New(tx.CreatedOn),
	}
}

//
// ---------- UPDATE ----------
//

func UpdateReqToBiz(req *pb.UpdateTransactionRequest) (*model.Transaction, error) {

	tx := &model.Transaction{}

	transactionID, err := strconv.ParseInt(req.TransactionId, 10, 64)
	if err != nil {
		return nil, err
	}
	tx.TransactionID = transactionID

	if req.Status != pb.TransactionStatus_TRANSACTION_STATUS_UNSPECIFIED {
		tx.Status = mapTransactionStatus(req.Status)
	}

	if req.Description != "" {
		tx.Description = req.Description
	}

	return tx, nil
}

func ModelToUpdateResponse(tx *model.Transaction) *pb.UpdateTransactionResponse {
	return &pb.UpdateTransactionResponse{
		TransactionId: strconv.FormatInt(tx.TransactionID, 10),
		Status:        mapStatusToProto(tx.Status),
	}
}
