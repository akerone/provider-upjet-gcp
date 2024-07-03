// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Function.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Function) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.BuildConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("artifact.gcp.upbound.io", "v1beta2", "RegistryRepository", "RegistryRepositoryList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BuildConfig.DockerRepository),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.BuildConfig.DockerRepositoryRef,
				Selector:     mg.Spec.ForProvider.BuildConfig.DockerRepositorySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.BuildConfig.DockerRepository")
		}
		mg.Spec.ForProvider.BuildConfig.DockerRepository = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.BuildConfig.DockerRepositoryRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.BuildConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BuildConfig.ServiceAccount),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.BuildConfig.ServiceAccountRef,
				Selector:     mg.Spec.ForProvider.BuildConfig.ServiceAccountSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.BuildConfig.ServiceAccount")
		}
		mg.Spec.ForProvider.BuildConfig.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.BuildConfig.ServiceAccountRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.BuildConfig != nil {
		if mg.Spec.ForProvider.BuildConfig.Source != nil {
			if mg.Spec.ForProvider.BuildConfig.Source.StorageSource != nil {
				{
					m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta2", "Bucket", "BucketList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BuildConfig.Source.StorageSource.Bucket),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.ForProvider.BuildConfig.Source.StorageSource.BucketRef,
						Selector:     mg.Spec.ForProvider.BuildConfig.Source.StorageSource.BucketSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.BuildConfig.Source.StorageSource.Bucket")
				}
				mg.Spec.ForProvider.BuildConfig.Source.StorageSource.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.BuildConfig.Source.StorageSource.BucketRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.ForProvider.BuildConfig != nil {
		if mg.Spec.ForProvider.BuildConfig.Source != nil {
			if mg.Spec.ForProvider.BuildConfig.Source.StorageSource != nil {
				{
					m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta2", "BucketObject", "BucketObjectList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BuildConfig.Source.StorageSource.Object),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.ForProvider.BuildConfig.Source.StorageSource.ObjectRef,
						Selector:     mg.Spec.ForProvider.BuildConfig.Source.StorageSource.ObjectSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.BuildConfig.Source.StorageSource.Object")
				}
				mg.Spec.ForProvider.BuildConfig.Source.StorageSource.Object = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.BuildConfig.Source.StorageSource.ObjectRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.ForProvider.BuildConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudbuild.gcp.upbound.io", "v1beta2", "WorkerPool", "WorkerPoolList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BuildConfig.WorkerPool),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.BuildConfig.WorkerPoolRef,
				Selector:     mg.Spec.ForProvider.BuildConfig.WorkerPoolSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.BuildConfig.WorkerPool")
		}
		mg.Spec.ForProvider.BuildConfig.WorkerPool = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.BuildConfig.WorkerPoolRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.EventTrigger != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.EventTrigger.EventFilters); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta2", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventTrigger.EventFilters[i4].Value),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.EventTrigger.EventFilters[i4].ValueRef,
					Selector:     mg.Spec.ForProvider.EventTrigger.EventFilters[i4].ValueSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EventTrigger.EventFilters[i4].Value")
			}
			mg.Spec.ForProvider.EventTrigger.EventFilters[i4].Value = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EventTrigger.EventFilters[i4].ValueRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.EventTrigger != nil {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventTrigger.PubsubTopic),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.EventTrigger.PubsubTopicRef,
				Selector:     mg.Spec.ForProvider.EventTrigger.PubsubTopicSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EventTrigger.PubsubTopic")
		}
		mg.Spec.ForProvider.EventTrigger.PubsubTopic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EventTrigger.PubsubTopicRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.EventTrigger != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventTrigger.ServiceAccountEmail),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.ForProvider.EventTrigger.ServiceAccountEmailRef,
				Selector:     mg.Spec.ForProvider.EventTrigger.ServiceAccountEmailSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EventTrigger.ServiceAccountEmail")
		}
		mg.Spec.ForProvider.EventTrigger.ServiceAccountEmail = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EventTrigger.ServiceAccountEmailRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.ServiceConfig != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ServiceConfig.SecretEnvironmentVariables); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta2", "Secret", "SecretList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceConfig.SecretEnvironmentVariables[i4].Secret),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.ServiceConfig.SecretEnvironmentVariables[i4].SecretRef,
					Selector:     mg.Spec.ForProvider.ServiceConfig.SecretEnvironmentVariables[i4].SecretSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ServiceConfig.SecretEnvironmentVariables[i4].Secret")
			}
			mg.Spec.ForProvider.ServiceConfig.SecretEnvironmentVariables[i4].Secret = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ServiceConfig.SecretEnvironmentVariables[i4].SecretRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.ServiceConfig != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ServiceConfig.SecretVolumes); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta2", "Secret", "SecretList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceConfig.SecretVolumes[i4].Secret),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.ServiceConfig.SecretVolumes[i4].SecretRef,
					Selector:     mg.Spec.ForProvider.ServiceConfig.SecretVolumes[i4].SecretSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ServiceConfig.SecretVolumes[i4].Secret")
			}
			mg.Spec.ForProvider.ServiceConfig.SecretVolumes[i4].Secret = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ServiceConfig.SecretVolumes[i4].SecretRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.ServiceConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceConfig.ServiceAccountEmail),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.ForProvider.ServiceConfig.ServiceAccountEmailRef,
				Selector:     mg.Spec.ForProvider.ServiceConfig.ServiceAccountEmailSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ServiceConfig.ServiceAccountEmail")
		}
		mg.Spec.ForProvider.ServiceConfig.ServiceAccountEmail = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ServiceConfig.ServiceAccountEmailRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.BuildConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("artifact.gcp.upbound.io", "v1beta2", "RegistryRepository", "RegistryRepositoryList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BuildConfig.DockerRepository),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.BuildConfig.DockerRepositoryRef,
				Selector:     mg.Spec.InitProvider.BuildConfig.DockerRepositorySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.BuildConfig.DockerRepository")
		}
		mg.Spec.InitProvider.BuildConfig.DockerRepository = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.BuildConfig.DockerRepositoryRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.BuildConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BuildConfig.ServiceAccount),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.BuildConfig.ServiceAccountRef,
				Selector:     mg.Spec.InitProvider.BuildConfig.ServiceAccountSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.BuildConfig.ServiceAccount")
		}
		mg.Spec.InitProvider.BuildConfig.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.BuildConfig.ServiceAccountRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.BuildConfig != nil {
		if mg.Spec.InitProvider.BuildConfig.Source != nil {
			if mg.Spec.InitProvider.BuildConfig.Source.StorageSource != nil {
				{
					m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta2", "Bucket", "BucketList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BuildConfig.Source.StorageSource.Bucket),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.InitProvider.BuildConfig.Source.StorageSource.BucketRef,
						Selector:     mg.Spec.InitProvider.BuildConfig.Source.StorageSource.BucketSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.BuildConfig.Source.StorageSource.Bucket")
				}
				mg.Spec.InitProvider.BuildConfig.Source.StorageSource.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.BuildConfig.Source.StorageSource.BucketRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.InitProvider.BuildConfig != nil {
		if mg.Spec.InitProvider.BuildConfig.Source != nil {
			if mg.Spec.InitProvider.BuildConfig.Source.StorageSource != nil {
				{
					m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta2", "BucketObject", "BucketObjectList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BuildConfig.Source.StorageSource.Object),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.InitProvider.BuildConfig.Source.StorageSource.ObjectRef,
						Selector:     mg.Spec.InitProvider.BuildConfig.Source.StorageSource.ObjectSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.BuildConfig.Source.StorageSource.Object")
				}
				mg.Spec.InitProvider.BuildConfig.Source.StorageSource.Object = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.BuildConfig.Source.StorageSource.ObjectRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.InitProvider.BuildConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudbuild.gcp.upbound.io", "v1beta2", "WorkerPool", "WorkerPoolList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BuildConfig.WorkerPool),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.BuildConfig.WorkerPoolRef,
				Selector:     mg.Spec.InitProvider.BuildConfig.WorkerPoolSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.BuildConfig.WorkerPool")
		}
		mg.Spec.InitProvider.BuildConfig.WorkerPool = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.BuildConfig.WorkerPoolRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.EventTrigger != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.EventTrigger.EventFilters); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta2", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventTrigger.EventFilters[i4].Value),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.EventTrigger.EventFilters[i4].ValueRef,
					Selector:     mg.Spec.InitProvider.EventTrigger.EventFilters[i4].ValueSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.EventTrigger.EventFilters[i4].Value")
			}
			mg.Spec.InitProvider.EventTrigger.EventFilters[i4].Value = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.EventTrigger.EventFilters[i4].ValueRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.EventTrigger != nil {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventTrigger.PubsubTopic),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.EventTrigger.PubsubTopicRef,
				Selector:     mg.Spec.InitProvider.EventTrigger.PubsubTopicSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EventTrigger.PubsubTopic")
		}
		mg.Spec.InitProvider.EventTrigger.PubsubTopic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EventTrigger.PubsubTopicRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.EventTrigger != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventTrigger.ServiceAccountEmail),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.InitProvider.EventTrigger.ServiceAccountEmailRef,
				Selector:     mg.Spec.InitProvider.EventTrigger.ServiceAccountEmailSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EventTrigger.ServiceAccountEmail")
		}
		mg.Spec.InitProvider.EventTrigger.ServiceAccountEmail = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EventTrigger.ServiceAccountEmailRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.ServiceConfig != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ServiceConfig.SecretEnvironmentVariables); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta2", "Secret", "SecretList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceConfig.SecretEnvironmentVariables[i4].Secret),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.ServiceConfig.SecretEnvironmentVariables[i4].SecretRef,
					Selector:     mg.Spec.InitProvider.ServiceConfig.SecretEnvironmentVariables[i4].SecretSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.ServiceConfig.SecretEnvironmentVariables[i4].Secret")
			}
			mg.Spec.InitProvider.ServiceConfig.SecretEnvironmentVariables[i4].Secret = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.ServiceConfig.SecretEnvironmentVariables[i4].SecretRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.ServiceConfig != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ServiceConfig.SecretVolumes); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta2", "Secret", "SecretList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceConfig.SecretVolumes[i4].Secret),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.ServiceConfig.SecretVolumes[i4].SecretRef,
					Selector:     mg.Spec.InitProvider.ServiceConfig.SecretVolumes[i4].SecretSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.ServiceConfig.SecretVolumes[i4].Secret")
			}
			mg.Spec.InitProvider.ServiceConfig.SecretVolumes[i4].Secret = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.ServiceConfig.SecretVolumes[i4].SecretRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.ServiceConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceConfig.ServiceAccountEmail),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.InitProvider.ServiceConfig.ServiceAccountEmailRef,
				Selector:     mg.Spec.InitProvider.ServiceConfig.ServiceAccountEmailSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ServiceConfig.ServiceAccountEmail")
		}
		mg.Spec.InitProvider.ServiceConfig.ServiceAccountEmail = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ServiceConfig.ServiceAccountEmailRef = rsp.ResolvedReference

	}

	return nil
}
