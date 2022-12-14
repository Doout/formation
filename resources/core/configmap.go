package core

import (
	"github.com/Doout/formation/types"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ConfigMap struct {
	cm            *v1.ConfigMap
	DisableUpdate bool
}

func NewConfigMap(cm *v1.ConfigMap) *ConfigMap {
	return &ConfigMap{cm: cm}
}

func (c *ConfigMap) Type() string { return "configmap" }

func (c *ConfigMap) Name() string { return c.cm.Name }

func (c *ConfigMap) Runtime() client.Object { return &v1.ConfigMap{} }

func (c *ConfigMap) Create() (client.Object, error) {
	if c.DisableUpdate {
		if c.cm.Annotations == nil {
			c.cm.Annotations = make(map[string]string)
		}
		c.cm.Annotations[types.UpdateKey] = "disabled"
	}
	return c.cm, nil

}
