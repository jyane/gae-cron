# ~GAE Cron for Cloud PubSub~ Obsolete
GCP now provides a cron service.

Original: https://github.com/GoogleCloudPlatform/reliable-task-scheduling-compute-engine-sample

The Google App Engine Application which publishes 4 publications to 4 topics,
every mintes, every hours, every days and every weeks.

Using this service for scheduling and Google Cloud Pub/Sub for distributed messaging, you can build an application to reliably schedule tasks which can trigger Google Cloud Functions.

## Topics

```
projects/:project_id/topics/daily-tick
projects/:project_id/topics/hourly-tick
projects/:project_id/topics/minutely-tick
projects/:project_id/topics/weekly-tick
```

## Usage

``` sh
gcloud app deploy app.yaml \cron.yaml

# check
gcloud app browse
```
