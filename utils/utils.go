package utils

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	strConst = "abcdefghijklmnoprstuvwxyz0123456789"
)

func Test() {
	///
}

func CreateFile(filePath string, data interface{}, fileType string) (string, error) {

	if file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
		return "open file failed", err
	} else {
		switch fileType {
		case "yaml":
			yamlEncoder := yaml.NewEncoder(file)
			yamlEncoder.SetIndent(2)
			if err := yamlEncoder.Encode(&data); err != nil {
				return "encode failed", err
			}
		case "json":
			jsonEncoder := json.NewEncoder(file)
			if err := jsonEncoder.Encode(&data); err != nil {
				return "encode failed", err
			}

		}

		return filePath, nil
	}

}

func FileConv(filepath string, out interface{}) error {

	if inFile, err := os.Open(filepath); err != nil {

		return err
	} else {
		if stats, err := inFile.Stat(); err != nil {
			return err

		} else {
			byteArr := make([]byte, stats.Size())
			if _, err := inFile.Read(byteArr); err != nil {
				return err
			} else {
				yaml.Unmarshal(byteArr, out)
			}
		}

	}

	return nil
}

func GetHostIp(netInteface string) (string, error) {
	var ipAddr string
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err

	}
	for _, intf := range interfaces {
		if intf.Name == netInteface {
			addrs, err := intf.Addrs()
			if err != nil {
				return "", err
			}
			for _, addr := range addrs {
				ipnet, ok := addr.(*net.IPNet)
				if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					ipAddr = ipnet.IP.String()
					//fmt.Println(ipAddr.IP)
				}
			}
		}
		continue
	}
	//strIp := ipAddr.IP.String()
	return ipAddr, nil
}

func GetPoNamePgName() (string, string, error) {
	//var podName string

	var pgInsName string

	podName, err := os.Hostname()
	if err != nil {

		return "", "", err

	}
	pgInsName = strings.Split(podName, "-")[3]

	return podName, pgInsName, nil
}
func GetRandomName(n int) string {
	btArr := make([]byte, n)
	//rand.Intn(len(strConst))
	for i, _ := range btArr {
		//rand.Int(len(strConst))

		btArr[i] = strConst[rand.Intn(len(strConst))]

	}
	randStr := string(btArr)
	return randStr
}

func ExecCmd(cmdName string, cmdOpt string, cmd_args ...string) (string, error) {

	cmd := exec.Command(cmdName, append([]string{cmdOpt}, cmd_args...)...)
	//fmt.Println(cmd)
	//cmd := exec.Command("kubectl", "exec", cmd_args[0], "--", "sed", "-i", sed_rex, sed_file)
	var output = bytes.NewBuffer(nil)
	cmd.Stdout = output
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {

		return output.String(), err

	}

	return output.String(), nil

}
