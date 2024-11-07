[开发者文档](https://solana.com/docs/intro/quick-start)

[JSON RPC Method](https://solana.com/docs/rpc)

# 地址

Solana 区块链上的每个公共地址都使用长度介于 32 到 44 个字符之间的字符串。 这些地址中的每一个都遵循 BIP44 标准，因此使用 base-58 字符集。

对于 Solana 而言，助记词通过 [Ed25519 算法](https://en.wikipedia.org/wiki/EdDSA?ref=defiplot.com#Ed25519)“派生”（或“计算”）私钥。 然后可以使用该私钥来派生公钥。

Solana 上的派生路径通常按 BIP44 标准来生成网络上的地址。

- `m/44'/501'` 由 Solana CLI 用于生成“根密钥”
- `m/44'/501'/0'/0'` 用于大多数基于网络或基于浏览器的钱包

在 Solana 链上，通过派生路径从私钥派生出的公钥即为钱包地址。

# 账户模型

Solana 上所有数据都存储在账户中，**帐户需要 SOL 押金**，而押金在帐户关闭时可全额退还。每个账户可以通过其唯一地址来识别，该地址为 32 字节的 Ed25519 公钥格式。

> 要在链上存储数据，必须将一定量的 SOL 转移到账户中，SOL 数量与帐户中存储的数据大小成正比。
> 
> 程序（即智能合约）是存储可执行代码的无状态账户。


每个帐户最多存储 10MB 的数据，这些数据可以包含可执行程序代码或程序状态。AccountInfo 数据结构如下

| Field      | type            | desc                                            |
|------------|-----------------|-------------------------------------------------|
| data       | bytes           | 如果账户是程序（智能合约），则存储可执行程序代码。此字段通常称为 “account data” |
| executable | boolean         | 标识账户是否是程序                                       |
| lamports   | number          | 以 lamport 为单位记录的 SOL 余额（1 SOL = 10^9 lamport）   |
| owner      | program address | 指定拥有该账户的程序的公有密钥（程序 ID）                          |

# 交易&指令

在 Solana 上通过 transaction 来和网络交互，transaction 包含一个或多个 instructions（指令），每个 instruction 处理特定操作，一笔 transaction 中的指令按其添加顺序执行，且具备原子性。

