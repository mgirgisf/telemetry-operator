package metricstorage

import (
	telemetryv1 "github.com/openstack-k8s-operators/telemetry-operator/api/v1beta1"
	monv1 "github.com/rhobs/obo-prometheus-operator/pkg/apis/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PrometheusProxy(
	instance *telemetryv1.MetricStorage,
) monv1.Prometheus {
	prom := monv1.Prometheus{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Prometheus",
			APIVersion: "monitoring.rhobs/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
		},
		Spec: monv1.PrometheusSpec{
			Containers: []corev1.Container{
				{
					Name:  "kube-rbac-proxy",
					Image: "quay.io/openstack-k8s-operators/kube-rbac-proxy:v0.16.0",
					Args: []string{
						"--secure-listen-address=0.0.0.0:8443",
						"--upstream=https://metric-storage-prometheus.openstack.svc:9090/",
						"--logtostderr=true",
						"--v=10",
					},
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 8443,
							Name:          "https",
						},
					},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "ca-cert",
							MountPath: "/etc/tls",
							ReadOnly:  true,
						},
					},
				},
			},
			Volumes: []corev1.Volume{
				{
					Name: "ca-cert",
					VolumeSource: corev1.VolumeSource{
						Secret: &corev1.SecretVolumeSource{
							SecretName: "cert-metric-storage-prometheus-svc",
							Items: []corev1.KeyToPath{
								{
									Key:  "ca.crt",
									Path: "ca.crt",
								},
							},
						},
					},
				},
			},
		},
	}
	return prom
}