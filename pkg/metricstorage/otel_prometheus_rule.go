/*
Copyright 2024.

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

package metricstorage

import (
	telemetryv1 "github.com/openstack-k8s-operators/telemetry-operator/api/v1beta1"
	monv1 "github.com/rhobs/obo-prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func OTelPrometheusRule(
	instance *telemetryv1.MetricStorage,
	labels map[string]string,
) *monv1.PrometheusRule {

	prometheusRule := &monv1.PrometheusRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
			Labels:    labels,
		},
		Spec: monv1.PrometheusRuleSpec{
			Groups: []monv1.RuleGroup{
				{
					Name: "osp-node-exporter-OTel.rules",
					Rules: []monv1.Rule{
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `count without (cpu, mode) (node_cpu_seconds_total{mode="idle"})`,
							},
							Record: "otel:system_cpu_physical_count",
						},
						{
							Expr: intstr.IntOrString{
								Type: intstr.String,
								// Labeling matches semantics, not otel-collector, which has cpu=cpu0 instead of logical_number=0
								StrVal: `sum without(cpu, mode) (label_replace(label_replace(node_cpu_seconds_total{mode!~".*irq"}, "logical_number", "$1", "cpu", "(.*)"), "state", "$1", "mode", "(.*)"))`,
							},
							Record: "otel:system_cpu_time_seconds_total",
						},
						{
							Expr: intstr.IntOrString{
								Type: intstr.String,
								// Labeling matches semantics, not otel-collector, which has cpu=cpu0 instead of logical_number=0
								StrVal: `sum without(cpu) (label_replace(label_replace(node_cpu_seconds_total{mode="irq"} + ignoring(mode) node_cpu_seconds_total{mode="softirq"}, "state", "interrupt", "", ""), "logical_number", "$1", "cpu", "(.*)"))`,
							},
							Record: "otel:system_cpu_time_seconds_total",
						},
						{
							Expr: intstr.IntOrString{
								Type: intstr.String,
								// Using irate because the semantics say "since the last measurement..." not an average rate over the minute
								StrVal: `irate(node_cpu_seconds_total[1m])`,
							},
							Record: "otel:system_cpu_utilization",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `node_load1`,
							},
							// Naming matches semantics, not otel-collector, which has system_cpu_load_average_1m
							Record: "otel:system_linux_cpu_load_1m",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `node_load5`,
							},
							// Naming matches semantics, not otel-collector, which has system_cpu_load_average_5m
							Record: "otel:system_linux_cpu_load_5m",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `node_load15`,
							},
							// Naming matches semantics, not otel-collector, which has system_cpu_load_average_15m
							Record: "otel:system_linux_cpu_load_15m",
						},
						{
							Expr: intstr.IntOrString{
								Type: intstr.String,
								// Calculation (Total - Available) matches output of linux `free` command
								StrVal: `label_replace(node_memory_MemTotal_bytes - node_memory_MemAvailable_bytes, "state", "used", "", "")`,
							},
							Record: "otel:system_memory_usage_bytes",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_memory_MemFree_bytes, "state", "free", "", "")`,
							},
							Record: "otel:system_memory_usage_bytes",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_memory_Buffers_bytes, "state", "buffers", "", "")`,
							},
							Record: "otel:system_memory_usage_bytes",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_memory_Cached_bytes, "state", "cached", "", "")`,
							},
							Record: "otel:system_memory_usage_bytes",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes), "state", "used", "", "")`,
							},
							Record: "otel:system_memory_utilization",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_memory_MemFree_bytes / node_memory_MemTotal_bytes, "state", "free", "", "")`,
							},
							Record: "otel:system_memory_utilization",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_memory_Buffers_bytes / node_memory_MemTotal_bytes, "state", "buffers", "", "")`,
							},
							Record: "otel:system_memory_utilization",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_memory_Cached_bytes / node_memory_MemTotal_bytes, "state", "cached", "", "")`,
							},
							Record: "otel:system_memory_utilization",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `node_memory_Shmem_bytes`,
							},
							Record: "otel:system_memory_shared_bytes",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `node_memory_MemTotal_bytes`,
							},
							Record: "otel:system_memory_limit_bytes",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_memory_SwapFree_bytes, "state", "free", "", "")`,
							},
							Record: "otel:system_paging_usage_bytes",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_memory_SwapTotal_bytes - node_memory_SwapFree_bytes, "state", "used", "", "")`,
							},
							Record: "otel:system_paging_usage_bytes",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_vmstat_pgmajfault, "type", "major", "", "")`,
							},
							Record: "otel:system_paging_faults_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `sum without (fstype) (label_replace(label_replace(node_filesystem_size_bytes - node_filesystem_free_bytes, "type", "$1", "fstype", "(.*)"), "state", "used", "", ""))`,
							},
							Record: "otel:system_filesystem_usage_bytes",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `sum without (fstype) (label_replace(label_replace(node_filesystem_free_bytes, "type", "$1", "fstype", "(.*)"), "state", "free", "", ""))`,
							},
							Record: "otel:system_filesystem_usage_bytes",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `sum without (fstype) (label_replace(label_replace((node_filesystem_size_bytes - node_filesystem_free_bytes) / node_filesystem_size_bytes, "type", "$1", "fstype", "(.*)"), "state", "used", "", ""))`,
							},
							Record: "otel:system_filesystem_utilization",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `sum without (fstype) (label_replace(label_replace(node_filesystem_free_bytes / node_filesystem_size_bytes, "type", "$1", "fstype", "(.*)"), "state", "free", "", ""))`,
							},
							Record: "otel:system_filesystem_utilization",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `sum without (fstype) (label_replace(node_filesystem_device_error, "type", "$1", "fstype", "(.*)"))`,
							},
							Record: "otel:system_filesystem_errors_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_disk_read_bytes_total, "direction", "read", "", "")`,
							},
							Record: "otel:system_disk_io_bytes_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_disk_written_bytes_total, "direction", "write", "", "")`,
							},
							Record: "otel:system_disk_io_bytes_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_disk_reads_completed_total, "direction", "read", "", "")`,
							},
							Record: "otel:system_disk_operations_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_disk_writes_completed_total, "direction", "write", "", "")`,
							},
							Record: "otel:system_disk_operations_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `node_disk_io_now`,
							},
							Record: "otel:system_disk_pending_operations",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `node_disk_io_time_seconds_total`,
							},
							Record: "otel:system_disk_io_time_seconds_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_disk_read_time_seconds_total, "direction", "read", "", "")`,
							},
							Record: "otel:system_disk_operation_time_seconds_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_disk_write_time_seconds_total, "direction", "write", "", "")`,
							},
							Record: "otel:system_disk_operation_time_seconds_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_network_receive_bytes_total, "direction", "receive", "", "")`,
							},
							Record: "otel:system_network_io_bytes_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_network_transmit_bytes_total, "direction", "transmit", "", "")`,
							},
							Record: "otel:system_network_io_bytes_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_network_receive_packets_total, "direction", "receive", "", "")`,
							},
							Record: "otel:system_network_packets_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_network_transmit_packets_total, "direction", "transmit", "", "")`,
							},
							Record: "otel:system_network_packets_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_network_receive_drop_total, "direction", "receive", "", "")`,
							},
							Record: "otel:system_network_dropped_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_network_transmit_drop_total, "direction", "transmit", "", "")`,
							},
							Record: "otel:system_network_dropped_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_network_receive_errs_total, "direction", "receive", "", "")`,
							},
							Record: "otel:system_network_errors_total",
						},
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `label_replace(node_network_transmit_errs_total, "direction", "transmit", "", "")`,
							},
							Record: "otel:system_network_errors_total",
						},
					},
				},
				{
					Name: "osp-ceilometer-OTel.rules",
					Rules: []monv1.Rule{
						{
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: `sum without(unit, type) (label_replace(rate(ceilometer_cpu[1m]), "vm_name", "$1", "resource_name", "(.*):.*") / 1000000000)`,
							},
							Record: "otel:system_cpu_time_seconds_total",
						},
					},
				},
			},
		},
	}
	return prometheusRule
}
