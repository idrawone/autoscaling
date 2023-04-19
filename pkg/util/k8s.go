package util

// Kubernetes-specific utility functions

import (
	corev1 "k8s.io/api/core/v1"
)

// PodReady returns true iff the pod is marked as ready (as determined by the pod's
// Status.Conditions)
func PodReady(pod *corev1.Pod) bool {
	for _, c := range pod.Status.Conditions {
		if c.Type == corev1.PodReady {
			return c.Status == corev1.ConditionTrue
		}
	}

	return false
}

// PodStartedBefore returns true iff Pod p started before Pod q
func PodStartedBefore(p, q *corev1.Pod) bool {
	return p.Status.StartTime.Before(q.Status.StartTime)
}