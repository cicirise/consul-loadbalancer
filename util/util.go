package util

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"os/exec"
	"time"
)

const (
	CLOUD_AWS = "aws"
	CLOUD_ALI = "aliyun"
	CLOUD_HW  = "huawei"

	API_AWS_META_DATA = "http://169.254.169.254/latest/meta-data/placement/availability-zone"
	API_ALI_META_DATA = "http://100.100.100.200/latest/meta-data/zone-id"
	API_HW_META_DATA  = "http://169.254.169.254/latest/meta-data/placement/availability-zone"
)

// Logger for log
type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

func Zone(cloud string) string {
	var api string
	switch cloud {
	case CLOUD_AWS:
		api = API_AWS_META_DATA
	case CLOUD_ALI:
		api = API_ALI_META_DATA
	case API_HW_META_DATA:
		api = API_HW_META_DATA
	default:
		api = API_AWS_META_DATA
	}

	out, err := exec.Command("curl", "-s", api).Output()
	if err != nil {
		return "unknown"
	}

	return string(out)
}

func IntPseudoRandom(min, max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(max-min+1) + min
}

func IntGenuineRandom(min, max int64) int64 {
	res, _ := crand.Int(crand.Reader, big.NewInt(max-min+1))
	return res.Int64() + min
}

func FloatPseudoRandom() float64 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Float64()
}
