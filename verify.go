package main

import (
	"log"
	"strconv"
)

func verify(minConn string, maxConn string, min *int, max *int) {

	vMinConn := 0
	vMaxConn := 0
	vMinCond := false
	vMaxCond := false

	if minConn != "" {
		v, err := strconv.Atoi(minConn)
		if err != nil {
			log.Println("warning: DB_MIN_COONN has not a valid value, we will set min connections to", min)
			vMinCond = false
		} else {
			vMinConn = v
			vMinCond = true
		}
	}

	if maxConn != "" {
		v, err := strconv.Atoi(maxConn)
		if err != nil {
			log.Println("warning: DB_MAX_COONN has not a valid value, we will set max connections to", max)
			vMaxCond = false
		} else {
			vMaxConn = v
			vMaxCond = true
		}
	}

	if vMinCond && vMaxCond {
		if vMinConn <= vMaxConn {
			if vMinConn >= *min && vMinConn <= *max {
				*min = vMinConn
			} else {
				log.Printf("warning: DB_MIN_CONN has not a valid value: [%v, %v]", *min, *max)
			}
			if vMaxConn >= *min && vMaxConn <= *max {
				*max = vMaxConn
			} else {
				log.Printf("warning: DB_MAX_CONN has not a valid value: [%v, %v]", *min, *max)
			}
		} else {
			log.Printf("warning: DB_MAX_CONN: %v must be greater than DB_MIN_CONN: %v", vMaxConn, vMinConn)
		}
	}
}
