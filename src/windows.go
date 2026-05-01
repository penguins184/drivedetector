//go:build windows
package drivedetector;

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"strings"
	"syscall"
);

func Detect() ([]string, error) {
	var drives []string;
	dMap := make(map[string]bool);

	command := "Get-CimInstance -ClassName Win32_LogicalDisk | Where-Object { $_.DriveType -eq 2 } | Select-Object -ExpandProperty DeviceID";
	
	cmd := exec.Command("powershell", "-NoProfile", "-Command", command);
	cmd.SysProcAttr = &syscall.SysProcAttr{ HideWindow: true };
	
	out, err := cmd.Output();
	if err != nil {
		return nil, err;
	};

	s := bufio.NewScanner(bytes.NewReader(out));
	for s.Scan() {
		line := strings.TrimSpace(s.Text());
		if line != "" && strings.Contains(line, ":") {

			path := line + string(os.PathSeparator);
			dMap[path] = true;
		}
	}

	for k := range dMap { //Ensure Mounted
		if _, err := os.Stat(k); err == nil {
			drives = append(drives, k);
		};
	};

	return drives, nil;
};
