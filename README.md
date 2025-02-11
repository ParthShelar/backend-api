### A gRPC-based Employee Management Service that uses Google Cloud Spanner for database storage. This service allows clients to:
-- Create Employees

### Project Structure
```
backendapi/
|-- cmd/                             # Application entry point
|   |-- main.go                      # Main file to start the gRPC server and Prometheus server
|-- handler/                         # gRPC handlers (middleware layer)
|   |-- handler.go                   # Handler implementation for gRPC methods
|-- repository/                      # Core data persistence and operations
|   |-- repository.go                # Implementation of data operations and spanner db interaction
|-- proto/                           # gRPC proto definition and generated files
|   |-- backend-api.proto            # Proto file defining gRPC services and messages
|   |-- generate_proto.sh            # Script to generate gRPC code using protoc compiler
|   |-- backend-api/                 # Generated code from .proto file
|   |   |-- backend-api.pb.go        # Generated protobuf code
|   |   |-- backend-api_grpc.pb.go   # Generated gRPC service code
|-- go.mod                           # Go module dependencies
|-- go.sum                           # Checksum file for dependencies
|-- README.md                        # Documentation
```
Sample request and response body for inserting employee in database:

### Request: 

```bash
{
  "first_name": "Bill",
  "last_name": "Gates",
  "joining_date": "2024-01-10",
  "salary": 95000
}
```

### Respose:

```bash
{
    "employee": {
        "id": "2305843009213693952",
        "first_name": "Bill",
        "last_name": "Gates",
        "joining_date": "2024-01-10",
        "salary": 95000
    }
}
```


# Bazel Integration with gRPC, Golang, and Spanner DB

## Bazel Version
**8.0.1**
## Go Version
**1.23.5**

## Project Overview
In this project, we have tried to integrate **gRPC, Golang, and Spanner DB** with **Bazel** using **Gazelle**.

## Manual repository setup
# Download the package
wget https://proxy.golang.org/github.com/prometheus/client_golang/@v/v1.20.5.zip

# Extract to a temporary location
unzip v1.20.5.zip -d temp_prometheus/

# Navigate into the extracted package
cd temp_prometheus/github.com/prometheus/client_golang/

# Repackage the patched module
zip -r patched_prometheus.zip temp_prometheus

# Compute the SHA256 checksum
sha256sum patched_prometheus.zip

## Generating Build Files with Gazelle
To generate Bazel build files using Gazelle, use the following command:

```sh
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_repositories
```

## Running Bazel Commands
To clean, build, and run the project using Bazel, use the following commands:

```sh
bazel clean --expunge
bazel build //cmd:backend-api
bazel run //cmd:backend-api
```

## Error Encountered
After executing `bazel run`, the following error occurs:
```
bazel run //cmd:backend-api                                                                      
Starting local Bazel server (8.0.1) and connecting to it...
ERROR: no such package '@@[unknown repo 'com_github_klauspost_compress' requested from @@+prometheus_deps+com_github_prometheus_client_golang]//zstd': The repository '@@[unknown repo 'com_github_klauspost_compress' requested from @@+prometheus_deps+com_github_prometheus_client_golang]' could not be resolved: No repository visible as '@com_github_klauspost_compress' from repository '@@+prometheus_deps+com_github_prometheus_client_golang'
ERROR: /private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/external/+prometheus_deps+com_github_prometheus_client_golang/prometheus/promhttp/BUILD.bazel:3:11: no such package '@@[unknown repo 'com_github_klauspost_compress' requested from @@+prometheus_deps+com_github_prometheus_client_golang]//zstd': The repository '@@[unknown repo 'com_github_klauspost_compress' requested from @@+prometheus_deps+com_github_prometheus_client_golang]' could not be resolved: No repository visible as '@com_github_klauspost_compress' from repository '@@+prometheus_deps+com_github_prometheus_client_golang' and referenced by '@@+prometheus_deps+com_github_prometheus_client_golang//prometheus/promhttp:promhttp'
ERROR: Analysis of target '//cmd:backend-api' failed; build aborted: Analysis failed
INFO: Elapsed time: 8.529s, Critical Path: 0.01s
INFO: 1 process: 1 internal.
ERROR: Build did NOT complete successfully
ERROR: Build failed. Not running target
FAILED: 
    Fetching repository @@+google_cloud_deps+com_google_cloud_go_spanner; starting
    Fetching repository @@+grpc_deps+org_golang_google_grpc; starting
    Fetching ...te/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/external/+google_cloud_deps+com_google_cloud_go_spanner; Extracting v1.73.0.zip
    Fetching /private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/external/+grpc_deps+org_golang_google_grpc; Extracting v1.69.2.zip
```
After executing `bazel build`, the following error occurs:

