package trans

/*
  Config
  @Description: 翻译配置
*/
type Config struct {
	// 翻译文本配置文件路径[csv格式]
	FilePath string `yaml:"file_path"`
	// 语言代码表链接：http://www.lingoes.net/zh/translator/langcode.htm
	// 翻译主语言，例如：以中文为准，输入中文，翻译其他语言，默认值: zh
	SourceLan string `yaml:"source_lan"`
}

// defaultConfig
func defaultConfig(config *Config) *Config {
	if config == nil {
		return &Config{
			FilePath:  "./trans.csv",
			SourceLan: "zh",
		}
	}
	if len(config.FilePath) <= 0 {
		config.FilePath = "./trans.csv"
	}
	if len(config.SourceLan) <= 0 {
		config.SourceLan = "zh"
	}

	return config
}

/*
  Trans
  @Description: 翻译数据
*/
type Trans struct {
	// config
	Config *Config
	// data
	Data map[string]map[string]string
	// 主语言
	sourceLanIndex int
	// 语言代码表
	langCodes []string
}

var (
	trans *Trans
)

/*
  Reload
  @Desc: 加载配置表
  @param: config
*/
func Reload(config *Config) error {
	trans = &Trans{
		Config: defaultConfig(config),
		Data:   make(map[string]map[string]string),
	}

	// reload data
	return trans.reload()
}

/*
  Translate
  @Desc: 翻译
  @param: lanCode 需要翻译的语言代码
  @param: text 需要翻译的文本
  @return: string
*/
func Translate(lanCode, text string) string {
	return trans.translate(lanCode, text)
}
