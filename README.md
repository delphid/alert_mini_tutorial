# Do you wonder how Prometheus, Alertmanager and Grafana work together?

### `alert_mini_tutorial` aims to combine these 3 and a demo app into one docker-compose, and then show the relations in a minimal manner.

### This demo app acts as two roles: both the metric producer and the alert receiver.

Clone this repo, cd into the repo directory, and run:
```shell
docker-compose up
```

In your browser, open [Prometheus](http://localhost:9090), [Alertmanager](http://localhost:9093) and [Grafana](http://localhost:3000).

Tweak with your gauges, for example:

- tune gauge_**A** **UP** with label proc=**1**
```shell
make a-up proc=1
```

- tune gauge_**B** **DOWN**
```shell
make b-down
```
...

- Check on Prometheus how your gauge value and rule status are changing.
- Check on Alertmanager about your alert status and labels.
- Check out your log produced by app.go, and see what alert message it's receiving from Alertmanager.
```shell
less log/log
```

In Grafana, add your Prometheus as a data source, with URL `http://prometheus:9090`, and with `Manage alerts via Alerting UI` enabled.

Go to [Grafana Alert Page](http://localhost:3000/alerting/list), and VOILA, your Prometheus rules are also displayed on your Grafana.
