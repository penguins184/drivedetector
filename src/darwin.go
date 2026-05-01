//go:build darwin
package drivedetector;

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"regexp"
);

func Detect() ([]string, error) {
	var drives []string;
	dMap := make(map[string]bool);

	pattern := regexp.MustCompile(`Mount Point: (.+)$`);

	cmd := exec.Command("system_profiler", "SPUSBDataType");
	out, err := cmd.Output();

	if err != nil {
		return nil, err;
	};

	s := bufio.NewScanner(bytes.NewReader(out));
	for s.Scan() {
		line := s.Text();
		if pattern.MatchString(line) {
			match := pattern.FindStringSubmatch(line);
			if len(match) > 1 {
				d := match[1];
				dMap[d] = true;
			};
		};
	};

	for k := range dMap {
		if _, err := os.Stat(k); err == nil {
			drives = append(drives, k);
		};
	};

	return drives, nil;
};
