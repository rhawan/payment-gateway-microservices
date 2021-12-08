package factory

import "github.com/rhawan/payment-gateway-microservices/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
