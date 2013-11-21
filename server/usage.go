package server

import (
	"strings"
)

// usage defines the message shown when a help flag is passed to etcd.
var usage = `
etcd

Usage:
  etcd -name <name>
  etcd -name <name> [-data-dir=<path>]
  etcd -h | -help
  etcd -version

Options:
  -h -help          Show this screen.
  --version         Show version.
  -f -force         Force a new configuration to be used.
  -config=<path>    Path to configuration file.
  -name=<name>      Name of this node in the etcd cluster.
  -data-dir=<path>  Path to the data directory.
  -cors=<origins>   Comma-separated list of CORS origins.
  -v                Enabled verbose logging.
  -vv               Enabled very verbose logging.

Cluster Configuration Options:
  -peers=<peers>      Comma-separated list of peers (ip + port) in the cluster.
  -peers-file=<path>  Path to a file containing the peer list.

Client Communication Options:
  -addr=<host:port>   The public host:port used for client communication.
  -bind-addr=<host>   The listening hostname used for client communication.
  -ca-file=<path>     Path to the client CA file.
  -cert-file=<path>   Path to the client cert file.
  -key-file=<path>    Path to the client key file.

Peer Communication Options:
  -peer-addr=<host:port>  The public host:port used for peer communication.
  -peer-bind-addr=<host>  The listening hostname used for peer communication.
  -peer-ca-file=<path>    Path to the peer CA file.
  -peer-cert-file=<path>  Path to the peer cert file.
  -peer-key-file=<path>   Path to the peer key file.

Other Options:
  -max-result-buffer   Max size of the result buffer.
  -max-retry-attempts  Number of times a node will try to join a cluster.
  -max-cluster-size    Maximum number of nodes in the cluster.
  -snapshot            Open or close the snapshot.
  -snapshot-count      Number of transactions before issuing a snapshot.
`

// Usage returns the usage message for etcd.
func Usage() string {
	return strings.TrimSpace(usage)
}
