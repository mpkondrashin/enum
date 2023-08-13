# Enum - autmatically generate enums for Go

Enum prvide ability to create enumirated values with unmarshaling from JSON and YAML. 

## Example

Command 
```commandline
enum -package main -type TrafficLight -values=Green,Yellow,Red
```
will generate following code:
```golang
type TrafficLight int

const (
    TrafficLightGreen TrafficLight = iota
    TrafficLightYellow
    TrafficLightRed
)

// String - return string representation for TrafficLight value
func (v TrafficLight)String() string {
   ...
}

// UnmarshalJSON implements the Unmarshaler interface of the json package for TrafficLight.
func (s *TrafficLight) UnmarshalJSON(data []byte) error {
    ...
}

// MarshalJSON implements the Marshaler interface of the json package for TrafficLight.
func (s TrafficLight) MarshalJSON() ([]byte, error) {
    ...
}

// UnmarshalYAML implements the Unmarshaler interface of the yaml.v3 package for TrafficLight.
func (s *TrafficLight) UnmarshalYAML(unmarshal func(interface{}) error) error {
    ...
}
```

## Installation
```commandline
go install github.com/mpkondrashin/enum
```

## Using go generate

Add following line to your golang code:
```golang
//enum -package=<name of the package> -type=<name of the enum type> -values=<coma separated list of the values>
```
Then run following command:
```commandline
go generate
```
