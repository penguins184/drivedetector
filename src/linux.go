//go:build linux
package drivedetector;

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"strings"
);

func Detect() ([]string, error) {
	var drives []string;
	dMap := make(map[string]bool);
	pattern := regexp.MustCompile(`^(/[^\s]+)\s+.*?\s+([/].*)$`);

	cmd := exec.Command("df");
	out, err := cmd.Output();
	if err != nil {
		return nil, err;
	};

	s := bufio.NewScanner(bytes.NewReader(out));
	for s.Scan() {
		line := s.Text();
		if pattern.MatchString(line) {
			match := pattern.FindStringSubmatch(line);
			device := match[1];
			path := match[2];

			if check(device) {
				dMap[path] = true;
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

func check(device string) bool {
	verify := "ID_USB_DRIVER=usb-storage";

	cmd := exec.Command("udevadm", "info", "-q", "property", "-n", device);
	out, err := cmd.Output();

	if err != nil {
		return false;
	};

	if strings.Contains(string(out), verify) {
		return true;
	};

	return false;
};
