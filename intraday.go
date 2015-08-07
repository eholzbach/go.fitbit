package fitbit

import (
	"encoding/json"
	"fmt"
	"time"
)

type IntradayCalories struct {
	ActivitiesCalories []struct {
		Datetime string `json:"dateTime"`
		Value string `json:"value"`
	} `json:"activities-calories"`
	ActivitiesCaloriesIntraday struct {
		Dataset []struct {
			Level int `json:"level"`
			Mets int `json:"mets"`
			Time string `json:"time"`
			Value float64 `json:"value"`
		} `json:"dataset"`
		Datasetinterval int `json:"datasetInterval"`
		Datasettype string `json:"datasetType"`
	} `json:"activities-calories-intraday"`
}

type IntradayDistance struct {
	ActivitiesDistance []struct {
		Datetime string `json:"dateTime"`
		Value string `json:"value"`
	} `json:"activities-distance"`
	ActivitiesDistanceIntraday struct {
		Dataset []struct {
			Time string `json:"time"`
			Value float64 `json:"value"`
		} `json:"dataset"`
		Datasetinterval int `json:"datasetInterval"`
		Datasettype string `json:"datasetType"`
	} `json:"activities-distance-intraday"`
}

type IntradayElevation struct {
        ActivitiesElevation []struct {
                Datetime string `json:"dateTime"`
                Value string `json:"value"`
        } `json:"activities-elevation"`
        ActivitiesElevationIntraday struct {
                Dataset []struct {
                        Time string `json:"time"`
                        Value float64 `json:"value"`
                } `json:"dataset"`
                Datasetinterval int `json:"datasetInterval"`
                Datasettype string `json:"datasetType"`
        } `json:"activities-elevation-intraday"`
}

type IntradayFloors struct {
        ActivitiesFloors []struct {
                Datetime string `json:"dateTime"`
                Value string `json:"value"`
        } `json:"activities-floors"`
        ActivitiesFloorsIntraday struct {
                Dataset []struct {
                        Time string `json:"time"`
                        Value int `json:"value"`
                } `json:"dataset"`
                Datasetinterval int `json:"datasetInterval"`
                Datasettype string `json:"datasetType"`
        } `json:"activities-floors-intraday"`
}

type IntradaySteps struct {
        ActivitiesSteps []struct {
                Datetime string `json:"dateTime"`
                Value string `json:"value"`
        } `json:"activities-steps"`
        ActivitiesStepsIntraday struct {
                Dataset []struct {
                        Time string `json:"time"`
                        Value int `json:"value"`
                } `json:"dataset"`
                Datasetinterval int `json:"datasetInterval"`
                Datasettype string `json:"datasetType"`
        } `json:"activities-steps-intraday"`
}

func (c *Client) GetIntradayCalories(date time.Time) (*IntradayCalories, error) {
	requestURL := fmt.Sprintf("user/-/activities/calories/date/%s/1d/1min.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	intradayCaloriesData := &IntradayCalories{}
	err = json.NewDecoder(responseBody).Decode(intradayCaloriesData)
	if err != nil {
		return nil, err
	}

	return intradayCaloriesData, nil
}

func (c *Client) GetIntradayDistance(date time.Time) (*IntradayDistance, error) {
        requestURL := fmt.Sprintf("user/-/activities/distance/date/%s/1d/1min.json", date.Format("2006-01-02"))
        responseBody, err := c.getData(requestURL)
        if err != nil {
                return nil, err
        }

        intradayDistanceData := &IntradayDistance{}
        err = json.NewDecoder(responseBody).Decode(intradayDistanceData)
        if err != nil {
                return nil, err
        }

        return intradayDistanceData, nil
}

func (c *Client) GetIntradayElevation(date time.Time) (*IntradayElevation, error) {
        requestURL := fmt.Sprintf("user/-/activities/elevation/date/%s/1d/1min.json", date.Format("2006-01-02"))
        responseBody, err := c.getData(requestURL)
        if err != nil {
                return nil, err
        }

        intradayElevationData := &IntradayElevation{}
        err = json.NewDecoder(responseBody).Decode(intradayElevationData)
        if err != nil {
                return nil, err
        }

        return intradayElevationData, nil
}

func (c *Client) GetIntradayFloors(date time.Time) (*IntradayFloors, error) {
        requestURL := fmt.Sprintf("user/-/activities/floors/date/%s/1d/1min.json", date.Format("2006-01-02"))
        responseBody, err := c.getData(requestURL)
        if err != nil {
                return nil, err
        }

        intradayFloorsData := &IntradayFloors{}
        err = json.NewDecoder(responseBody).Decode(intradayFloorsData)
        if err != nil {
                return nil, err
        }

        return intradayFloorsData, nil
}

func (c *Client) GetIntradaySteps(date time.Time) (*IntradaySteps, error) {
        requestURL := fmt.Sprintf("user/-/activities/steps/date/%s/1d/1min.json", date.Format("2006-01-02"))
        responseBody, err := c.getData(requestURL)
        if err != nil {
                return nil, err
        }

        intradayStepsData := &IntradaySteps{}
        err = json.NewDecoder(responseBody).Decode(intradayStepsData)
        if err != nil {
                return nil, err
        }

        return intradayStepsData, nil
}

