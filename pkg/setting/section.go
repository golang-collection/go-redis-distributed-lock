package setting

/**
* @Author: super
* @Date: 2020-09-18 08:32
* @Description: 读取配置
**/

type CacheSettingS struct {
	UserName  string
	Password  string
	Host      string
	MaxIdle   int
	MaxActive int
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
