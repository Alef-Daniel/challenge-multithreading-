package ports

import "context"

type GetAddressAPIExternal interface {
	GetAddress(ctx context.Context, cep string) (string, error)
}
