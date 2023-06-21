package resources

// CDICRDs is a map containing yaml strings of all CRDs
var CDICRDs map[string]string = map[string]string{
	"cdi": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: cdis.cdi.kubevirt.io
spec:
  group: cdi.kubevirt.io
  names:
    kind: CDI
    listKind: CDIList
    plural: cdis
    shortNames:
    - cdi
    - cdis
    singular: cdi
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    - jsonPath: .status.phase
      name: Phase
      type: string
    name: v1alpha1
    schema:
      openAPIV3Schema:
        description: CDI is the CDI Operator CRD
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: CDISpec defines our specification for the CDI installation
            properties:
              certConfig:
                description: certificate configuration
                properties:
                  ca:
                    description: CA configuration CA certs are kept in the CA bundle
                      as long as they are valid
                    properties:
                      duration:
                        description: The requested 'duration' (i.e. lifetime) of the
                          Certificate.
                        type: string
                      renewBefore:
                        description: The amount of time before the currently issued
                          certificate's ` + "`" + `notAfter` + "`" + ` time that we will begin to attempt
                          to renew the certificate.
                        type: string
                    type: object
                  server:
                    description: Server configuration Certs are rotated and discarded
                    properties:
                      duration:
                        description: The requested 'duration' (i.e. lifetime) of the
                          Certificate.
                        type: string
                      renewBefore:
                        description: The amount of time before the currently issued
                          certificate's ` + "`" + `notAfter` + "`" + ` time that we will begin to attempt
                          to renew the certificate.
                        type: string
                    type: object
                type: object
              cloneStrategyOverride:
                description: 'Clone strategy override: should we use a host-assisted
                  copy even if snapshots are available?'
                enum:
                - copy
                - snapshot
                - csi-clone
                type: string
              config:
                description: CDIConfig at CDI level
                properties:
                  dataVolumeTTLSeconds:
                    description: DataVolumeTTLSeconds is the time in seconds after
                      DataVolume completion it can be garbage collected. Disabled
                      by default.
                    format: int32
                    type: integer
                  featureGates:
                    description: FeatureGates are a list of specific enabled feature
                      gates
                    items:
                      type: string
                    type: array
                  filesystemOverhead:
                    description: FilesystemOverhead describes the space reserved for
                      overhead when using Filesystem volumes. A value is between 0
                      and 1, if not defined it is 0.055 (5.5% overhead)
                    properties:
                      global:
                        description: Global is how much space of a Filesystem volume
                          should be reserved for overhead. This value is used unless
                          overridden by a more specific value (per storageClass)
                        pattern: ^(0(?:\.\d{1,3})?|1)$
                        type: string
                      storageClass:
                        additionalProperties:
                          description: 'Percent is a string that can only be a value
                            between [0,1) (Note: we actually rely on reconcile to
                            reject invalid values)'
                          pattern: ^(0(?:\.\d{1,3})?|1)$
                          type: string
                        description: StorageClass specifies how much space of a Filesystem
                          volume should be reserved for safety. The keys are the storageClass
                          and the values are the overhead. This value overrides the
                          global value
                        type: object
                    type: object
                  imagePullSecrets:
                    description: The imagePullSecrets used to pull the container images
                    items:
                      description: LocalObjectReference contains enough information
                        to let you locate the referenced object inside the same namespace.
                      properties:
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                            TODO: Add other useful fields. apiVersion, kind, uid?'
                          type: string
                      type: object
                      x-kubernetes-map-type: atomic
                    type: array
                  importProxy:
                    description: ImportProxy contains importer pod proxy configuration.
                    properties:
                      HTTPProxy:
                        description: HTTPProxy is the URL http://<username>:<pswd>@<ip>:<port>
                          of the import proxy for HTTP requests.  Empty means unset
                          and will not result in the import pod env var.
                        type: string
                      HTTPSProxy:
                        description: HTTPSProxy is the URL https://<username>:<pswd>@<ip>:<port>
                          of the import proxy for HTTPS requests.  Empty means unset
                          and will not result in the import pod env var.
                        type: string
                      noProxy:
                        description: NoProxy is a comma-separated list of hostnames
                          and/or CIDRs for which the proxy should not be used. Empty
                          means unset and will not result in the import pod env var.
                        type: string
                      trustedCAProxy:
                        description: "TrustedCAProxy is the name of a ConfigMap in
                          the cdi namespace that contains a user-provided trusted
                          certificate authority (CA) bundle. The TrustedCAProxy ConfigMap
                          is consumed by the DataImportCron controller for creating
                          cronjobs, and by the import controller referring a copy
                          of the ConfigMap in the import namespace. Here is an example
                          of the ConfigMap (in yaml): \n apiVersion: v1 kind: ConfigMap
                          metadata: name: my-ca-proxy-cm namespace: cdi data: ca.pem:
                          | -----BEGIN CERTIFICATE----- ... <base64 encoded cert>
                          ... -----END CERTIFICATE-----"
                        type: string
                    type: object
                  insecureRegistries:
                    description: InsecureRegistries is a list of TLS disabled registries
                    items:
                      type: string
                    type: array
                  podResourceRequirements:
                    description: ResourceRequirements describes the compute resource
                      requirements.
                    properties:
                      claims:
                        description: "Claims lists the names of resources, defined
                          in spec.resourceClaims, that are used by this container.
                          \n This is an alpha field and requires enabling the DynamicResourceAllocation
                          feature gate. \n This field is immutable."
                        items:
                          description: ResourceClaim references one entry in PodSpec.ResourceClaims.
                          properties:
                            name:
                              description: Name must match the name of one entry in
                                pod.spec.resourceClaims of the Pod where this field
                                is used. It makes that resource available inside a
                                container.
                              type: string
                          required:
                          - name
                          type: object
                        type: array
                        x-kubernetes-list-map-keys:
                        - name
                        x-kubernetes-list-type: map
                      limits:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Limits describes the maximum amount of compute
                          resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                      requests:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Requests describes the minimum amount of compute
                          resources required. If Requests is omitted for a container,
                          it defaults to Limits if that is explicitly specified, otherwise
                          to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                    type: object
                  preallocation:
                    description: Preallocation controls whether storage for DataVolumes
                      should be allocated in advance.
                    type: boolean
                  scratchSpaceStorageClass:
                    description: 'Override the storage class to used for scratch space
                      during transfer operations. The scratch space storage class
                      is determined in the following order: 1. value of scratchSpaceStorageClass,
                      if that doesn''t exist, use the default storage class, if there
                      is no default storage class, use the storage class of the DataVolume,
                      if no storage class specified, use no storage class for scratch
                      space'
                    type: string
                  tlsSecurityProfile:
                    description: TLSSecurityProfile is used by operators to apply
                      cluster-wide TLS security settings to operands.
                    properties:
                      custom:
                        description: "custom is a user-defined TLS security profile.
                          Be extremely careful using a custom profile as invalid configurations
                          can be catastrophic. An example custom profile looks like
                          this: \n ciphers: - ECDHE-ECDSA-CHACHA20-POLY1305 - ECDHE-RSA-CHACHA20-POLY1305
                          - ECDHE-RSA-AES128-GCM-SHA256 - ECDHE-ECDSA-AES128-GCM-SHA256
                          minTLSVersion: TLSv1.1"
                        nullable: true
                        properties:
                          ciphers:
                            description: "ciphers is used to specify the cipher algorithms
                              that are negotiated during the TLS handshake.  Operators
                              may remove entries their operands do not support.  For
                              example, to use DES-CBC3-SHA  (yaml): \n ciphers: -
                              DES-CBC3-SHA"
                            items:
                              type: string
                            type: array
                          minTLSVersion:
                            description: "minTLSVersion is used to specify the minimal
                              version of the TLS protocol that is negotiated during
                              the TLS handshake. For example, to use TLS versions
                              1.1, 1.2 and 1.3 (yaml): \n minTLSVersion: TLSv1.1 \n
                              NOTE: currently the highest minTLSVersion allowed is
                              VersionTLS12"
                            enum:
                            - VersionTLS10
                            - VersionTLS11
                            - VersionTLS12
                            - VersionTLS13
                            type: string
                        type: object
                      intermediate:
                        description: "intermediate is a TLS security profile based
                          on: \n https://wiki.mozilla.org/Security/Server_Side_TLS#Intermediate_compatibility_.28recommended.29
                          \n and looks like this (yaml): \n ciphers: - TLS_AES_128_GCM_SHA256
                          - TLS_AES_256_GCM_SHA384 - TLS_CHACHA20_POLY1305_SHA256
                          - ECDHE-ECDSA-AES128-GCM-SHA256 - ECDHE-RSA-AES128-GCM-SHA256
                          - ECDHE-ECDSA-AES256-GCM-SHA384 - ECDHE-RSA-AES256-GCM-SHA384
                          - ECDHE-ECDSA-CHACHA20-POLY1305 - ECDHE-RSA-CHACHA20-POLY1305
                          - DHE-RSA-AES128-GCM-SHA256 - DHE-RSA-AES256-GCM-SHA384
                          minTLSVersion: TLSv1.2"
                        nullable: true
                        type: object
                      modern:
                        description: "modern is a TLS security profile based on: \n
                          https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility
                          \n and looks like this (yaml): \n ciphers: - TLS_AES_128_GCM_SHA256
                          - TLS_AES_256_GCM_SHA384 - TLS_CHACHA20_POLY1305_SHA256
                          minTLSVersion: TLSv1.3 \n NOTE: Currently unsupported."
                        nullable: true
                        type: object
                      old:
                        description: "old is a TLS security profile based on: \n https://wiki.mozilla.org/Security/Server_Side_TLS#Old_backward_compatibility
                          \n and looks like this (yaml): \n ciphers: - TLS_AES_128_GCM_SHA256
                          - TLS_AES_256_GCM_SHA384 - TLS_CHACHA20_POLY1305_SHA256
                          - ECDHE-ECDSA-AES128-GCM-SHA256 - ECDHE-RSA-AES128-GCM-SHA256
                          - ECDHE-ECDSA-AES256-GCM-SHA384 - ECDHE-RSA-AES256-GCM-SHA384
                          - ECDHE-ECDSA-CHACHA20-POLY1305 - ECDHE-RSA-CHACHA20-POLY1305
                          - DHE-RSA-AES128-GCM-SHA256 - DHE-RSA-AES256-GCM-SHA384
                          - DHE-RSA-CHACHA20-POLY1305 - ECDHE-ECDSA-AES128-SHA256
                          - ECDHE-RSA-AES128-SHA256 - ECDHE-ECDSA-AES128-SHA - ECDHE-RSA-AES128-SHA
                          - ECDHE-ECDSA-AES256-SHA384 - ECDHE-RSA-AES256-SHA384 -
                          ECDHE-ECDSA-AES256-SHA - ECDHE-RSA-AES256-SHA - DHE-RSA-AES128-SHA256
                          - DHE-RSA-AES256-SHA256 - AES128-GCM-SHA256 - AES256-GCM-SHA384
                          - AES128-SHA256 - AES256-SHA256 - AES128-SHA - AES256-SHA
                          - DES-CBC3-SHA minTLSVersion: TLSv1.0"
                        nullable: true
                        type: object
                      type:
                        description: "type is one of Old, Intermediate, Modern or
                          Custom. Custom provides the ability to specify individual
                          TLS security profile parameters. Old, Intermediate and Modern
                          are TLS security profiles based on: \n https://wiki.mozilla.org/Security/Server_Side_TLS#Recommended_configurations
                          \n The profiles are intent based, so they may change over
                          time as new ciphers are developed and existing ciphers are
                          found to be insecure.  Depending on precisely which ciphers
                          are available to a process, the list may be reduced. \n
                          Note that the Modern profile is currently not supported
                          because it is not yet well adopted by common software libraries."
                        enum:
                        - Old
                        - Intermediate
                        - Modern
                        - Custom
                        type: string
                    type: object
                  uploadProxyURLOverride:
                    description: Override the URL used when uploading to a DataVolume
                    type: string
                type: object
              imagePullPolicy:
                description: PullPolicy describes a policy for if/when to pull a container
                  image
                enum:
                - Always
                - IfNotPresent
                - Never
                type: string
              infra:
                description: Rules on which nodes CDI infrastructure pods will be
                  scheduled
                properties:
                  affinity:
                    description: affinity enables pod affinity/anti-affinity placement
                      expanding the types of constraints that can be expressed with
                      nodeSelector. affinity is going to be applied to the relevant
                      kind of pods in parallel with nodeSelector See https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity
                    properties:
                      nodeAffinity:
                        description: Describes node affinity scheduling rules for
                          the pod.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node matches the corresponding matchExpressions;
                              the node(s) with the highest sum are the most preferred.
                            items:
                              description: An empty preferred scheduling term matches
                                all objects with implicit weight 0 (i.e. it's a no-op).
                                A null preferred scheduling term matches no objects
                                (i.e. is also a no-op).
                              properties:
                                preference:
                                  description: A node selector term, associated with
                                    the corresponding weight.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                weight:
                                  description: Weight associated with matching the
                                    corresponding nodeSelectorTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - preference
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to an update), the system
                              may or may not try to eventually evict the pod from
                              its node.
                            properties:
                              nodeSelectorTerms:
                                description: Required. A list of node selector terms.
                                  The terms are ORed.
                                items:
                                  description: A null or empty node selector term
                                    matches no objects. The requirements of them are
                                    ANDed. The TopologySelectorTerm type implements
                                    a subset of the NodeSelectorTerm.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                type: array
                            required:
                            - nodeSelectorTerms
                            type: object
                            x-kubernetes-map-type: atomic
                        type: object
                      podAffinity:
                        description: Describes pod affinity scheduling rules (e.g.
                          co-locate this pod in the same node, zone, etc. as some
                          other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to a pod label update),
                              the system may or may not try to eventually evict the
                              pod from its node. When there are multiple elements,
                              the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                      podAntiAffinity:
                        description: Describes pod anti-affinity scheduling rules
                          (e.g. avoid putting this pod in the same node, zone, etc.
                          as some other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the anti-affinity expressions
                              specified by this field, but it may choose a node that
                              violates one or more of the expressions. The node that
                              is most preferred is the one with the greatest sum of
                              weights, i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              anti-affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the anti-affinity requirements specified
                              by this field are not met at scheduling time, the pod
                              will not be scheduled onto the node. If the anti-affinity
                              requirements specified by this field cease to be met
                              at some point during pod execution (e.g. due to a pod
                              label update), the system may or may not try to eventually
                              evict the pod from its node. When there are multiple
                              elements, the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                    type: object
                  nodeSelector:
                    additionalProperties:
                      type: string
                    description: 'nodeSelector is the node selector applied to the
                      relevant kind of pods It specifies a map of key-value pairs:
                      for the pod to be eligible to run on a node, the node must have
                      each of the indicated key-value pairs as labels (it can have
                      additional labels as well). See https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector'
                    type: object
                  tolerations:
                    description: tolerations is a list of tolerations applied to the
                      relevant kind of pods See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
                      for more info. These are additional tolerations other than default
                      ones.
                    items:
                      description: The pod this Toleration is attached to tolerates
                        any taint that matches the triple <key,value,effect> using
                        the matching operator <operator>.
                      properties:
                        effect:
                          description: Effect indicates the taint effect to match.
                            Empty means match all taint effects. When specified, allowed
                            values are NoSchedule, PreferNoSchedule and NoExecute.
                          type: string
                        key:
                          description: Key is the taint key that the toleration applies
                            to. Empty means match all taint keys. If the key is empty,
                            operator must be Exists; this combination means to match
                            all values and all keys.
                          type: string
                        operator:
                          description: Operator represents a key's relationship to
                            the value. Valid operators are Exists and Equal. Defaults
                            to Equal. Exists is equivalent to wildcard for value,
                            so that a pod can tolerate all taints of a particular
                            category.
                          type: string
                        tolerationSeconds:
                          description: TolerationSeconds represents the period of
                            time the toleration (which must be of effect NoExecute,
                            otherwise this field is ignored) tolerates the taint.
                            By default, it is not set, which means tolerate the taint
                            forever (do not evict). Zero and negative values will
                            be treated as 0 (evict immediately) by the system.
                          format: int64
                          type: integer
                        value:
                          description: Value is the taint value the toleration matches
                            to. If the operator is Exists, the value should be empty,
                            otherwise just a regular string.
                          type: string
                      type: object
                    type: array
                type: object
              priorityClass:
                description: PriorityClass of the CDI control plane
                type: string
              uninstallStrategy:
                description: CDIUninstallStrategy defines the state to leave CDI on
                  uninstall
                enum:
                - RemoveWorkloads
                - BlockUninstallIfWorkloadsExist
                type: string
              workload:
                description: Restrict on which nodes CDI workload pods will be scheduled
                properties:
                  affinity:
                    description: affinity enables pod affinity/anti-affinity placement
                      expanding the types of constraints that can be expressed with
                      nodeSelector. affinity is going to be applied to the relevant
                      kind of pods in parallel with nodeSelector See https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity
                    properties:
                      nodeAffinity:
                        description: Describes node affinity scheduling rules for
                          the pod.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node matches the corresponding matchExpressions;
                              the node(s) with the highest sum are the most preferred.
                            items:
                              description: An empty preferred scheduling term matches
                                all objects with implicit weight 0 (i.e. it's a no-op).
                                A null preferred scheduling term matches no objects
                                (i.e. is also a no-op).
                              properties:
                                preference:
                                  description: A node selector term, associated with
                                    the corresponding weight.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                weight:
                                  description: Weight associated with matching the
                                    corresponding nodeSelectorTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - preference
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to an update), the system
                              may or may not try to eventually evict the pod from
                              its node.
                            properties:
                              nodeSelectorTerms:
                                description: Required. A list of node selector terms.
                                  The terms are ORed.
                                items:
                                  description: A null or empty node selector term
                                    matches no objects. The requirements of them are
                                    ANDed. The TopologySelectorTerm type implements
                                    a subset of the NodeSelectorTerm.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                type: array
                            required:
                            - nodeSelectorTerms
                            type: object
                            x-kubernetes-map-type: atomic
                        type: object
                      podAffinity:
                        description: Describes pod affinity scheduling rules (e.g.
                          co-locate this pod in the same node, zone, etc. as some
                          other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to a pod label update),
                              the system may or may not try to eventually evict the
                              pod from its node. When there are multiple elements,
                              the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                      podAntiAffinity:
                        description: Describes pod anti-affinity scheduling rules
                          (e.g. avoid putting this pod in the same node, zone, etc.
                          as some other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the anti-affinity expressions
                              specified by this field, but it may choose a node that
                              violates one or more of the expressions. The node that
                              is most preferred is the one with the greatest sum of
                              weights, i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              anti-affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the anti-affinity requirements specified
                              by this field are not met at scheduling time, the pod
                              will not be scheduled onto the node. If the anti-affinity
                              requirements specified by this field cease to be met
                              at some point during pod execution (e.g. due to a pod
                              label update), the system may or may not try to eventually
                              evict the pod from its node. When there are multiple
                              elements, the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                    type: object
                  nodeSelector:
                    additionalProperties:
                      type: string
                    description: 'nodeSelector is the node selector applied to the
                      relevant kind of pods It specifies a map of key-value pairs:
                      for the pod to be eligible to run on a node, the node must have
                      each of the indicated key-value pairs as labels (it can have
                      additional labels as well). See https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector'
                    type: object
                  tolerations:
                    description: tolerations is a list of tolerations applied to the
                      relevant kind of pods See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
                      for more info. These are additional tolerations other than default
                      ones.
                    items:
                      description: The pod this Toleration is attached to tolerates
                        any taint that matches the triple <key,value,effect> using
                        the matching operator <operator>.
                      properties:
                        effect:
                          description: Effect indicates the taint effect to match.
                            Empty means match all taint effects. When specified, allowed
                            values are NoSchedule, PreferNoSchedule and NoExecute.
                          type: string
                        key:
                          description: Key is the taint key that the toleration applies
                            to. Empty means match all taint keys. If the key is empty,
                            operator must be Exists; this combination means to match
                            all values and all keys.
                          type: string
                        operator:
                          description: Operator represents a key's relationship to
                            the value. Valid operators are Exists and Equal. Defaults
                            to Equal. Exists is equivalent to wildcard for value,
                            so that a pod can tolerate all taints of a particular
                            category.
                          type: string
                        tolerationSeconds:
                          description: TolerationSeconds represents the period of
                            time the toleration (which must be of effect NoExecute,
                            otherwise this field is ignored) tolerates the taint.
                            By default, it is not set, which means tolerate the taint
                            forever (do not evict). Zero and negative values will
                            be treated as 0 (evict immediately) by the system.
                          format: int64
                          type: integer
                        value:
                          description: Value is the taint value the toleration matches
                            to. If the operator is Exists, the value should be empty,
                            otherwise just a regular string.
                          type: string
                      type: object
                    type: array
                type: object
            type: object
          status:
            description: CDIStatus defines the status of the installation
            properties:
              conditions:
                description: A list of current conditions of the resource
                items:
                  description: Condition represents the state of the operator's reconciliation
                    functionality.
                  properties:
                    lastHeartbeatTime:
                      format: date-time
                      type: string
                    lastTransitionTime:
                      format: date-time
                      type: string
                    message:
                      type: string
                    reason:
                      type: string
                    status:
                      type: string
                    type:
                      description: ConditionType is the state of the operator's reconciliation
                        functionality.
                      type: string
                  required:
                  - status
                  - type
                  type: object
                type: array
              observedVersion:
                description: The observed version of the resource
                type: string
              operatorVersion:
                description: The version of the resource as defined by the operator
                type: string
              phase:
                description: Phase is the current phase of the deployment
                type: string
              targetVersion:
                description: The desired version of the resource
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: false
    subresources: {}
  - additionalPrinterColumns:
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    - jsonPath: .status.phase
      name: Phase
      type: string
    name: v1beta1
    schema:
      openAPIV3Schema:
        description: CDI is the CDI Operator CRD
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: CDISpec defines our specification for the CDI installation
            properties:
              certConfig:
                description: certificate configuration
                properties:
                  ca:
                    description: CA configuration CA certs are kept in the CA bundle
                      as long as they are valid
                    properties:
                      duration:
                        description: The requested 'duration' (i.e. lifetime) of the
                          Certificate.
                        type: string
                      renewBefore:
                        description: The amount of time before the currently issued
                          certificate's ` + "`" + `notAfter` + "`" + ` time that we will begin to attempt
                          to renew the certificate.
                        type: string
                    type: object
                  server:
                    description: Server configuration Certs are rotated and discarded
                    properties:
                      duration:
                        description: The requested 'duration' (i.e. lifetime) of the
                          Certificate.
                        type: string
                      renewBefore:
                        description: The amount of time before the currently issued
                          certificate's ` + "`" + `notAfter` + "`" + ` time that we will begin to attempt
                          to renew the certificate.
                        type: string
                    type: object
                type: object
              cloneStrategyOverride:
                description: 'Clone strategy override: should we use a host-assisted
                  copy even if snapshots are available?'
                enum:
                - copy
                - snapshot
                - csi-clone
                type: string
              config:
                description: CDIConfig at CDI level
                properties:
                  dataVolumeTTLSeconds:
                    description: DataVolumeTTLSeconds is the time in seconds after
                      DataVolume completion it can be garbage collected. Disabled
                      by default.
                    format: int32
                    type: integer
                  featureGates:
                    description: FeatureGates are a list of specific enabled feature
                      gates
                    items:
                      type: string
                    type: array
                  filesystemOverhead:
                    description: FilesystemOverhead describes the space reserved for
                      overhead when using Filesystem volumes. A value is between 0
                      and 1, if not defined it is 0.055 (5.5% overhead)
                    properties:
                      global:
                        description: Global is how much space of a Filesystem volume
                          should be reserved for overhead. This value is used unless
                          overridden by a more specific value (per storageClass)
                        pattern: ^(0(?:\.\d{1,3})?|1)$
                        type: string
                      storageClass:
                        additionalProperties:
                          description: 'Percent is a string that can only be a value
                            between [0,1) (Note: we actually rely on reconcile to
                            reject invalid values)'
                          pattern: ^(0(?:\.\d{1,3})?|1)$
                          type: string
                        description: StorageClass specifies how much space of a Filesystem
                          volume should be reserved for safety. The keys are the storageClass
                          and the values are the overhead. This value overrides the
                          global value
                        type: object
                    type: object
                  imagePullSecrets:
                    description: The imagePullSecrets used to pull the container images
                    items:
                      description: LocalObjectReference contains enough information
                        to let you locate the referenced object inside the same namespace.
                      properties:
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                            TODO: Add other useful fields. apiVersion, kind, uid?'
                          type: string
                      type: object
                      x-kubernetes-map-type: atomic
                    type: array
                  importProxy:
                    description: ImportProxy contains importer pod proxy configuration.
                    properties:
                      HTTPProxy:
                        description: HTTPProxy is the URL http://<username>:<pswd>@<ip>:<port>
                          of the import proxy for HTTP requests.  Empty means unset
                          and will not result in the import pod env var.
                        type: string
                      HTTPSProxy:
                        description: HTTPSProxy is the URL https://<username>:<pswd>@<ip>:<port>
                          of the import proxy for HTTPS requests.  Empty means unset
                          and will not result in the import pod env var.
                        type: string
                      noProxy:
                        description: NoProxy is a comma-separated list of hostnames
                          and/or CIDRs for which the proxy should not be used. Empty
                          means unset and will not result in the import pod env var.
                        type: string
                      trustedCAProxy:
                        description: "TrustedCAProxy is the name of a ConfigMap in
                          the cdi namespace that contains a user-provided trusted
                          certificate authority (CA) bundle. The TrustedCAProxy ConfigMap
                          is consumed by the DataImportCron controller for creating
                          cronjobs, and by the import controller referring a copy
                          of the ConfigMap in the import namespace. Here is an example
                          of the ConfigMap (in yaml): \n apiVersion: v1 kind: ConfigMap
                          metadata: name: my-ca-proxy-cm namespace: cdi data: ca.pem:
                          | -----BEGIN CERTIFICATE----- ... <base64 encoded cert>
                          ... -----END CERTIFICATE-----"
                        type: string
                    type: object
                  insecureRegistries:
                    description: InsecureRegistries is a list of TLS disabled registries
                    items:
                      type: string
                    type: array
                  podResourceRequirements:
                    description: ResourceRequirements describes the compute resource
                      requirements.
                    properties:
                      claims:
                        description: "Claims lists the names of resources, defined
                          in spec.resourceClaims, that are used by this container.
                          \n This is an alpha field and requires enabling the DynamicResourceAllocation
                          feature gate. \n This field is immutable."
                        items:
                          description: ResourceClaim references one entry in PodSpec.ResourceClaims.
                          properties:
                            name:
                              description: Name must match the name of one entry in
                                pod.spec.resourceClaims of the Pod where this field
                                is used. It makes that resource available inside a
                                container.
                              type: string
                          required:
                          - name
                          type: object
                        type: array
                        x-kubernetes-list-map-keys:
                        - name
                        x-kubernetes-list-type: map
                      limits:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Limits describes the maximum amount of compute
                          resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                      requests:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Requests describes the minimum amount of compute
                          resources required. If Requests is omitted for a container,
                          it defaults to Limits if that is explicitly specified, otherwise
                          to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                    type: object
                  preallocation:
                    description: Preallocation controls whether storage for DataVolumes
                      should be allocated in advance.
                    type: boolean
                  scratchSpaceStorageClass:
                    description: 'Override the storage class to used for scratch space
                      during transfer operations. The scratch space storage class
                      is determined in the following order: 1. value of scratchSpaceStorageClass,
                      if that doesn''t exist, use the default storage class, if there
                      is no default storage class, use the storage class of the DataVolume,
                      if no storage class specified, use no storage class for scratch
                      space'
                    type: string
                  tlsSecurityProfile:
                    description: TLSSecurityProfile is used by operators to apply
                      cluster-wide TLS security settings to operands.
                    properties:
                      custom:
                        description: "custom is a user-defined TLS security profile.
                          Be extremely careful using a custom profile as invalid configurations
                          can be catastrophic. An example custom profile looks like
                          this: \n ciphers: - ECDHE-ECDSA-CHACHA20-POLY1305 - ECDHE-RSA-CHACHA20-POLY1305
                          - ECDHE-RSA-AES128-GCM-SHA256 - ECDHE-ECDSA-AES128-GCM-SHA256
                          minTLSVersion: TLSv1.1"
                        nullable: true
                        properties:
                          ciphers:
                            description: "ciphers is used to specify the cipher algorithms
                              that are negotiated during the TLS handshake.  Operators
                              may remove entries their operands do not support.  For
                              example, to use DES-CBC3-SHA  (yaml): \n ciphers: -
                              DES-CBC3-SHA"
                            items:
                              type: string
                            type: array
                          minTLSVersion:
                            description: "minTLSVersion is used to specify the minimal
                              version of the TLS protocol that is negotiated during
                              the TLS handshake. For example, to use TLS versions
                              1.1, 1.2 and 1.3 (yaml): \n minTLSVersion: TLSv1.1 \n
                              NOTE: currently the highest minTLSVersion allowed is
                              VersionTLS12"
                            enum:
                            - VersionTLS10
                            - VersionTLS11
                            - VersionTLS12
                            - VersionTLS13
                            type: string
                        type: object
                      intermediate:
                        description: "intermediate is a TLS security profile based
                          on: \n https://wiki.mozilla.org/Security/Server_Side_TLS#Intermediate_compatibility_.28recommended.29
                          \n and looks like this (yaml): \n ciphers: - TLS_AES_128_GCM_SHA256
                          - TLS_AES_256_GCM_SHA384 - TLS_CHACHA20_POLY1305_SHA256
                          - ECDHE-ECDSA-AES128-GCM-SHA256 - ECDHE-RSA-AES128-GCM-SHA256
                          - ECDHE-ECDSA-AES256-GCM-SHA384 - ECDHE-RSA-AES256-GCM-SHA384
                          - ECDHE-ECDSA-CHACHA20-POLY1305 - ECDHE-RSA-CHACHA20-POLY1305
                          - DHE-RSA-AES128-GCM-SHA256 - DHE-RSA-AES256-GCM-SHA384
                          minTLSVersion: TLSv1.2"
                        nullable: true
                        type: object
                      modern:
                        description: "modern is a TLS security profile based on: \n
                          https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility
                          \n and looks like this (yaml): \n ciphers: - TLS_AES_128_GCM_SHA256
                          - TLS_AES_256_GCM_SHA384 - TLS_CHACHA20_POLY1305_SHA256
                          minTLSVersion: TLSv1.3 \n NOTE: Currently unsupported."
                        nullable: true
                        type: object
                      old:
                        description: "old is a TLS security profile based on: \n https://wiki.mozilla.org/Security/Server_Side_TLS#Old_backward_compatibility
                          \n and looks like this (yaml): \n ciphers: - TLS_AES_128_GCM_SHA256
                          - TLS_AES_256_GCM_SHA384 - TLS_CHACHA20_POLY1305_SHA256
                          - ECDHE-ECDSA-AES128-GCM-SHA256 - ECDHE-RSA-AES128-GCM-SHA256
                          - ECDHE-ECDSA-AES256-GCM-SHA384 - ECDHE-RSA-AES256-GCM-SHA384
                          - ECDHE-ECDSA-CHACHA20-POLY1305 - ECDHE-RSA-CHACHA20-POLY1305
                          - DHE-RSA-AES128-GCM-SHA256 - DHE-RSA-AES256-GCM-SHA384
                          - DHE-RSA-CHACHA20-POLY1305 - ECDHE-ECDSA-AES128-SHA256
                          - ECDHE-RSA-AES128-SHA256 - ECDHE-ECDSA-AES128-SHA - ECDHE-RSA-AES128-SHA
                          - ECDHE-ECDSA-AES256-SHA384 - ECDHE-RSA-AES256-SHA384 -
                          ECDHE-ECDSA-AES256-SHA - ECDHE-RSA-AES256-SHA - DHE-RSA-AES128-SHA256
                          - DHE-RSA-AES256-SHA256 - AES128-GCM-SHA256 - AES256-GCM-SHA384
                          - AES128-SHA256 - AES256-SHA256 - AES128-SHA - AES256-SHA
                          - DES-CBC3-SHA minTLSVersion: TLSv1.0"
                        nullable: true
                        type: object
                      type:
                        description: "type is one of Old, Intermediate, Modern or
                          Custom. Custom provides the ability to specify individual
                          TLS security profile parameters. Old, Intermediate and Modern
                          are TLS security profiles based on: \n https://wiki.mozilla.org/Security/Server_Side_TLS#Recommended_configurations
                          \n The profiles are intent based, so they may change over
                          time as new ciphers are developed and existing ciphers are
                          found to be insecure.  Depending on precisely which ciphers
                          are available to a process, the list may be reduced. \n
                          Note that the Modern profile is currently not supported
                          because it is not yet well adopted by common software libraries."
                        enum:
                        - Old
                        - Intermediate
                        - Modern
                        - Custom
                        type: string
                    type: object
                  uploadProxyURLOverride:
                    description: Override the URL used when uploading to a DataVolume
                    type: string
                type: object
              imagePullPolicy:
                description: PullPolicy describes a policy for if/when to pull a container
                  image
                enum:
                - Always
                - IfNotPresent
                - Never
                type: string
              infra:
                description: Rules on which nodes CDI infrastructure pods will be
                  scheduled
                properties:
                  affinity:
                    description: affinity enables pod affinity/anti-affinity placement
                      expanding the types of constraints that can be expressed with
                      nodeSelector. affinity is going to be applied to the relevant
                      kind of pods in parallel with nodeSelector See https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity
                    properties:
                      nodeAffinity:
                        description: Describes node affinity scheduling rules for
                          the pod.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node matches the corresponding matchExpressions;
                              the node(s) with the highest sum are the most preferred.
                            items:
                              description: An empty preferred scheduling term matches
                                all objects with implicit weight 0 (i.e. it's a no-op).
                                A null preferred scheduling term matches no objects
                                (i.e. is also a no-op).
                              properties:
                                preference:
                                  description: A node selector term, associated with
                                    the corresponding weight.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                weight:
                                  description: Weight associated with matching the
                                    corresponding nodeSelectorTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - preference
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to an update), the system
                              may or may not try to eventually evict the pod from
                              its node.
                            properties:
                              nodeSelectorTerms:
                                description: Required. A list of node selector terms.
                                  The terms are ORed.
                                items:
                                  description: A null or empty node selector term
                                    matches no objects. The requirements of them are
                                    ANDed. The TopologySelectorTerm type implements
                                    a subset of the NodeSelectorTerm.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                type: array
                            required:
                            - nodeSelectorTerms
                            type: object
                            x-kubernetes-map-type: atomic
                        type: object
                      podAffinity:
                        description: Describes pod affinity scheduling rules (e.g.
                          co-locate this pod in the same node, zone, etc. as some
                          other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to a pod label update),
                              the system may or may not try to eventually evict the
                              pod from its node. When there are multiple elements,
                              the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                      podAntiAffinity:
                        description: Describes pod anti-affinity scheduling rules
                          (e.g. avoid putting this pod in the same node, zone, etc.
                          as some other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the anti-affinity expressions
                              specified by this field, but it may choose a node that
                              violates one or more of the expressions. The node that
                              is most preferred is the one with the greatest sum of
                              weights, i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              anti-affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the anti-affinity requirements specified
                              by this field are not met at scheduling time, the pod
                              will not be scheduled onto the node. If the anti-affinity
                              requirements specified by this field cease to be met
                              at some point during pod execution (e.g. due to a pod
                              label update), the system may or may not try to eventually
                              evict the pod from its node. When there are multiple
                              elements, the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                    type: object
                  nodeSelector:
                    additionalProperties:
                      type: string
                    description: 'nodeSelector is the node selector applied to the
                      relevant kind of pods It specifies a map of key-value pairs:
                      for the pod to be eligible to run on a node, the node must have
                      each of the indicated key-value pairs as labels (it can have
                      additional labels as well). See https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector'
                    type: object
                  tolerations:
                    description: tolerations is a list of tolerations applied to the
                      relevant kind of pods See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
                      for more info. These are additional tolerations other than default
                      ones.
                    items:
                      description: The pod this Toleration is attached to tolerates
                        any taint that matches the triple <key,value,effect> using
                        the matching operator <operator>.
                      properties:
                        effect:
                          description: Effect indicates the taint effect to match.
                            Empty means match all taint effects. When specified, allowed
                            values are NoSchedule, PreferNoSchedule and NoExecute.
                          type: string
                        key:
                          description: Key is the taint key that the toleration applies
                            to. Empty means match all taint keys. If the key is empty,
                            operator must be Exists; this combination means to match
                            all values and all keys.
                          type: string
                        operator:
                          description: Operator represents a key's relationship to
                            the value. Valid operators are Exists and Equal. Defaults
                            to Equal. Exists is equivalent to wildcard for value,
                            so that a pod can tolerate all taints of a particular
                            category.
                          type: string
                        tolerationSeconds:
                          description: TolerationSeconds represents the period of
                            time the toleration (which must be of effect NoExecute,
                            otherwise this field is ignored) tolerates the taint.
                            By default, it is not set, which means tolerate the taint
                            forever (do not evict). Zero and negative values will
                            be treated as 0 (evict immediately) by the system.
                          format: int64
                          type: integer
                        value:
                          description: Value is the taint value the toleration matches
                            to. If the operator is Exists, the value should be empty,
                            otherwise just a regular string.
                          type: string
                      type: object
                    type: array
                type: object
              priorityClass:
                description: PriorityClass of the CDI control plane
                type: string
              uninstallStrategy:
                description: CDIUninstallStrategy defines the state to leave CDI on
                  uninstall
                enum:
                - RemoveWorkloads
                - BlockUninstallIfWorkloadsExist
                type: string
              workload:
                description: Restrict on which nodes CDI workload pods will be scheduled
                properties:
                  affinity:
                    description: affinity enables pod affinity/anti-affinity placement
                      expanding the types of constraints that can be expressed with
                      nodeSelector. affinity is going to be applied to the relevant
                      kind of pods in parallel with nodeSelector See https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity
                    properties:
                      nodeAffinity:
                        description: Describes node affinity scheduling rules for
                          the pod.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node matches the corresponding matchExpressions;
                              the node(s) with the highest sum are the most preferred.
                            items:
                              description: An empty preferred scheduling term matches
                                all objects with implicit weight 0 (i.e. it's a no-op).
                                A null preferred scheduling term matches no objects
                                (i.e. is also a no-op).
                              properties:
                                preference:
                                  description: A node selector term, associated with
                                    the corresponding weight.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                weight:
                                  description: Weight associated with matching the
                                    corresponding nodeSelectorTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - preference
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to an update), the system
                              may or may not try to eventually evict the pod from
                              its node.
                            properties:
                              nodeSelectorTerms:
                                description: Required. A list of node selector terms.
                                  The terms are ORed.
                                items:
                                  description: A null or empty node selector term
                                    matches no objects. The requirements of them are
                                    ANDed. The TopologySelectorTerm type implements
                                    a subset of the NodeSelectorTerm.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                type: array
                            required:
                            - nodeSelectorTerms
                            type: object
                            x-kubernetes-map-type: atomic
                        type: object
                      podAffinity:
                        description: Describes pod affinity scheduling rules (e.g.
                          co-locate this pod in the same node, zone, etc. as some
                          other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to a pod label update),
                              the system may or may not try to eventually evict the
                              pod from its node. When there are multiple elements,
                              the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                      podAntiAffinity:
                        description: Describes pod anti-affinity scheduling rules
                          (e.g. avoid putting this pod in the same node, zone, etc.
                          as some other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the anti-affinity expressions
                              specified by this field, but it may choose a node that
                              violates one or more of the expressions. The node that
                              is most preferred is the one with the greatest sum of
                              weights, i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              anti-affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the anti-affinity requirements specified
                              by this field are not met at scheduling time, the pod
                              will not be scheduled onto the node. If the anti-affinity
                              requirements specified by this field cease to be met
                              at some point during pod execution (e.g. due to a pod
                              label update), the system may or may not try to eventually
                              evict the pod from its node. When there are multiple
                              elements, the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                    type: object
                  nodeSelector:
                    additionalProperties:
                      type: string
                    description: 'nodeSelector is the node selector applied to the
                      relevant kind of pods It specifies a map of key-value pairs:
                      for the pod to be eligible to run on a node, the node must have
                      each of the indicated key-value pairs as labels (it can have
                      additional labels as well). See https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector'
                    type: object
                  tolerations:
                    description: tolerations is a list of tolerations applied to the
                      relevant kind of pods See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
                      for more info. These are additional tolerations other than default
                      ones.
                    items:
                      description: The pod this Toleration is attached to tolerates
                        any taint that matches the triple <key,value,effect> using
                        the matching operator <operator>.
                      properties:
                        effect:
                          description: Effect indicates the taint effect to match.
                            Empty means match all taint effects. When specified, allowed
                            values are NoSchedule, PreferNoSchedule and NoExecute.
                          type: string
                        key:
                          description: Key is the taint key that the toleration applies
                            to. Empty means match all taint keys. If the key is empty,
                            operator must be Exists; this combination means to match
                            all values and all keys.
                          type: string
                        operator:
                          description: Operator represents a key's relationship to
                            the value. Valid operators are Exists and Equal. Defaults
                            to Equal. Exists is equivalent to wildcard for value,
                            so that a pod can tolerate all taints of a particular
                            category.
                          type: string
                        tolerationSeconds:
                          description: TolerationSeconds represents the period of
                            time the toleration (which must be of effect NoExecute,
                            otherwise this field is ignored) tolerates the taint.
                            By default, it is not set, which means tolerate the taint
                            forever (do not evict). Zero and negative values will
                            be treated as 0 (evict immediately) by the system.
                          format: int64
                          type: integer
                        value:
                          description: Value is the taint value the toleration matches
                            to. If the operator is Exists, the value should be empty,
                            otherwise just a regular string.
                          type: string
                      type: object
                    type: array
                type: object
            type: object
          status:
            description: CDIStatus defines the status of the installation
            properties:
              conditions:
                description: A list of current conditions of the resource
                items:
                  description: Condition represents the state of the operator's reconciliation
                    functionality.
                  properties:
                    lastHeartbeatTime:
                      format: date-time
                      type: string
                    lastTransitionTime:
                      format: date-time
                      type: string
                    message:
                      type: string
                    reason:
                      type: string
                    status:
                      type: string
                    type:
                      description: ConditionType is the state of the operator's reconciliation
                        functionality.
                      type: string
                  required:
                  - status
                  - type
                  type: object
                type: array
              observedVersion:
                description: The observed version of the resource
                type: string
              operatorVersion:
                description: The version of the resource as defined by the operator
                type: string
              phase:
                description: Phase is the current phase of the deployment
                type: string
              targetVersion:
                description: The desired version of the resource
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"cdiconfig": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: cdiconfigs.cdi.kubevirt.io
spec:
  group: cdi.kubevirt.io
  names:
    kind: CDIConfig
    listKind: CDIConfigList
    plural: cdiconfigs
    singular: cdiconfig
  scope: Cluster
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: CDIConfig provides a user configuration for CDI
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: CDIConfigSpec defines specification for user configuration
            properties:
              dataVolumeTTLSeconds:
                description: DataVolumeTTLSeconds is the time in seconds after DataVolume
                  completion it can be garbage collected. Disabled by default.
                format: int32
                type: integer
              featureGates:
                description: FeatureGates are a list of specific enabled feature gates
                items:
                  type: string
                type: array
              filesystemOverhead:
                description: FilesystemOverhead describes the space reserved for overhead
                  when using Filesystem volumes. A value is between 0 and 1, if not
                  defined it is 0.055 (5.5% overhead)
                properties:
                  global:
                    description: Global is how much space of a Filesystem volume should
                      be reserved for overhead. This value is used unless overridden
                      by a more specific value (per storageClass)
                    pattern: ^(0(?:\.\d{1,3})?|1)$
                    type: string
                  storageClass:
                    additionalProperties:
                      description: 'Percent is a string that can only be a value between
                        [0,1) (Note: we actually rely on reconcile to reject invalid
                        values)'
                      pattern: ^(0(?:\.\d{1,3})?|1)$
                      type: string
                    description: StorageClass specifies how much space of a Filesystem
                      volume should be reserved for safety. The keys are the storageClass
                      and the values are the overhead. This value overrides the global
                      value
                    type: object
                type: object
              imagePullSecrets:
                description: The imagePullSecrets used to pull the container images
                items:
                  description: LocalObjectReference contains enough information to
                    let you locate the referenced object inside the same namespace.
                  properties:
                    name:
                      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                        TODO: Add other useful fields. apiVersion, kind, uid?'
                      type: string
                  type: object
                  x-kubernetes-map-type: atomic
                type: array
              importProxy:
                description: ImportProxy contains importer pod proxy configuration.
                properties:
                  HTTPProxy:
                    description: HTTPProxy is the URL http://<username>:<pswd>@<ip>:<port>
                      of the import proxy for HTTP requests.  Empty means unset and
                      will not result in the import pod env var.
                    type: string
                  HTTPSProxy:
                    description: HTTPSProxy is the URL https://<username>:<pswd>@<ip>:<port>
                      of the import proxy for HTTPS requests.  Empty means unset and
                      will not result in the import pod env var.
                    type: string
                  noProxy:
                    description: NoProxy is a comma-separated list of hostnames and/or
                      CIDRs for which the proxy should not be used. Empty means unset
                      and will not result in the import pod env var.
                    type: string
                  trustedCAProxy:
                    description: "TrustedCAProxy is the name of a ConfigMap in the
                      cdi namespace that contains a user-provided trusted certificate
                      authority (CA) bundle. The TrustedCAProxy ConfigMap is consumed
                      by the DataImportCron controller for creating cronjobs, and
                      by the import controller referring a copy of the ConfigMap in
                      the import namespace. Here is an example of the ConfigMap (in
                      yaml): \n apiVersion: v1 kind: ConfigMap metadata: name: my-ca-proxy-cm
                      namespace: cdi data: ca.pem: | -----BEGIN CERTIFICATE----- ...
                      <base64 encoded cert> ... -----END CERTIFICATE-----"
                    type: string
                type: object
              insecureRegistries:
                description: InsecureRegistries is a list of TLS disabled registries
                items:
                  type: string
                type: array
              podResourceRequirements:
                description: ResourceRequirements describes the compute resource requirements.
                properties:
                  claims:
                    description: "Claims lists the names of resources, defined in
                      spec.resourceClaims, that are used by this container. \n This
                      is an alpha field and requires enabling the DynamicResourceAllocation
                      feature gate. \n This field is immutable."
                    items:
                      description: ResourceClaim references one entry in PodSpec.ResourceClaims.
                      properties:
                        name:
                          description: Name must match the name of one entry in pod.spec.resourceClaims
                            of the Pod where this field is used. It makes that resource
                            available inside a container.
                          type: string
                      required:
                      - name
                      type: object
                    type: array
                    x-kubernetes-list-map-keys:
                    - name
                    x-kubernetes-list-type: map
                  limits:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: 'Limits describes the maximum amount of compute resources
                      allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                    type: object
                  requests:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: 'Requests describes the minimum amount of compute
                      resources required. If Requests is omitted for a container,
                      it defaults to Limits if that is explicitly specified, otherwise
                      to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                    type: object
                type: object
              preallocation:
                description: Preallocation controls whether storage for DataVolumes
                  should be allocated in advance.
                type: boolean
              scratchSpaceStorageClass:
                description: 'Override the storage class to used for scratch space
                  during transfer operations. The scratch space storage class is determined
                  in the following order: 1. value of scratchSpaceStorageClass, if
                  that doesn''t exist, use the default storage class, if there is
                  no default storage class, use the storage class of the DataVolume,
                  if no storage class specified, use no storage class for scratch
                  space'
                type: string
              tlsSecurityProfile:
                description: TLSSecurityProfile is used by operators to apply cluster-wide
                  TLS security settings to operands.
                properties:
                  custom:
                    description: "custom is a user-defined TLS security profile. Be
                      extremely careful using a custom profile as invalid configurations
                      can be catastrophic. An example custom profile looks like this:
                      \n ciphers: - ECDHE-ECDSA-CHACHA20-POLY1305 - ECDHE-RSA-CHACHA20-POLY1305
                      - ECDHE-RSA-AES128-GCM-SHA256 - ECDHE-ECDSA-AES128-GCM-SHA256
                      minTLSVersion: TLSv1.1"
                    nullable: true
                    properties:
                      ciphers:
                        description: "ciphers is used to specify the cipher algorithms
                          that are negotiated during the TLS handshake.  Operators
                          may remove entries their operands do not support.  For example,
                          to use DES-CBC3-SHA  (yaml): \n ciphers: - DES-CBC3-SHA"
                        items:
                          type: string
                        type: array
                      minTLSVersion:
                        description: "minTLSVersion is used to specify the minimal
                          version of the TLS protocol that is negotiated during the
                          TLS handshake. For example, to use TLS versions 1.1, 1.2
                          and 1.3 (yaml): \n minTLSVersion: TLSv1.1 \n NOTE: currently
                          the highest minTLSVersion allowed is VersionTLS12"
                        enum:
                        - VersionTLS10
                        - VersionTLS11
                        - VersionTLS12
                        - VersionTLS13
                        type: string
                    type: object
                  intermediate:
                    description: "intermediate is a TLS security profile based on:
                      \n https://wiki.mozilla.org/Security/Server_Side_TLS#Intermediate_compatibility_.28recommended.29
                      \n and looks like this (yaml): \n ciphers: - TLS_AES_128_GCM_SHA256
                      - TLS_AES_256_GCM_SHA384 - TLS_CHACHA20_POLY1305_SHA256 - ECDHE-ECDSA-AES128-GCM-SHA256
                      - ECDHE-RSA-AES128-GCM-SHA256 - ECDHE-ECDSA-AES256-GCM-SHA384
                      - ECDHE-RSA-AES256-GCM-SHA384 - ECDHE-ECDSA-CHACHA20-POLY1305
                      - ECDHE-RSA-CHACHA20-POLY1305 - DHE-RSA-AES128-GCM-SHA256 -
                      DHE-RSA-AES256-GCM-SHA384 minTLSVersion: TLSv1.2"
                    nullable: true
                    type: object
                  modern:
                    description: "modern is a TLS security profile based on: \n https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility
                      \n and looks like this (yaml): \n ciphers: - TLS_AES_128_GCM_SHA256
                      - TLS_AES_256_GCM_SHA384 - TLS_CHACHA20_POLY1305_SHA256 minTLSVersion:
                      TLSv1.3 \n NOTE: Currently unsupported."
                    nullable: true
                    type: object
                  old:
                    description: "old is a TLS security profile based on: \n https://wiki.mozilla.org/Security/Server_Side_TLS#Old_backward_compatibility
                      \n and looks like this (yaml): \n ciphers: - TLS_AES_128_GCM_SHA256
                      - TLS_AES_256_GCM_SHA384 - TLS_CHACHA20_POLY1305_SHA256 - ECDHE-ECDSA-AES128-GCM-SHA256
                      - ECDHE-RSA-AES128-GCM-SHA256 - ECDHE-ECDSA-AES256-GCM-SHA384
                      - ECDHE-RSA-AES256-GCM-SHA384 - ECDHE-ECDSA-CHACHA20-POLY1305
                      - ECDHE-RSA-CHACHA20-POLY1305 - DHE-RSA-AES128-GCM-SHA256 -
                      DHE-RSA-AES256-GCM-SHA384 - DHE-RSA-CHACHA20-POLY1305 - ECDHE-ECDSA-AES128-SHA256
                      - ECDHE-RSA-AES128-SHA256 - ECDHE-ECDSA-AES128-SHA - ECDHE-RSA-AES128-SHA
                      - ECDHE-ECDSA-AES256-SHA384 - ECDHE-RSA-AES256-SHA384 - ECDHE-ECDSA-AES256-SHA
                      - ECDHE-RSA-AES256-SHA - DHE-RSA-AES128-SHA256 - DHE-RSA-AES256-SHA256
                      - AES128-GCM-SHA256 - AES256-GCM-SHA384 - AES128-SHA256 - AES256-SHA256
                      - AES128-SHA - AES256-SHA - DES-CBC3-SHA minTLSVersion: TLSv1.0"
                    nullable: true
                    type: object
                  type:
                    description: "type is one of Old, Intermediate, Modern or Custom.
                      Custom provides the ability to specify individual TLS security
                      profile parameters. Old, Intermediate and Modern are TLS security
                      profiles based on: \n https://wiki.mozilla.org/Security/Server_Side_TLS#Recommended_configurations
                      \n The profiles are intent based, so they may change over time
                      as new ciphers are developed and existing ciphers are found
                      to be insecure.  Depending on precisely which ciphers are available
                      to a process, the list may be reduced. \n Note that the Modern
                      profile is currently not supported because it is not yet well
                      adopted by common software libraries."
                    enum:
                    - Old
                    - Intermediate
                    - Modern
                    - Custom
                    type: string
                type: object
              uploadProxyURLOverride:
                description: Override the URL used when uploading to a DataVolume
                type: string
            type: object
          status:
            description: CDIConfigStatus provides the most recently observed status
              of the CDI Config resource
            properties:
              defaultPodResourceRequirements:
                description: ResourceRequirements describes the compute resource requirements.
                properties:
                  claims:
                    description: "Claims lists the names of resources, defined in
                      spec.resourceClaims, that are used by this container. \n This
                      is an alpha field and requires enabling the DynamicResourceAllocation
                      feature gate. \n This field is immutable."
                    items:
                      description: ResourceClaim references one entry in PodSpec.ResourceClaims.
                      properties:
                        name:
                          description: Name must match the name of one entry in pod.spec.resourceClaims
                            of the Pod where this field is used. It makes that resource
                            available inside a container.
                          type: string
                      required:
                      - name
                      type: object
                    type: array
                    x-kubernetes-list-map-keys:
                    - name
                    x-kubernetes-list-type: map
                  limits:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: 'Limits describes the maximum amount of compute resources
                      allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                    type: object
                  requests:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: 'Requests describes the minimum amount of compute
                      resources required. If Requests is omitted for a container,
                      it defaults to Limits if that is explicitly specified, otherwise
                      to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                    type: object
                type: object
              filesystemOverhead:
                description: FilesystemOverhead describes the space reserved for overhead
                  when using Filesystem volumes. A percentage value is between 0 and
                  1
                properties:
                  global:
                    description: Global is how much space of a Filesystem volume should
                      be reserved for overhead. This value is used unless overridden
                      by a more specific value (per storageClass)
                    pattern: ^(0(?:\.\d{1,3})?|1)$
                    type: string
                  storageClass:
                    additionalProperties:
                      description: 'Percent is a string that can only be a value between
                        [0,1) (Note: we actually rely on reconcile to reject invalid
                        values)'
                      pattern: ^(0(?:\.\d{1,3})?|1)$
                      type: string
                    description: StorageClass specifies how much space of a Filesystem
                      volume should be reserved for safety. The keys are the storageClass
                      and the values are the overhead. This value overrides the global
                      value
                    type: object
                type: object
              imagePullSecrets:
                description: The imagePullSecrets used to pull the container images
                items:
                  description: LocalObjectReference contains enough information to
                    let you locate the referenced object inside the same namespace.
                  properties:
                    name:
                      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                        TODO: Add other useful fields. apiVersion, kind, uid?'
                      type: string
                  type: object
                  x-kubernetes-map-type: atomic
                type: array
              importProxy:
                description: ImportProxy contains importer pod proxy configuration.
                properties:
                  HTTPProxy:
                    description: HTTPProxy is the URL http://<username>:<pswd>@<ip>:<port>
                      of the import proxy for HTTP requests.  Empty means unset and
                      will not result in the import pod env var.
                    type: string
                  HTTPSProxy:
                    description: HTTPSProxy is the URL https://<username>:<pswd>@<ip>:<port>
                      of the import proxy for HTTPS requests.  Empty means unset and
                      will not result in the import pod env var.
                    type: string
                  noProxy:
                    description: NoProxy is a comma-separated list of hostnames and/or
                      CIDRs for which the proxy should not be used. Empty means unset
                      and will not result in the import pod env var.
                    type: string
                  trustedCAProxy:
                    description: "TrustedCAProxy is the name of a ConfigMap in the
                      cdi namespace that contains a user-provided trusted certificate
                      authority (CA) bundle. The TrustedCAProxy ConfigMap is consumed
                      by the DataImportCron controller for creating cronjobs, and
                      by the import controller referring a copy of the ConfigMap in
                      the import namespace. Here is an example of the ConfigMap (in
                      yaml): \n apiVersion: v1 kind: ConfigMap metadata: name: my-ca-proxy-cm
                      namespace: cdi data: ca.pem: | -----BEGIN CERTIFICATE----- ...
                      <base64 encoded cert> ... -----END CERTIFICATE-----"
                    type: string
                type: object
              preallocation:
                description: Preallocation controls whether storage for DataVolumes
                  should be allocated in advance.
                type: boolean
              scratchSpaceStorageClass:
                description: The calculated storage class to be used for scratch space
                type: string
              uploadProxyURL:
                description: The calculated upload proxy URL
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"dataimportcron": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: dataimportcrons.cdi.kubevirt.io
spec:
  group: cdi.kubevirt.io
  names:
    categories:
    - all
    kind: DataImportCron
    listKind: DataImportCronList
    plural: dataimportcrons
    shortNames:
    - dic
    - dics
    singular: dataimportcron
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: DataImportCron defines a cron job for recurring polling/importing
          disk images as PVCs into a golden image namespace
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: DataImportCronSpec defines specification for DataImportCron
            properties:
              garbageCollect:
                description: GarbageCollect specifies whether old PVCs should be cleaned
                  up after a new PVC is imported. Options are currently "Outdated"
                  and "Never", defaults to "Outdated".
                type: string
              importsToKeep:
                description: Number of import PVCs to keep when garbage collecting.
                  Default is 3.
                format: int32
                type: integer
              managedDataSource:
                description: ManagedDataSource specifies the name of the corresponding
                  DataSource this cron will manage. DataSource has to be in the same
                  namespace.
                type: string
              retentionPolicy:
                description: RetentionPolicy specifies whether the created DataVolumes
                  and DataSources are retained when their DataImportCron is deleted.
                  Default is RatainAll.
                type: string
              schedule:
                description: Schedule specifies in cron format when and how often
                  to look for new imports
                type: string
              template:
                description: Template specifies template for the DVs to be created
                properties:
                  apiVersion:
                    description: 'APIVersion defines the versioned schema of this
                      representation of an object. Servers should convert recognized
                      schemas to the latest internal value, and may reject unrecognized
                      values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
                    type: string
                  kind:
                    description: 'Kind is a string value representing the REST resource
                      this object represents. Servers may infer this from the endpoint
                      the client submits requests to. Cannot be updated. In CamelCase.
                      More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                    type: string
                  metadata:
                    type: object
                  spec:
                    description: DataVolumeSpec defines the DataVolume type specification
                    properties:
                      checkpoints:
                        description: Checkpoints is a list of DataVolumeCheckpoints,
                          representing stages in a multistage import.
                        items:
                          description: DataVolumeCheckpoint defines a stage in a warm
                            migration.
                          properties:
                            current:
                              description: Current is the identifier of the snapshot
                                created for this checkpoint.
                              type: string
                            previous:
                              description: Previous is the identifier of the snapshot
                                from the previous checkpoint.
                              type: string
                          required:
                          - current
                          - previous
                          type: object
                        type: array
                      contentType:
                        description: 'DataVolumeContentType options: "kubevirt", "archive"'
                        enum:
                        - kubevirt
                        - archive
                        type: string
                      finalCheckpoint:
                        description: FinalCheckpoint indicates whether the current
                          DataVolumeCheckpoint is the final checkpoint.
                        type: boolean
                      preallocation:
                        description: Preallocation controls whether storage for DataVolumes
                          should be allocated in advance.
                        type: boolean
                      priorityClassName:
                        description: PriorityClassName for Importer, Cloner and Uploader
                          pod
                        type: string
                      pvc:
                        description: PVC is the PVC specification
                        properties:
                          accessModes:
                            description: 'accessModes contains the desired access
                              modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1'
                            items:
                              type: string
                            type: array
                          dataSource:
                            description: 'dataSource field can be used to specify
                              either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)
                              * An existing PVC (PersistentVolumeClaim) If the provisioner
                              or an external controller can support the specified
                              data source, it will create a new volume based on the
                              contents of the specified data source. When the AnyVolumeDataSource
                              feature gate is enabled, dataSource contents will be
                              copied to dataSourceRef, and dataSourceRef contents
                              will be copied to dataSource when dataSourceRef.namespace
                              is not specified. If the namespace is specified, then
                              dataSourceRef will not be copied to dataSource.'
                            properties:
                              apiGroup:
                                description: APIGroup is the group for the resource
                                  being referenced. If APIGroup is not specified,
                                  the specified Kind must be in the core API group.
                                  For any other third-party types, APIGroup is required.
                                type: string
                              kind:
                                description: Kind is the type of resource being referenced
                                type: string
                              name:
                                description: Name is the name of resource being referenced
                                type: string
                            required:
                            - kind
                            - name
                            type: object
                            x-kubernetes-map-type: atomic
                          dataSourceRef:
                            description: 'dataSourceRef specifies the object from
                              which to populate the volume with data, if a non-empty
                              volume is desired. This may be any object from a non-empty
                              API group (non core object) or a PersistentVolumeClaim
                              object. When this field is specified, volume binding
                              will only succeed if the type of the specified object
                              matches some installed volume populator or dynamic provisioner.
                              This field will replace the functionality of the dataSource
                              field and as such if both fields are non-empty, they
                              must have the same value. For backwards compatibility,
                              when namespace isn''t specified in dataSourceRef, both
                              fields (dataSource and dataSourceRef) will be set to
                              the same value automatically if one of them is empty
                              and the other is non-empty. When namespace is specified
                              in dataSourceRef, dataSource isn''t set to the same
                              value and must be empty. There are three important differences
                              between dataSource and dataSourceRef: * While dataSource
                              only allows two specific types of objects, dataSourceRef
                              allows any non-core object, as well as PersistentVolumeClaim
                              objects. * While dataSource ignores disallowed values
                              (dropping them), dataSourceRef preserves all values,
                              and generates an error if a disallowed value is specified.
                              * While dataSource only allows local objects, dataSourceRef
                              allows objects in any namespaces. (Beta) Using this
                              field requires the AnyVolumeDataSource feature gate
                              to be enabled. (Alpha) Using the namespace field of
                              dataSourceRef requires the CrossNamespaceVolumeDataSource
                              feature gate to be enabled.'
                            properties:
                              apiGroup:
                                description: APIGroup is the group for the resource
                                  being referenced. If APIGroup is not specified,
                                  the specified Kind must be in the core API group.
                                  For any other third-party types, APIGroup is required.
                                type: string
                              kind:
                                description: Kind is the type of resource being referenced
                                type: string
                              name:
                                description: Name is the name of resource being referenced
                                type: string
                              namespace:
                                description: Namespace is the namespace of resource
                                  being referenced Note that when a namespace is specified,
                                  a gateway.networking.k8s.io/ReferenceGrant object
                                  is required in the referent namespace to allow that
                                  namespace's owner to accept the reference. See the
                                  ReferenceGrant documentation for details. (Alpha)
                                  This field requires the CrossNamespaceVolumeDataSource
                                  feature gate to be enabled.
                                type: string
                            required:
                            - kind
                            - name
                            type: object
                          resources:
                            description: 'resources represents the minimum resources
                              the volume should have. If RecoverVolumeExpansionFailure
                              feature is enabled users are allowed to specify resource
                              requirements that are lower than previous value but
                              must still be higher than capacity recorded in the status
                              field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources'
                            properties:
                              claims:
                                description: "Claims lists the names of resources,
                                  defined in spec.resourceClaims, that are used by
                                  this container. \n This is an alpha field and requires
                                  enabling the DynamicResourceAllocation feature gate.
                                  \n This field is immutable."
                                items:
                                  description: ResourceClaim references one entry
                                    in PodSpec.ResourceClaims.
                                  properties:
                                    name:
                                      description: Name must match the name of one
                                        entry in pod.spec.resourceClaims of the Pod
                                        where this field is used. It makes that resource
                                        available inside a container.
                                      type: string
                                  required:
                                  - name
                                  type: object
                                type: array
                                x-kubernetes-list-map-keys:
                                - name
                                x-kubernetes-list-type: map
                              limits:
                                additionalProperties:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                  x-kubernetes-int-or-string: true
                                description: 'Limits describes the maximum amount
                                  of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                                type: object
                              requests:
                                additionalProperties:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                  x-kubernetes-int-or-string: true
                                description: 'Requests describes the minimum amount
                                  of compute resources required. If Requests is omitted
                                  for a container, it defaults to Limits if that is
                                  explicitly specified, otherwise to an implementation-defined
                                  value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                                type: object
                            type: object
                          selector:
                            description: selector is a label query over volumes to
                              consider for binding.
                            properties:
                              matchExpressions:
                                description: matchExpressions is a list of label selector
                                  requirements. The requirements are ANDed.
                                items:
                                  description: A label selector requirement is a selector
                                    that contains values, a key, and an operator that
                                    relates the key and values.
                                  properties:
                                    key:
                                      description: key is the label key that the selector
                                        applies to.
                                      type: string
                                    operator:
                                      description: operator represents a key's relationship
                                        to a set of values. Valid operators are In,
                                        NotIn, Exists and DoesNotExist.
                                      type: string
                                    values:
                                      description: values is an array of string values.
                                        If the operator is In or NotIn, the values
                                        array must be non-empty. If the operator is
                                        Exists or DoesNotExist, the values array must
                                        be empty. This array is replaced during a
                                        strategic merge patch.
                                      items:
                                        type: string
                                      type: array
                                  required:
                                  - key
                                  - operator
                                  type: object
                                type: array
                              matchLabels:
                                additionalProperties:
                                  type: string
                                description: matchLabels is a map of {key,value} pairs.
                                  A single {key,value} in the matchLabels map is equivalent
                                  to an element of matchExpressions, whose key field
                                  is "key", the operator is "In", and the values array
                                  contains only "value". The requirements are ANDed.
                                type: object
                            type: object
                            x-kubernetes-map-type: atomic
                          storageClassName:
                            description: 'storageClassName is the name of the StorageClass
                              required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1'
                            type: string
                          volumeMode:
                            description: volumeMode defines what type of volume is
                              required by the claim. Value of Filesystem is implied
                              when not included in claim spec.
                            type: string
                          volumeName:
                            description: volumeName is the binding reference to the
                              PersistentVolume backing this claim.
                            type: string
                        type: object
                      source:
                        description: Source is the src of the data for the requested
                          DataVolume
                        properties:
                          blank:
                            description: DataVolumeBlankImage provides the parameters
                              to create a new raw blank image for the PVC
                            type: object
                          gcs:
                            description: DataVolumeSourceGCS provides the parameters
                              to create a Data Volume from an GCS source
                            properties:
                              secretRef:
                                description: SecretRef provides the secret reference
                                  needed to access the GCS source
                                type: string
                              url:
                                description: URL is the url of the GCS source
                                type: string
                            required:
                            - url
                            type: object
                          http:
                            description: DataVolumeSourceHTTP can be either an http
                              or https endpoint, with an optional basic auth user
                              name and password, and an optional configmap containing
                              additional CAs
                            properties:
                              certConfigMap:
                                description: CertConfigMap is a configmap reference,
                                  containing a Certificate Authority(CA) public key,
                                  and a base64 encoded pem certificate
                                type: string
                              extraHeaders:
                                description: ExtraHeaders is a list of strings containing
                                  extra headers to include with HTTP transfer requests
                                items:
                                  type: string
                                type: array
                              secretExtraHeaders:
                                description: SecretExtraHeaders is a list of Secret
                                  references, each containing an extra HTTP header
                                  that may include sensitive information
                                items:
                                  type: string
                                type: array
                              secretRef:
                                description: SecretRef A Secret reference, the secret
                                  should contain accessKeyId (user name) base64 encoded,
                                  and secretKey (password) also base64 encoded
                                type: string
                              url:
                                description: URL is the URL of the http(s) endpoint
                                type: string
                            required:
                            - url
                            type: object
                          imageio:
                            description: DataVolumeSourceImageIO provides the parameters
                              to create a Data Volume from an imageio source
                            properties:
                              certConfigMap:
                                description: CertConfigMap provides a reference to
                                  the CA cert
                                type: string
                              diskId:
                                description: DiskID provides id of a disk to be imported
                                type: string
                              secretRef:
                                description: SecretRef provides the secret reference
                                  needed to access the ovirt-engine
                                type: string
                              url:
                                description: URL is the URL of the ovirt-engine
                                type: string
                            required:
                            - diskId
                            - url
                            type: object
                          pvc:
                            description: DataVolumeSourcePVC provides the parameters
                              to create a Data Volume from an existing PVC
                            properties:
                              name:
                                description: The name of the source PVC
                                type: string
                              namespace:
                                description: The namespace of the source PVC
                                type: string
                            required:
                            - name
                            - namespace
                            type: object
                          registry:
                            description: DataVolumeSourceRegistry provides the parameters
                              to create a Data Volume from an registry source
                            properties:
                              certConfigMap:
                                description: CertConfigMap provides a reference to
                                  the Registry certs
                                type: string
                              imageStream:
                                description: ImageStream is the name of image stream
                                  for import
                                type: string
                              pullMethod:
                                description: PullMethod can be either "pod" (default
                                  import), or "node" (node docker cache based import)
                                type: string
                              secretRef:
                                description: SecretRef provides the secret reference
                                  needed to access the Registry source
                                type: string
                              url:
                                description: 'URL is the url of the registry source
                                  (starting with the scheme: docker, oci-archive)'
                                type: string
                            type: object
                          s3:
                            description: DataVolumeSourceS3 provides the parameters
                              to create a Data Volume from an S3 source
                            properties:
                              certConfigMap:
                                description: CertConfigMap is a configmap reference,
                                  containing a Certificate Authority(CA) public key,
                                  and a base64 encoded pem certificate
                                type: string
                              secretRef:
                                description: SecretRef provides the secret reference
                                  needed to access the S3 source
                                type: string
                              url:
                                description: URL is the url of the S3 source
                                type: string
                            required:
                            - url
                            type: object
                          snapshot:
                            description: DataVolumeSourceSnapshot provides the parameters
                              to create a Data Volume from an existing VolumeSnapshot
                            properties:
                              name:
                                description: The name of the source VolumeSnapshot
                                type: string
                              namespace:
                                description: The namespace of the source VolumeSnapshot
                                type: string
                            required:
                            - name
                            - namespace
                            type: object
                          upload:
                            description: DataVolumeSourceUpload provides the parameters
                              to create a Data Volume by uploading the source
                            type: object
                          vddk:
                            description: DataVolumeSourceVDDK provides the parameters
                              to create a Data Volume from a Vmware source
                            properties:
                              backingFile:
                                description: BackingFile is the path to the virtual
                                  hard disk to migrate from vCenter/ESXi
                                type: string
                              initImageURL:
                                description: InitImageURL is an optional URL to an
                                  image containing an extracted VDDK library, overrides
                                  v2v-vmware config map
                                type: string
                              secretRef:
                                description: SecretRef provides a reference to a secret
                                  containing the username and password needed to access
                                  the vCenter or ESXi host
                                type: string
                              thumbprint:
                                description: Thumbprint is the certificate thumbprint
                                  of the vCenter or ESXi host
                                type: string
                              url:
                                description: URL is the URL of the vCenter or ESXi
                                  host with the VM to migrate
                                type: string
                              uuid:
                                description: UUID is the UUID of the virtual machine
                                  that the backing file is attached to in vCenter/ESXi
                                type: string
                            type: object
                        type: object
                      sourceRef:
                        description: SourceRef is an indirect reference to the source
                          of data for the requested DataVolume
                        properties:
                          kind:
                            description: The kind of the source reference, currently
                              only "DataSource" is supported
                            type: string
                          name:
                            description: The name of the source reference
                            type: string
                          namespace:
                            description: The namespace of the source reference, defaults
                              to the DataVolume namespace
                            type: string
                        required:
                        - kind
                        - name
                        type: object
                      storage:
                        description: Storage is the requested storage specification
                        properties:
                          accessModes:
                            description: 'AccessModes contains the desired access
                              modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1'
                            items:
                              type: string
                            type: array
                          dataSource:
                            description: 'This field can be used to specify either:
                              * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)
                              * An existing PVC (PersistentVolumeClaim) * An existing
                              custom resource that implements data population (Alpha)
                              In order to use custom resource types that implement
                              data population, the AnyVolumeDataSource feature gate
                              must be enabled. If the provisioner or an external controller
                              can support the specified data source, it will create
                              a new volume based on the contents of the specified
                              data source. If the AnyVolumeDataSource feature gate
                              is enabled, this field will always have the same contents
                              as the DataSourceRef field.'
                            properties:
                              apiGroup:
                                description: APIGroup is the group for the resource
                                  being referenced. If APIGroup is not specified,
                                  the specified Kind must be in the core API group.
                                  For any other third-party types, APIGroup is required.
                                type: string
                              kind:
                                description: Kind is the type of resource being referenced
                                type: string
                              name:
                                description: Name is the name of resource being referenced
                                type: string
                            required:
                            - kind
                            - name
                            type: object
                            x-kubernetes-map-type: atomic
                          dataSourceRef:
                            description: 'Specifies the object from which to populate
                              the volume with data, if a non-empty volume is desired.
                              This may be any local object from a non-empty API group
                              (non core object) or a PersistentVolumeClaim object.
                              When this field is specified, volume binding will only
                              succeed if the type of the specified object matches
                              some installed volume populator or dynamic provisioner.
                              This field will replace the functionality of the DataSource
                              field and as such if both fields are non-empty, they
                              must have the same value. For backwards compatibility,
                              both fields (DataSource and DataSourceRef) will be set
                              to the same value automatically if one of them is empty
                              and the other is non-empty. There are two important
                              differences between DataSource and DataSourceRef: *
                              While DataSource only allows two specific types of objects,
                              DataSourceRef allows any non-core object, as well as
                              PersistentVolumeClaim objects. * While DataSource ignores
                              disallowed values (dropping them), DataSourceRef preserves
                              all values, and generates an error if a disallowed value
                              is specified. (Beta) Using this field requires the AnyVolumeDataSource
                              feature gate to be enabled.'
                            properties:
                              apiGroup:
                                description: APIGroup is the group for the resource
                                  being referenced. If APIGroup is not specified,
                                  the specified Kind must be in the core API group.
                                  For any other third-party types, APIGroup is required.
                                type: string
                              kind:
                                description: Kind is the type of resource being referenced
                                type: string
                              name:
                                description: Name is the name of resource being referenced
                                type: string
                              namespace:
                                description: Namespace is the namespace of resource
                                  being referenced Note that when a namespace is specified,
                                  a gateway.networking.k8s.io/ReferenceGrant object
                                  is required in the referent namespace to allow that
                                  namespace's owner to accept the reference. See the
                                  ReferenceGrant documentation for details. (Alpha)
                                  This field requires the CrossNamespaceVolumeDataSource
                                  feature gate to be enabled.
                                type: string
                            required:
                            - kind
                            - name
                            type: object
                          resources:
                            description: 'Resources represents the minimum resources
                              the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources'
                            properties:
                              claims:
                                description: "Claims lists the names of resources,
                                  defined in spec.resourceClaims, that are used by
                                  this container. \n This is an alpha field and requires
                                  enabling the DynamicResourceAllocation feature gate.
                                  \n This field is immutable."
                                items:
                                  description: ResourceClaim references one entry
                                    in PodSpec.ResourceClaims.
                                  properties:
                                    name:
                                      description: Name must match the name of one
                                        entry in pod.spec.resourceClaims of the Pod
                                        where this field is used. It makes that resource
                                        available inside a container.
                                      type: string
                                  required:
                                  - name
                                  type: object
                                type: array
                                x-kubernetes-list-map-keys:
                                - name
                                x-kubernetes-list-type: map
                              limits:
                                additionalProperties:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                  x-kubernetes-int-or-string: true
                                description: 'Limits describes the maximum amount
                                  of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                                type: object
                              requests:
                                additionalProperties:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                  x-kubernetes-int-or-string: true
                                description: 'Requests describes the minimum amount
                                  of compute resources required. If Requests is omitted
                                  for a container, it defaults to Limits if that is
                                  explicitly specified, otherwise to an implementation-defined
                                  value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                                type: object
                            type: object
                          selector:
                            description: A label query over volumes to consider for
                              binding.
                            properties:
                              matchExpressions:
                                description: matchExpressions is a list of label selector
                                  requirements. The requirements are ANDed.
                                items:
                                  description: A label selector requirement is a selector
                                    that contains values, a key, and an operator that
                                    relates the key and values.
                                  properties:
                                    key:
                                      description: key is the label key that the selector
                                        applies to.
                                      type: string
                                    operator:
                                      description: operator represents a key's relationship
                                        to a set of values. Valid operators are In,
                                        NotIn, Exists and DoesNotExist.
                                      type: string
                                    values:
                                      description: values is an array of string values.
                                        If the operator is In or NotIn, the values
                                        array must be non-empty. If the operator is
                                        Exists or DoesNotExist, the values array must
                                        be empty. This array is replaced during a
                                        strategic merge patch.
                                      items:
                                        type: string
                                      type: array
                                  required:
                                  - key
                                  - operator
                                  type: object
                                type: array
                              matchLabels:
                                additionalProperties:
                                  type: string
                                description: matchLabels is a map of {key,value} pairs.
                                  A single {key,value} in the matchLabels map is equivalent
                                  to an element of matchExpressions, whose key field
                                  is "key", the operator is "In", and the values array
                                  contains only "value". The requirements are ANDed.
                                type: object
                            type: object
                            x-kubernetes-map-type: atomic
                          storageClassName:
                            description: 'Name of the StorageClass required by the
                              claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1'
                            type: string
                          volumeMode:
                            description: volumeMode defines what type of volume is
                              required by the claim. Value of Filesystem is implied
                              when not included in claim spec.
                            type: string
                          volumeName:
                            description: VolumeName is the binding reference to the
                              PersistentVolume backing this claim.
                            type: string
                        type: object
                    type: object
                  status:
                    description: DataVolumeStatus contains the current status of the
                      DataVolume
                    properties:
                      claimName:
                        description: ClaimName is the name of the underlying PVC used
                          by the DataVolume.
                        type: string
                      conditions:
                        items:
                          description: DataVolumeCondition represents the state of
                            a data volume condition.
                          properties:
                            lastHeartbeatTime:
                              format: date-time
                              type: string
                            lastTransitionTime:
                              format: date-time
                              type: string
                            message:
                              type: string
                            reason:
                              type: string
                            status:
                              type: string
                            type:
                              description: DataVolumeConditionType is the string representation
                                of known condition types
                              type: string
                          required:
                          - status
                          - type
                          type: object
                        type: array
                      phase:
                        description: Phase is the current phase of the data volume
                        type: string
                      progress:
                        description: DataVolumeProgress is the current progress of
                          the DataVolume transfer operation. Value between 0 and 100
                          inclusive, N/A if not available
                        type: string
                      restartCount:
                        description: RestartCount is the number of times the pod populating
                          the DataVolume has restarted
                        format: int32
                        type: integer
                    type: object
                required:
                - spec
                type: object
            required:
            - managedDataSource
            - schedule
            - template
            type: object
          status:
            description: DataImportCronStatus provides the most recently observed
              status of the DataImportCron
            properties:
              conditions:
                items:
                  description: DataImportCronCondition represents the state of a data
                    import cron condition
                  properties:
                    lastHeartbeatTime:
                      format: date-time
                      type: string
                    lastTransitionTime:
                      format: date-time
                      type: string
                    message:
                      type: string
                    reason:
                      type: string
                    status:
                      type: string
                    type:
                      description: DataImportCronConditionType is the string representation
                        of known condition types
                      type: string
                  required:
                  - status
                  - type
                  type: object
                type: array
              currentImports:
                description: CurrentImports are the imports in progress. Currently
                  only a single import is supported.
                items:
                  description: ImportStatus of a currently in progress import
                  properties:
                    DataVolumeName:
                      description: DataVolumeName is the currently in progress import
                        DataVolume
                      type: string
                    Digest:
                      description: Digest of the currently imported image
                      type: string
                  required:
                  - DataVolumeName
                  - Digest
                  type: object
                type: array
              lastExecutionTimestamp:
                description: LastExecutionTimestamp is the time of the last polling
                format: date-time
                type: string
              lastImportTimestamp:
                description: LastImportTimestamp is the time of the last import
                format: date-time
                type: string
              lastImportedPVC:
                description: LastImportedPVC is the last imported PVC
                properties:
                  name:
                    description: The name of the source PVC
                    type: string
                  namespace:
                    description: The namespace of the source PVC
                    type: string
                required:
                - name
                - namespace
                type: object
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"datasource": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: datasources.cdi.kubevirt.io
spec:
  group: cdi.kubevirt.io
  names:
    categories:
    - all
    kind: DataSource
    listKind: DataSourceList
    plural: datasources
    shortNames:
    - das
    singular: datasource
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: DataSource references an import/clone source for a DataVolume
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: DataSourceSpec defines specification for DataSource
            properties:
              source:
                description: Source is the source of the data referenced by the DataSource
                properties:
                  pvc:
                    description: DataVolumeSourcePVC provides the parameters to create
                      a Data Volume from an existing PVC
                    properties:
                      name:
                        description: The name of the source PVC
                        type: string
                      namespace:
                        description: The namespace of the source PVC
                        type: string
                    required:
                    - name
                    - namespace
                    type: object
                  snapshot:
                    description: DataVolumeSourceSnapshot provides the parameters
                      to create a Data Volume from an existing VolumeSnapshot
                    properties:
                      name:
                        description: The name of the source VolumeSnapshot
                        type: string
                      namespace:
                        description: The namespace of the source VolumeSnapshot
                        type: string
                    required:
                    - name
                    - namespace
                    type: object
                type: object
            required:
            - source
            type: object
          status:
            description: DataSourceStatus provides the most recently observed status
              of the DataSource
            properties:
              conditions:
                items:
                  description: DataSourceCondition represents the state of a data
                    source condition
                  properties:
                    lastHeartbeatTime:
                      format: date-time
                      type: string
                    lastTransitionTime:
                      format: date-time
                      type: string
                    message:
                      type: string
                    reason:
                      type: string
                    status:
                      type: string
                    type:
                      description: DataSourceConditionType is the string representation
                        of known condition types
                      type: string
                  required:
                  - status
                  - type
                  type: object
                type: array
              source:
                description: Source is the current source of the data referenced by
                  the DataSource
                properties:
                  pvc:
                    description: DataVolumeSourcePVC provides the parameters to create
                      a Data Volume from an existing PVC
                    properties:
                      name:
                        description: The name of the source PVC
                        type: string
                      namespace:
                        description: The namespace of the source PVC
                        type: string
                    required:
                    - name
                    - namespace
                    type: object
                  snapshot:
                    description: DataVolumeSourceSnapshot provides the parameters
                      to create a Data Volume from an existing VolumeSnapshot
                    properties:
                      name:
                        description: The name of the source VolumeSnapshot
                        type: string
                      namespace:
                        description: The namespace of the source VolumeSnapshot
                        type: string
                    required:
                    - name
                    - namespace
                    type: object
                type: object
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"datavolume": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: datavolumes.cdi.kubevirt.io
spec:
  group: cdi.kubevirt.io
  names:
    categories:
    - all
    kind: DataVolume
    listKind: DataVolumeList
    plural: datavolumes
    shortNames:
    - dv
    - dvs
    singular: datavolume
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - description: The phase the data volume is in
      jsonPath: .status.phase
      name: Phase
      type: string
    - description: Transfer progress in percentage if known, N/A otherwise
      jsonPath: .status.progress
      name: Progress
      type: string
    - description: The number of times the transfer has been restarted.
      jsonPath: .status.restartCount
      name: Restarts
      type: integer
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1beta1
    schema:
      openAPIV3Schema:
        description: DataVolume is an abstraction on top of PersistentVolumeClaims
          to allow easy population of those PersistentVolumeClaims with relation to
          VirtualMachines
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: DataVolumeSpec defines the DataVolume type specification
            properties:
              checkpoints:
                description: Checkpoints is a list of DataVolumeCheckpoints, representing
                  stages in a multistage import.
                items:
                  description: DataVolumeCheckpoint defines a stage in a warm migration.
                  properties:
                    current:
                      description: Current is the identifier of the snapshot created
                        for this checkpoint.
                      type: string
                    previous:
                      description: Previous is the identifier of the snapshot from
                        the previous checkpoint.
                      type: string
                  required:
                  - current
                  - previous
                  type: object
                type: array
              contentType:
                description: 'DataVolumeContentType options: "kubevirt", "archive"'
                enum:
                - kubevirt
                - archive
                type: string
              finalCheckpoint:
                description: FinalCheckpoint indicates whether the current DataVolumeCheckpoint
                  is the final checkpoint.
                type: boolean
              preallocation:
                description: Preallocation controls whether storage for DataVolumes
                  should be allocated in advance.
                type: boolean
              priorityClassName:
                description: PriorityClassName for Importer, Cloner and Uploader pod
                type: string
              pvc:
                description: PVC is the PVC specification
                properties:
                  accessModes:
                    description: 'accessModes contains the desired access modes the
                      volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1'
                    items:
                      type: string
                    type: array
                  dataSource:
                    description: 'dataSource field can be used to specify either:
                      * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)
                      * An existing PVC (PersistentVolumeClaim) If the provisioner
                      or an external controller can support the specified data source,
                      it will create a new volume based on the contents of the specified
                      data source. When the AnyVolumeDataSource feature gate is enabled,
                      dataSource contents will be copied to dataSourceRef, and dataSourceRef
                      contents will be copied to dataSource when dataSourceRef.namespace
                      is not specified. If the namespace is specified, then dataSourceRef
                      will not be copied to dataSource.'
                    properties:
                      apiGroup:
                        description: APIGroup is the group for the resource being
                          referenced. If APIGroup is not specified, the specified
                          Kind must be in the core API group. For any other third-party
                          types, APIGroup is required.
                        type: string
                      kind:
                        description: Kind is the type of resource being referenced
                        type: string
                      name:
                        description: Name is the name of resource being referenced
                        type: string
                    required:
                    - kind
                    - name
                    type: object
                    x-kubernetes-map-type: atomic
                  dataSourceRef:
                    description: 'dataSourceRef specifies the object from which to
                      populate the volume with data, if a non-empty volume is desired.
                      This may be any object from a non-empty API group (non core
                      object) or a PersistentVolumeClaim object. When this field is
                      specified, volume binding will only succeed if the type of the
                      specified object matches some installed volume populator or
                      dynamic provisioner. This field will replace the functionality
                      of the dataSource field and as such if both fields are non-empty,
                      they must have the same value. For backwards compatibility,
                      when namespace isn''t specified in dataSourceRef, both fields
                      (dataSource and dataSourceRef) will be set to the same value
                      automatically if one of them is empty and the other is non-empty.
                      When namespace is specified in dataSourceRef, dataSource isn''t
                      set to the same value and must be empty. There are three important
                      differences between dataSource and dataSourceRef: * While dataSource
                      only allows two specific types of objects, dataSourceRef allows
                      any non-core object, as well as PersistentVolumeClaim objects.
                      * While dataSource ignores disallowed values (dropping them),
                      dataSourceRef preserves all values, and generates an error if
                      a disallowed value is specified. * While dataSource only allows
                      local objects, dataSourceRef allows objects in any namespaces.
                      (Beta) Using this field requires the AnyVolumeDataSource feature
                      gate to be enabled. (Alpha) Using the namespace field of dataSourceRef
                      requires the CrossNamespaceVolumeDataSource feature gate to
                      be enabled.'
                    properties:
                      apiGroup:
                        description: APIGroup is the group for the resource being
                          referenced. If APIGroup is not specified, the specified
                          Kind must be in the core API group. For any other third-party
                          types, APIGroup is required.
                        type: string
                      kind:
                        description: Kind is the type of resource being referenced
                        type: string
                      name:
                        description: Name is the name of resource being referenced
                        type: string
                      namespace:
                        description: Namespace is the namespace of resource being
                          referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant
                          object is required in the referent namespace to allow that
                          namespace's owner to accept the reference. See the ReferenceGrant
                          documentation for details. (Alpha) This field requires the
                          CrossNamespaceVolumeDataSource feature gate to be enabled.
                        type: string
                    required:
                    - kind
                    - name
                    type: object
                  resources:
                    description: 'resources represents the minimum resources the volume
                      should have. If RecoverVolumeExpansionFailure feature is enabled
                      users are allowed to specify resource requirements that are
                      lower than previous value but must still be higher than capacity
                      recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources'
                    properties:
                      claims:
                        description: "Claims lists the names of resources, defined
                          in spec.resourceClaims, that are used by this container.
                          \n This is an alpha field and requires enabling the DynamicResourceAllocation
                          feature gate. \n This field is immutable."
                        items:
                          description: ResourceClaim references one entry in PodSpec.ResourceClaims.
                          properties:
                            name:
                              description: Name must match the name of one entry in
                                pod.spec.resourceClaims of the Pod where this field
                                is used. It makes that resource available inside a
                                container.
                              type: string
                          required:
                          - name
                          type: object
                        type: array
                        x-kubernetes-list-map-keys:
                        - name
                        x-kubernetes-list-type: map
                      limits:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Limits describes the maximum amount of compute
                          resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                      requests:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Requests describes the minimum amount of compute
                          resources required. If Requests is omitted for a container,
                          it defaults to Limits if that is explicitly specified, otherwise
                          to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                    type: object
                  selector:
                    description: selector is a label query over volumes to consider
                      for binding.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                    x-kubernetes-map-type: atomic
                  storageClassName:
                    description: 'storageClassName is the name of the StorageClass
                      required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1'
                    type: string
                  volumeMode:
                    description: volumeMode defines what type of volume is required
                      by the claim. Value of Filesystem is implied when not included
                      in claim spec.
                    type: string
                  volumeName:
                    description: volumeName is the binding reference to the PersistentVolume
                      backing this claim.
                    type: string
                type: object
              source:
                description: Source is the src of the data for the requested DataVolume
                properties:
                  blank:
                    description: DataVolumeBlankImage provides the parameters to create
                      a new raw blank image for the PVC
                    type: object
                  gcs:
                    description: DataVolumeSourceGCS provides the parameters to create
                      a Data Volume from an GCS source
                    properties:
                      secretRef:
                        description: SecretRef provides the secret reference needed
                          to access the GCS source
                        type: string
                      url:
                        description: URL is the url of the GCS source
                        type: string
                    required:
                    - url
                    type: object
                  http:
                    description: DataVolumeSourceHTTP can be either an http or https
                      endpoint, with an optional basic auth user name and password,
                      and an optional configmap containing additional CAs
                    properties:
                      certConfigMap:
                        description: CertConfigMap is a configmap reference, containing
                          a Certificate Authority(CA) public key, and a base64 encoded
                          pem certificate
                        type: string
                      extraHeaders:
                        description: ExtraHeaders is a list of strings containing
                          extra headers to include with HTTP transfer requests
                        items:
                          type: string
                        type: array
                      secretExtraHeaders:
                        description: SecretExtraHeaders is a list of Secret references,
                          each containing an extra HTTP header that may include sensitive
                          information
                        items:
                          type: string
                        type: array
                      secretRef:
                        description: SecretRef A Secret reference, the secret should
                          contain accessKeyId (user name) base64 encoded, and secretKey
                          (password) also base64 encoded
                        type: string
                      url:
                        description: URL is the URL of the http(s) endpoint
                        type: string
                    required:
                    - url
                    type: object
                  imageio:
                    description: DataVolumeSourceImageIO provides the parameters to
                      create a Data Volume from an imageio source
                    properties:
                      certConfigMap:
                        description: CertConfigMap provides a reference to the CA
                          cert
                        type: string
                      diskId:
                        description: DiskID provides id of a disk to be imported
                        type: string
                      secretRef:
                        description: SecretRef provides the secret reference needed
                          to access the ovirt-engine
                        type: string
                      url:
                        description: URL is the URL of the ovirt-engine
                        type: string
                    required:
                    - diskId
                    - url
                    type: object
                  pvc:
                    description: DataVolumeSourcePVC provides the parameters to create
                      a Data Volume from an existing PVC
                    properties:
                      name:
                        description: The name of the source PVC
                        type: string
                      namespace:
                        description: The namespace of the source PVC
                        type: string
                    required:
                    - name
                    - namespace
                    type: object
                  registry:
                    description: DataVolumeSourceRegistry provides the parameters
                      to create a Data Volume from an registry source
                    properties:
                      certConfigMap:
                        description: CertConfigMap provides a reference to the Registry
                          certs
                        type: string
                      imageStream:
                        description: ImageStream is the name of image stream for import
                        type: string
                      pullMethod:
                        description: PullMethod can be either "pod" (default import),
                          or "node" (node docker cache based import)
                        type: string
                      secretRef:
                        description: SecretRef provides the secret reference needed
                          to access the Registry source
                        type: string
                      url:
                        description: 'URL is the url of the registry source (starting
                          with the scheme: docker, oci-archive)'
                        type: string
                    type: object
                  s3:
                    description: DataVolumeSourceS3 provides the parameters to create
                      a Data Volume from an S3 source
                    properties:
                      certConfigMap:
                        description: CertConfigMap is a configmap reference, containing
                          a Certificate Authority(CA) public key, and a base64 encoded
                          pem certificate
                        type: string
                      secretRef:
                        description: SecretRef provides the secret reference needed
                          to access the S3 source
                        type: string
                      url:
                        description: URL is the url of the S3 source
                        type: string
                    required:
                    - url
                    type: object
                  snapshot:
                    description: DataVolumeSourceSnapshot provides the parameters
                      to create a Data Volume from an existing VolumeSnapshot
                    properties:
                      name:
                        description: The name of the source VolumeSnapshot
                        type: string
                      namespace:
                        description: The namespace of the source VolumeSnapshot
                        type: string
                    required:
                    - name
                    - namespace
                    type: object
                  upload:
                    description: DataVolumeSourceUpload provides the parameters to
                      create a Data Volume by uploading the source
                    type: object
                  vddk:
                    description: DataVolumeSourceVDDK provides the parameters to create
                      a Data Volume from a Vmware source
                    properties:
                      backingFile:
                        description: BackingFile is the path to the virtual hard disk
                          to migrate from vCenter/ESXi
                        type: string
                      initImageURL:
                        description: InitImageURL is an optional URL to an image containing
                          an extracted VDDK library, overrides v2v-vmware config map
                        type: string
                      secretRef:
                        description: SecretRef provides a reference to a secret containing
                          the username and password needed to access the vCenter or
                          ESXi host
                        type: string
                      thumbprint:
                        description: Thumbprint is the certificate thumbprint of the
                          vCenter or ESXi host
                        type: string
                      url:
                        description: URL is the URL of the vCenter or ESXi host with
                          the VM to migrate
                        type: string
                      uuid:
                        description: UUID is the UUID of the virtual machine that
                          the backing file is attached to in vCenter/ESXi
                        type: string
                    type: object
                type: object
              sourceRef:
                description: SourceRef is an indirect reference to the source of data
                  for the requested DataVolume
                properties:
                  kind:
                    description: The kind of the source reference, currently only
                      "DataSource" is supported
                    type: string
                  name:
                    description: The name of the source reference
                    type: string
                  namespace:
                    description: The namespace of the source reference, defaults to
                      the DataVolume namespace
                    type: string
                required:
                - kind
                - name
                type: object
              storage:
                description: Storage is the requested storage specification
                properties:
                  accessModes:
                    description: 'AccessModes contains the desired access modes the
                      volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1'
                    items:
                      type: string
                    type: array
                  dataSource:
                    description: 'This field can be used to specify either: * An existing
                      VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)
                      * An existing PVC (PersistentVolumeClaim) * An existing custom
                      resource that implements data population (Alpha) In order to
                      use custom resource types that implement data population, the
                      AnyVolumeDataSource feature gate must be enabled. If the provisioner
                      or an external controller can support the specified data source,
                      it will create a new volume based on the contents of the specified
                      data source. If the AnyVolumeDataSource feature gate is enabled,
                      this field will always have the same contents as the DataSourceRef
                      field.'
                    properties:
                      apiGroup:
                        description: APIGroup is the group for the resource being
                          referenced. If APIGroup is not specified, the specified
                          Kind must be in the core API group. For any other third-party
                          types, APIGroup is required.
                        type: string
                      kind:
                        description: Kind is the type of resource being referenced
                        type: string
                      name:
                        description: Name is the name of resource being referenced
                        type: string
                    required:
                    - kind
                    - name
                    type: object
                    x-kubernetes-map-type: atomic
                  dataSourceRef:
                    description: 'Specifies the object from which to populate the
                      volume with data, if a non-empty volume is desired. This may
                      be any local object from a non-empty API group (non core object)
                      or a PersistentVolumeClaim object. When this field is specified,
                      volume binding will only succeed if the type of the specified
                      object matches some installed volume populator or dynamic provisioner.
                      This field will replace the functionality of the DataSource
                      field and as such if both fields are non-empty, they must have
                      the same value. For backwards compatibility, both fields (DataSource
                      and DataSourceRef) will be set to the same value automatically
                      if one of them is empty and the other is non-empty. There are
                      two important differences between DataSource and DataSourceRef:
                      * While DataSource only allows two specific types of objects,
                      DataSourceRef allows any non-core object, as well as PersistentVolumeClaim
                      objects. * While DataSource ignores disallowed values (dropping
                      them), DataSourceRef preserves all values, and generates an
                      error if a disallowed value is specified. (Beta) Using this
                      field requires the AnyVolumeDataSource feature gate to be enabled.'
                    properties:
                      apiGroup:
                        description: APIGroup is the group for the resource being
                          referenced. If APIGroup is not specified, the specified
                          Kind must be in the core API group. For any other third-party
                          types, APIGroup is required.
                        type: string
                      kind:
                        description: Kind is the type of resource being referenced
                        type: string
                      name:
                        description: Name is the name of resource being referenced
                        type: string
                      namespace:
                        description: Namespace is the namespace of resource being
                          referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant
                          object is required in the referent namespace to allow that
                          namespace's owner to accept the reference. See the ReferenceGrant
                          documentation for details. (Alpha) This field requires the
                          CrossNamespaceVolumeDataSource feature gate to be enabled.
                        type: string
                    required:
                    - kind
                    - name
                    type: object
                  resources:
                    description: 'Resources represents the minimum resources the volume
                      should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources'
                    properties:
                      claims:
                        description: "Claims lists the names of resources, defined
                          in spec.resourceClaims, that are used by this container.
                          \n This is an alpha field and requires enabling the DynamicResourceAllocation
                          feature gate. \n This field is immutable."
                        items:
                          description: ResourceClaim references one entry in PodSpec.ResourceClaims.
                          properties:
                            name:
                              description: Name must match the name of one entry in
                                pod.spec.resourceClaims of the Pod where this field
                                is used. It makes that resource available inside a
                                container.
                              type: string
                          required:
                          - name
                          type: object
                        type: array
                        x-kubernetes-list-map-keys:
                        - name
                        x-kubernetes-list-type: map
                      limits:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Limits describes the maximum amount of compute
                          resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                      requests:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Requests describes the minimum amount of compute
                          resources required. If Requests is omitted for a container,
                          it defaults to Limits if that is explicitly specified, otherwise
                          to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                    type: object
                  selector:
                    description: A label query over volumes to consider for binding.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                    x-kubernetes-map-type: atomic
                  storageClassName:
                    description: 'Name of the StorageClass required by the claim.
                      More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1'
                    type: string
                  volumeMode:
                    description: volumeMode defines what type of volume is required
                      by the claim. Value of Filesystem is implied when not included
                      in claim spec.
                    type: string
                  volumeName:
                    description: VolumeName is the binding reference to the PersistentVolume
                      backing this claim.
                    type: string
                type: object
            type: object
          status:
            description: DataVolumeStatus contains the current status of the DataVolume
            properties:
              claimName:
                description: ClaimName is the name of the underlying PVC used by the
                  DataVolume.
                type: string
              conditions:
                items:
                  description: DataVolumeCondition represents the state of a data
                    volume condition.
                  properties:
                    lastHeartbeatTime:
                      format: date-time
                      type: string
                    lastTransitionTime:
                      format: date-time
                      type: string
                    message:
                      type: string
                    reason:
                      type: string
                    status:
                      type: string
                    type:
                      description: DataVolumeConditionType is the string representation
                        of known condition types
                      type: string
                  required:
                  - status
                  - type
                  type: object
                type: array
              phase:
                description: Phase is the current phase of the data volume
                type: string
              progress:
                description: DataVolumeProgress is the current progress of the DataVolume
                  transfer operation. Value between 0 and 100 inclusive, N/A if not
                  available
                type: string
              restartCount:
                description: RestartCount is the number of times the pod populating
                  the DataVolume has restarted
                format: int32
                type: integer
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"objecttransfer": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: objecttransfers.cdi.kubevirt.io
spec:
  group: cdi.kubevirt.io
  names:
    kind: ObjectTransfer
    listKind: ObjectTransferList
    plural: objecttransfers
    shortNames:
    - ot
    - ots
    singular: objecttransfer
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    - description: The phase of the ObjectTransfer
      jsonPath: .status.phase
      name: Phase
      type: string
    name: v1beta1
    schema:
      openAPIV3Schema:
        description: ObjectTransfer is the cluster scoped object transfer resource
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ObjectTransferSpec specifies the source/target of the transfer
            properties:
              parentName:
                type: string
              source:
                description: TransferSource is the source of a ObjectTransfer
                properties:
                  apiVersion:
                    type: string
                  kind:
                    type: string
                  name:
                    type: string
                  namespace:
                    type: string
                  requiredAnnotations:
                    additionalProperties:
                      type: string
                    type: object
                required:
                - kind
                - name
                - namespace
                type: object
              target:
                description: TransferTarget is the target of an ObjectTransfer
                properties:
                  name:
                    type: string
                  namespace:
                    type: string
                type: object
            required:
            - source
            - target
            type: object
          status:
            description: ObjectTransferStatus is the status of the ObjectTransfer
            properties:
              conditions:
                items:
                  description: ObjectTransferCondition contains condition data
                  properties:
                    lastHeartbeatTime:
                      format: date-time
                      type: string
                    lastTransitionTime:
                      format: date-time
                      type: string
                    message:
                      type: string
                    reason:
                      type: string
                    status:
                      type: string
                    type:
                      description: ObjectTransferConditionType is the type of ObjectTransferCondition
                      type: string
                  required:
                  - status
                  - type
                  type: object
                type: array
              data:
                additionalProperties:
                  type: string
                description: Data is a place for intermediary state.  Or anything
                  really.
                type: object
              phase:
                description: Phase is the current phase of the transfer
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"storageprofile": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: storageprofiles.cdi.kubevirt.io
spec:
  group: cdi.kubevirt.io
  names:
    kind: StorageProfile
    listKind: StorageProfileList
    plural: storageprofiles
    singular: storageprofile
  scope: Cluster
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: StorageProfile provides a CDI specific recommendation for storage
          parameters
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: StorageProfileSpec defines specification for StorageProfile
            properties:
              claimPropertySets:
                description: ClaimPropertySets is a provided set of properties applicable
                  to PVC
                items:
                  description: ClaimPropertySet is a set of properties applicable
                    to PVC
                  properties:
                    accessModes:
                      description: 'AccessModes contains the desired access modes
                        the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1'
                      items:
                        type: string
                      type: array
                    volumeMode:
                      description: VolumeMode defines what type of volume is required
                        by the claim. Value of Filesystem is implied when not included
                        in claim spec.
                      type: string
                  type: object
                type: array
              cloneStrategy:
                description: CloneStrategy defines the preferred method for performing
                  a CDI clone
                type: string
              dataImportCronSourceFormat:
                description: DataImportCronSourceFormat defines the format of the
                  DataImportCron-created disk image sources
                type: string
            type: object
          status:
            description: StorageProfileStatus provides the most recently observed
              status of the StorageProfile
            properties:
              claimPropertySets:
                description: ClaimPropertySets computed from the spec and detected
                  in the system
                items:
                  description: ClaimPropertySet is a set of properties applicable
                    to PVC
                  properties:
                    accessModes:
                      description: 'AccessModes contains the desired access modes
                        the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1'
                      items:
                        type: string
                      type: array
                    volumeMode:
                      description: VolumeMode defines what type of volume is required
                        by the claim. Value of Filesystem is implied when not included
                        in claim spec.
                      type: string
                  type: object
                type: array
              cloneStrategy:
                description: CloneStrategy defines the preferred method for performing
                  a CDI clone
                type: string
              dataImportCronSourceFormat:
                description: DataImportCronSourceFormat defines the format of the
                  DataImportCron-created disk image sources
                type: string
              provisioner:
                description: The Storage class provisioner plugin name
                type: string
              storageClass:
                description: The StorageClass name for which capabilities are defined
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"volumeclonesource": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: volumeclonesources.cdi.kubevirt.io
spec:
  group: cdi.kubevirt.io
  names:
    kind: VolumeCloneSource
    listKind: VolumeCloneSourceList
    plural: volumeclonesources
    singular: volumeclonesource
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: VolumeCloneSource refers to a PVC/VolumeSnapshot of any storageclass/volumemode
          to be used as the source of a new PVC
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: VolumeCloneSourceSpec defines the Spec field for VolumeCloneSource
            properties:
              preallocation:
                description: Preallocation controls whether storage for the target
                  PVC should be allocated in advance.
                type: boolean
              priorityClassName:
                description: PriorityClassName is the priorityclass for the claim
                type: string
              source:
                description: Source is the src of the data to be cloned to the target
                  PVC
                properties:
                  apiGroup:
                    description: APIGroup is the group for the resource being referenced.
                      If APIGroup is not specified, the specified Kind must be in
                      the core API group. For any other third-party types, APIGroup
                      is required.
                    type: string
                  kind:
                    description: Kind is the type of resource being referenced
                    type: string
                  name:
                    description: Name is the name of resource being referenced
                    type: string
                required:
                - kind
                - name
                type: object
                x-kubernetes-map-type: atomic
            required:
            - source
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"volumeimportsource": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: volumeimportsources.cdi.kubevirt.io
spec:
  group: cdi.kubevirt.io
  names:
    kind: VolumeImportSource
    listKind: VolumeImportSourceList
    plural: volumeimportsources
    singular: volumeimportsource
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: VolumeImportSource works as a specification to populate PersistentVolumeClaims
          with data imported from an HTTP/S3/Registry/Blank/ImageIO/VDDK source
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: VolumeImportSourceSpec defines the Spec field for VolumeImportSource
            properties:
              contentType:
                description: ContentType represents the type of the imported data
                  (Kubevirt or archive)
                type: string
              preallocation:
                description: Preallocation controls whether storage for the target
                  PVC should be allocated in advance.
                type: boolean
              source:
                description: Source is the src of the data to be imported in the target
                  PVC
                properties:
                  blank:
                    description: DataVolumeBlankImage provides the parameters to create
                      a new raw blank image for the PVC
                    type: object
                  gcs:
                    description: DataVolumeSourceGCS provides the parameters to create
                      a Data Volume from an GCS source
                    properties:
                      secretRef:
                        description: SecretRef provides the secret reference needed
                          to access the GCS source
                        type: string
                      url:
                        description: URL is the url of the GCS source
                        type: string
                    required:
                    - url
                    type: object
                  http:
                    description: DataVolumeSourceHTTP can be either an http or https
                      endpoint, with an optional basic auth user name and password,
                      and an optional configmap containing additional CAs
                    properties:
                      certConfigMap:
                        description: CertConfigMap is a configmap reference, containing
                          a Certificate Authority(CA) public key, and a base64 encoded
                          pem certificate
                        type: string
                      extraHeaders:
                        description: ExtraHeaders is a list of strings containing
                          extra headers to include with HTTP transfer requests
                        items:
                          type: string
                        type: array
                      secretExtraHeaders:
                        description: SecretExtraHeaders is a list of Secret references,
                          each containing an extra HTTP header that may include sensitive
                          information
                        items:
                          type: string
                        type: array
                      secretRef:
                        description: SecretRef A Secret reference, the secret should
                          contain accessKeyId (user name) base64 encoded, and secretKey
                          (password) also base64 encoded
                        type: string
                      url:
                        description: URL is the URL of the http(s) endpoint
                        type: string
                    required:
                    - url
                    type: object
                  imageio:
                    description: DataVolumeSourceImageIO provides the parameters to
                      create a Data Volume from an imageio source
                    properties:
                      certConfigMap:
                        description: CertConfigMap provides a reference to the CA
                          cert
                        type: string
                      diskId:
                        description: DiskID provides id of a disk to be imported
                        type: string
                      secretRef:
                        description: SecretRef provides the secret reference needed
                          to access the ovirt-engine
                        type: string
                      url:
                        description: URL is the URL of the ovirt-engine
                        type: string
                    required:
                    - diskId
                    - url
                    type: object
                  registry:
                    description: DataVolumeSourceRegistry provides the parameters
                      to create a Data Volume from an registry source
                    properties:
                      certConfigMap:
                        description: CertConfigMap provides a reference to the Registry
                          certs
                        type: string
                      imageStream:
                        description: ImageStream is the name of image stream for import
                        type: string
                      pullMethod:
                        description: PullMethod can be either "pod" (default import),
                          or "node" (node docker cache based import)
                        type: string
                      secretRef:
                        description: SecretRef provides the secret reference needed
                          to access the Registry source
                        type: string
                      url:
                        description: 'URL is the url of the registry source (starting
                          with the scheme: docker, oci-archive)'
                        type: string
                    type: object
                  s3:
                    description: DataVolumeSourceS3 provides the parameters to create
                      a Data Volume from an S3 source
                    properties:
                      certConfigMap:
                        description: CertConfigMap is a configmap reference, containing
                          a Certificate Authority(CA) public key, and a base64 encoded
                          pem certificate
                        type: string
                      secretRef:
                        description: SecretRef provides the secret reference needed
                          to access the S3 source
                        type: string
                      url:
                        description: URL is the url of the S3 source
                        type: string
                    required:
                    - url
                    type: object
                  vddk:
                    description: DataVolumeSourceVDDK provides the parameters to create
                      a Data Volume from a Vmware source
                    properties:
                      backingFile:
                        description: BackingFile is the path to the virtual hard disk
                          to migrate from vCenter/ESXi
                        type: string
                      initImageURL:
                        description: InitImageURL is an optional URL to an image containing
                          an extracted VDDK library, overrides v2v-vmware config map
                        type: string
                      secretRef:
                        description: SecretRef provides a reference to a secret containing
                          the username and password needed to access the vCenter or
                          ESXi host
                        type: string
                      thumbprint:
                        description: Thumbprint is the certificate thumbprint of the
                          vCenter or ESXi host
                        type: string
                      url:
                        description: URL is the URL of the vCenter or ESXi host with
                          the VM to migrate
                        type: string
                      uuid:
                        description: UUID is the UUID of the virtual machine that
                          the backing file is attached to in vCenter/ESXi
                        type: string
                    type: object
                type: object
            type: object
          status:
            description: VolumeImportSourceStatus provides the most recently observed
              status of the VolumeImportSource
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
	"volumeuploadsource": `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: volumeuploadsources.cdi.kubevirt.io
spec:
  group: cdi.kubevirt.io
  names:
    kind: VolumeUploadSource
    listKind: VolumeUploadSourceList
    plural: volumeuploadsources
    singular: volumeuploadsource
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: VolumeUploadSource is a specification to populate PersistentVolumeClaims
          with upload data
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: VolumeUploadSourceSpec defines specification for VolumeUploadSource
            properties:
              contentType:
                description: ContentType represents the type of the upload data (Kubevirt
                  or archive)
                type: string
              preallocation:
                description: Preallocation controls whether storage for the target
                  PVC should be allocated in advance.
                type: boolean
            type: object
          status:
            description: VolumeUploadSourceStatus provides the most recently observed
              status of the VolumeUploadSource
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`,
}
