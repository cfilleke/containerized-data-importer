// Code generated by swagger-doc. DO NOT EDIT.

package v1beta1

func (DataVolume) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "DataVolume is an abstraction on top of PersistentVolumeClaims to allow easy population of those PersistentVolumeClaims with relation to VirtualMachines\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion\n+kubebuilder:resource:shortName=dv;dvs,categories=all\n+kubebuilder:subresource:status\n+kubebuilder:printcolumn:name=\"Phase\",type=\"string\",JSONPath=\".status.phase\",description=\"The phase the data volume is in\"\n+kubebuilder:printcolumn:name=\"Progress\",type=\"string\",JSONPath=\".status.progress\",description=\"Transfer progress in percentage if known, N/A otherwise\"\n+kubebuilder:printcolumn:name=\"Restarts\",type=\"integer\",JSONPath=\".status.restartCount\",description=\"The number of times the transfer has been restarted.\"\n+kubebuilder:printcolumn:name=\"Age\",type=\"date\",JSONPath=\".metadata.creationTimestamp\"",
		"status": "+optional",
	}
}

func (DataVolumeSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                  "DataVolumeSpec defines the DataVolume type specification",
		"source":            "Source is the src of the data for the requested DataVolume\n+optional",
		"sourceRef":         "SourceRef is an indirect reference to the source of data for the requested DataVolume\n+optional",
		"pvc":               "PVC is the PVC specification",
		"storage":           "Storage is the requested storage specification",
		"priorityClassName": "PriorityClassName for Importer, Cloner and Uploader pod",
		"contentType":       "DataVolumeContentType options: \"kubevirt\", \"archive\"\n+kubebuilder:validation:Enum=\"kubevirt\";\"archive\"",
		"checkpoints":       "Checkpoints is a list of DataVolumeCheckpoints, representing stages in a multistage import.",
		"finalCheckpoint":   "FinalCheckpoint indicates whether the current DataVolumeCheckpoint is the final checkpoint.",
		"preallocation":     "Preallocation controls whether storage for DataVolumes should be allocated in advance.",
	}
}

func (StorageSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                 "StorageSpec defines the Storage type specification",
		"accessModes":      "AccessModes contains the desired access modes the volume should have.\nMore info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1\n+optional",
		"selector":         "A label query over volumes to consider for binding.\n+optional",
		"resources":        "Resources represents the minimum resources the volume should have.\nMore info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources\n+optional",
		"volumeName":       "VolumeName is the binding reference to the PersistentVolume backing this claim.\n+optional",
		"storageClassName": "Name of the StorageClass required by the claim.\nMore info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1\n+optional",
		"volumeMode":       "volumeMode defines what type of volume is required by the claim.\nValue of Filesystem is implied when not included in claim spec.\n+optional",
		"dataSource":       "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) * An existing custom resource that implements data population (Alpha) In order to use custom resource types that implement data population, the AnyVolumeDataSource feature gate must be enabled. If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source.\nIf the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.\n+optional",
		"dataSourceRef":    "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner.\nThis field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty.\nThere are two important differences between DataSource and DataSourceRef:\n* While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects.\n* While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified.\n(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.\n+optional",
	}
}

func (DataVolumeCheckpoint) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "DataVolumeCheckpoint defines a stage in a warm migration.",
		"previous": "Previous is the identifier of the snapshot from the previous checkpoint.",
		"current":  "Current is the identifier of the snapshot created for this checkpoint.",
	}
}

func (DataVolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataVolumeSource represents the source for our Data Volume, this can be HTTP, Imageio, S3, GCS, Registry or an existing PVC",
	}
}

func (DataVolumeSourcePVC) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "DataVolumeSourcePVC provides the parameters to create a Data Volume from an existing PVC",
		"namespace": "The namespace of the source PVC",
		"name":      "The name of the source PVC",
	}
}

