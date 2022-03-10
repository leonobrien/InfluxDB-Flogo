package influxdbv2

import (
	"fmt"

	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Host         string                 `md:"Host.required"`
	Organisation string                 `md:"Organisation.required"`
	Bucket       string                 `md:"Bucket.required"`
	Token        string                 `md:"Token.required"`
	Measurement  string                 `md:"Measurement.required"`
	Value        map[string]interface{} `md:"Value"`
}

type Output struct {
	Output string `md:"Output"`
}

func (r *Input) FromMap(values map[string]interface{}) error {

	Val1, _ := coerce.ToString(values["Host"])
	r.Host = Val1

	fmt.Println(values["Organisation"])
	Val2, _ := coerce.ToString(values["Organisation"])
	r.Organisation = Val2

	Val3, _ := coerce.ToString(values["Bucket"])
	r.Bucket = Val3

	Val4, _ := coerce.ToString(values["Token"])
	r.Token = Val4

	Val6, _ := coerce.ToString(values["Measurement"])
	r.Measurement = Val6

	Val5, _ := coerce.ToObject(values["Value"])
	//Val4, _ := coerce.ToParams(values["values"])
	r.Value = Val5
	fmt.Println(Val5)
	fmt.Println(r.Value)

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Host":         r.Host,
		"Organisation": r.Organisation,
		"Table":        r.Bucket,
		"Token":        r.Token,
		"Measurement":  r.Measurement,
		"Value":        r.Value,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["Output"])
	o.Output = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Output": o.Output,
	}
}
