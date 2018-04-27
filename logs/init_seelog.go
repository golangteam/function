package logs

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"strings"
	"github.com/cihub/seelog"
)

const (
	logConfig = `<seelog type="asynctimer" asyncinterval="5000000" minlevel="debug">
	<outputs formatid="main">
		<console/>
		<rollingfile type="size" filename="__log_url__" maxsize="1024000" maxrolls="10" />
	</outputs>
	<formats>
		<format id="main" format="%Date(2006-01-02 15:04:05) [%Level] %RelFile line:%Line %Msg%n"/>
	</formats>
</seelog>`
)
// init seelog
//
// @param pathConfig string  path of log.xml
// @param pathLog    string  path of logs
func InitSeeLog(pathConfig, pathLog string) {
	tmp := filepath.Dir(pathConfig)
	if _, err := os.Stat(tmp); os.IsNotExist(err) {
		os.MkdirAll(tmp, 0764)
	}
	tmp = filepath.Dir(pathLog)
	if _, err := os.Stat(tmp); os.IsNotExist(err) {
		os.MkdirAll(tmp, 0764)
	}
	if _, err := os.Stat(pathConfig); os.IsNotExist(err) {
		ioutil.WriteFile(pathConfig, []byte(strings.Replace(logConfig, "__log_url__", pathLog, 1)), 0764)
	}
	if logger, err := seelog.LoggerFromConfigAsFile(pathConfig); err == nil {
		seelog.ReplaceLogger(logger)
	}
}
// flush the seelog
func FlushSeelog()  {
	seelog.Flush()
}