func (DataVolumeSourceSnapshot) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "DataVolumeSourceSnapshot provides the parameters to create a Data Volume from an existing VolumeSnapshot",
		"namespace": "The namespace of the source VolumeSnapshot",
		"name":      "The name of the source VolumeSnapshot",
	}
}

func (DataVolumeBlankImage) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataVolumeBlankImage provides the parameters to create a new raw blank image for the PVC",
	}
}

func (DataVolumeSourceUpload) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataVolumeSourceUpload provides the parameters to create a Data Volume by uploading the source",
	}
}

func (DataVolumeSourceS3) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "DataVolumeSourceS3 provides the parameters to create a Data Volume from an S3 source",
		"url":           "URL is the url of the S3 source",
		"secretRef":     "SecretRef provides the secret reference needed to access the S3 source",
		"certConfigMap": "CertConfigMap is a configmap reference, containing a Certificate Authority(CA) public key, and a base64 encoded pem certificate\n+optional",
	}
}

func (DataVolumeSourceGCS) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "DataVolumeSourceGCS provides the parameters to create a Data Volume from an GCS source",
		"url":       "URL is the url of the GCS source",
		"secretRef": "SecretRef provides the secret reference needed to access the GCS source",
	}
}

func (DataVolumeSourceRegistry) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "DataVolumeSourceRegistry provides the parameters to create a Data Volume from an registry source",
		"url":           "URL is the url of the registry source (starting with the scheme: docker, oci-archive)\n+optional",
		"imageStream":   "ImageStream is the name of image stream for import\n+optional",
		"pullMethod":    "PullMethod can be either \"pod\" (default import), or \"node\" (node docker cache based import)\n+optional",
		"secretRef":     "SecretRef provides the secret reference needed to access the Registry source\n+optional",
		"certConfigMap": "CertConfigMap provides a reference to the Registry certs\n+optional",
	}
}

func (DataVolumeSourceHTTP) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "DataVolumeSourceHTTP can be either an http or https endpoint, with an optional basic auth user name and password, and an optional configmap containing additional CAs",
		"url":                "URL is the URL of the http(s) endpoint",
		"secretRef":          "SecretRef A Secret reference, the secret should contain accessKeyId (user name) base64 encoded, and secretKey (password) also base64 encoded\n+optional",
		"certConfigMap":      "CertConfigMap is a configmap reference, containing a Certificate Authority(CA) public key, and a base64 encoded pem certificate\n+optional",
		"extraHeaders":       "ExtraHeaders is a list of strings containing extra headers to include with HTTP transfer requests\n+optional",
		"secretExtraHeaders": "SecretExtraHeaders is a list of Secret references, each containing an extra HTTP header that may include sensitive information\n+optional",
	}
}

func (DataVolumeSourceImageIO) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "DataVolumeSourceImageIO provides the parameters to create a Data Volume from an imageio source",
		"url":           "URL is the URL of the ovirt-engine",
		"diskId":        "DiskID provides id of a disk to be imported",
		"secretRef":     "SecretRef provides the secret reference needed to access the ovirt-engine",
		"certConfigMap": "CertConfigMap provides a reference to the CA cert",
	}
}

func (DataVolumeSourceVDDK) SwaggerDoc() map[string]string {
	return map[string]string{
		"":             "DataVolumeSourceVDDK provides the parameters to create a Data Volume from a Vmware source",
		"url":          "URL is the URL of the vCenter or ESXi host with the VM to migrate",
		"uuid":         "UUID is the UUID of the virtual machine that the backing file is attached to in vCenter/ESXi",
		"backingFile":  "BackingFile is the path to the virtual hard disk to migrate from vCenter/ESXi",
		"thumbprint":   "Thumbprint is the certificate thumbprint of the vCenter or ESXi host",
		"secretRef":    "SecretRef provides a reference to a secret containing the username and password needed to access the vCenter or ESXi host",
		"initImageURL": "InitImageURL is an optional URL to an image containing an extracted VDDK library, overrides v2v-vmware config map",
	}
}

