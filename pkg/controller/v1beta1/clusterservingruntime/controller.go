/*
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

package clusterservingruntime

import (
	"context"
	"fmt"

	v1beta1api "github.com/kserve/kserve/pkg/apis/serving/v1beta1"

	appsv1 "k8s.io/api/apps/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// ClusterServiceRuntimeReconciler reconciles a ClusterServiceRuntime object
type ClusterServiceRuntimeReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=serving.kserve.io,resources=clusterservingruntimes;clusterservingruntimes/finalizers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=serving.kserve.io,resources=clusterservingruntimes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments;deployments/finalizers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=services;configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch

func (r *ClusterServiceRuntimeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//Fetch the ServingRuntime instance
	csrt := &v1beta1api.ClusterServingRuntime{}
	if err := r.Client.Get(ctx, req.NamespacedName, csrt); err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, nil
	}

	// If invalid ServerType is provided in rt.Spec or if this value doesn't match with that of the specified container, delete the deployment
	if err := validateServingRuntimeSpec(csrt); err != nil {
		werr := fmt.Errorf("Invalid ServingRuntime Spec: %w", err)
		return ctrl.Result{}, werr
	}
	return ctrl.Result{}, nil
}

func validateServingRuntimeSpec(csrt *v1beta1api.ClusterServingRuntime) error {
	return nil
}

func (r *ClusterServiceRuntimeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1api.ClusterServingRuntime{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
