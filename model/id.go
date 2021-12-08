package model

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ID int64

func StringID(s string) (ID, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return ID(i), nil
}

func (id ID) Int64() int64 {
	return int64(id)
}

func (id ID) String() string {
	return strconv.FormatInt(id.Int64(), 10)
}

func (id *ID) UnmarshalJSON(b []byte) (err error) {
	var i int
	if err = json.Unmarshal(b, &i); err == nil {
		*id = ID(i)
		return nil
	}

	var s string
	if err = json.Unmarshal(b, &s); err == nil {
		if s == "" {
			return nil
		}
		i, err = strconv.Atoi(s)
		if err != nil {
			return err
		}
		*id = ID(i)
		return
	}

	return fmt.Errorf("error otrs TypeID(%v) parse err=%v", string(b), err)
}


type PendingTime struct{
	Diff int `json:"Diff"`
}


type Time int

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	var i int
	if err = json.Unmarshal(b, &i); err == nil {
		*t = Time(i)
		return nil
	}

	var s string
	if err = json.Unmarshal(b, &s); err == nil {
		if s == "" {
			return nil
		}
		i, err = strconv.Atoi(s)
		if err != nil {
			return err
		}
		*t = Time(i)
		return
	}

	return fmt.Errorf("error otrs TypeTime(%v) parse err=%v", string(b), err)
}