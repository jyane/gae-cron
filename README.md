# GAE Cron for Cloud PubSub
The Google App Engine Application which publishes 4 publications to 4 topics,
every mintes, every hours, every days and every weeks.

## Topics

```
projects/:project_id/topics/daily-tick
projects/:project_id/topics/hourly-tick
projects/:project_id/topics/minutely-tick
projects/:project_id/topics/weekly-tick
```

## Usage

``` sh
# use pyenv, virtualenv if you want
pip install -t lib -r requirements.txt
gcloud app deploy app.yaml \cron.yaml

# check
gcloud app browse
```

## LICENSE
Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
