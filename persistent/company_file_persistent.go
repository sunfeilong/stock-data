package persistent

import (
	"encoding/json"
	"errors"
	"github.com/xiaotian/stock/config"
	"github.com/xiaotian/stock/model"
	"github.com/xiaotian/stock/s-logger"
	"github.com/xiaotian/stock/tool"
	"io/ioutil"
)

var logger = s_logger.New()

type CompanyFilePreserver struct{}

func (c CompanyFilePreserver) Save(data []model.Company) error {
	pathName := config.GetDataSaveFilePath()
	marshal, err := json.Marshal(data)
	if nil != err {
		logger.Infow("保存数据到文件,数据格式化异常", "pathName", pathName, "err", err)
		return errors.New("保存数据到文件,数据格式化异常")
	}
	err = ioutil.WriteFile(pathName+c.getFullFileName(tool.NowDate()), marshal, 0664)
	if nil != err {
		logger.Infow("保存数据到文件,写入文件数据异常", "pathName", pathName, "err", err)
		return errors.New("保存数据到文件,写入文件数据异常")
	}
	return nil
}

func (c CompanyFilePreserver) Read() ([]model.Company, error) {
	pathName := config.GetDataSaveFilePath()
	file, err := ioutil.ReadFile(pathName + c.getFullFileName(tool.NowDate()))
	if nil != err {
		logger.Infow("从文件读取数据,读取数据异常", "pathName", pathName, "err", err)
		return nil, errors.New("从文件读取数据,读取数据异常")
	}

	d := &[]model.Company{}
	err = json.Unmarshal(file, &d)
	if nil != err {
		logger.Infow("从文件读取数据,解析数据异常", "pathName", pathName, "err", err)
		return nil, errors.New("从文件读取数据,读取数据异常")
	}
	return *d, nil
}

func (c CompanyFilePreserver) getFullFileName(append string) string {
	return c.getPrefix() + append + c.getSuffix()
}

func (c CompanyFilePreserver) getPrefix() string {
	return "company-"
}

func (c CompanyFilePreserver) getSuffix() string {
	return ".json"
}
