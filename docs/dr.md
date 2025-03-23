# DR Commands

The dr command supports the following sub-commands:

* [init](#init)
* [test](#test)

## init

The init command crates a configuration file required for all other dr commands.

```bash
$ odf dr init
✅ Created config file "config.yaml" - please modify for your clusters
```

> [!IMPORTANT]
> Before using the config file, you need to edit it and configure your
> clusters and storage.

### Sample configuration file

```yaml
## odf dr configuration file

## Clusters configuration.
# - Modify clusters "kubeconfig" and "name" to match your hub and managed
#   clusters names and path to the kubeconfig file.
clusters:
  hub:
    name: hub
    kubeconfig: /Users/nir/.config/drenv/rdr/kubeconfigs/hub
  c1:
    name: dr1
    kubeconfig: /Users/nir/.config/drenv/rdr/kubeconfigs/dr1
  c2:
    name: dr2
    kubeconfig: /Users/nir/.config/drenv/rdr/kubeconfigs/dr2

## Kubernetes distribution
# - Modify to specify your clusters kubernetes distribution (k8s or ocp).
distro: k8s

## Git repository for test command.
# - Modify "url" to use your own Git repository.
# - Modify "branch" to test a different branch.
repo:
  url: https://github.com/RamenDR/ocm-ramen-samples.git
  branch: main

## DRPolicy for test command.
# - Modify to match actual DRPolicy in the hub cluster.
drPolicy: dr-policy

## ClusterSet for test command.
# - Modify to match your Open Cluster Management configuration.
clusterSet: default

## PVC specifications for test command.
# - Modify items "storageclassname" to match the actual storage classes in the
#   managed clusters.
# - Add new items for testing more storage types.
PVCSpecs:
- name: rbd
  storageClassName: rook-ceph-block
  accessModes: ReadWriteOnce
- name: cephfs
  storageClassName: rook-cephfs-fs1
  accessModes: ReadWriteMany

## Tests cases for test command.
# - Modify the test for your preferred workload or deployment type.
# - Add new tests for testing more combinations in parallel.
# - Available workloads: deploy
# - Available deployers: appset, subscr, disapp
tests:
- workload: deploy
  deployer: appset
  pvcSpec: rbd
```

## test

The command supports the following sub-commands:

* [run](#test-run)
* [clean](#test-clean)

### test run

The test command tests complete disaster recovery flow with a tiny application.

```bash
% ramenctl test run -o odf-dr-test
 ⭐ Using report "test"
 ⭐ Using config "config.yaml"

 🔎 Setup environment ...
    ✅ Environment setup

 🔎 Run tests ...
    ✅ Application "appset-deploy-rbd" deployed
    ✅ Application "appset-deploy-cephfs" deployed
    ✅ Application "appset-deploy-cephfs" protected
    ✅ Application "appset-deploy-rbd" protected
    ✅ Application "appset-deploy-cephfs" failed over
    ✅ Application "appset-deploy-rbd" failed over
    ✅ Application "appset-deploy-cephfs" relocated
    ✅ Application "appset-deploy-rbd" relocated
    ✅ Application "appset-deploy-cephfs" unprotected
    ✅ Application "appset-deploy-cephfs" undeployed
    ✅ Application "appset-deploy-rbd" unprotected
    ✅ Application "appset-deploy-rbd" undeployed

 ✅ passed (2 passed, 0 failed, 0 skipped)
```

The command stores `test-run.yaml` and `test-run.log` in the specified directory:

```bash
% tree odf-dr-test/
odf-dr-test/
├── test-run.log
└── test-run.yaml
```

> [!IMPORTANT]
> When reporting DR related issues, please create an archive with the output
> directory and upload it to the issue tracker.

### test clean

The clean command delete resources created by the [run](#test-run) command.

```bash
% ramenctl test clean -o odf-dr-test
⭐ Using report "test"
⭐ Using config "config.yaml"

🔎 Clean tests ...
   ✅ Application "appset-deploy-cephfs" unprotected
   ✅ Application "appset-deploy-rbd" unprotected
   ✅ Application "appset-deploy-rbd" undeployed
   ✅ Application "appset-deploy-cephfs" undeployed

🔎 Clean environment ...
   ✅ Environment cleaned

✅ passed (2 passed, 0 failed, 0 skipped)
```

The command adds test-clean.yaml and test-clean.log to the specified directory.