func (DataVolumeSourceRef) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "DataVolumeSourceRef defines an indirect reference to the source of data for the DataVolume",
		"kind":      "The kind of the source reference, currently only \"DataSource\" is supported",
		"namespace": "The namespace of the source reference, defaults to the DataVolume namespace\n+optional",
		"name":      "The name of the source reference",
	}
}

func (DataVolumeStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":             "DataVolumeStatus contains the current status of the DataVolume",
		"claimName":    "ClaimName is the name of the underlying PVC used by the DataVolume.",
		"phase":        "Phase is the current phase of the data volume",
		"restartCount": "RestartCount is the number of times the pod populating the DataVolume has restarted",
	}
}

func (DataVolumeList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "DataVolumeList provides the needed parameters to do request a list of Data Volumes from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of DataVolumes",
	}
}

func (DataVolumeCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataVolumeCondition represents the state of a data volume condition.",
	}
}

func (StorageProfile) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "StorageProfile provides a CDI specific recommendation for storage parameters\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion\n+kubebuilder:resource:scope=Cluster",
	}
}

func (StorageProfileSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                           "StorageProfileSpec defines specification for StorageProfile",
		"cloneStrategy":              "CloneStrategy defines the preferred method for performing a CDI clone",
		"claimPropertySets":          "ClaimPropertySets is a provided set of properties applicable to PVC",
		"dataImportCronSourceFormat": "DataImportCronSourceFormat defines the format of the DataImportCron-created disk image sources",
		"snapshotClass":              "SnapshotClass is optional specific VolumeSnapshotClass for CloneStrategySnapshot. If not set, a VolumeSnapshotClass is chosen according to the provisioner.",
	}
}

func (StorageProfileStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                           "StorageProfileStatus provides the most recently observed status of the StorageProfile",
		"storageClass":               "The StorageClass name for which capabilities are defined",
		"provisioner":                "The Storage class provisioner plugin name",
		"cloneStrategy":              "CloneStrategy defines the preferred method for performing a CDI clone",
		"claimPropertySets":          "ClaimPropertySets computed from the spec and detected in the system",
		"dataImportCronSourceFormat": "DataImportCronSourceFormat defines the format of the DataImportCron-created disk image sources",
		"snapshotClass":              "SnapshotClass is optional specific VolumeSnapshotClass for CloneStrategySnapshot. If not set, a VolumeSnapshotClass is chosen according to the provisioner.",
	}
}

func (ClaimPropertySet) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "ClaimPropertySet is a set of properties applicable to PVC",
		"accessModes": "AccessModes contains the desired access modes the volume should have.\nMore info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1\n+optional",
		"volumeMode":  "VolumeMode defines what type of volume is required by the claim.\nValue of Filesystem is implied when not included in claim spec.\n+optional",
	}
}

func (StorageProfileList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "StorageProfileList provides the needed parameters to request a list of StorageProfile from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of StorageProfile",
	}
}

func (DataSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataSource references an import/clone source for a DataVolume\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion\n+kubebuilder:resource:shortName=das,categories=all",
	}
}

func (DataSourceSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "DataSourceSpec defines specification for DataSource",
		"source": "Source is the source of the data referenced by the DataSource",
	}
}

func (DataSourceSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "DataSourceSource represents the source for our DataSource",
		"pvc":      "+optional",
		"snapshot": "+optional",
	}
}

func (DataSourceStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "DataSourceStatus provides the most recently observed status of the DataSource",
		"source": "Source is the current source of the data referenced by the DataSource",
	}
}

func (DataSourceCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataSourceCondition represents the state of a data source condition",
	}
}

func (ConditionState) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "ConditionState represents the state of a condition",
	}
}

func (DataSourceList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "DataSourceList provides the needed parameters to do request a list of Data Sources from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of DataSources",
	}
}

