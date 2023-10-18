package main

import "time"

// endpoint request structs

type ListWorkingHoursRequest struct {
	PayloadId string `json:"payloadId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type ListBlockHoursRequest struct {
	PayloadId string `json:"payloadId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type ListBookingRequest struct {
	PayloadId string `json:"payloadId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

// entities

type Duration struct {
	Seconds int64 `json:"seconds"`
}

type Payload struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type WorkingHour struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	Quantity   int64     `json:"quantity"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type BlockHour struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type Booking struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	Quantity   int64     `json:"quantity"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}


// helper functions

func TimeToString(tm time.Time) string {
	return tm.Format(time.RFC3339)
}

func StringToTime(timeStr string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
