## 综述

> 工具地址：https://github.com/siovanus/ontology-tool/tree/ontology-tool-demo
> 所需包：github.com/ontio/ontology-tool，是https://github.com/siovanus/ontology-tool的ontology-tool分支

使用该工具包，首先选择功能，修改params目录里面对应名称的json文件，运行./main -t <方法名> 即可

## 方法

1.RegIdWithPublicKey

注册一个绑定钱包的ontid

RegIdWithPublicKey.json:

```json
{
    "Path": "wallets/admin/wallet.dat"     //钱包路径
}
```

2.AssignFuncsToRole

将注册节点的权限授予角色，只有ontid管理员可以调用，该步骤只需执行一次

AssignFuncsToRole.json:

```json
{
  "Path": "wallets/admin/wallet.dat"   //ontid管理员钱包路径
}
```

3.AssignOntIDsToRole

将某个ontid指定为角色，使其可以拥有注册节点的权限， 只有ontid管理员可以调用

AssignOntIDsToRole.json

```json
{
  "Path1": "wallets/admin/wallet.dat",   //ontid管理员钱包路径
  "Ontid": ["did:ont:AMAx993nE6NEqZjwBssUfopxnnvTdob9ij"]   //待授权的ontid
}
```

4.RegisterCandidate

注册成为网络节点，需要授权的ontid和对应的钱包，参数为数组，可以一次提交多个注册请求

RegisterCandidate.json

```json
{
   "Path": ["wallets/peer1/wallet.dat"],  //钱包路径
   "PeerPubkey": ["03acea758e49b87a03b3dbf29e3055857ce7a4673ea864e640ed8f13d43861da41"], //节点公钥
   "InitPos": [10000]  //节点权重（质押的ont数量）
}
```

5.UnRegisterCandidate

取消注册请求

UnRegisterCandidate.json

```json
{
  "Path": "wallets/peer1-owner/wallet.dat",    //钱包路径
  "PeerPubkey": "0253ccfd439b29eca0fe90ca7c6eaa1f98572a054aa2d1d56e72ad96c466107a85"    //节点公钥
}
```

6.ApproveCandidate

理事会同意节点注册请求，需要理事会至少五个钱包签名

ApproveCandidate.json

```json
{
  "Path": ["wallets/peer1/wallet.dat","wallets/peer2/wallet.dat","wallets/peer3/wallet.dat","wallets/peer4/wallet.dat","wallets/peer5/wallet.dat","wallets/peer6/wallet.dat","wallets/peer7/wallet.dat"],   //理事会钱包路径
  "PeerPubkey": ["03338d1313cb78653bda3580474da8d1981c124b254ef9d474ada926784e49539c"]  //要同意的节点公钥
}
```

7.RejectCandidate

理事会拒绝节点注册请求，需要理事会至少五个钱包签名

RejectCandidate.json

```json
{
  "Path": ["wallets/peer1/wallet.dat","wallets/peer2/wallet.dat","wallets/peer3/wallet.dat","wallets/peer4/wallet.dat","wallets/peer5/wallet.dat","wallets/peer6/wallet.dat","wallets/peer7/wallet.dat"],   //理事会钱包路径
  "PeerPubkey": "03338d1313cb78653bda3580474da8d1981c124b254ef9d474ada926784e49539c"   //要同意的节点公钥
}
```

8.QuitNode

已加入网络的节点退出网络

QuitNode.json

```json
{
  "Path": ["wallets/peer1-owner/wallet.dat"],   //节点钱包
  "PeerPubkey": ["0253ccfd439b29eca0fe90ca7c6eaa1f98572a054aa2d1d56e72ad96c466107a85"]  //节点公钥
}
```

9.BlackNode

理事会强行拉黑某个节点，踢出网络并不允许再注册，需要理事会至少五个钱包签名

BlackNode.json

```json
{
  "Path": ["wallets/peer1/wallet.dat","wallets/peer2/wallet.dat","wallets/peer3/wallet.dat","wallets/peer4/wallet.dat","wallets/peer5/wallet.dat","wallets/peer6/wallet.dat","wallets/peer7/wallet.dat"],   //理事会钱包路径
  "PeerPubkeyList": ["0253ccfd439b29eca0fe90ca7c6eaa1f98572a054aa2d1d56e72ad96c466107a85"]   //要拉黑的节点公钥
}
```

10.DestroyContract

理事会强行删除某本智能合约，需要理事会至少五个钱包签名

DestroyContract.json

```json
{
  "Path": ["wallets/peer1/wallet.dat","wallets/peer2/wallet.dat","wallets/peer3/wallet.dat","wallets/peer4/wallet.dat","wallets/peer5/wallet.dat","wallets/peer6/wallet.dat","wallets/peer7/wallet.dat"],   //理事会钱包路径
  "ContractAddress": "0100000000000000000000000000000000000000"          //要删除的智能合约地址
}
```

11.GetPeerPoolMap

获取当前网络中的节点，状态0代表正在申请中的节点，1代表候选节点，2代表共识节点，3、4代表退出中的节点，该方法无需json

## 流程

首先通过RegIdWithPublicKey注册ontid管理员的ontid，再通过AssignFuncsToRole将权限赋予角色。

用户通过RegIdWithPublicKey注册自己的ontid，ontid管理员通过AssignOntIDsToRole授予其权限，用户通过RegisterCandidate注册成为节点，理事会通过ApproveCandidate同意该申请。