func (DataImportCron) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataImportCron defines a cron job for recurring polling/importing disk images as PVCs into a golden image namespace\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion\n+kubebuilder:resource:shortName=dic;dics,categories=all\n+kubebuilder:printcolumn:name=\"Format\",type=\"string\",JSONPath=\".status.sourceFormat\",description=\"The format in which created sources are saved\"",
	}
}

func (DataImportCronSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                  "DataImportCronSpec defines specification for DataImportCron",
		"template":          "Template specifies template for the DVs to be created",
		"schedule":          "Schedule specifies in cron format when and how often to look for new imports",
		"garbageCollect":    "GarbageCollect specifies whether old PVCs should be cleaned up after a new PVC is imported.\nOptions are currently \"Outdated\" and \"Never\", defaults to \"Outdated\".\n+optional",
		"importsToKeep":     "Number of import PVCs to keep when garbage collecting. Default is 3.\n+optional",
		"managedDataSource": "ManagedDataSource specifies the name of the corresponding DataSource this cron will manage.\nDataSource has to be in the same namespace.",
		"retentionPolicy":   "RetentionPolicy specifies whether the created DataVolumes and DataSources are retained when their DataImportCron is deleted. Default is RatainAll.\n+optional",
	}
}

func (DataImportCronStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                       "DataImportCronStatus provides the most recently observed status of the DataImportCron",
		"currentImports":         "CurrentImports are the imports in progress. Currently only a single import is supported.",
		"lastImportedPVC":        "LastImportedPVC is the last imported PVC",
		"lastExecutionTimestamp": "LastExecutionTimestamp is the time of the last polling",
		"lastImportTimestamp":    "LastImportTimestamp is the time of the last import",
		"sourceFormat":           "SourceFormat defines the format of the DataImportCron-created disk image sources",
	}
}

func (ImportStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":               "ImportStatus of a currently in progress import",
		"DataVolumeName": "DataVolumeName is the currently in progress import DataVolume",
		"Digest":         "Digest of the currently imported image",
	}
}

func (DataImportCronCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataImportCronCondition represents the state of a data import cron condition",
	}
}

func (DataImportCronList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "DataImportCronList provides the needed parameters to do request a list of DataImportCrons from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of DataImportCrons",
	}
}

func (VolumeImportSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VolumeImportSource works as a specification to populate PersistentVolumeClaims with data\nimported from an HTTP/S3/Registry/Blank/ImageIO/VDDK source\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion",
		"status": "+optional",
	}
}

func (VolumeImportSourceSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "VolumeImportSourceSpec defines the Spec field for VolumeImportSource",
		"source":          "Source is the src of the data to be imported in the target PVC",
		"preallocation":   "Preallocation controls whether storage for the target PVC should be allocated in advance.",
		"contentType":     "ContentType represents the type of the imported data (Kubevirt or archive)",
		"targetClaim":     "TargetClaim the name of the specific claim to be populated with a multistage import.",
		"checkpoints":     "Checkpoints is a list of DataVolumeCheckpoints, representing stages in a multistage import.",
		"finalCheckpoint": "FinalCheckpoint indicates whether the current DataVolumeCheckpoint is the final checkpoint.",
	}
}

func (ImportSourceType) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "ImportSourceType contains each one of the source types allowed in a VolumeImportSource",
	}
}

func (VolumeImportSourceStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VolumeImportSourceStatus provides the most recently observed status of the VolumeImportSource",
	}
}

func (VolumeImportSourceList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "VolumeImportSourceList provides the needed parameters to do request a list of Import Sources from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of DataSources",
	}
}

func (VolumeUploadSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VolumeUploadSource is a specification to populate PersistentVolumeClaims with upload data\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion",
		"status": "+optional",
	}
}

func (VolumeUploadSourceSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "VolumeUploadSourceSpec defines specification for VolumeUploadSource",
		"contentType":   "ContentType represents the type of the upload data (Kubevirt or archive)",
		"preallocation": "Preallocation controls whether storage for the target PVC should be allocated in advance.",
	}
}

