#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

script_dir=$(cd $(dirname $0); pwd)
wget -nv https://raw.githubusercontent.com/OPCFoundation/UA-Nodeset/master/Schema/NodeIds.csv -O "${script_dir}/NodeIds.csv"
wget -nv https://raw.githubusercontent.com/OPCFoundation/UA-Nodeset/master/Schema/StatusCode.csv -O "${script_dir}/StatusCode.csv"
wget -nv https://raw.githubusercontent.com/OPCFoundation/UA-Nodeset/master/Schema/Opc.Ua.Types.bsd -O "${script_dir}/Opc.Ua.Types.bsd"
wget -nv https://raw.githubusercontent.com/OPCFoundation/UA-Nodeset/master/Schema/Opc.Ua.PredefinedNodes.xml -O "${script_dir}/Opc.Ua.PredefinedNodes.xml"
