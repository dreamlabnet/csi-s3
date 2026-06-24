# CSI for S3

This is a Container Storage Interface ([CSI](https://github.com/container-storage-interface/spec/blob/master/spec.md)) for S3 (or S3 compatible) storage. This can dynamically allocate buckets and mount them via a fuse mount into any container.

## Configuration

Volume behaviour is controlled through the volume parameters passed to the driver: `mounter`, `options`, and `bucket`.

### Bucket

By default, csi-s3 will create a new bucket per volume. The bucket name will match that of the volume ID. If you want your volumes to live in a precreated bucket, you can set the `bucket` parameter:

```
mounter: s3fs
bucket: some-existing-bucket-name
```

If the bucket is specified, it will still be created if it does not exist on the backend. Every volume will get its own prefix within the bucket which matches the volume ID. When deleting a volume, also just the prefix will be deleted.

### Mounter

The default mounter is [s3fs](https://github.com/s3fs-fuse/s3fs-fuse).
There is also support for [rclone](https://rclone.org/commands/rclone_mount).

The mounter can be set with the `mounter` parameter.

As S3 is not a real file system there are some limitations to consider here.
Depending on what mounter you are using, you will have different levels of POSIX compability.
Also depending on what S3 storage backend you are using there are not always [consistency guarantees](https://github.com/gaul/are-we-consistent-yet#observed-consistency).

#### s3fs

* Almost full POSIX compatibility
* Good performance for big files, poor performance for small files
* Very slow for directories with a large number of files

#### rclone

* Poor POSIX compatibility
* Bad performance for big files, okayish performance for small files
* Doesn't create directory objects like s3fs
* May hang :-)

## Development

This project can be built like any other go application.

```bash
go get -u github.com/dreamlabnet/csi-s3
```

### Build executable

```bash
make build
```

### Tests

Currently the driver is tested by the [CSI Sanity Tester](https://github.com/kubernetes-csi/csi-test/tree/master/pkg/sanity). As end-to-end tests require S3 storage and a mounter like s3fs, this is best done in a docker container. A Dockerfile and the test script are in the `test` directory. The easiest way to run the tests is to just use the make command:

```bash
make test
```
