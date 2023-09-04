package models

type Worker struct {
	WorkerId        int    `json:"worker_id"`
	JobsInQueue     int    `json:"jobs_in_queue"`
	JobsUnackBuffer int    `json:"jobs_unack_buffer"`
	LastUnack       string `json:"last_unack"`
	LastAck         string `json:"last_ack"`
	Count           int    `json:"count"`
}
