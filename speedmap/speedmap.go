package speedmap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Point struct {
	Temperature float32 `json:"temperature"`
	Speed       float32 `json:"speed"`
}

func (p *Point) String() string {
	return fmt.Sprintf("t: %0.2f, s: %0.2f", p.Temperature, p.Speed)
}

var points = []Point{}

func insertPointAtIndex(p Point, i int) {
	points = append(points[:i+1], points[i:]...)
	points[i] = p
}

func updatePointAtIndex(p Point, i int) error {
	if points[i].Temperature != p.Temperature {
		return errors.New("bad speed map update")
	}
	if points[i].Speed < p.Speed {
		points[i].Speed = p.Speed
	}
	return nil
}

func init() {
	Reset()
}

func Clear() {
	points = []Point{}
}

func Reset() {
	points = []Point{
		{0.0, 0.0},
		{80.0, 100.0},
	}
}

func AddPoint(point Point) error {
	for i := 0; i < len(points); i++ {
		if point.Temperature < points[i].Temperature {
			insertPointAtIndex(point, i)
			return nil
		}
		if point.Temperature == points[i].Temperature {
			return updatePointAtIndex(point, i)
		}
	}
	points = append(points, point)
	return nil
}

func SetPointMap(newPointMap []Point) error {

	if len(newPointMap) == 0 {
		Reset()
		return errors.New("empty point map")
	}

	Clear()
	var err error = nil

	for _, point := range newPointMap {
		err = AddPoint(point)
		if err != nil {
			Reset()
			return err
		}
	}

	return nil
}

func Points() []Point {
	return points
}

func LoadFromFile(path string) error {
	argFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer argFile.Close()
	byteValue, err := ioutil.ReadAll(argFile)
	if err != nil {
		return err
	}
	var argSpeedMap []Point
	err = json.Unmarshal(byteValue, &argSpeedMap)
	if err != nil {
		return err
	}
	return SetPointMap(argSpeedMap)
}

func interpolateSpeed(temperature float32, l, r Point) float32 {
	m := (r.Speed - l.Speed) / (r.Temperature - l.Temperature)
	return (m * temperature) + (l.Speed - (m * l.Temperature))
}
func GetFanSpeed(temperature float32) float32 {
	split := 0

	for ; split < len(points); split++ {
		if temperature < points[split].Temperature {
			break
		}
	}

	if split == 0 {
		fmt.Println("bottom")
		return points[split].Speed
	}

	return interpolateSpeed(temperature, points[split-1], points[split])
}
