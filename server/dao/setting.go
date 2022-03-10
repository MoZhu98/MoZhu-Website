/*
Package dao
@Author: MoZhu
@File: setting
@Software: GoLand
*/
package dao

type Setting struct {
	ID    uint   `gorm:"primary_key"`
	Key   string `json:"key"`
	Value string `gorm:"type:text"`
}

func (Setting) TableName() string {
	return "settings"
}

func NewSettingDAO(client *MySQLClient) SettingDAOIF {
	return &settingDAO{
		client: client,
	}
}

type SettingDAOIF interface {
	GetSettings() (map[string]string, error)
}

type settingDAO struct {
	client Client
}

func (s *settingDAO) GetSettings() (map[string]string, error) {
	var settings []Setting
	err := s.client.DB().Order("id desc").Find(&settings).Error
	result := map[string]string{}
	for _, s := range settings {
		result[s.Key] = s.Value
	}
	return result, err
}

func GetSettingsByKeys(keys []string) (map[string]string, error) {
	var settings []Setting
	result := map[string]string{}
	for _, s := range settings {
		result[s.Key] = s.Value
	}
	return result, nil
}
