# Media Offload Notes

An analysis of media offload options for serverless WordPress sites running on Knative.

## Plugins
There are a few plugins that appear to be able to offload media uploads like [S3 Uploads](https://github.com/humanmade/S3-Uploads).

Many plugins appear to require uploads to be persisted to the filesystem and are offloaded later. This would be possible to achieve with an empty volume mounted to the uploads directory but that would expose the possibility of a race with Knative service scale downs.

### CDN
Honestly though, instead of a vendor specific solution, I'm left to wonder what a nice open source (possibly Knative) CDN might look like. Maybe with a reverse proxy in the nginx config for the WordPress service so files appear to serve from the same domain?

## NFS
We could mount an nfs volume to the uploads directory but that would also open up the posibility of writing to the filesystem from multiple containers at once which is kind of the antithesis of this example.