func (VolumeUploadSourceStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VolumeUploadSourceStatus provides the most recently observed status of the VolumeUploadSource",
	}
}

func (VolumeUploadSourceList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "VolumeUploadSourceList provides the needed parameters to do request a list of Upload Sources from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of DataSources",
	}
}

func (VolumeCloneSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VolumeCloneSource refers to a PVC/VolumeSnapshot of any storageclass/volumemode\nto be used as the source of a new PVC\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion",
	}
}

func (VolumeCloneSourceSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                  "VolumeCloneSourceSpec defines the Spec field for VolumeCloneSource",
		"source":            "Source is the src of the data to be cloned to the target PVC",
		"preallocation":     "Preallocation controls whether storage for the target PVC should be allocated in advance.\n+optional",
		"priorityClassName": "PriorityClassName is the priorityclass for the claim\n+optional",
	}
}

func (VolumeCloneSourceList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "VolumeCloneSourceList provides the needed parameters to do request a list of VolumeCloneSources from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of DataSources",
	}
}

func (CDI) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "CDI is the CDI Operator CRD\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion\n+kubebuilder:resource:shortName=cdi;cdis,scope=Cluster\n+kubebuilder:printcolumn:name=\"Age\",type=\"date\",JSONPath=\".metadata.creationTimestamp\"\n+kubebuilder:printcolumn:name=\"Phase\",type=\"string\",JSONPath=\".status.phase\"",
		"status": "+optional",
	}
}

func (CertConfig) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "CertConfig contains the tunables for TLS certificates",
		"duration":    "The requested 'duration' (i.e. lifetime) of the Certificate.",
		"renewBefore": "The amount of time before the currently issued certificate's `notAfter`\ntime that we will begin to attempt to renew the certificate.",
	}
}

func (CDICertConfig) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "CDICertConfig has the CertConfigs for CDI",
		"ca":     "CA configuration\nCA certs are kept in the CA bundle as long as they are valid",
		"server": "Server configuration\nCerts are rotated and discarded",
	}
}

func (CDISpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                      "CDISpec defines our specification for the CDI installation",
		"imagePullPolicy":       "+kubebuilder:validation:Enum=Always;IfNotPresent;Never\nPullPolicy describes a policy for if/when to pull a container image",
		"uninstallStrategy":     "+kubebuilder:validation:Enum=RemoveWorkloads;BlockUninstallIfWorkloadsExist\nCDIUninstallStrategy defines the state to leave CDI on uninstall",
		"infra":                 "Rules on which nodes CDI infrastructure pods will be scheduled",
		"workload":              "Restrict on which nodes CDI workload pods will be scheduled",
		"cloneStrategyOverride": "Clone strategy override: should we use a host-assisted copy even if snapshots are available?\n+kubebuilder:validation:Enum=\"copy\";\"snapshot\";\"csi-clone\"",
		"config":                "CDIConfig at CDI level",
		"certConfig":            "certificate configuration",
		"priorityClass":         "PriorityClass of the CDI control plane",
	}
}

func (CDIStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "CDIStatus defines the status of the installation",
	}
}

func (CDIList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "CDIList provides the needed parameters to do request a list of CDIs from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of CDIs",
	}
}

func (CDIConfig) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "CDIConfig provides a user configuration for CDI\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion\n+kubebuilder:resource:scope=Cluster",
	}
}

func (FilesystemOverhead) SwaggerDoc() map[string]string {
	return map[string]string{
		"":             "FilesystemOverhead defines the reserved size for PVCs with VolumeMode: Filesystem",
		"global":       "Global is how much space of a Filesystem volume should be reserved for overhead. This value is used unless overridden by a more specific value (per storageClass)",
		"storageClass": "StorageClass specifies how much space of a Filesystem volume should be reserved for safety. The keys are the storageClass and the values are the overhead. This value overrides the global value",
	}
}

