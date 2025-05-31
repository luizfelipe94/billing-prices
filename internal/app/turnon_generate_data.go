package app

import (
	"context"
	"log"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type GenerateDataCommand struct {
	Count int `json:"count"`
}

type TurnOnGenerateDataHandler struct {
	Clientset *kubernetes.Clientset
}

func NewTurnOnGenerateDataHandler(clientset *kubernetes.Clientset) *TurnOnGenerateDataHandler {
	return &TurnOnGenerateDataHandler{
		Clientset: clientset,
	}
}

func (h *TurnOnGenerateDataHandler) Handle(ctx context.Context, command GenerateDataCommand) error {
	jobName := "billing-data-generator"
	namespace := "default"

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobName,
			Namespace: namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": jobName,
					},
				},
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicyNever,
					Containers: []corev1.Container{
						{
							Name:  jobName,
							Image: "billing.localhost:5000/billing-data-generator:latest",
							Env: []corev1.EnvVar{
								{
									Name:  "KAFKA_BROKERS",
									Value: "billing-kafka-bootstrap:9092",
								},
								{
									Name:  "DATABASE_URL",
									Value: "to-do",
								},
								{
									Name:  "NUM_WORKERS",
									Value: "1",
								},
								{
									Name:  "NUM_MESSAGES",
									Value: "100",
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := h.Clientset.BatchV1().Jobs(namespace).Create(ctx, job, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Failed to create Job: %v", err)
		return err
	}

	log.Printf("Job %s created successfully", jobName)
	return nil
}