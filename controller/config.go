package controller

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/bo"
	"h-ui/model/dto"
	"h-ui/model/vo"
	"h-ui/service"
)

func UpdateConfig(c *gin.Context) {
	configUpdateDto, err := validateField(c, dto.ConfigUpdateDto{})
	if err != nil {
		return
	}
	if err = service.UpdateConfig(*configUpdateDto.Key, *configUpdateDto.Value); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func ListConfig(c *gin.Context) {
	configDto, err := validateField(c, dto.ConfigDto{})
	if err != nil {
		return
	}
	configs, err := service.ListConfig(configDto.Keys)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(configs, c)
}

func GetHysteria2Config(c *gin.Context) {
	config, err := service.GetHysteria2Config()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(config, c)
}

func UpdateHysteria2Config(c *gin.Context) {
	hysteria2ServerConfig, err := validateField(c, bo.Hysteria2ServerConfig{})
	if err != nil {
		return
	}
	if err = service.UpdateHysteria2Config(hysteria2ServerConfig); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}
