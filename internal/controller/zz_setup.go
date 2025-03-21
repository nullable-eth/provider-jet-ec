/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/crossplane/terrajet/pkg/terraform"

	elasticsearchkeystore "github.com/crossplane-contrib/provider-jet-ec/internal/controller/deployment/elasticsearchkeystore"
	extension "github.com/crossplane-contrib/provider-jet-ec/internal/controller/deployment/extension"
	trafficfilter "github.com/crossplane-contrib/provider-jet-ec/internal/controller/deployment/trafficfilter"
	trafficfilterassociation "github.com/crossplane-contrib/provider-jet-ec/internal/controller/deployment/trafficfilterassociation"
	deployment "github.com/crossplane-contrib/provider-jet-ec/internal/controller/ec/deployment"
	providerconfig "github.com/crossplane-contrib/provider-jet-ec/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		elasticsearchkeystore.Setup,
		extension.Setup,
		trafficfilter.Setup,
		trafficfilterassociation.Setup,
		deployment.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
