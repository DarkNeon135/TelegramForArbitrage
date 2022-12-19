package server

import (
	message "TelegramForArbitrage/api/proto"
	"TelegramForArbitrage/pkg/telegram"
	"context"
	"fmt"
	"github.com/posipaka-trade/posipaka-trade-cmn/log"
	"google.golang.org/grpc"
	"net"
)

type MessageSender struct {
	message.UnimplementedTelegramMessageSenderServer
	telegramApi *telegram.Telegram
}

func (m *MessageSender) SendMessageToTelegram(ctx context.Context, req *message.MessageRequest) (*message.MessageResponse, error) {
	err := m.telegramApi.SendMessagesToChannel(req.GetMessage())
	if err != nil {
		return &message.MessageResponse{ResponseMessage: err.Error()}, err
	}

	return &message.MessageResponse{ResponseMessage: "Message successfully sent to Telegram!"}, nil
}
func StartGrpcServer(ipAddress net.IP, telegramApi *telegram.Telegram) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ipAddress, 9000))
	if err != nil {
		return fmt.Errorf("network listener error. Error: %s", err)
	}

	s := grpc.NewServer()

	message.RegisterTelegramMessageSenderServer(s, &MessageSender{
		telegramApi: telegramApi,
	})
	log.Info.Println("GRPC server started successfully!")
	if err = s.Serve(lis); err != nil {
		return fmt.Errorf("failed to start GRPC server. Error: %s", err)
	}

	return nil
}
