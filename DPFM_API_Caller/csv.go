package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-function-csv-generator-from-json-rmq-kube/DPFM_API_Input_Formatter"
	dpfm_api_output_formatter "data-platform-function-csv-generator-from-json-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-function-csv-generator-from-json-rmq-kube/config"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func (c *DPFMAPICaller) process(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) (interface{}, *string) {
	var header *[]dpfm_api_output_formatter.Header
	var records *string
	var mountPath *string

	for _, fn := range accepter {
		switch fn {
		case "CSV":
			func() {
				header, records, mountPath, _ = c.GenerateCSV(input, errs, log, conf)
			}()
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:  header,
		Records: records,
	}

	return data, mountPath
}

func (c *DPFMAPICaller) GenerateCSV(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) (*[]dpfm_api_output_formatter.Header, *string, *string, error) {
	var data []dpfm_api_output_formatter.Header

	randomString := generateRandomString(10)
	mountPath := ""

	var err error

	err = mkdirP(fmt.Sprintf(
		"%s/%s",
		conf.MountPath,
		randomString,
	))
	if err != nil {
		return nil, nil, nil, err
	}

	outputPath := fmt.Sprintf(
		"%s/%s/%s%s",
		conf.MountPath,
		randomString,
		fmt.Sprintf("output"),
		".csv",
	)

	mountPath = outputPath

	file, err := os.Create(outputPath)
	if err != nil {
		return nil, nil, nil, err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(input.Header)
	if err != nil {
		return nil, nil, nil, err
	}

	for _, record := range input.Records {
		var row []string
		for _, key := range input.Header {
			switch value := record[key].(type) {
			case int:
				row = append(row, strconv.Itoa(value))
			case float64:
				row = append(row, strconv.FormatFloat(value, 'f', -1, 64))
			case string:
				row = append(row, value)
			case nil:
				row = append(row, "")
			default:
				return nil, nil, nil, err
			}
		}
		if err := writer.Write(row); err != nil {
			return nil, nil, nil, err
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, nil, nil, err
	}

	records, err := json.Marshal(input.Records)
	if err != nil {
		return nil, nil, nil, err
	}

	responseRecord := string(records)

	return &data, &responseRecord, &mountPath, nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func mkdirP(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func copyPDF(srcPath, destPath string) error {
	srcFile, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destPath, srcFile, 0644)
	if err != nil {
		return err
	}

	return nil
}
