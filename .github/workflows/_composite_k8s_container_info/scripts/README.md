# get k8s image tag

## Description

標準入力に渡したK8sのマニフェストの中から、コンテナの情報を抜き出しJSON形式で出力します。

## Usage

```shell
$ kustomize build ./app/server/overlay/prod/estate | ./scripts/get_k8s_container_info/get_k8s_container_info

[{"Namespace":"estate","Name":"estate-deployment","Containers":[{"name":"server","image":"asia-northeast1-docker.pkg.dev/canary-dev-214513/server/estate:20230823_175520","args":["serverApp"],"ports":[{"name":"server","containerPort":9000,"protocol":"TCP"}],"env":[{"name":"CONFIG_FILE_PATH","value":"/etc/config/config.yaml"},{"name":"ELASTIC_SEARCH_USERNAME","valueFrom":{"secretKeyRef":{"name":"estate-elastic-search-secret","key":"username"}}},{"name":"ELASTIC_SEARCH_PASSWORD","valueFrom":{"secretKeyRef":{"name":"estate-elastic-search-secret","key":"password"}}},{"name":"S3_SECRET","valueFrom":{"secretKeyRef":{"name":"estate-s3-secret","key":"secret"}}},{"name":"FTP_APAMAN_PASSWORD","valueFrom":{"secretKeyRef":{"name":"estate-ftp-secret","key":"apamanPassword"}}},{"name":"BLUAGE_APP","value":"estate"},{"name":"BLUAGE_APP_VERSION","value":"20230823_175520"}],"resources":{"limits":{"cpu":"1250m","memory":"5000Mi"},"requests":{"cpu":"1250m","memory":"5000Mi"}},"volumeMounts":[{"name":"config-volume","readOnly":true,"mountPath":"/etc/config"},{"name":"id-rsa-ftp","readOnly":true,"mountPath":"/etc/id_rsa_ftp"}],"livenessProbe":{"grpc":{"port":9000,"service":null},"initialDelaySeconds":10,"periodSeconds":5},"readinessProbe":{"grpc":{"port":9000,"service":null},"initialDelaySeconds":10,"periodSeconds":5},"imagePullPolicy":"IfNotPresent"}]}]
```