```
Computing main repo mapping: 
Loading: 
Loading: 0 packages loaded
Analyzing: target //cmd:backend-api (0 packages loaded, 0 targets configured)
Analyzing: target //cmd:backend-api (0 packages loaded, 0 targets configured)
[0 / 1] [Prepa] BazelWorkspaceStatusAction stable-status.txt
INFO: Analyzed target //cmd:backend-api (0 packages loaded, 0 targets configured).
ERROR: /private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/BUILD.bazel:4:11: GoCompilePkg external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/go_default_library.a failed: (Exit 1): builder failed: error executing GoCompilePkg command (from target @@_main~google_cloud_deps~com_google_cloud_go_secretmanager//:go_default_library) 
  (cd /private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main && \
  exec env - \
    CGO_ENABLED=1 \
    GOARCH=arm64 \
    GODEBUG='winsymlink=0' \
    GOEXPERIMENT=nocoverageredesign \
    GOOS=darwin \
    GOPATH='' \
    GOROOT_FINAL=GOROOT \
    GOTOOLCHAIN=local \
    PATH=/usr/bin:external/bazel_tools~cc_configure_extension~local_config_cc:/bin \
    ZERO_AR_DATE=1 \
  bazel-out/darwin_arm64-opt-exec-ST-d57f47055a04/bin/external/rules_go~~go_sdk~backend_api__download_0/builder_reset/builder compilepkg -sdk external/rules_go~~go_sdk~backend_api__download_0 -goroot bazel-out/darwin_arm64-fastbuild/bin/external/rules_go~/stdlib_ -installsuffix darwin_arm64 -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/aliasshim/aliasshim.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/auxiliary.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/auxiliary_go123.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/doc.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/iam.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/iam_example_test.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client_example_go123_test.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client_example_test.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/resources.pb.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/service.pb.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/version.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/auxiliary.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/auxiliary_go123.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/doc.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client_example_go123_test.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client_example_test.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/resources.pb.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/service.pb.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/version.go -src external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/internal/version.go -importpath cloud.google.com/go/secretmanager -p cloud.google.com/go/secretmanager -package_list bazel-out/darwin_arm64-opt-exec-ST-d57f47055a04/bin/external/rules_go~~go_sdk~backend_api__download_0/packages.txt -embedroot '' -embedroot bazel-out/darwin_arm64-fastbuild/bin -embedlookupdir external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/aliasshim -embedlookupdir external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1 -embedlookupdir external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb -embedlookupdir external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2 -embedlookupdir external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb -embedlookupdir external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/internal -embedlookupdir external/_main~google_cloud_deps~com_google_cloud_go_secretmanager -lo bazel-out/darwin_arm64-fastbuild/bin/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/go_default_library.a -o bazel-out/darwin_arm64-fastbuild/bin/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/go_default_library.x -gcflags '')
# Configuration: 096dcc84165363e69a851ebe8131b032f5448c94ddc4951775429dc78e79f898
# Execution platform: @@platforms//host:host

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
compilepkg: missing strict dependencies:
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/auxiliary.go: import of "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/auxiliary.go: import of "google.golang.org/api/iterator"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/auxiliary.go: import of "google.golang.org/genproto/googleapis/cloud/location"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/auxiliary_go123.go: import of "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/auxiliary_go123.go: import of "github.com/googleapis/gax-go/v2/iterator"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/auxiliary_go123.go: import of "google.golang.org/genproto/googleapis/cloud/location"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/doc.go: import of "google.golang.org/api/option"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/iam.go: import of "cloud.google.com/go/iam"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/iam_example_test.go: import of "cloud.google.com/go/secretmanager/apiv1"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "cloud.google.com/go/iam/apiv1/iampb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "github.com/googleapis/gax-go/v2"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "google.golang.org/api/googleapi"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "google.golang.org/api/iterator"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "google.golang.org/api/option"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "google.golang.org/api/option/internaloption"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "google.golang.org/api/transport/grpc"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "google.golang.org/api/transport/http"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "google.golang.org/genproto/googleapis/cloud/location"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "google.golang.org/grpc"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "google.golang.org/grpc/codes"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "google.golang.org/protobuf/encoding/protojson"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client.go: import of "google.golang.org/protobuf/proto"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client_example_go123_test.go: import of "cloud.google.com/go/secretmanager/apiv1"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client_example_go123_test.go: import of "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client_example_go123_test.go: import of "google.golang.org/genproto/googleapis/cloud/location"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client_example_test.go: import of "cloud.google.com/go/iam/apiv1/iampb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client_example_test.go: import of "cloud.google.com/go/secretmanager/apiv1"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client_example_test.go: import of "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client_example_test.go: import of "google.golang.org/api/iterator"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secret_manager_client_example_test.go: import of "google.golang.org/genproto/googleapis/cloud/location"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/resources.pb.go: import of "google.golang.org/genproto/googleapis/api/annotations"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/resources.pb.go: import of "google.golang.org/protobuf/reflect/protoreflect"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/resources.pb.go: import of "google.golang.org/protobuf/runtime/protoimpl"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/resources.pb.go: import of "google.golang.org/protobuf/types/known/durationpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/resources.pb.go: import of "google.golang.org/protobuf/types/known/timestamppb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/service.pb.go: import of "cloud.google.com/go/iam/apiv1/iampb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/service.pb.go: import of "google.golang.org/genproto/googleapis/api/annotations"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/service.pb.go: import of "google.golang.org/grpc"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/service.pb.go: import of "google.golang.org/grpc/codes"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/service.pb.go: import of "google.golang.org/grpc/status"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/service.pb.go: import of "google.golang.org/protobuf/reflect/protoreflect"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/service.pb.go: import of "google.golang.org/protobuf/runtime/protoimpl"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/service.pb.go: import of "google.golang.org/protobuf/types/known/emptypb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/secretmanagerpb/service.pb.go: import of "google.golang.org/protobuf/types/known/fieldmaskpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1/version.go: import of "cloud.google.com/go/secretmanager/internal"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/auxiliary.go: import of "cloud.google.com/go/secretmanager/apiv1beta2/secretmanagerpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/auxiliary.go: import of "google.golang.org/api/iterator"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/auxiliary.go: import of "google.golang.org/genproto/googleapis/cloud/location"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/auxiliary_go123.go: import of "cloud.google.com/go/secretmanager/apiv1beta2/secretmanagerpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/auxiliary_go123.go: import of "github.com/googleapis/gax-go/v2/iterator"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/auxiliary_go123.go: import of "google.golang.org/genproto/googleapis/cloud/location"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/doc.go: import of "google.golang.org/api/option"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "cloud.google.com/go/iam/apiv1/iampb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "cloud.google.com/go/secretmanager/apiv1beta2/secretmanagerpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "github.com/googleapis/gax-go/v2"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "google.golang.org/api/googleapi"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "google.golang.org/api/iterator"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "google.golang.org/api/option"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "google.golang.org/api/option/internaloption"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "google.golang.org/api/transport/grpc"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "google.golang.org/api/transport/http"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "google.golang.org/genproto/googleapis/cloud/location"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "google.golang.org/grpc"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "google.golang.org/grpc/codes"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "google.golang.org/protobuf/encoding/protojson"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client.go: import of "google.golang.org/protobuf/proto"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client_example_go123_test.go: import of "cloud.google.com/go/secretmanager/apiv1beta2"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client_example_go123_test.go: import of "cloud.google.com/go/secretmanager/apiv1beta2/secretmanagerpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client_example_go123_test.go: import of "google.golang.org/genproto/googleapis/cloud/location"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client_example_test.go: import of "cloud.google.com/go/iam/apiv1/iampb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client_example_test.go: import of "cloud.google.com/go/secretmanager/apiv1beta2"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client_example_test.go: import of "cloud.google.com/go/secretmanager/apiv1beta2/secretmanagerpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client_example_test.go: import of "google.golang.org/api/iterator"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secret_manager_client_example_test.go: import of "google.golang.org/genproto/googleapis/cloud/location"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/resources.pb.go: import of "google.golang.org/genproto/googleapis/api/annotations"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/resources.pb.go: import of "google.golang.org/protobuf/reflect/protoreflect"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/resources.pb.go: import of "google.golang.org/protobuf/runtime/protoimpl"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/resources.pb.go: import of "google.golang.org/protobuf/types/known/durationpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/resources.pb.go: import of "google.golang.org/protobuf/types/known/timestamppb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/service.pb.go: import of "cloud.google.com/go/iam/apiv1/iampb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/service.pb.go: import of "google.golang.org/genproto/googleapis/api/annotations"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/service.pb.go: import of "google.golang.org/grpc"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/service.pb.go: import of "google.golang.org/grpc/codes"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/service.pb.go: import of "google.golang.org/grpc/status"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/service.pb.go: import of "google.golang.org/protobuf/reflect/protoreflect"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/service.pb.go: import of "google.golang.org/protobuf/runtime/protoimpl"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/service.pb.go: import of "google.golang.org/protobuf/types/known/emptypb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/secretmanagerpb/service.pb.go: import of "google.golang.org/protobuf/types/known/fieldmaskpb"
	/private/var/tmp/_bazel_parth/7fb6f4592b46e7a9ca8f8e14c870cf04/sandbox/darwin-sandbox/184/execroot/_main/external/_main~google_cloud_deps~com_google_cloud_go_secretmanager/temp_secretmanager/cloud.google.com/go/cloud.google.com/go/secretmanager@v1.14.1/apiv1beta2/version.go: import of "cloud.google.com/go/secretmanager/internal"
No dependencies were provided.
Check that imports in Go sources match importpath attributes in deps.
Target //cmd:backend-api failed to build
INFO: Elapsed time: 0.209s, Critical Path: 0.08s
INFO: 10 processes: 10 internal.
ERROR: Build did NOT complete successfully
```

## Issue
After fixing the above error by manually adding a `BUILD` file in that repository, another error pops up related to a different library stating:

```
no such package "package_name" found in directory
```
