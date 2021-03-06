# Changelog

## 0.16.0

* New Optional REST API
* New Optional Policy Reporter UI Helm SubChart

## 0.15.1

* Add a checksum for the target configuration secret to the deployment. This enforces a pod recreation when the configuration changed by a Helm upgrade.

## 0.15.0

* Customizable Dashboards via new Helm values for the Monitoring Subchart.
## 0.14.0

* Internal refactoring
    * Improved test coverage
    * Removed duplicated caching
* Updated Dashboard
    * Filter zero values from Policy Report Detail after Policies / Resources are deleted

## 0.13.0

* Split the Monitoring out in a Sub Helm chart
    * Changed naming from `metrics` to `monitoring`
* Make Annotations for the Deployment configurable
* Add two new Grafana Dashboard (PolicyReport Details, ClusterPolicyReport Details)

## 0.12.0

* Add support for a special `default` key in the Policy Priority. The `default` key can be used to configure a global default priority instead of `error`

## 0.11.1

* Use a Secret instead of ConfigMap to persist target configurations

## 0.11.0

* Helm Chart Value `metrics.serviceMonitor` changed to `metrics.serviceMonitor.enabled`
* New Helm Chart Value `metrics.serviceMonitor.labels` can be used to add additional `labels` to the `SeriveMonitor`. This helps to fullfil the `serviceMonitorSelector` of the `Prometheus` Resource in the MonitoringStack.

## 0.10.0

* Implement Discord as Target for PolicyReportResults

## 0.9.0

* Implement Slack as Target for PolicyReportResults

## 0.8.0

* Implement Elasticsearch as Target for PolicyReportResults
* Replace CLI flags with a single `config.yaml` to manage target-configurations as separate `ConfigMap`
* Set `loki.skipExistingOnStartup` default value to `true`
