/*
Copyright 2022.

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

package v2

import (
	v1 "k8s.io/api/batch/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// represents a Cron field specifier.
type CronField string

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CronSchedule describes a Cron schedule.
type CronSchedule struct {
	// specifies the minute during which the job executes.
	// +optional
	Minute *CronField `json:"minute,omitempty"`
	// specifies the hour during which the job executes.
	// +optional
	Hour *CronField `json:"hour,omitempty"`
	// specifies the day of the month during which the job executes.
	// +optional
	DayOfMonth *CronField `json:"dayOfMonth,omitempty"`
	// specifies the month during which the job executes.
	// +optional
	Month *CronField `json:"month,omitempty"`
	// specifies the day of the week during which the job executes.
	// +optional
	DayOfWeek *CronField `json:"dayOfWeek,omitempty"`
}

// CronJobSpec defines the desired state of CronJob
type CronJobSpec struct {
	Schedule                  CronSchedule       `json:"schedule"`
	StartingDeadlineSeconds   *int64             `json:"startingDeadlineSeconds,omitempty"`
	ConcurrencyPolicy         ConcurrencyPolicy  `json:"concurrencyPolicy,omitempty"`
	Suspend                   *bool              `json:"suspend,omitempty"`
	JobTemplate               v1.JobTemplateSpec `json:"jobTemplate"`
	SuccessfulJobHistoryLimit *int32             `json:"successfulJobHistoryLimit,omitempty"`
	FailedJobsHistoryLimit    *int32             `json:"failedJobsHistoryLimit,omitempty"`
}

// CronJobStatus defines the observed state of CronJob
type CronJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Active []v12.ObjectReference `json:"active,omitempty"`

	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

type ConcurrencyPolicy string

const (
	AllowConcurrent   ConcurrencyPolicy = "Allow"
	ForbidConcurrent  ConcurrencyPolicy = "Forbid"
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

func init() {
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}
