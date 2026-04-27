package service

import (
	"github.com/6dainas/6dainas-app-management/codegen"
	"github.com/6dainas/6dainas-app-management/common"
	"github.com/6dainas/6dainas-os-Common/utils/logger"
	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
)

type App types.ServiceConfig

func (a *App) StoreInfo() (codegen.AppStoreInfo, error) {
	var storeInfo codegen.AppStoreInfo

	ex, ok := a.Extensions[common.ComposeExtensionNameXCasaOS]
	if !ok {
		logger.Error("extension `x-6dainas` not found")
		// return storeInfo, ErrComposeExtensionNameXCasaOSNotFound
	}

	// add image to store info for check stable version function.
	storeInfo.Image = a.Image

	if err := loader.Transform(ex, &storeInfo); err != nil {
		return storeInfo, err
	}

	return storeInfo, nil
}
