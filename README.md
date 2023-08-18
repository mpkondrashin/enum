# Enum - autmatically generate enums for Go

Enum provides ability to create enumirated values with support of unmarshaling from JSON and YAML. 

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
Then run the following command:
```commandline
go generate
```

## Commandline options
- -noprefix - do not add type name as prefix to values
- -output string - output filename (default enum_<type name>.go)
- -package string - package name
- -type string - name of the enum type. It will be alias to int
- -values string - comma-separated list of values names

## Bugs

Default file name is enum_<type name lower case>.go, so for two types differ by characters case, will overwrite same file. 
Use  -output option to avoid this problem.