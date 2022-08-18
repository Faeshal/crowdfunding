package payment

import (
	"crowdfunding/user"
	"strconv"

	"github.com/spf13/viper"
	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	// midclient.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	// midclient.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midclient.ServerKey = viper.Get("MIDTRANS_SERVER_KEY").(string)
	midclient.ClientKey = viper.Get("MIDTRANS_CLIENT_KEY").(string)
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
