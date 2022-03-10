package influxdbv2

import (
	"fmt"
	"time"

	_ "github.com/influxdata/influxdb-client-go"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}

	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	//TODO - InfluxDB v2 code for using the influx library.

	fmt.Println("Initiating push to InfluxDB")

	//initiate client
	client := influxdb2.NewClient(input.Host, input.Token)
	defer client.Close()

	//create writeAPI object for Data Point writing.
	writeAPI := client.WriteAPI(input.Organisation, input.Bucket)
	//initiate point for measurement
	point := influxdb2.NewPoint(input.Measurement,
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())

	writeAPI.WritePoint(point)
	writeAPI.Flush()

	output := &Output{Output: "ok"}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
