# These paths must be absolute

# METASHARD_ID will be used to identify a shard ID as metachain
export METASHARD_ID=4294967295

# Path to elrond-go. Determined automatically. Do not change.
export ELRONDDIR=$(dirname $(dirname $ELRONDTESTNETSCRIPTSDIR))

# Enable the Elrond Proxy. Note that this is a private repository
# (elrond-proxy-go).
export USE_PROXY=1

# Enable the Elrond Transaction Generator. Note that this is a private
# repository (elrond-txgen-go).
export USE_TXGEN=0

# Path where the testnet will be instantiated. This folder is assumed to not
# exist, but it doesn't matter if it already does. It will be created if not,
# anyway.
export TESTNETDIR="$HOME/YollloNetwork/testnet"

# Path to elrond-deploy-go, branch: master. Default: near elrond-go.
export CONFIGGENERATORDIR="$(dirname $ELRONDDIR)/yolllo-deploy-go/cmd/filegen"
export CONFIGGENERATOR="$CONFIGGENERATORDIR/filegen"    # Leave unchanged.
export CONFIGGENERATOROUTPUTDIR="output"

# Path to the executable node. Leave unchanged unless well justified.
export NODEDIR="$ELRONDDIR/cmd/node"
export NODE="$NODEDIR/node"     # Leave unchanged

# Path to the executable seednode. Leave unchanged unless well justified.
export SEEDNODEDIR="$ELRONDDIR/cmd/seednode"
export SEEDNODE="$SEEDNODEDIR/seednode"   # Leave unchanged.

# Niceness value of the Seednode, Observer Nodes and Validator Nodes. Leave
# blank to not adjust niceness.
export NODE_NICENESS=10

# Start a watcher daemon for each validator node, which restarts the node if it
# is suffled out of its shard.
export NODE_WATCHER=0

# Delays after running executables.
export SEEDNODE_DELAY=5
export GENESIS_DELAY=30
export HARDFORK_DELAY=900 #15 minutes enough to take export and gracefully close
export NODE_DELAY=60

export GENESIS_STAKE_TYPE="direct" #'delegated' or 'direct' as in direct stake

#if set to 1, each observer will turn off the antiflooding capability, allowing spam in our network
export OBSERVERS_ANTIFLOOD_DISABLE=0

# Shard structure
export SHARDCOUNT=1
export SHARD_VALIDATORCOUNT=3
export SHARD_OBSERVERCOUNT=1
export SHARD_CONSENSUS_SIZE=3

# Metashard structure
export META_VALIDATORCOUNT=3
export META_OBSERVERCOUNT=1
export META_CONSENSUS_SIZE=$META_VALIDATORCOUNT

# ALWAYS_NEW_CHAINID will generate a fresh new chain ID each time start.sh/config.sh is called
export ALWAYS_NEW_CHAINID=1

# HYSTERESIS defines the hysteresis value for number of nodes in shard
export HYSTERESIS=0.0

# ALWAYS_NEW_APP_VERSION will set a new version each time the node will be compiled
export ALWAYS_NEW_APP_VERSION=0

# ALWAYS_UPDATE_CONFIGS will re-generate configs (toml + json) each time ./start.sh
# Set this variable to 0 when testing bootstrap from storage or other edge cases where you do not want a fresh new config
# each time.
export ALWAYS_UPDATE_CONFIGS=1

# Always rebuild Arwen from its sources and copy the executable to the testnet folder.
export ALWAYS_BUILD_ARWEN=1

# Ports used by the Nodes
export PORT_SEEDNODE="9999"
export PORT_ORIGIN_OBSERVER="21100"
export PORT_ORIGIN_OBSERVER_REST="10000"
export PORT_ORIGIN_VALIDATOR="21500"
export PORT_ORIGIN_VALIDATOR_REST="9500"

# Address of the Seednode. Will be written to the p2p.toml file of the Nodes
export P2P_SEEDNODE_ADDRESS="/ip4/127.0.0.1/tcp/$PORT_SEEDNODE/p2p/16Uiu2HAkw5SNNtSvH1zJiQ6Gc3WoGNSxiyNueRKe6fuAuh57G3Bk"


# UI configuration profiles

# Use tmux or not. If set to 1, only 2 terminal windows will be opened, and
# tmux will be used to display the running executables using split windows.
# Recommended. Tmux needs to be installed.
export USETMUX=1

# Log level for the logger in the Node.
export LOGLEVEL="*:INFO"


if [ "$TESTNETMODE" == "debug" ]; then
  LOGLEVEL="*:DEBUG,api:INFO"
fi

if [ "$TESTNETMODE" == "trace" ]; then
  LOGLEVEL="*:TRACE"
fi

########################################################################
# Proxy configuration

# Path to elrond-proxy-go, branch: master. Default: near elrond-go.
export PROXYDIR="$(dirname $ELRONDDIR)/yolllo-proxy-go/cmd/proxy"
export PROXY=$PROXYDIR/proxy    # Leave unchanged.

export PORT_PROXY="7950"
export PROXY_DELAY=10



########################################################################
# TxGen configuration

# Path to elrond-txgen-go. Default: near elrond-go.
export TXGENDIR="$(dirname $ELRONDDIR)/elrond-txgen-go/cmd/txgen"
export TXGEN=$TXGENDIR/txgen    # Leave unchanged.

export PORT_TXGEN="7951"

export TXGEN_SCENARIOS_LINE='Scenarios = ["basic", "erc20", "esdt"]'

# Number of accounts to be generated by txgen
export NUMACCOUNTS="250"

# Whether txgen should regenerate its accounts when starting, or not.
# Recommended value is 1, but 0 is useful to run the txgen a second time, to
# continue a testing session on the same accounts.
export TXGEN_REGENERATE_ACCOUNTS=0

# COPY_BACK_CONFIGS when set to 1 will copy back the configs and keys to the ./cmd/node/config directory
# in order to have a node in the IDE that can run a node in debug mode but in the same network with the rest of the nodes
# this option greatly helps the debugging process when running a small system test
export COPY_BACK_CONFIGS=0
# SKIP_VALIDATOR_IDX when setting a value greater than -1 will not launch the validator with the provided index
export SKIP_VALIDATOR_IDX=-1
# SKIP_OBSERVER_IDX when setting a value greater than -1 will not launch the observer with the provided index
export SKIP_OBSERVER_IDX=-1

# USE_HARDFORK will prepare the nodes to run the hardfork process, if needed
export USE_HARDFORK=1

# Load local overrides, .gitignored
LOCAL_OVERRIDES="$ELRONDTESTNETSCRIPTSDIR/local.sh"
if [ -f "$LOCAL_OVERRIDES" ]; then
  source "$ELRONDTESTNETSCRIPTSDIR/local.sh"
fi

# Leave unchanged.
let "total_observer_count = $SHARD_OBSERVERCOUNT * $SHARDCOUNT + $META_OBSERVERCOUNT"
export TOTAL_OBSERVERCOUNT=$total_observer_count

# to enable the full archive feature on the observers, please use the --full-archive flag
export EXTRA_OBSERVERS_FLAGS=""

# Leave unchanged.
let "total_node_count = $SHARD_VALIDATORCOUNT * $SHARDCOUNT + $META_VALIDATORCOUNT + $TOTAL_OBSERVERCOUNT"
export TOTAL_NODECOUNT=$total_node_count
