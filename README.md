# sockstat_exporter

Linux listening ports Prometheus exporter

## Usage

```
  -listen string
    	Listen metrics server address. [env: LISTEN] (default ":9997")
  -v	Print version
```

## Metrics example

```
sockstat_listen{proto="tcp",address="0.0.0.0",port="22",path="/usr/sbin/sshd"} 1
sockstat_listen{proto="tcp",address="0.0.0.0",port="50007",path="/usr/local/bin/cloud-torrent"} 1
sockstat_listen{proto="tcp",address="0.0.0.0",port="80",path="/usr/sbin/nginx"} 1
sockstat_listen{proto="tcp",address="0.0.0.0",port="81",path="/usr/sbin/nginx"} 1
sockstat_listen{proto="tcp",address="0.0.0.0",port="88",path="/usr/sbin/nginx"} 1
sockstat_listen{proto="tcp",address="0.0.0.0",port="888",path="/usr/sbin/nginx"} 1
sockstat_listen{proto="tcp",address="127.0.0.53",port="53",path="/usr/lib/systemd/systemd-resolved"} 1
sockstat_listen{proto="tcp6",address="2000::1",port="888",path="/usr/sbin/nginx"} 1
sockstat_listen{proto="tcp6",address="::",port="22",path="/usr/sbin/sshd"} 1
sockstat_listen{proto="tcp6",address="::",port="3000",path="/usr/local/bin/cloud-torrent"} 1
sockstat_listen{proto="tcp6",address="::",port="4000",path="/usr/share/grafana/bin/grafana-server"} 1
sockstat_listen{proto="tcp6",address="::",port="50007",path="/usr/local/bin/cloud-torrent"} 1
sockstat_listen{proto="tcp6",address="::",port="7879",path="/usr/local/bin/rclone"} 1
sockstat_listen{proto="tcp6",address="::",port="80",path="/usr/sbin/nginx"} 1
sockstat_listen{proto="tcp6",address="::",port="9090",path="/bin/prometheus"} 1
sockstat_listen{proto="tcp6",address="::",port="9100",path="/bin/node_exporter"} 1
sockstat_listen{proto="tcp6",address="::",port="9997",path="/bin/sockstat_exporter"} 1
sockstat_listen{proto="udp",address="0.0.0.0",port="1900",path="/usr/local/bin/rclone"} 1
sockstat_listen{proto="udp",address="0.0.0.0",port="50007",path="/usr/local/bin/cloud-torrent"} 1
sockstat_listen{proto="udp",address="127.0.0.53",port="53",path="/usr/lib/systemd/systemd-resolved"} 1
sockstat_listen{proto="udp",address="192.168.1.100",port="68",path="/usr/lib/systemd/systemd-networkd"} 1
sockstat_listen{proto="udp6",address="::",port="50007",path="/usr/local/bin/cloud-torrent"} 1
sockstat_listen{proto="udp6",address="FE80::F64D:30FF:FE60:E13B",port="546",path="/usr/lib/systemd/systemd-networkd"} 1
```
