package fileimpl

import (
	"encoding/json"

	"github.com/aide-family/sovereign/pkg/enum"
	namespacev1 "github.com/aide-family/sovereign/pkg/repo/namespace/v1"
	"github.com/aide-family/sovereign/pkg/repo/namespace/v1/fileimpl/model"
)

func convertNamespaceModel(namespaceModel *model.NamespaceModel) *namespacev1.NamespaceModel {
	return &namespacev1.NamespaceModel{
		Id:        namespaceModel.ID,
		Uid:       namespaceModel.UID,
		Name:      namespaceModel.Name,
		Metadata:  namespaceModel.Metadata,
		Status:    namespaceModel.Status,
		CreatedAt: namespaceModel.CreatedAt,
		UpdatedAt: namespaceModel.UpdatedAt,
		DeletedAt: namespaceModel.DeletedAt,
		Creator:   namespaceModel.Creator,
	}
}

func convertNamespaceItemSelect(namespaceModel *model.NamespaceModel) *namespacev1.NamespaceItemSelect {
	metadata, _ := json.Marshal(namespaceModel.Metadata)
	return &namespacev1.NamespaceItemSelect{
		Value:    namespaceModel.UID,
		Label:    namespaceModel.Name,
		Disabled: namespaceModel.DeletedAt != 0 || namespaceModel.Status != enum.GlobalStatus_ENABLED,
		Tooltip:  string(metadata),
	}
}
