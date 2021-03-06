#!/bin/bash
# run clients and collect stats locally
set -eu

LABEL=${LABEL:-"lola-flows"}
CONF=${CONF:-flows.env}
DB=${DB:-lola}

base=$(dirname $0)
. ${base}/${CONF}

STATS=${STATS:-/root/share/stats/$LABEL}

schedule clients \
	--stats-enabled \
	--stats-dir=${STATS} \
	--log-tag=C \
	--flows-dirs=${FLOWS}
