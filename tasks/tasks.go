package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeEmailDelivery = "email:deliver"
	TypeImageResize   = "image:resize"
)

type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
}

type ImageResizePayload struct {
	SourceURL string
}

func NewEmailDeliveryTask(userId int, templateId string) (*asynq.Task, error) {
	emailData := &EmailDeliveryPayload{UserID: userId, TemplateID: templateId}
	payload, err := json.Marshal(emailData)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func NewImageResizeTask(src string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageResizePayload{SourceURL: src})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeImageResize, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

// HandleEmailDeliveryTask handles email delivery task
func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("failed to parse email delivery task payload: %v", err)
	}

	log.Printf("Sending email to User: userid=%v, templateid=%s", p.UserID, p.TemplateID)
	// TODO: send email to userId
	return nil
}

type ImageProcessor struct {
	// fields for struct
}

func (p *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed to parse image resize payload: %v", err)
	}
	log.Printf("Resizing image src=%s", payload.SourceURL)
	// TODO: resize image code
	return nil
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}
