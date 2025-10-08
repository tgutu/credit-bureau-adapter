package apicode

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Define custom gRPC error codes and errors.
// Custom codes should not be reused in different
// packages or functions.
var (
	// internal/cba 1001-1999
	ErrCbaCreditBureausNotFound = status.Errorf(codes.Code(1001), "no credit bureaus found")

	// internal/database 2001-2999

	// internal/repository 3001-3999
	ErrCreditRepoListBureausFailed     = status.Errorf(codes.Code(3001), "failed to list credit bureaus")
	ErrCreditRepoGetBureauByNameFailed = status.Errorf(codes.Code(3002), "failed to get credit bureau by name")

	// internal/server 4001-4999
)
