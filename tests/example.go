package main;

import (
	"github.com/penguins184/drivedetector"
);

func main() {
	drives, err := drivedetector.Detect();

	if err != nil {
		panic(err);
	};

	for _, d := range drives {
		println(d);
	};
};