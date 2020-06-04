package persistent

import (
    "encoding/json"
    "errors"
    "github.com/xiaotian/stock/pkg/model"
    "github.com/xiaotian/stock/pkg/tool"
    "io/ioutil"
    "os"
)

type DataFilePreserver struct{}

func (c DataFilePreserver) Save(data []model.Data) error {
    path, err := tool.GetPath(pathName, maxLevel, maxLevel)
    if nil != err {
        logger.Infow("保存数据到文件,未找到配置路径", "pathName", pathName, "err", err)
        return errors.New("保存数据到文件,未找到配置路径")
    }
    marshal, err := json.MarshalIndent(data, "", "  ")
    if nil != err {
        logger.Infow("保存数据到文件,数据格式化异常", "pathName", pathName, "err", err)
        return errors.New("保存数据到文件,数据格式化异常")
    }
    err = ioutil.WriteFile(path+c.getFullFileName(tool.NowDate()), marshal, os.ModeAppend)
    if nil != err {
        logger.Infow("保存数据到文件,写入文件数据异常", "pathName", pathName, "err", err)
        return errors.New("保存数据到文件,写入文件数据异常")
    }
    return nil
}

func (c DataFilePreserver) Read() ([]model.Data, error) {
    path, err := tool.GetPath(pathName, maxLevel, maxLevel)
    if nil != err {
        logger.Infow("从文件读取数据,未找到配置路径", "pathName", pathName, "err", err)
        return nil, errors.New("从文件读取数据,未找到配置路径")
    }
    file, err := ioutil.ReadFile(path + c.getFullFileName(tool.NowDate()))
    if nil != err {
        logger.Infow("从文件读取数据,读取数据异常", "pathName", pathName, "err", err)
        return nil, errors.New("从文件读取数据,读取数据异常")
    }

    d := &[]model.Data{}
    err = json.Unmarshal(file, d)
    if nil != err {
        logger.Infow("从文件读取数据,解析数据异常", "pathName", pathName, "err", err)
        return nil, errors.New("从文件读取数据,读取数据异常")
    }
    return *d, nil
}

func (c DataFilePreserver) getFullFileName(append string) string {
    return c.getPrefix() + append + c.getSuffix()
}

func (c DataFilePreserver) getPrefix() string {
    return "data-"
}

func (c DataFilePreserver) getSuffix() string {
    return ".json"
}