func (CDIConfigSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                         "CDIConfigSpec defines specification for user configuration",
		"uploadProxyURLOverride":   "Override the URL used when uploading to a DataVolume",
		"importProxy":              "ImportProxy contains importer pod proxy configuration.\n+optional",
		"scratchSpaceStorageClass": "Override the storage class to used for scratch space during transfer operations. The scratch space storage class is determined in the following order: 1. value of scratchSpaceStorageClass, if that doesn't exist, use the default storage class, if there is no default storage class, use the storage class of the DataVolume, if no storage class specified, use no storage class for scratch space",
		"podResourceRequirements":  "ResourceRequirements describes the compute resource requirements.",
		"featureGates":             "FeatureGates are a list of specific enabled feature gates",
		"filesystemOverhead":       "FilesystemOverhead describes the space reserved for overhead when using Filesystem volumes. A value is between 0 and 1, if not defined it is 0.055 (5.5% overhead)",
		"preallocation":            "Preallocation controls whether storage for DataVolumes should be allocated in advance.",
		"insecureRegistries":       "InsecureRegistries is a list of TLS disabled registries",
		"dataVolumeTTLSeconds":     "DataVolumeTTLSeconds is the time in seconds after DataVolume completion it can be garbage collected. Disabled by default.\n+optional",
		"tlsSecurityProfile":       "TLSSecurityProfile is used by operators to apply cluster-wide TLS security settings to operands.",
		"imagePullSecrets":         "The imagePullSecrets used to pull the container images",
		"logVerbosity":             "LogVerbosity overrides the default verbosity level used to initialize loggers\n+optional",
	}
}

func (CDIConfigStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                               "CDIConfigStatus provides the most recently observed status of the CDI Config resource",
		"uploadProxyURL":                 "The calculated upload proxy URL",
		"importProxy":                    "ImportProxy contains importer pod proxy configuration.\n+optional",
		"scratchSpaceStorageClass":       "The calculated storage class to be used for scratch space",
		"defaultPodResourceRequirements": "ResourceRequirements describes the compute resource requirements.",
		"filesystemOverhead":             "FilesystemOverhead describes the space reserved for overhead when using Filesystem volumes. A percentage value is between 0 and 1",
		"preallocation":                  "Preallocation controls whether storage for DataVolumes should be allocated in advance.",
		"imagePullSecrets":               "The imagePullSecrets used to pull the container images",
	}
}

func (CDIConfigList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "CDIConfigList provides the needed parameters to do request a list of CDIConfigs from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of CDIConfigs",
	}
}

func (ImportProxy) SwaggerDoc() map[string]string {
	return map[string]string{
		"":               "ImportProxy provides the information on how to configure the importer pod proxy.",
		"HTTPProxy":      "HTTPProxy is the URL http://<username>:<pswd>@<ip>:<port> of the import proxy for HTTP requests.  Empty means unset and will not result in the import pod env var.\n+optional",
		"HTTPSProxy":     "HTTPSProxy is the URL https://<username>:<pswd>@<ip>:<port> of the import proxy for HTTPS requests.  Empty means unset and will not result in the import pod env var.\n+optional",
		"noProxy":        "NoProxy is a comma-separated list of hostnames and/or CIDRs for which the proxy should not be used. Empty means unset and will not result in the import pod env var.\n+optional",
		"trustedCAProxy": "TrustedCAProxy is the name of a ConfigMap in the cdi namespace that contains a user-provided trusted certificate authority (CA) bundle.\nThe TrustedCAProxy ConfigMap is consumed by the DataImportCron controller for creating cronjobs, and by the import controller referring a copy of the ConfigMap in the import namespace.\nHere is an example of the ConfigMap (in yaml):\n\napiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: my-ca-proxy-cm\n  namespace: cdi\ndata:\n  ca.pem: |",
	}
}
